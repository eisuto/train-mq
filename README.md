<h1 align="center">
  <a href="/" alt="logo" >
  <img src="https://images.mingming.dev/file/c8e9aa86f88ef5b4fe1b2.png" width="200" />
  </a>
  <br>
    TrainMQ
  <br>
</h1>
<h4 align="center">一个轻量级的高速消息队列</h4>

<p align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/version-0.0.5-blue">
  </a>
  <a href="#">
      <img src="https://img.shields.io/badge/build-passing-brightgreen">
    </a>
</p>

## ✨特性
- 轻量级
- 高吞吐量
- 开箱即用
## 消息模型
TrainMQ 的设计核心在于实现高效、无锁的多消费者消息队列。每个主题维护一个无锁队列，使用 Compare-And-Swap (CAS) 机制确保操作的原子性和线程安全。消费者在消费消息之前，必须先订阅相应的队列， 未订阅的队列无法进行消费操作 。 消费时通过记录每个消费者在订阅队列上的偏移量（offset），确保每次消费后偏移量自动递增，从而共享读取该队列的数据，实现无需复制队列内容的高效消费机制。

![MessageModel](https://images.mingming.dev/file/05e0ea4bc921e3c28af94.jpg)
## 📝 TODO
- [x] 发布-订阅模式
- [x] 并发安全：无锁队列实现
- [ ] 消息确认机制
- [ ] 监控：提供 Web 监控界面
- [ ] 消息优先级
- [ ] 多节点：动态节点加入移除，节点间复制与同步
- [ ] 一致性：多节点消息保持一致性
- [ ] 去中心化：无主节点的拓扑结构
- [ ] 延迟队列
- [ ] 消息持久化
- [ ] 客户端SDK - JAVA
- [ ] 客户端SDK - GO
- [ ] 客户端SDK - Python
- [ ] 客户端SDK - C
- [ ] 客户端SDK - C++
- [ ] 客户端SDK - Rust
- [ ] 客户端SDK - Node.js

## 🛠部署
下载适用于您目标机器操作系统的主程序，直接运行即可。
```shell script
# 解压
tar -zxvf train-mq-amd64-linux-v0.0.2.tar.gz

# 赋予执行权限
chmod +x ./train-mq-0.0.2

# 启动
./train-mq-0.0.2
```

## 🚀 使用方法
###  📤发布消息
发送 POST 请求到 /publish ，使用以下 JSON 请求体：
```json
{
  "content": "Hello,TrainMQ!",
  "topic": "test_topic"
}
```
###  📥消费消息
发送 GET 请求到 /consume ，并在查询参数中指定主题
```java
GET /consume?topic=test_topic
```


## 🤝 贡献者
<a href="https://github.com/eisuto/train-mq/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=eisuto/train-mq" />
</a>

## 📄 许可证
本项目基于 MIT 许可证开源 - 详细信息请参阅 LICENSE 文件。