package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type Socket struct {
	Server_NetWorkType string
	Server_Address     string
	Delimiter          string
}

var socket = Socket{
	Server_NetWorkType: "tcp",
	Server_Address:     "127.0.0.1:8085",
	Delimiter:          "\t",
}

func StartUdpServer() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("本地地址: <%s> \n", listener.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("中转服务器退出,仍不影响peers间通信")
			return
		}
	}
}

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		fmt.Printf("收到的数据：%v\n", recv)

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func StartTCPServer() {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "0.0.0.0:9981")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	log.Printf("本地地址: <%s> \n", listen.Addr().String())
	peers := make([]net.Addr, 0, 2)
	conns := make([]net.Conn, 0, 2)
	// data := make([]byte, 1024)
	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen failed, err:%v\n", err)
			return
		}
		remoteAddr := conn.RemoteAddr()
		log.Printf("<%s>\n", remoteAddr.String())
		peers = append(peers, remoteAddr)
		conns = append(conns, conn)
		if len(peers) >= 2 {
			conns[0].Write([]byte(peers[1].String()))
			conns[1].Write([]byte(peers[0].String()))
			time.Sleep(time.Second * 8)
			listen.Close()
		}
	}
}
