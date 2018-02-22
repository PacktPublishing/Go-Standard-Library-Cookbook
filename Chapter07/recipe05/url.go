package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {

	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "localhost"
	u.Path = "index.html"
	u.RawQuery = "id=1&name=John"
	u.User = url.UserPassword("admin", "1234")

	fmt.Printf("Assembled URL:\n%v\n\n\n", u)

	parsedURL, err := url.Parse(u.String())
	if err != nil {
		panic(err)
	}
	jsonURL, err := json.Marshal(parsedURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed URL:")
	fmt.Println(string(jsonURL))

}
