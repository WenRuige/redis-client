package main

import (
	"github.com/redis-cli/client"
	"fmt"
)

func main() {

	conn := client.NewConn("127.0.0.1", 6379)
	conn.NewClient()
	_, err := conn.DoRequest("get", "keys")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	res, _ := conn.GetResponse()
	fmt.Println(string(res))

}
