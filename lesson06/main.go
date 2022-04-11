/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:17
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数 kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是 error 类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气！", nil
	}
	// 定义模板
	t := template.New("f.tmpl") // 创建一个名字是 f 的模板对象
	// 告诉模板引擎，我现在多了一个自定义的函数 kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	name := "小王子"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
