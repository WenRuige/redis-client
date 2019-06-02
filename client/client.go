package client

import (
	"net"
	"fmt"
	"strconv"
)

var host string
var port int

//func init() {
//	flag.StringVar(&host, "h", "localhost", "--localhost")
//	flag.IntVar(&port, "p", 6379, "--port")
//
//}

// 实例化一个RedisClient
type RedisClient struct {
	Host string
	Port int
	Conn *net.TCPConn
}

// 新建一个连接
func NewConn(host string, port int) RedisClient {
	return RedisClient{
		Host: host,
		Port: port,
	}

}

// 新建一个RedisClient
func (r *RedisClient) NewClient() error {
	tcpAddr := &net.TCPAddr{IP: net.ParseIP(r.Host), Port: r.Port}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	r.Conn = conn
	return nil
}

// 关闭RedisClient
func (r *RedisClient) Close() error {
	return r.Conn.Close()
}

// 实现一次请求
func (r *RedisClient) DoRequest(args ...string) (int, error) {
	return r.Conn.Write([]byte(MultiBulkMarshal(args...)))
}

// 拼凑redis协议
func MultiBulkMarshal(args ...string) string {
	str := fmt.Sprintf("*%v\r\n", strconv.Itoa(len(args)))
	// 命令所有参数
	for _, v := range args {
		str += fmt.Sprintf("$%v\r\n%v\r\n", strconv.Itoa(len(v)), v)
	}
	return str
}

func (r *RedisClient) GetResponse() ([]byte, error) {
	p := make([]byte, 1024)
	_, err := r.Conn.Read(p)
	if err != nil {
		return nil, err
	}
	switch p[0] {
	case '-':
		fmt.Println("error")
	case '+':
		fmt.Println("this is status")
	case ':':
		fmt.Println("this is 整数")
	case '$':

	}
	return p, nil
}
