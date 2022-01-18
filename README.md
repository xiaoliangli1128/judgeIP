# judgeIP
判断字符串中的ip类型


#举个例子  ip.txt 有以下内容
```txt
sw.xxx.com,120.92.9.150,securitytrails
xx.xxx.com,10.100.59.140,securitytrails
m.xx.xxx.com,10.69.34.166,securitytrails

```
> 执行该程序后，会在每行结尾做出一个标志

```bash
cat ip.txt |main.exe
```

```txt
sw.xxx.com,120.92.9.150,securitytrails,public    >公网ip
xx.xxx.com,10.100.59.140,securitytrails,private  > 内网ip
m.xx.xxx.com,10.69.34.166,securitytrails,private > 内网ip

```
