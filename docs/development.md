# 开发指南

## 环境准备

- Go 1.24+
- Node.js 18+
- Wails CLI v2

```bash
# 安装 Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 检查环境依赖
wails doctor
```

## 开发模式

```bash
# 安装前端依赖
cd cmd/desktop/frontend && npm install

# 启动开发模式（支持热重载）
cd ..
wails dev
```

## 构建发布

```bash
# 桌面端前端修改后先生成 dist
cd cmd/desktop/frontend && npm run build

# 构建桌面应用
cd ..
wails build -platform darwin/universal
wails build -platform windows/amd64

# 构建服务器模式
cd ../../cmd/server && go build -ldflags="-s -w" -o ainexus-server .
```

桌面端构建产物位于 `cmd/desktop/build/bin/` 目录。

## 项目结构

```
AINexus/
├── cmd/
│   ├── desktop/            # Wails 桌面应用
│   │   ├── app.go          # 桌面端核心应用逻辑
│   │   ├── main.go         # 桌面端入口
│   │   └── frontend/       # Vue/原生模块化前端
│   ├── server/             # 无头服务器模式
│   └── license-server/     # 在线授权服务
├── internal/
│   ├── proxy/              # HTTP 代理核心
│   ├── transformer/        # API 格式转换器
│   ├── storage/            # SQLite 数据存储
│   ├── service/            # 端点、统计、备份、更新等业务服务
│   ├── config/             # 配置管理
│   ├── onlinelicense/      # 在线授权、卡密和设备激活
│   ├── webdav/             # WebDAV 同步
│   ├── logger/             # 日志系统
│   └── tray/               # 系统托盘
└── docs/                   # 配置、开发、FAQ 和模型接口文档
```
