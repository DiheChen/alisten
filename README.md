# Alisten

一个音乐播放和房间管理系统，支持从 Bilibili、网易云音乐、QQ音乐等平台获取音乐。

这个仓库是一个 fork 的版本, 部署原版请参考: <https://github.com/bihua-university/alisten>

## 部署

1. 复制配置文件模板

```bash
cp .env.example .env
```

2. 编辑你的配置文件 `.env`, 填写你的配置, 其中配置项 `ALISTEN_PERSIST` 是一个 JSON 字符串, 它看上去就像:

```env
ALISTEN_PERSIST = '[{"id": "你的房间id", "name": "你的房间名", "desc": "", "password": "你的房间密码"}]'
```

> 房间 id 是唯一标识符
>
> 只有写入配置文件的房间才没有限制

3. 假设 `example.com` 是你的域名, 将 `alisten.example.com` 和 `oss.example.com` 解析到你的服务器

4. 执行以下命令完成部署:

```bash
docker compose up -d
```

5. 访问 `https://alisten.example.com/house/search`, 此时应该返回你的房间信息。

前端请参考 <https://github.com/bihua-university/alisten-ui> 自行部署。