package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"testing"
)

func TestRpcRegistry(t *testing.T) {
	discovery := NewRpcRegistry()
	if err := discovery.Start(":8080"); err != nil {
		t.Fatal(err)
	}
}

func TestRcpServer(t *testing.T) {
	server := RpcServer{
		server: rpc.NewServer(),
	}
	if err := server.Register(&MathService{}); err != nil {
		log.Fatal(err)
	}
	if err := server.RegisterName("math", &MathService2{}); err != nil {
		log.Fatal(err)
	}
	if err := server.RegisterName("math.0", &MathService3{}); err != nil {
		log.Fatal(err)
	}

	if err := server.RegisterToDiscovery("http://localhost:8080"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("rpc server start")
	err := server.ServeTCP(":1234")
	if err != nil {
		panic(err)
	}
}

func TestRpcClient(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/getProvidersByType?node_type=rpc")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	resp.Body.Close()

	fmt.Println(string(buf))

	providers := []*Provider{}
	if err := json.Unmarshal(buf, &providers); err != nil {
		t.Fatal(err)
	}

	if len(providers) == 0 {
		t.Fatal("no providers")
	}

	c := RpcClient{}
	if err := c.DialTCPWithCodec(providers[0].Addr); err != nil {
		log.Fatal(err)
	}

	var result int
	if err := c.Call("MathService.Add", &Args{X: 1, Y: 2}, &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result)

	var res Result
	call := c.Go("MathService.Divide", &Args{X: 1, Y: 2}, &res)
	done := <-call.Done
	if done.Error != nil {
		log.Fatal(done.Error)
	}

	fmt.Println("result is", res.Quotient, res.Remainder)

	var rand int
	if err := c.Call("math.Rand", &Args{X: 1, Y: 100}, &rand); err != nil {
		log.Fatal(err)
	}
	fmt.Println("random number is", rand)

	if err := c.Call("math.NoReply", &Args{X: 1, Y: 2}, &struct{}{}); err != nil {
		log.Fatal(err)
	}

	if err := c.Call("math.0.Rand", &Args{X: 1, Y: 100}, &rand); err != nil {
		log.Fatal(err)
	}
	fmt.Println("random number is", rand)
}

func TestXxx(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	fmt.Println(ip)
}
