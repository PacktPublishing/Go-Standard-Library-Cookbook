package main

import "net/http"
import "html/template"
import "fmt"

func main() {
	fmt.Println("Server is starting...")
	tpl, err := template.ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, "John Doe")
		if err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
