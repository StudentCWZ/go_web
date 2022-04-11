/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:23
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 渲染模板
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 渲染模板
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板(模板继承)
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 渲染模板
	name := "小王子"
	err = t.ExecuteTemplate(w, "index2.tmpl", name)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}
func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板(模板继承)
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 渲染模板
	name := "七米"
	err = t.ExecuteTemplate(w, "home2.tmpl", name)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
