package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main(){
	http.Handle("/css/",http.StripPrefix("/css/",http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/success",success)
	http.HandleFunc("/",add)
	http.ListenAndServe(":8080", nil)//サーバー起動&ハンドラ登録
}

func success(w http.ResponseWriter, r *http.Request){
	t :=template.Must(template.ParseFiles("success.gtpl"))

	//値を取得している
	values :=map[string]string{
		"name":r.FormValue("wordname"),
		"mean":r.FormValue("mean"),
	}

	if err :=t.ExecuteTemplate(w,"success.gtpl",values); err !=nil{
		fmt.Println(err)
	}
}

func add(w http.ResponseWriter, r *http.Request){
	t :=template.Must(template.ParseFiles("add.gtpl"))

	values :=map[string]string{}

	if err :=t.ExecuteTemplate(w,"add.gtpl",values); err !=nil{
		fmt.Println(err)
	}
}
