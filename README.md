# AgentCanvas — Visualize How AI Agents Think

> Submit a task to an AI agent. Watch every step of its reasoning unfold in an interactive tree diagram.

## Why AgentCanvas?

AI agents are powerful but opaque. When an agent breaks down a task, calls tools, makes decisions, and recovers from errors — all of that happens in a black box. If you're learning about AI agents (or building one), you can't see what's actually going on.

**AgentCanvas** opens the black box. It renders an agent's full execution trace as an interactive visualization, so you can understand exactly how agents think, decide, and act.

## Features

| Feature | What It Shows |
|---------|---------------|
| **Task Decomposition** | How the agent breaks a complex goal into smaller steps |
| **Chain of Thought** | The agent's reasoning at each step — what it's thinking and why |
| **Tool Calls** | Which tools were invoked, with what parameters, and what came back |
| **Decision Branches** | Why the agent chose path A over path B, including backtracking |
| **Error Recovery** | How the agent detected failures and corrected course |
| **Teaching Mode** | Step-by-step playback that pauses between actions for explanation |
| **Layered View** | Beginners see a simplified view; developers drill into technical details |
| **Multi-View** | Switch between tree diagram, timeline, and flowchart layouts |

## How It Looks

Imagine a tree structure where each node represents one step the agent took:

```
Task: "Research best Go web frameworks"
├── 🔍 Search: "Go web frameworks 2026" → Found 15 results
│   ├── 📖 Fetch: gin-gonic.com → Summary of features
│   ├── 📖 Fetch: echo.labstack.com → Summary of features
│   └── 📖 Fetch: fiber.wiki → Summary of features
├── 🤔 Analyze: Comparing framework popularity...
│   ├── 🏷️ Gin: most GitHub stars, largest community
│   ├── 🏷️ Echo: lightweight, good performance
│   └── 🏷️ Fiber: Express-like API, fast
└── ✅ Conclusion: Recommend Gin for production, Echo for microservices
```

## Tech Stack

| Environment | Backend | Database | Frontend | Special |
|-------------|---------|----------|----------|---------|
| Online (Oxelia51) | Go | PostgreSQL | React + D3.js/Canvas | — |
| Desktop (exe) | Go | SQLite | Embedded React + D3.js | — |

## Data Source

- **MVP**: Built-in demo agents with preset tasks — no external setup needed
- **Future**: Connect your own agents (LangChain, AutoGPT, CrewAI, or custom HTTP callbacks)

## Getting Started

### Online (via Oxelia51)

1. Visit [oxelia51.com](https://oxelia51.com) and sign in
2. Open AgentCanvas from the tools menu
3. Select a demo task and watch the agent's execution unfold

### Desktop (exe)

1. Download `AgentCanvas.exe` from [GitHub Releases](https://github.com/XiaoleC05/AgentCanvas/releases)
2. Run the executable
3. Same experience as online, fully local

## Status

Concept phase. Development not yet started.
