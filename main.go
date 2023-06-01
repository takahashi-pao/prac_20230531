package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 入力フォーム画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-form.gtpl"))

	// テンプレートに出力する値をマップにセット
	values := map[string]string{}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "user-form.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

// 入力内容の確認画面
func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-confirm.gtpl"))

	// テンプレートに出力する値をマップにセット
	values := map[string]string{
		"account": req.FormValue("account"),
		"name":    req.FormValue("name"),
		"passwd":  req.FormValue("passwd"),
	}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "user-confirm.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

// テンプレートハンドラーテスト
func gtplHandler(w http.ResponseWriter, r *http.Request) {

	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/sample.gtpl"))

	// テンプレートに出力する値をマップにセット
	values := map[string]string{
		"account": "user-0001",
		"name":    "山田太郎",
		"passwd":  "sample-pass",
	}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "sample.gtpl", values); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// ルートへのリクエストを"gtplHandler"関数で処理する
	http.HandleFunc("/", gtplHandler)

	// "user-form"へのリクエストを関数で処理する
	http.HandleFunc("/user-form", HandlerUserForm)

	// "user-confirm"へのリクエストを関数で処理する
	http.HandleFunc("/user-confirm", HandlerUserConfirm)

	// localhost:8080でサーバー処理開始
	http.ListenAndServe(":8080", nil)
}
