//curl -T main.go -H "Content-Type: text/plain" http://localhost:18888
// Go言語で任意のボディをPOST送信
package main

import (
	"log"
	"net/http"
	"os"
)

// ファイルから読み込んだ任意のコンテンツを送信
func main() {
	file, err := os.Open("3-8_POST.go")
	if err != nil {
		panic(err)
	}
	resp, _ := http.Post("http://localhost:18888", "text/plain", file)
	log.Println("Status:", resp.Status)
}
