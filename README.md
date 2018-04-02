### GRPC简介
gRPC  是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计。目前提供 C、Java 和 Go 语言版本，分别是：grpc, grpc-java, grpc-go. 其中 C 版本支持 C, C++, Node.js, Python, Ruby, Objective-C, PHP 和 C# 支持.

gRPC 基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。

### CONSUL简介
Consul有多个组件，但总体而言，它是发现和配置基础架构中的服务的工具。它提供了几个关键功能：

- 服务发现：Consul可以提供服务，比如 api或mysql，其他客户可以使用领事来发现给定服务的提供者。使用DNS或HTTP，应用程序可以轻松找到他们所依赖的服务。

- 健康检查：Consul客户端可以提供任何数量的健康检查，或者与给定服务相关联（“是Web服务器返回200 OK”），还是与本地节点（“内存使用率低于90％”）。操作员可以使用此信息来监视群集运行状况，服务发现组件使用此信息将流量从不健康的主机中引导出去。

- KV Store：应用程序可以将Consul的分层键/值存储用于许多目的，包括动态配置，功能标记，协调，master选举等等。简单的HTTP API使其易于使用。

- 多数据中心：Consul支持多个数据中心。这意味着Consul的用户不必担心构建额外的抽象层以扩展到多个区域。

Consul旨在与DevOps社区和应用程序开发人员保持友好，使其成为现代化，弹性基础架构的完美选择。

### GRPC 服务搭建(golang)
详见：http://www.yuuuu.wang/post/grpc-wei-fu-wu-shi-li-xie-redis-huan-cun-fu-yuan-ma.html

### consul 
下载consul 地址：https://www.consul.io/downloads.html
解压将consul放到 /usr/local/bin
起consul服务

    consul agent -server -bootstrap -advertise=127.0.0.1 -client=127.0.0.1 -data-dir=/tmp/node -ui -node=s1

-server 指定为server模式
-bootstrap 指定为引导模式
-advertise 设置监听地址
-client 指定ip可以访问ui
-data-dir 数据存放路径
-ui 是否开启web
-node 指定当前节点名称

更多参数说明

    -advertise=<value>
     Sets the advertise address to use.
    -advertise-wan=<value>
     Sets address to advertise on WAN instead of -advertise address.
    -bind=<value>
     Sets the bind address for cluster communication.
    -bootstrap
     Sets server to bootstrap mode.
    -bootstrap-expect=<value>
     Sets server to expect bootstrap mode.
    -client=<value>
     Sets the address to bind for client access. This includes RPC, DNS,
     HTTP and HTTPS (if configured).
    -config-dir=<value>
     Path to a directory to read configuration files from. This
     will read every file ending in '.json' as configuration in this
     directory in alphabetical order. Can be specified multiple times.
    -config-file=<value>
     Path to a JSON file to read configuration from. Can be specified
     multiple times.
    -config-format=<value>
     Config files are in this format irrespective of their extension.
     Must be 'hcl' or 'json'
    -data-dir=<value>
     Path to a data directory to store agent state.
    -dev
     Starts the agent in development mode.
    -disable-host-node-id
     Setting this to true will prevent Consul from using information
     from the host to generate a node ID, and will cause Consul to
     generate a random node ID instead.
    -disable-keyring-file
     Disables the backing up of the keyring to a file.
    -dns-port=<value>
     DNS port to use.
    -domain=<value>
     Domain to use for DNS interface.
    -enable-script-checks
     Enables health check scripts.
    -encrypt=<value>
     Provides the gossip encryption key.
    -hcl=<value>
     hcl config fragment. Can be specified multiple times.
    -http-port=<value>
     Sets the HTTP API port to listen on.
    -join=<value>
     Address of an agent to join at start time. Can be specified
     multiple times.
    -join-wan=<value>
     Address of an agent to join -wan at start time. Can be specified
     multiple times.
    -log-level=<value>
     Log level of the agent.
    -node=<value>
     Name of this node. Must be unique in the cluster.
    -node-id=<value>
     A unique ID for this node across space and time. Defaults to a
     randomly-generated ID that persists in the data-dir.
    -node-meta=<key:value>
     An arbitrary metadata key/value pair for this node, of the format
     `key:value`. Can be specified multiple times.
    -non-voting-server
     (Enterprise-only) This flag is used to make the server not
     participate in the Raft quorum, and have it only receive the data
     replication stream. This can be used to add read scalability to
     a cluster in cases where a high volume of reads to servers are
     needed.
    -pid-file=<value>
     Path to file to store agent PID.
    -protocol=<value>
     Sets the protocol version. Defaults to latest.
    -raft-protocol=<value>
     Sets the Raft protocol version. Defaults to latest.
    -recursor=<value>
     Address of an upstream DNS server. Can be specified multiple times.
    -rejoin
     Ignores a previous leave and attempts to rejoin the cluster.
    -retry-interval=<value>
     Time to wait between join attempts.
    -retry-interval-wan=<value>
     Time to wait between join -wan attempts.
    -retry-join=<value>
     Address of an agent to join at start time with retries enabled. Can
     be specified multiple times.
    -retry-join-wan=<value>
     Address of an agent to join -wan at start time with retries
     enabled. Can be specified multiple times.
    -retry-max=<value>
     Maximum number of join attempts. Defaults to 0, which will retry
     indefinitely.
    -retry-max-wan=<value>
     Maximum number of join -wan attempts. Defaults to 0, which will
     retry indefinitely.
    -segment=<value>
     (Enterprise-only) Sets the network segment to join.
    -serf-lan-bind=<value>
     Address to bind Serf LAN listeners to.
    -serf-wan-bind=<value>
     Address to bind Serf WAN listeners to.
    -server
     Switches agent to server mode.
    -syslog
     Enables logging to syslog.
    -ui
     Enables the built-in static web UI server.
    -ui-dir=<value>
     Path to directory containing the web UI resources.


### 将GRPC服务注册到consul （服务发现,GOLANG） 关键代码

    var (
        serv = flag.String("service", "Cache", "elcLog RedisCache")
        port = flag.Int("port", 50051, "listening port") //本地服务的端口
        reg  = flag.String("reg", "127.0.0.1:8500", "register address") //consul的地址
        ttl = 15
    )
    func main() {
        flag.Parse()
        base := services.BaseService{}
        //起RPC服务
        gRPCport := base.GetVal("rpcserve")
        base.LogInfo("index" , gRPCport)
        lis , err := net.Listen("tcp", gRPCport)
        if err != nil {
            log.Fatal("failed to listen: %v", err)
        }
        //注册服务到consul
        err = consul.Register(*serv, "127.0.0.1", *port, *reg, time.Second*10, ttl) //将该grpc服务注册到consul
        if err != nil {
            panic(err)
        }
        s := grpc.NewServer()
        pb.RegisterGreeterServer(s, &server{})
        s.Serve(lis)
    }

