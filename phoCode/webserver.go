package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	content := `<!DOCTYPE html>
				<html>
					<head>
						<title>Sample Go Web Server</title>
					</head>
					<body>
						<h1>It Worked!</h1>
						<h1>Hello Go Backend!</h1>
					</body>
				</html>`
	io.WriteString(res, content)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}
