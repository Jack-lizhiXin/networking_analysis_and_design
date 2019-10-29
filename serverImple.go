package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = `<html><body>
        <form action="#" method="post" />
            <p>用户名：<input type="text" name="name" /></p>
            <p>&ensp;&ensp;密码：<input type="text" name="pass" /></p>
            <input type="submit" value="登录" />
        </form>
        </body></html>`

var (
	hasLogin bool // 标记是否已登录过
	name, pass string
)

func show(w http.ResponseWriter) {
	io.WriteString(w, name)
	io.WriteString(w, "<p></p>")
	io.WriteString(w, pass)
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func handleFunc (w http.ResponseWriter, r *http.Request) {
	// 设置编码格式以防web端显示乱码
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if hasLogin {
		show(w)
	} else {
		//Method 代表指定的 http 方法
		switch r.Method {
			case "GET":
				// 将表单写入到w中
				io.WriteString(w, form)
			case "POST":
				name = r.PostFormValue("name")
				pass = r.PostFormValue("pass")
				hasLogin = true // 取到用户名/密码后修改标记
				show(w)
		}
	}
}

func main () {
	// 设置路由，如果访问 / ，则调用handleFunc方法
	http.HandleFunc("/", handleFunc)

	// 启动web服务，监听8080端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen server failed, error:", err)
		return
	}
}
