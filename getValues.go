package main

// Goには標準搭載されているライブラリが多い。そこから抜き出し。
import (
	"fmt" //formatの略
	"log"
	"net/http" //Goのnet/httpパッケージは、HTTP通信に関するさまざまな処理を提供してくれる。
)

// リクエスト（r）とレスポンス（w）
 func hello(w http.ResponseWriter, r *http.Request) { 
     fmt.Fprint(w, "Hello, Web application!") 
 }

 func main() {
     http.HandleFunc("/", hello) //"/"というパスにリクエストが来たら、hello関数を実行
     err := http.ListenAndServe(":8080", nil) //httpサーバーを起動する関数、ポート番号を8080に設定
     //nil: デフォルトのマルチプレクサ（http.DefaultServeMux）を使用することを意味します。
     if err != nil {
         log.Fatal("failed to start : ", err)
     }
 }
