<h1 align="center">
  <a href="/" alt="logo" >
  <img src="https://github.com/eisuto/TrainMQ/blob/main/static/logo2.png?raw=true" width="200" />
  </a>
  <br>
    TrainMQ
  <br>
</h1>
<h4 align="center">一个轻量级的高速消息队列</h4>

<p align="center">
  <a href="#">
    <img src="https://img.shields.io/badge/version-0.0.2-blue">
  </a>
  <a href="#">
      <img src="https://img.shields.io/badge/build-passing-brightgreen">
    </a>
</p>

## ✨特性
- 轻量级
- 高吞吐量
- 开箱即用

## 📝 TODO
- [ ] 消息确认机制
- [ ] 消息优先级
- [ ] 多节点：动态节点加入移除，节点间复制与同步
- [ ] 一致性：多节点消息保持一致性
- [ ] 去中心化：无主节点的拓扑结构
- [ ] 监控：提供 Web 监控界面
- [ ] 延迟队列
- [ ] 消息持久化

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


## 🤝 贡献
欢迎贡献！请 fork 本仓库并提交 pull request。

## 📄 许可证
本项目基于 MIT 许可证开源 - 详细信息请参阅 LICENSE 文件。