# AgentCanvas

可视化 AI Agent 的任务分解、推理链路与工具调用过程。

## Features

- 将 Agent 执行链路渲染为交互式树形图
- 支持逐步播放，每步可暂停查看详细推理内容
- 展示工具调用名称、参数和返回结果
- 区分初学者简版视图和开发者详细视图
- 支持树形图、时间线、流程图三种切换模式
- MVP 阶段内置预设演示 Agent，开箱可用

## Architecture

```text
Browser
  ↓
React Frontend (D3.js rendered graph)
  ↓
Go API Layer (agent task dispatch, result aggregation)
  ↓
PostgreSQL / SQLite (execution traces)
```

在线版运行于 Oxelia51 平台，由平台前端统一渲染界面，Go 后端负责 Agent 任务调度和结果聚合。桌面版使用 SQLite 替代 PostgreSQL，通过 Go 二进制内嵌 React 前端运行。

## Requirements

在线版依赖 Oxelia51 平台（Go + PostgreSQL + React）。桌面版为独立可执行文件，无需运行时依赖。

## Installation

### 桌面版

从 [GitHub Releases](https://github.com/XiaoleC05/AgentCanvas/releases) 下载 `AgentCanvas.exe`。

### 在线版

在线版集成于 Oxelia51 平台，参见 [Oxelia51 部署指南](https://github.com/XiaoleC05/Oxelia51)。

## Usage

### 在线

1. 访问 [oxelia51.com](https://oxelia51.com) 注册并登录
2. 进入 AgentCanvas 工具页
3. 选择演示任务，观察 Agent 的执行过程

### 桌面

1. 双击 `AgentCanvas.exe` 启动
2. 选择内置演示任务，无需联网

## Roadmap

- [ ] 内置演示 Agent（MVP）
- [ ] 对接自定义 Agent 框架（LangChain, AutoGPT, CrewAI）

## Contributing

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/xxx`)
3. 提交变更 (`git commit -m 'Add xxx'`)
4. 推送分支 (`git push origin feature/xxx`)
5. 提交 Pull Request

## License

This project is licensed under the MIT License.
