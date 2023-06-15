package main

import (
	"encoding/json"
	"fmt"
	"github.com/acheong08/OpenAIAuth/auth"
	"log"
	"net/http"
	"os"
)

func requestToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	username, eu := r.Form["username"]
	password, ep := r.Form["password"]
	if eu && ep {
		token := getToken(username[0], password[0])
		fmt.Fprintf(w, token) //这个写入到w的是输出到客户端的
	} else {
		fmt.Fprintf(w, "ERROR: BAD RESPONSE") //这个写入到w的是输出到客户端的
	}
}

func getToken(username string, password string) string {
	auth := auth.NewAuthenticator(username, password, os.Getenv("PROXY"))
	err := auth.Begin()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return fmt.Sprint("ERROR: ", err.Error.Error())
	}
	// if os.Getenv("PROXY") != "" {
	puid, err := auth.GetPUID()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return fmt.Sprint("ERROR: ", err.Error.Error())
	}
	println("PUID: " + puid)
	// }
	// JSON encode auth.GetAuthResult()
	result := auth.GetAuthResult()
	result_json, _ := json.Marshal(result)
	println(string(result_json))
	return token
}

func main() {
	http.HandleFunc("/token", requestToken) //设置访问的路由
	println("服务已启动... ")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
