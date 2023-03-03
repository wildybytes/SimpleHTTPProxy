# Simple HTTP Proxy
This is a simple HTTP proxy program implemented in Go using the "elazarl/goproxy" library. It allows users to intercept HTTP requests and forward them to a remote server through the proxy server. It also provides basic authentication using a username and password, and logs each connection to the server.

## Installation
Before running this program, make sure you have the following dependencies installed:
  - Go programming language
  - "elazarl/goproxy" library
To install the "elazarl/goproxy" library, run the following command in your terminal:

## Usage
To use this program, follow these steps:
1. Open your terminal and navigate to the directory where the program is located.
2. Run the following command to start the proxy server:
3. The server will start running on port 8080.
4. To test the server, configure your web browser to use the proxy server with the following settings:
  - IP address: [ your server ip ]
  - Port: 8080 [ can change in main.go ]
  - Username: admin [ can change in main.go ]
  - Password: admin [ can change in main.go ]
5. Open your web browser and visit any website. The proxy server will intercept the request and ask for authentication.
6. Enter the correct username and password (admin/admin) to access the website through the proxy server.
7. The server will log each connection in the terminal window.

## Contributing
Contributions are welcome! If you find any bugs or have suggestions for improvement, please create an issue or pull request.

## License
Simple HTTP Proxy is licensed under the GNU General Public License V3. See the LICENSE file for more information.
