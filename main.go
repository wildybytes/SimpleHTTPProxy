// ----------------------------------------------------------------------------
//  Name         : Simple HTTP Proxy
//  Desc         : Simple HTTP Proxy Using golang and Library from goproxy
//  Author       : Wildy Sheverando [WildyBytes]
//  Date         : 03-03-2023
//  License      : GNU General Public License V3
//  License Link : https://raw.githubusercontent.com/wildybytes/lcn/main/gplv3
// ----------------------------------------------------------------------------

package main

import (
	"fmt"
	"net/http"
	"github.com/elazarl/goproxy"
)

func main() {
	// Proxy Server Configuration
	username := "admin"
	password := "admin"
	port := 8080

	// Create new http proxy instance
	proxy := goproxy.NewProxyHttpServer()

	// Set verbose to active
	proxy.Verbose = true

	// Setting Basic auth dan logging untuk proxy connection
	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			user, pass, ok := req.BasicAuth()
			if !ok || user != username || pass != password {
				res := goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusProxyAuthRequired, "HTTP 401 Unauthorized")
				return req, res
			}

			// Add log for each connection
			fmt.Printf("Connection from %s to %s\n", ctx.Req.RemoteAddr, req.URL.Host)
			return req, nil
		})

	// Start the proxy server
	fmt.Printf("Proxy Server Listening on : %d...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), proxy); err != nil {
		panic(err)
	}
}
