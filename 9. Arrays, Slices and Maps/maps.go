package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}
	fmt.Println(websites)
	websites["Twitter"] = "https://www.twitter.com"
	delete(websites, "Facebook")
	fmt.Println(websites)
}
