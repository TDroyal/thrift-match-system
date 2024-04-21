package main

import (
    "context"
    "fmt"

    "github.com/apache/thrift/lib/go/thrift"
    //"github.com/apache/thrift/tutorial/go/gen-go/tutorial"
    "game/src/match_client/match"
    
)

var defaultCtx = context.Background()

func handleClient(client *match.MatchClient) (err error) {
    //client.Ping(defaultCtx)
    // fmt.Println("ping()")

    user := match.User{1, "royal_111", 1500}

    res, _ := client.Add(defaultCtx, &user)
    fmt.Print("add user: ", user.Username, ".   the result is: ",  res, "\n")

    return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration) error {
    var transport thrift.TTransport
    if secure {
        transport = thrift.NewTSSLSocketConf(addr, cfg)
    } else {
        transport = thrift.NewTSocketConf(addr, cfg)
    }
    transport, err := transportFactory.GetTransport(transport)
    if err != nil {
        return err
    }
    defer transport.Close()
    if err := transport.Open(); err != nil {
        return err
    }
    iprot := protocolFactory.GetProtocol(transport)
    oprot := protocolFactory.GetProtocol(transport)
    return handleClient(match.NewMatchClient(thrift.NewTStandardClient(iprot, oprot)))
}

func main() {
    transportFactory := thrift.NewTTransportFactory()
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    addr := "localhost:9090" // 设置服务器地址
    secure := false          // 设置是否使用安全连接
    cfg := &thrift.TConfiguration{} // 可选的配置
    err := runClient(transportFactory, protocolFactory, addr, secure, cfg)
    if err != nil {
        fmt.Println("Error running client:", err)
    }
}
