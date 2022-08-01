# Windows SSH 远程登录工具
- 解决 windows 下 bat 批处理命令无法处理输入密码的问题。
- 解决 vbs 运行效率低下，反应慢的问题。

## 安装／运行

### step 1: 拉取代码
~~~
git clone https://github.com/liangtao-top/win-ssh.git
~~~

### step 2: 编译启动
~~~
cd win-ssh
./dist/win_ssh.exe -h xxx.xxx.xx.xx -p 123456 -sh "cd /iot && docker-compose ps"
~~~
## 运行参数
~~~
  -cmd string
        登录后预执行指令
  -h string
        远程Linux主机IP地址 (default "127.0.0.1")
  -host string
        远程Linux主机IP地址 (default "127.0.0.1")
  -p string
        远程Linux主机密码 (default "root")
  -password string
        远程Linux主机密码 (default "root")
  -port int
        远程Linux主机端口 (default 22)
  -sh string
        登录后预执行指令
  -u string
        远程Linux主机用户名 (default "root")
  -user string
        远程Linux主机用户名 (default "root")

~~~