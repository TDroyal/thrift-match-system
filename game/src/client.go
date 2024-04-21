package main

import (
    "context"
    "fmt"

    "github.com/apache/thrift/lib/go/thrift"
    //"github.com/apache/thrift/tutorial/go/gen-go/tutorial"
    "game/src/match_client/match"
    "os"
    "bufio"
    "strings"
    "strconv"
)

var defaultCtx = context.Background()

func handleClient(client *match.MatchClient, user *match.User, task_type string) (err error) {
    // client.Ping(defaultCtx)
    // fmt.Println("ping()")
    // user := match.User{1, "royal_111", 1500}
    var res int32

    if task_type == "add" {
        res, err = client.Add(defaultCtx, user)
        fmt.Println("userid: ", user.ID, " username: ", user.Username, " enter match pool",  res)
    }else if task_type == "remove" {
        res, err = client.Remove(defaultCtx, user)
        fmt.Println("userid: ", user.ID, " username: ", user.Username, " exit match pool",  res)
    }
    // res, _ := client.Add(defaultCtx, &user)
    // fmt.Print("add user: ", user.Username, ".   the result is: ",  res, "\n")
    return
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration, user *match.User, op string) error {
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
    return handleClient(match.NewMatchClient(thrift.NewTStandardClient(iprot, oprot)), user, op)
}

func main() {
    transportFactory := thrift.NewTTransportFactory()
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    addr := "localhost:9090" // 设置服务器地址
    secure := false          // 设置是否使用安全连接
    cfg := &thrift.TConfiguration{} // 可选的配置
    
    // 从控制台读入
    reader := bufio.NewReader(os.Stdin)
    for {
        line, _ := reader.ReadString('\n')
        line = strings.TrimRight(line, "\n")
        parts := strings.Split(line, " ")
        op, username:= parts[0], parts[2]
        id, _ := strconv.ParseInt(parts[1], 10, 32)
        score, _ := strconv.ParseInt(parts[3], 10, 32)
        // fmt.Printf("%T\n", id) int64
        user := match.User{int32(id), username, int32(score)}

        if err := runClient(transportFactory, protocolFactory, addr, secure, cfg, &user, op); err != nil {
            fmt.Println("Error running client:", err)
        }
    }
}
