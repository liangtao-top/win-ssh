package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

var (
	host     string
	port     int
	user     string
	password string
	sh       string
)

func initCommand() {
	flag.StringVar(&host, "h", "127.0.0.1", "远程Linux主机IP地址")
	flag.StringVar(&host, "host", "127.0.0.1", "远程Linux主机IP地址")
	flag.IntVar(&port, "port", 22, "远程Linux主机端口")
	flag.StringVar(&user, "user", "root", "远程Linux主机用户名")
	flag.StringVar(&user, "u", "root", "远程Linux主机用户名")
	flag.StringVar(&password, "password", "root", "远程Linux主机密码")
	flag.StringVar(&password, "p", "root", "远程Linux主机密码")
	flag.StringVar(&sh, "sh", "", "登录后预执行指令")
	flag.StringVar(&sh, "cmd", "", "登录后预执行指令")
}

func main() {
	initCommand()
	flag.Parse()
	fmt.Println("Windows SSH 远程登录工具 Author: TaoGe <liangtao.gz@foxmail.com>")
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("User:", user)
	var pwd string
	for range password {
		pwd += "*"
	}
	fmt.Println("Password:", pwd)
	fmt.Println("sh:", sh)

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         3 * time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", host, port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()
	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	//执行远程命令
	combo, err := session.CombinedOutput("whoami; " + sh)
	if err != nil {
		log.Fatal("远程执行cmd 失败", err)
	}
	log.Println(string(combo))
}
