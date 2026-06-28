# FAQ

## Installation and Startup

**Q: Windows shows "Windows protected your PC"?**

Click "More info" → "Run anyway". The app is not digitally signed, but it works fine.

**Q: macOS shows "Cannot be opened because the developer cannot be verified"?**

Right-click the app → Select "Open" → Click "Open". Or allow it in "System Preferences" → "Security & Privacy".

**Q: Port is in use?**

Click the port number at the top of the interface and change it to another port (e.g., 3001).

## Endpoint Configuration

**Q: How to choose a transformer?**

- Claude official or compatible services → `claude`
- OpenAI Chat Completions-compatible services → `openai`
- Codex CLI / OpenAI Responses API → `openai2`
- DeepSeek → `deepseek`
- Kimi/Moonshot → `kimi`
- Google Gemini → `gemini`

**Q: Why is the model field required for OpenAI/Gemini/DeepSeek/Kimi/Codex-related endpoints?**

Claude Code requests contain Claude model names. The proxy needs to know which target model to convert to.

**Q: What is the recommended Codex CLI setup?**

Use Responses API in Codex CLI, point the provider/base URL to `http://localhost:3000/v1`, and use the `openai2` transformer in AINexus. Use Codex Token Pool when you need to rotate ChatGPT Codex credentials.

**Q: Endpoint test succeeds but usage fails?**

Check: API key permissions, model name, API quota. View logs for detailed errors.

## Usage Issues

**Q: Is token statistics accurate?**

It's an estimate based on text length, may differ from actual billing.

**Q: How to backup configuration?**

1. Use WebDAV cloud sync
2. Use local backup or S3-compatible storage
3. Manually copy `~/.AINexus/ainexus.db`

**Q: Endpoint rotation order?**

In list order, can be adjusted by drag and drop.

**Q: Is data secure?**

Desktop data is stored locally in `~/.AINexus/` by default. API keys are never sent to third parties unless you explicitly configure WebDAV, a local backup directory, or S3-compatible storage for backup/sync.
