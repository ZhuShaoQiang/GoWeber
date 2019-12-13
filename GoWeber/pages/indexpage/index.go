package indexpage

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
)

// w是发送的数据，r是接受的数据，与客户端所有的交互都依赖于这两个变量
func Index(w http.ResponseWriter, r *http.Request) {

	type PageDatas struct{
		Name		string
	}

	pageDatas := PageDatas{
		Name: "朱少强",
	}

	//fmt.Fprintf(w, "你好，黑镜")
	fmt.Fprintf(os.Stdout, "[+] from %s Method is %s.\n", r.URL.Path, r.Method)
	t, err := template.ParseFiles("HTML/index/index.html")
	// 这个路径是根据运行 go run main.go 的路径确定的
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
	}
	t.Execute(w, pageDatas)
}
