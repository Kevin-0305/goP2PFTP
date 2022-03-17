package main

import (
	"flag"
	"goP2PFTP/client"
	"goP2PFTP/server"
)

func main() {
	var task string
	// var host string
	// var port int
	var cport int
	flag.StringVar(&task, "t", "client", "需要启动服务端还是客户端,如果为服务端请输入 s,客户端输入 c,默认为客户端")
	// flag.StringVar(&host, "h", "127.0.0.1", "主机地址默认地址为本地")
	// flag.IntVar(&port, "P", 9981, "服务端端口号，默认为9981")
	flag.IntVar(&cport, "CP", 9982, "客户端端口号，默认为9982")
	flag.Parse()
	if task == "s" {
		server.StartTCPServer()
	}
	if task == "c" {
		client.StartTCPClient(cport)
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/gogf/greuse"
// )

// // We can create two processes with this code.
// // Do some requests, then watch the output of the console.
// func main() {
// 	listener, err := greuse.Listen("tcp", ":8881")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer listener.Close()

// 	server := &http.Server{}
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "gid: %d, pid: %d\n", os.Getgid(), os.Getpid())
// 	})

// 	panic(server.Serve(listener))
// }
