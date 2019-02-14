// crul -G --data-urlencode "query=hello world" http://localhost:18888
// GETメソッドでクエリーを送信
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{ // クエリー文字列の作成
		"query": {"hello world"},
	}

	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode()) // 末尾にクエリーを追加
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
