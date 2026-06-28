# 常见问题

## 安装和启动

**Q: Windows 提示「Windows 已保护你的电脑」？**

点击「更多信息」→「仍要运行」。应用没有数字签名，不影响使用。

**Q: macOS 提示「无法打开，因为无法验证开发者」？**

右键点击应用 → 选择「打开」→ 点击「打开」。或在「系统偏好设置」→「安全性与隐私」中允许。

**Q: 端口被占用？**

点击界面顶部的端口号，修改为其他端口（如 3001）。

## 端点配置

**Q: 如何选择转换器？**

- Claude 官方或兼容服务 → `claude`
- OpenAI Chat Completions 兼容服务 → `openai`
- Codex CLI / OpenAI Responses API → `openai2`
- DeepSeek → `deepseek`
- Kimi/Moonshot → `kimi`
- Google Gemini → `gemini`

**Q: 为什么 OpenAI/Gemini/DeepSeek/Kimi/Codex 相关端点必须填模型？**

Claude Code 请求中包含 Claude 模型名，代理需要知道转换为哪个目标模型。

**Q: Codex CLI 推荐怎么接入？**

推荐在 Codex CLI 中使用 Responses API，把 provider/base URL 指向 `http://localhost:3000/v1`，并在 AINexus 里使用 `openai2` 转换器。需要轮换 ChatGPT Codex 凭证时，使用 Codex Token Pool。

**Q: 端点测试成功但使用失败？**

检查：API 密钥权限、模型名称、API 配额。查看日志获取详细错误。

## 使用问题

**Q: Token 统计准确吗？**

估算值，基于文本长度计算，与实际计费可能有差异。

**Q: 如何备份配置？**

1. 使用 WebDAV 云同步
2. 使用本地备份或 S3 兼容存储
3. 手动复制 `~/.AINexus/ainexus.db`

**Q: 端点轮换顺序？**

按列表顺序，可拖拽调整。

**Q: 数据安全吗？**

桌面端数据默认存储在本地 `~/.AINexus/`，API 密钥不会发送给第三方。远程备份或同步只会在你主动配置 WebDAV、本地备份目录或 S3 兼容存储后执行。
