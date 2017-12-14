### gRpc_example
    gRPC是一个高性能、通用的开源RPC框架，其由Google主要面向移动应用开发并基于HTTP/2协议标准而设计，基于ProtoBuf(Protocol
    Buffers)序列化协议开发，且支持众多开发语言。gRPC提供了一种简单的方法来精确地定义服务和为iOS、Android和后台支持服务自动生成可
    靠性很强的客户端功能库。客户端充分利用高级流和链接功能，从而有助于节省带宽、降低的TCP链接次数、节省CPU使用、和电池寿命。

### 项目功能
    项目有不同的branch，实现了简单的grpc调用、单向ssl认证、双向ssl认证
    master:    简单grpc调用
    tls:       单向ssl认证
    mutli_tls: 双向ssl认证

### Getting Started
因为grpc是基于ProtoBuf序列化协议开发的，所以需要编写proto文件，然后通过protoc命令生成对应的pb.go文件


##### 配置GOROOT、GOPATH
[golang下载地址][https://golang.org/dl/]
    *NOTE:*

    *GOROOT是golang安装的目录*

    *GOPATH是project目录，执行的的GOPATH目录下一般有{bin,src,pkg}三个目录，你的代码放置在src目录下*

    *多个项目时，GOPATH可以指定多个，但是每个项目用\"go get\"下载的包只在第一个指定的GOPATH路径下*


#### 配置protoc环境
    wget https://github.com/google/protobuf/releases/download/v2.6.1/protobuf-2.6.1.tar.gz
    ./configure
    make
    make install

    go get -u github.com/golang/protobuf/protoc-gen-go
    此命令会在$GOPATH/bin/下生成protoc-gen-go二进制文件，把$GOPATH/bin次路径加到$PATH中

    在项目目录执行：
    protoc -I cf/ cf/*.proto --go_out=plugins=grpc:cf
    生成cf.pb.go文件

通过ProtoBuf在cf.pb.go中提供的接口，即可实现grpc通信。