package route

import (
	"net/http"
	"log"
	"os"
	"regexp"

	"GoWeber/pages/indexpage"
	"GoWeber/pages/notFound404"
)
var (
	indexRE = regexp.MustCompile(`^/$`)		// 监听主页的正则表达式
)

func route(w http.ResponseWriter, r *http.Request) {
	log.Println("[+] now is:", r.URL)
	switch {		// 这里主要为了写各个网址
	case indexRE.MatchString(r.URL.Path):		// 匹配主页网址
		indexpage.Index(w, r)
	default:		// 如果都没有，就返回 404 页面
		notFound404.NotFound404(w, r)
	}
}

// go语言的每一个请求都是goroutine完成的
func handleURL() {
	// 注册一个路由，监听根路径，第二个参数是：func(ResponseWriter, *Request)，第三个参数时数据库对象
	http.HandleFunc("/", route)
}

func StartHTTPServer() {
	handleURL()
	log.Println("[+] 网站已经开启，访问任意机器网址访问，例如：\n\thttp://127.0.0.1:80/")
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("[-] 服务开启失败,因为: ", err)
		os.Exit(1)	//
	}
	os.Exit(0)	//
}
