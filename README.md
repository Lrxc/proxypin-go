proxypin-go

# 功能列表:

- 网络拦截重定向

# 打包

```shell
go install fyne.io/tools/cmd/fyne@latest # 安装 fyne cmd
fyne package --release --id lrxc.proxy -os windows -icon assets/logo.jpg # windows加入图标打包
```

包太大,剔除多余并压缩

```bash
#最小打包
#-ldflags=“参数”： 表示将引号里面的参数传给编译器
#-s：去掉符号信息（这样panic时，stack trace就没有任何文件名/行号信息了，这等价于普通C/C+=程序被strip的效果）
#-w：去掉DWARF调试信息 （得到的程序就不能用gdb调试了）
#-H windowsgui : 以windows gui形式打包，不带dos窗口。其中注意H是大写的
go build -ldflags="-s -w -H windowsgui" -o masking-upgrade.exe main.go 

#使用upx再次压缩(https://github.com/upx/upx/releases/tag/v4.1.0)
upx -9 masking-upgrade.exe
```

# 使用:

1. 安装 cert/server.crt 到 受信任的根证书颁发机构  
   <img src="docs/import.png" alt="import.png" style="zoom: 50%;" />
2. 配置[conf.yml](conf.yml)
3. 启动程序,查看系统代理是否生效  
   <img src="docs/sys_proxy.png" alt="sys_proxy.png" style="zoom: 50%;" />

# 自己生成证书[可选]

```shell
#生成私钥
openssl genrsa -out server.key 2048
#根据私钥生成证书申请文件csr
openssl req -new -key server.key -out server.csr
#使用私钥对证书申请进行签名从而生成证书(10年)
openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650
```