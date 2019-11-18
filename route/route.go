package route

import (
	"net/http"
	"log"
	"os"
	"regexp"

	"taskGoRebuild/pages/indexpage"
	"taskGoRebuild/pages/notFound404"
)
var (
	indexRE = regexp.MustCompile(`^/$`)		// 监听主页
)

func route(w http.ResponseWriter, r *http.Request) {
	switch {
	case indexRE.MatchString(r.URL.Path):		// 匹配主页网址
		indexpage.Index(w, r)
	default:
		notFound404.NotFound404(w, r)
	}
}

// go语言的每一个请求都是goroutine完成的
func handleURL() {
	// 注册一个路由，监听根路径，第二个参数是：func(ResponseWriter, *Request)
	http.HandleFunc("/", route)
}

func StartHTTPServer() {
	handleURL()
	log.Println("[+] 网站已经开启，点击访问网址访问：\nhttp://127.0.0.1:8080/")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("[-] 服务开启失败,因为: ", err)
		os.Exit(1)	//
	}
	os.Exit(0)	//
}
