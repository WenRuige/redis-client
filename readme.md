## golang 实现redis-client


> 背景

最近在熟悉redis相关命令,此次学习一下`redis client`与`redis server`如何通信


> `Redis-Client`与`Reids-Server`如何通信

redis-server与redis-client通过tcp连接进行数据交互,服务端默认端口号为6379,并且客户端和服务器发送的命令一律以`\r\n`(CRLF)结束






##### 请求协议

redis请求协议,在这个协议中,所有发至服务器的参数都是二进制安全的.

```
* <参数数量>
$<参数1的字节数量> CRLF
<参数1的数据> CRLF
```

比如:实现`redis get keys`操作如下:

```
*2
$3
get
$4
keys
```

##### 回复协议

* 状态回复`status reply`的第一个字节是`+`
* 错误回复`error  reply`的第一个字节是`-`
# redis-client
