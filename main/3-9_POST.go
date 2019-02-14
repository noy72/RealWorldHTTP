// Go言語で任意の文字列をPOST送信
package main

import (
	"log"
	"net/http"
	"strings"
)

// 任意の文字列を送信
func main() {
	reader := strings.NewReader("Text")
	resp, _ := http.Post("http://localhost:18888", "text/plain", reader)
	log.Println("Status:", resp.Status)
}
