package main

import (
	"fmt"
	"time"
)

// CoinJoinClient provides access to coin hours bank
type CoinJoinClient interface {
	PostSendTransaction(coin float64, input string) string
}

type httpClient struct {
	address string
}

func newHTTPClient(address string) CoinJoinClient {
	return &httpClient{address}
}

func (c *httpClient) PostSendTransaction(coin float64, input string) string {
	time.Sleep(time.Second * 2)
	return "transaction sent successfully"
}

func main() {
	c := newHTTPClient("/asdfsadf.com")

	fmt.Println("c", c)
	fmt.Println("c.PostSendTransaction", c.PostSendTransaction(23.23, "asdfaksj34534ksdf"))
}
