# AgentCanvas — AI Agent 思考过程可视化

> 为 Agent 提交任务，在交互式树形图中观察每一步的推理与决策。

## 背景

AI Agent 虽然强大，但运行过程是一个黑盒。Agent 如何拆解任务、调用工具、做出决策、从错误中恢复——这些细节对使用者和开发者都不可见。

AgentCanvas 将 Agent 的完整执行过程渲染为交互式树形图，可视化展示 Agent 的每一步思考、决策和行动。适用于学习和调试场景。

## 功能

| 展示环节 | 可视化内容 |
|---------|-----------|
| **任务分解** | Agent 如何将复杂目标拆分为子步骤 |
| **推理链条** | 每个步骤的推理内容与决策理由 |
| **工具调用** | 调用的工具名称、参数与返回结果 |
| **决策分支** | 方案选择逻辑与分支回退 |
| **错误恢复** | 错误检测与策略修正过程 |
| **教学模式** | 逐步播放，支持每一步间暂停讲解 |
| **分层视图** | 初学者简版视图 / 开发者详细视图 |
| **多视图切换** | 树形图 / 时间线 / 流程图 |

## 效果示意

```
任务：「调研最佳 Go Web 框架」
├── 🔍 搜索：「Go web frameworks 2026」→ 获取 15 条结果
│   ├── 📖 抓取 gin-gonic.com → 特性摘要
│   ├── 📖 抓取 echo.labstack.com → 特性摘要
│   └── 📖 抓取 fiber.wiki → 特性摘要
├── 🤔 分析：对比框架特性与社区活跃度……
│   ├── 🏷️ Gin：Star 最多，社区最大
│   ├── 🏷️ Echo：轻量，性能优秀
│   └── 🏷️ Fiber：类 Express API，开发效率高
└── ✅ 结论：生产环境推荐 Gin，微服务推荐 Echo
```

## 技术栈

| 环境 | 后端 | 数据库 | 前端 |
|------|------|--------|------|
| 在线（Oxelia51） | Go | PostgreSQL | React + D3.js |
| 桌面（exe） | Go | SQLite | 内嵌 React + D3.js |

## 数据来源

- **MVP**：内置预设演示 Agent，开箱即用
- **未来**：对接自定义 Agent（LangChain、AutoGPT、CrewAI 等）

## 使用方式

### 在线

1. 访问 [oxelia51.com](https://oxelia51.com) 注册登录
2. 从工具菜单打开 AgentCanvas
3. 选择演示任务，观察 Agent 执行过程

### 桌面

1. 从 [GitHub Releases](https://github.com/XiaoleC05/AgentCanvas/releases) 下载 `AgentCanvas.exe`
2. 双击运行，无需联网

## 开发状态

概念阶段，尚未开发。
