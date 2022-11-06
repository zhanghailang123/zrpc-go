package main

//
//import (
//	"fmt"
//	"log"
//	"net"
//	"sync"
//	"time"
//	"zrpc-go"
//)
//
//func startServer(addr chan string) {
//	l, err := net.Listen("tcp", ":0")
//	if err != nil {
//		log.Fatal("net work error:", err)
//	}
//	log.Println("start rpc server om :", l.Addr())
//	addr <- l.Addr().String()
//	zrpc_go.Accept(l)
//}
//
////5个RPC并发调用
//func main() {
//	log.SetFlags(0)
//	addr := make(chan string)
//	go startServer(addr)
//	client, _ := zrpc_go.Dial("tcp", <-addr)
//	defer func() { _ = client.Close() }()
//
//	time.Sleep(time.Second)
//
//	var wg sync.WaitGroup
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			args := fmt.Sprintf("zrpc seq %d", i)
//			var reply string
//			if err := client.Call("zhanghalang", args, &reply); err != nil {
//				log.Fatal("call zzz error:", err)
//			}
//			log.Println("reply:", reply)
//		}(i)
//	}
//	wg.Wait()
//}
