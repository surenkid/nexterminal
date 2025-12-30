# nexterminal

基于 [Next Terminal](https://github.com/dushixiang/next-terminal) 的二次开发项目。

## 项目说明

这是一个基于 AGPL-3.0 许可证的开源网页终端项目，支持 RDP、SSH、VNC、Telnet、Kubernetes 等协议。

## 许可证

本项目基于 AGPL-3.0 许可证开源。详情请参阅 [LICENSE](./LICENSE) 文件。

## 开发说明

### 环境要求

- Go 1.18 或以上版本
- Node.js 16 或以上版本
- npm 或 yarn

### 编译

1. 进入 `web` 目录，执行 `yarn` 或 `npm install` 安装前端依赖
2. 返回项目根目录，执行 `sh build.sh` 进行编译

## HTTPS 配置

生成自签证书并配置：

```bash
# 生成自签证书
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 3650 \
  -keyout /home/surenkid/localhost.key \
  -out /home/surenkid/localhost.crt \
  -subj "/CN=localhost"
```

在 `config.yml` 中配置：

```yaml
server:
  addr: 0.0.0.0:8443
  cert: /home/surenkid/localhost.crt
  key: /home/surenkid/localhost.key
```

## HTTP Basic Auth 配置

启用 HTTP Basic Authentication 可以保护所有资源，防止未授权访问：

```yaml
server:
  basic-auth:
    enable: true
    user: your_username
    pass: your_password
```

启用后，所有请求都需要先通过 Basic Auth 验证，未提供正确凭证将返回 401。

## 贡献

欢迎提交 Issue 和 Pull Request。
