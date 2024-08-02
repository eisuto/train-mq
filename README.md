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
    <img src="https://img.shields.io/badge/version-0.0.1-blue">
  </a>
  <a href="#">
      <img src="https://img.shields.io/badge/build-passing-brightgreen">
    </a>
</p>

## ✨特性
- 🪁轻量级
- 🛳高吞吐量
- 📦开箱即用
- 🖇双模式：🍕分食模式 📰订阅模式
## 🛠部署

下载适用于您目标机器操作系统的主程序，直接运行即可。

```shell script
# 解压
tar -zxvf trainMQ-amd64-linux-v0.0.1-alpha.tar.gz

# 赋予执行权限
chmod +x ./trainMQ-linux-amd64

# 启动
./TrainMQ-linux-amd64
```

## 🚀 使用方法
### 📥 发布消息
发送 POST 请求到 /publish 端点，使用以下 JSON 请求体：
```json
{
  "content": "Hello! TrainMQ!",
  "topic": "test_topic"
}
```
- 写入操作
```java
// 分食模式
TrainExecutor.send("{'msg':'内容，一般为json'}");
// 订阅模式
TrainExecutor.send("{'msg':'内容，一般为json'}","主题233");
```

- 读取操作
```java
// 分食模式
String jsonMsg = TrainExecutor.get();
// 订阅模式
String jsonMsg = TrainExecutor.get("主题233");
```
- 修改默认设置
```javas
// 修改端口 默认5757
TrainExecutor.setDefaultPort("2323");
// 修改地址 默认127.0.0.1
TrainExecutor.setDefaultIp("127.0.0.2");
```

## 🤝 贡献
欢迎贡献！请 fork 本仓库并提交 pull request。

## 📄 许可证
本项目基于 MIT 许可证开源 - 详细信息请参阅 LICENSE 文件。