/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:14
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模版
	t, err := template.ParseFiles("./hello.tmpl") // 请勿刻舟求剑
	if err != nil {
		fmt.Printf("Parse template failed, err: %v\n", err)
		return
	}
	// 3. 宣染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Render template failed, err: %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Http server start failed, err: %v\n", err)
		return
	}
}
