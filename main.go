package main

import (
	"html/template" // HTMLテンプレートを扱うためのパッケージ
	"net/http"      // HTTPサーバー機能を提供するパッケージ
)

// Todoリストを格納する空のスライス
var todoList []string

// "/todo" パスのハンドラ関数: Todoリストを表示
func handleTodo(w http.ResponseWriter, r *http.Request) {
    // templates/todo.htmlを解析してテンプレートを作成
    t, _ := template.ParseFiles("templates/todo.html")
    
    // todoListの内容をテンプレートに埋め込み、クライアントに返す
    t.Execute(w, todoList)
}

// "/add" パスのハンドラ関数: Todoリストに新しいタスクを追加
func handleAdd(w http.ResponseWriter, r *http.Request) {
    // フォームデータを解析
    r.ParseForm()
    
    // "todo"という名前のフォームの値を取得
    todo := r.Form.Get("todo")
    
    // 取得したタスクをtodoListに追加
    todoList = append(todoList, todo)
    
    // Todoリストのページを再表示
    handleTodo(w, r)
}

func main() {
    // ★ 静的ファイルのサーバー設定
    // "/static/" というURLパスを "static" ディレクトリに対応付け
    http.Handle("/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // ★ "/todo" パスにアクセスが来たら handleTodo 関数を実行
    http.HandleFunc("/todo", handleTodo)

    // ★ "/add" パスにアクセスが来たら handleAdd 関数を実行
    http.HandleFunc("/add", handleAdd)

    // ★ HTTPサーバーを起動し、ポート8080で待ち受け
    http.ListenAndServe(":8080", nil)
}
