package main

import "fmt"

func main() {
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }

	// client := &http.Client{Transport: tr}
	// _, err := client.Get("https://golang.org/")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	password := "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"

	fmt.Println("Doing something with: ", password)
}
