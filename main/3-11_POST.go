// 送信するファイルに任意のMIMEタイプを設定

package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer                // マルチパート部を組み立てた後のバイト列を格納するバッファ
	writer := multipart.NewWriter(&buffer) // マルチパートを組み立てるライター
	// バウンダリー文字列はmultipart.Writerオブジェクトが内部で作る
	writer.WriteField("name", "Michael Jackson") // ファイル以外のフィールドを登録

	// Read file-------------
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type:", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)
	//-----------------------

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
