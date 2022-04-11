/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: go_web
   @Date: 2022/4/11 22:07
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Printf("Read file failed, err: %v\n", err)
		return
	}
	_, _ = fmt.Fprintln(w, string(b))

}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Http server failed, err: %v\n", err)
		return
	}
}
