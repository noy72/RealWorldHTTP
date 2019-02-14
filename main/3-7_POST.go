// crul -d test=value http://localhost:18888
// x-www-form-urlencoded形式のPOSTメソッドの送信
package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, _ := http.PostForm("http://localhost:18888", values)
	log.Println("Status:", resp.Status)
}
