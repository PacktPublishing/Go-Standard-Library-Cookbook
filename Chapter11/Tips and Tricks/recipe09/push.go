package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	log.Println("Staring server...")
	// Adding to mani Mux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if p, ok := w.(http.Pusher); ok {
			if err := p.Push("/app.css", nil); err != nil {
				fmt.Printf("Push err : %v", err)
			}
		}
		io.WriteString(w,
			`<html>
				<head>
					<link rel="stylesheet" type="text/css" href="app.css">
				</head>
				<body>
						<p>Hello</p>
				</body>
			</html>`)
	})
	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,
			`p {
				text-align: left;
				color: red;
			}`)
	})

	if err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}

}
