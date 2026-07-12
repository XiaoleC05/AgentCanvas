# AgentCanvas

> Agent 执行流程可视化 — 将任务分解、推理链路、执行步骤渲染为交互式图表

[![Go](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)](https://go.dev/)
[![D3.js](https://img.shields.io/badge/D3.js-v7-F9A03C?logo=d3.js)](https://d3js.org/)
[![Oxelia51](https://img.shields.io/badge/Oxelia51-v2.3-6366f1)](https://oxelia51.com)
[![License](https://img.shields.io/badge/license-MIT-green)](./LICENSE)

## 简介

AgentCanvas 将 AI Agent 的任务执行过程可视化为交互式图表。支持树状图、时间线和流程图三种布局，可逐步回放每个节点的执行细节。

## 特性

- 🌳 **交互式树状图** — 任务分解与推理链路一目了然
- ⏱️ **逐步回放** — 每个节点的输入输出、耗时、状态可展开查看
- 🔀 **三种布局** — 树状图 / 时间线 / 流程图一键切换
- 🎬 **内置示例** — 预置演示任务，无需配置即可体验
- 📊 **SQLite 本地存储** — 执行追踪数据本地持久化

## 架构

```text
浏览器 (D3.js 渲染)
  → Go API (任务调度、结果聚合)
  → PostgreSQL / SQLite (执行追踪)
```

## 部署

集成于 [Oxelia51 平台](https://github.com/XiaoleC05/Oxelia51)。详见 [Oxelia51 部署指南](https://github.com/XiaoleC05/Oxelia51/blob/master/deploy/README.md)。

## 使用

1. 访问 [oxelia51.com](https://oxelia51.com)，注册登录
2. 打开 AgentCanvas 工具
3. 选择演示任务，查看可视化的执行流程

## 快速开始

```bash
git clone https://github.com/XiaoleC05/AgentCanvas.git
cd AgentCanvas
go build -o agentcanvas-server ./cmd/server
./agentcanvas-server
```

## 技术栈

| 层级 | 方案 |
|------|------|
| 后端 | Go + Gin |
| 数据库 | PostgreSQL / SQLite |
| 可视化 | D3.js |
| 前端 | React (Oxelia51 统一 UI) |

## License

MIT License
