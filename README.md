# AgentCanvas

Visualize task decomposition, reasoning chains, and execution steps in an interactive tree diagram.

## Features

- Render execution traces as interactive diagrams
- Step-by-step playback with detailed view for each node
- Toggle between tree diagram, timeline, and flowchart layouts
- Built-in demo tasks for getting started

## Architecture

```text
Browser (D3.js rendered graph)
  → Go API Layer (task dispatch, result aggregation)
  → PostgreSQL / SQLite (execution traces)
```

## Installation

Integrated into the Oxelia51 platform. See [Oxelia51 deployment guide](https://github.com/XiaoleC05/Oxelia51).

## Usage

1. Visit [oxelia51.com](https://oxelia51.com), register and sign in
2. Open AgentCanvas from the tools menu
3. Select a demo task to visualize the execution

## Contributing

1. Fork → 2. Feature branch → 3. Commit → 4. Push → 5. PR

## License

MIT License
