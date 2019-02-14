// curl http://localhost:18888
// GETメソッドでアクセスし，ボディを取得
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:18888")
	// resp : http.Response type
	defer resp.Body.Close() // pythonのwithみたいなもの
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
