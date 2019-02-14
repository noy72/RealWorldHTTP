// http.Client構造体を使用して，クッキーの送受信を行う
package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil) // cookie保存のインスタンス
	if err != nil {
		panic(err)
	}
	client := http.Client{ // cookie保存可能なインスタンス
		Jar: jar,
	}
	for i := 0; i < 2; i++ { // 2回目以降のアクセスにクッキーを使う
		// http.Getではなく，作成したクライアントのGet()メソッドでアクセスする
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
