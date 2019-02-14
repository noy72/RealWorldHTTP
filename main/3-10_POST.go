// curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888
// Go言語でマルチパートフォームをPOST送信

package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer                // マルチパート部を組み立てた後のバイト列を格納するバッファ
	writer := multipart.NewWriter(&buffer) // マルチパートを組み立てるライター
	// バウンダリー文字列はmultipart.Writerオブジェクトが内部で作る
	writer.WriteField("name", "Michael Jackson") // ファイル以外のフィールドを登録

	// Read file-------------
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile) // ファイルの全コンテンツをファイル書き込み用のio.Writerにコピー
	writer.Close()
	//-----------------------

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
