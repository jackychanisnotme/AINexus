# 详细配置

## 应用设置

| 设置项 | 说明 | 默认值 |
|--------|------|--------|
| 代理端口 | 本地代理监听端口 | `3000` |
| 日志级别 | 0= 调试，1= 信息，2= 警告，3= 错误 | `1` |
| 界面语言 | 中文 / English | `zh-CN` |
| 主题 | 12 种主题可选 | `light` |
| 自动主题 | 根据时间自动切换（7:00-19:00 浅色） | 关闭 |
| 窗口关闭行为 | 直接关闭 / 最小化到托盘 / 每次询问 | 每次询问 |

## 端点配置

### 认证模式

| 认证模式 | 说明 |
|----------|------|
| `api_key` | 使用普通 API Key 调用上游服务 |
| `token_pool` | 使用通用 Token Pool，适合需要轮换凭证的 OpenAI 兼容服务 |
| `codex_token_pool` | 使用 Codex Token Pool，自动适配 ChatGPT Codex 后端、凭证轮换、刷新和失效隔离 |

### 转换器类型

| 转换器 | 说明 |
|--------|------|
| `claude` | Claude / Anthropic 兼容 API |
| `openai` | OpenAI Chat Completions 兼容 API |
| `openai2` | OpenAI Responses API，推荐给 Codex CLI |
| `gemini` | Google Gemini API |
| `deepseek` | DeepSeek OpenAI Chat 兼容 API |
| `kimi` | Kimi/Moonshot OpenAI Chat 兼容 API |

### 配置示例

**Claude 端点：**
```json
{
  "name": "Claude 官方",
  "apiUrl": "https://api.anthropic.com",
  "apiKey": "sk-ant-api03-xxx",
  "enabled": true,
  "transformer": "claude"
}
```

**OpenAI 端点：**
```json
{
  "name": "OpenAI 代理",
  "apiUrl": "https://api.openai.com/v1",
  "apiKey": "sk-xxx",
  "enabled": true,
  "transformer": "openai",
  "model": "gpt-4-turbo"
}
```

**Codex CLI / OpenAI Responses 端点：**
```json
{
  "name": "OpenAI Responses",
  "apiUrl": "https://api.openai.com/v1",
  "apiKey": "sk-xxx",
  "enabled": true,
  "transformer": "openai2",
  "model": "gpt-5-codex"
}
```

**Codex Token Pool 端点：**
```json
{
  "name": "Codex Token Pool",
  "authMode": "codex_token_pool",
  "enabled": true,
  "transformer": "openai2",
  "model": "gpt-5-codex"
}
```

**Gemini 端点：**
```json
{
  "name": "Gemini",
  "apiUrl": "https://generativelanguage.googleapis.com",
  "apiKey": "AIza-xxx",
  "enabled": true,
  "transformer": "gemini",
  "model": "gemini-pro"
}
```

**DeepSeek 端点：**
```json
{
  "name": "DeepSeek",
  "apiUrl": "https://api.deepseek.com",
  "apiKey": "sk-xxx",
  "enabled": true,
  "transformer": "deepseek",
  "model": "deepseek-v4-pro"
}
```

**Kimi 端点：**
```json
{
  "name": "Kimi",
  "apiUrl": "https://api.moonshot.ai/v1",
  "apiKey": "sk-xxx",
  "enabled": true,
  "transformer": "kimi",
  "model": "kimi-k2.6"
}
```

## WebDAV 云同步

支持通过 WebDAV 协议同步配置和统计数据，兼容坚果云、NextCloud、ownCloud 等服务。AINexus 也支持本地备份和 S3 兼容存储，适合在多设备之间迁移配置。

**配置步骤：**
1. 点击界面上的「WebDAV 云备份」
2. 填写 WebDAV 服务器地址、用户名、密码
3. 点击「测试连接」确认配置正确
4. 使用「备份」和「恢复」功能管理数据

## 数据存储位置

- 桌面端数据库：`~/.AINexus/ainexus.db`
- 服务器模式可通过 `AINEXUS_DATA_DIR` 或 `AINEXUS_DB_PATH` 指定数据目录和数据库路径
