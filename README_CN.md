# AgentCanvas — AI Agent 思考过程可视化

> 提交任务给 AI Agent，在交互式树形图中观察它每一步的推理过程。

## 为什么需要 AgentCanvas？

AI Agent 很强大，但也很不透明。Agent 如何拆解任务、调用工具、做出决策、从错误中恢复——这一切都在黑盒中进行。无论你是在学习 Agent 的工作原理，还是在开发自己的 Agent，都看不清内部发生了什么。

**AgentCanvas** 打开了这个黑盒。它将 Agent 的完整执行过程渲染为交互式可视化图表，让你看清 Agent 如何思考、决策和行动。

## 功能

| 功能 | 展示内容 |
|------|---------|
| **任务分解** | Agent 如何将复杂目标拆分为子步骤 |
| **思考链** | 每个步骤的推理内容——它在想什么、为什么 |
| **工具调用** | 调用了哪些工具、参数、返回结果 |
| **决策分支** | 为什么选方案 A 不选 B，含分支回退 |
| **错误恢复** | Agent 如何检测错误并修正路线 |
| **教学模式** | 逐步播放，每个动作之间暂停，便于讲解 |
| **分层视图** | 初学者看简化版，开发者深入技术细节 |
| **多视图切换** | 树形图 / 时间线 / 流程图 自由切换 |

## 效果示意

```
任务：「调研最好的 Go Web 框架」
├── 🔍 搜索：「Go web frameworks 2026」→ 找到 15 条结果
│   ├── 📖 抓取：gin-gonic.com → 特性摘要
│   ├── 📖 抓取：echo.labstack.com → 特性摘要
│   └── 📖 抓取：fiber.wiki → 特性摘要
├── 🤔 分析：对比框架热度……
│   ├── 🏷️ Gin：GitHub Stars 最多，社区最大
│   ├── 🏷️ Echo：轻量，性能好
│   └── 🏷️ Fiber：类 Express API，速度快
└── ✅ 结论：生产推荐 Gin，微服务推荐 Echo
```

## 技术栈

| 环境 | 后端 | 数据库 | 前端 |
|------|------|--------|------|
| 在线版（Oxelia51） | Go | PostgreSQL | React + D3.js/Canvas |
| 桌面版（exe） | Go | SQLite | 内嵌 React + D3.js |

## 数据来源

- **MVP**：内置预设演示 Agent，无需外部配置即可体验
- **未来**：对接你自己的 Agent（LangChain、AutoGPT、CrewAI 或自定义回调）

## 使用方式

### 在线版（推荐）

1. 访问 [oxelia51.com](https://oxelia51.com) 注册登录
2. 从工具菜单打开 AgentCanvas
3. 选择一个演示任务，观察 Agent 的执行过程

### 桌面版（exe）

1. 从 [GitHub Releases](https://github.com/XiaoleC05/AgentCanvas/releases) 下载 `AgentCanvas.exe`
2. 运行即可，纯本地，无需联网

## 开发状态

概念阶段，尚未开发。
