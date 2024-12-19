package main

import (
	"html/template" // HTMLテンプレートを扱うためのパッケージ
	"log"           // ログ出力用パッケージ
	"net/http"      // HTTPサーバー機能を提供するパッケージ
)

// Todoリストを格納する変数（初期値は空のスライス＊JavaScriptの配列みたいなもの）
var todoList []string

// "/todo" にアクセスされた時に実行されるハンドラ関数
func handleTodo(w http.ResponseWriter, r *http.Request) {
	// "templates/todo.html" テンプレートファイルをパース（解析）してtに格納
	t, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		// テンプレートファイルの読み込みに失敗したらエラーログを出力し、処理を終了
		log.Println("Error loading template:", err)
		return
	}

	// テンプレートを実行し、HTTPレスポンスにtodoListの内容を埋め込んで返す
	// t.Executeは、第一引数にレスポンスライター、第二引数にデータ（todoList）を渡す
	err = t.Execute(w, todoList)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

func main() {
	// ★ 固定のTodoリストを作成
	// append関数は、スライス（配列のようなもの）に要素を追加する
	// ここでは、todoListに3つのタスクを追加している
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// ★ 静的ファイル（CSSやJS、画像など）を提供する設定
	// FileServer: 指定されたディレクトリのファイルをクライアントに返す
	// StripPrefix: URLの「/static/」部分を削除してディレクトリ内のファイルに対応付ける
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// ★ "/todo"パスにアクセスされた場合にhandleTodo関数を実行するよう設定
	http.HandleFunc("/todo", handleTodo)

	// ★ HTTPサーバーを起動し、ポート8080でリクエストを待ち受ける
	// 第二引数のnilはデフォルトのマルチプレクサ（http.DefaultServeMux）を使うことを意味する
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// サーバーの起動に失敗した場合にエラーログを出力して終了
		log.Fatal("failed to start : ", err)
	}
}
