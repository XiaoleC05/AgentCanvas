# AgentCanvas

Visualize AI agent task decomposition, reasoning chains, and tool calls in an interactive tree diagram.

## Features

- Render agent execution traces as an interactive tree diagram
- Step-by-step playback with pause for detailed reasoning view
- Display tool call names, parameters, and return values
- Layered views for beginners (simplified) and developers (detailed)
- Toggle between tree diagram, timeline, and flowchart layouts
- Built-in demo agents for MVP, no external setup required

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

The online version runs on the Oxelia51 platform with unified frontend rendering. The Go backend manages agent task dispatch and result aggregation. The desktop version uses SQLite and embeds the React frontend within the Go binary.

## Requirements

Online version depends on the Oxelia51 platform (Go, PostgreSQL, React). Desktop version is a standalone executable with no runtime dependencies.

## Installation

### Desktop

Download `AgentCanvas.exe` from [GitHub Releases](https://github.com/XiaoleC05/AgentCanvas/releases).

### Online

Integrated into the Oxelia51 platform. See [Oxelia51 deployment guide](https://github.com/XiaoleC05/Oxelia51).

## Usage

### Online

1. Visit [oxelia51.com](https://oxelia51.com), register and sign in
2. Open AgentCanvas from the tools menu
3. Select a demo task to visualize the agent execution

### Desktop

1. Double-click `AgentCanvas.exe` to start
2. Select a built-in demo task. No internet required.

## Roadmap

- [ ] Built-in demo agents (MVP)
- [ ] Custom agent framework integration (LangChain, AutoGPT, CrewAI)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/xxx`)
3. Commit your changes (`git commit -m 'Add xxx'`)
4. Push the branch (`git push origin feature/xxx`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
