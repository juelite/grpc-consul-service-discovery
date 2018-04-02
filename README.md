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


### 将GRPC服务注册到consul （服务发现,GOLANG）

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

consul目录结构如下
    
    ├── register.go
    ├── resolver.go
    └── watcher.go

文件源码：
- register.go

    package consul
    import (
        "fmt"
        "log"
        "os"
        "os/signal"
        "strconv"
        "syscall"
        "time"
        consul "github.com/hashicorp/consul/api"
    )
    // Register is the helper function to self-register service into Etcd/Consul server
    // name - service name
    // host - service host
    // port - service port
    // target - consul dial address, for example: "127.0.0.1:8500"
    // interval - interval of self-register to etcd
    // ttl - ttl of the register information
    func Register(name string, host string, port int, target string, interval time.Duration, ttl int) error {
        conf := &consul.Config{Scheme: "http", Address: target}
        client, err := consul.NewClient(conf)
        if err != nil {
            return fmt.Errorf("wonaming: create consul client error: %v", err)
        }
        serviceID := fmt.Sprintf("%s-%s-%d", name, host, port)
        //de-register if meet signhup
        go func() {
            ch := make(chan os.Signal, 1)
            signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
            x := <-ch
            log.Println("wonaming: receive signal: ", x)
            err := client.Agent().ServiceDeregister(serviceID)
            if err != nil {
                log.Println("wonaming: deregister service error: ", err.Error())
            } else {
                log.Println("wonaming: deregistered service from consul server.")
            }
            err = client.Agent().CheckDeregister(serviceID)
            if err != nil {
                log.Println("wonaming: deregister check error: ", err.Error())
            }
            s, _ := strconv.Atoi(fmt.Sprintf("%d", x))
            os.Exit(s)
        }()
        // routine to update ttl
        go func() {
            ticker := time.NewTicker(interval)
            for {
                <-ticker.C
                err = client.Agent().UpdateTTL(serviceID, "", "passing")
                if err != nil {
                    log.Println("wonaming: update ttl of service error: ", err.Error())
                }
            }
        }()
        // initial register service
        regis := &consul.AgentServiceRegistration{
            ID:      serviceID,
            Name:    name,
            Address: host,
            Port:    port,
        }
        err = client.Agent().ServiceRegister(regis)
        if err != nil {
            return fmt.Errorf("wonaming: initial register service '%s' host to consul error: %s", name, err.Error())
        }
        // initial register service check
        check := consul.AgentServiceCheck{TTL: fmt.Sprintf("%ds", ttl), Status: "passing"}
        err = client.Agent().CheckRegister(&consul.AgentCheckRegistration{ID: serviceID, Name: name, ServiceID: serviceID, AgentServiceCheck: check})
        if err != nil {
            return fmt.Errorf("wonaming: initial register service check to consul error: %s", err.Error())
        }
        return nil
    }

- resolver.go

    package consul
    import (
        "errors"
        "fmt"
        consul "github.com/hashicorp/consul/api"
        "google.golang.org/grpc/naming"
    )

    // ConsulResolver is the implementaion of grpc.naming.Resolver
    type ConsulResolver struct {
        ServiceName string //service name
    }

    // NewResolver return ConsulResolver with service name
    func NewResolver(serviceName string) *ConsulResolver {
        return &ConsulResolver{ServiceName: serviceName}
    }

    // Resolve to resolve the service from consul, target is the dial address of consul
    func (cr *ConsulResolver) Resolve(target string) (naming.Watcher, error) {
        if cr.ServiceName == "" {
            return nil, errors.New("wonaming: no service name provided")
        }

        // generate consul client, return if error
        conf := &consul.Config{
            Scheme:  "http",
            Address: target,
        }
        client, err := consul.NewClient(conf)
        if err != nil {
            return nil, fmt.Errorf("wonaming: creat consul error: %v", err)
        }

        // return ConsulWatcher
        watcher := &ConsulWatcher{
            cr: cr,
            cc: client,
        }
        return watcher, nil
    }

- watcher.go

    package consul

    import (
        "fmt"
        "time"

        consul "github.com/hashicorp/consul/api"
        "google.golang.org/grpc/naming"
    )

    // ConsulWatcher is the implementation of grpc.naming.Watcher
    type ConsulWatcher struct {
        // cr: ConsulResolver
        cr *ConsulResolver
        // cc: Consul Client
        cc *consul.Client

        // LastIndex to watch consul
        li uint64

        // addrs is the service address cache
        // before check: every value shoud be 1
        // after check: 1 - deleted  2 - nothing  3 - new added
        addrs []string
    }

    // Close do nonthing
    func (cw *ConsulWatcher) Close() {
    }

    // Next to return the updates
    func (cw *ConsulWatcher) Next() ([]*naming.Update, error) {
        // Nil cw.addrs means it is initial called
        // If get addrs, return to balancer
        // If no addrs, need to watch consul

        if len(cw.addrs) == 0 {
            // must return addrs to balancer, use ticker to query consul till data gotten
            addrs, li, _ := cw.queryConsul(nil)

            // got addrs, return
            if len(addrs) != 0 {
                cw.addrs = addrs
                cw.li = li
                return GenUpdates([]string{}, addrs), nil
            }
        }

        for {
            // watch consul
            addrs, li, err := cw.queryConsul(&consul.QueryOptions{WaitIndex: cw.li})
            if err != nil {
                time.Sleep(1 * time.Second)
                continue
            }

            // generate updates
            updates := GenUpdates(cw.addrs, addrs)

            // update addrs & last index
            cw.addrs = addrs
            cw.li = li

            if len(updates) != 0 {
                return updates, nil
            }
        }

        // should never come here
        return []*naming.Update{}, nil
    }

    // queryConsul is helper function to query consul
    func (cw *ConsulWatcher) queryConsul(q *consul.QueryOptions) ([]string, uint64, error) {
        cs, meta, err := cw.cc.Health().Service(cw.cr.ServiceName, "", true, q)
        if err != nil {
            return nil, 0, err
        }

        addrs := make([]string, 0)
        for _, s := range cs {
            // addr should like: 127.0.0.1:8001
            addrs = append(addrs, fmt.Sprintf("%s:%d", s.Service.Address, s.Service.Port))
        }

        return addrs, meta.LastIndex, nil
    }

    func GenUpdates(a, b []string) []*naming.Update {
        updates := []*naming.Update{}

        deleted := diff(a, b)
        for _, addr := range deleted {
            update := &naming.Update{Op: naming.Delete, Addr: addr}
            updates = append(updates, update)
        }

        added := diff(b, a)
        for _, addr := range added {
            update := &naming.Update{Op: naming.Add, Addr: addr}
            updates = append(updates, update)
        }
        return updates
    }

    // diff(a, b) = a - a(n)b
    func diff(a, b []string) []string {
        d := make([]string, 0)
        for _, va := range a {
            found := false
            for _, vb := range b {
                if va == vb {
                    found = true
                    break
                }
            }

            if !found {
                d = append(d, va)
            }
        }
        return d
    }


文件源码  