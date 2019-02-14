// curl --head http://localhost:18888
// HEADメソッドでヘッダーを取得
package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
