### gRpc_example
    gRPC是一个高性能、通用的开源RPC框架，其由Google主要面向移动应用开发并基于HTTP/2协议标准而设计，基于ProtoBuf(Protocol
    Buffers)序列化协议开发，且支持众多开发语言。gRPC提供了一种简单的方法来精确地定义服务和为iOS、Android和后台支持服务自动生成可
    靠性很强的客户端功能库。客户端充分利用高级流和链接功能，从而有助于节省带宽、降低的TCP链接次数、节省CPU使用、和电池寿命。

### 项目功能
    项目实现简单的grpc调用、单向ssl认证、双向ssl认证
    master: 简单grpc调用
    tls: 单向ssl认证
    mutli_tls: 双向ssl认证

### Getting Started
因为grpc是基于ProtoBuf序列化协议开发的，所以需要编写proto文件，然后通过protoc命令生成对应的pb.go文件
#### golang环境配置
配置GOROOT、GOPATH
