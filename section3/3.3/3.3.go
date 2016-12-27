package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func sayhelloName (w http.ResponseWriter, r *http.Request){
	r.ParseForm()  //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意：如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello World!")
}

func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("method:", r.Method)    //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("section3/3.3/login.gtpl")
		if err != nil {
			fmt.Println(err)
		}

		crutime := time.Now().Unix()
		fmt.Println(crutime)

		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t.Execute(w, token)
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println(r.Form)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func noicon(w http.ResponseWriter, r *http.Request){
	fmt.Println("no icon")
}

func main() {
	http.HandleFunc("/", sayhelloName)  //设置访问的路由
	http.HandleFunc("/favicon.ico", noicon)
	http.HandleFunc("/login", login)    //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
