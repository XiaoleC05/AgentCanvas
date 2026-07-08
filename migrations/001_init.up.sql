-- 001_init: Create agentcanvas schema and tables

CREATE SCHEMA IF NOT EXISTS agentcanvas;

CREATE TABLE IF NOT EXISTS agentcanvas.projects (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name TEXT NOT NULL DEFAULT '未命名项目',
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS agentcanvas.nodes (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES agentcanvas.projects(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    label TEXT DEFAULT '',
    config JSONB DEFAULT '{}',
    position_x FLOAT DEFAULT 0,
    position_y FLOAT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS agentcanvas.edges (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES agentcanvas.projects(id) ON DELETE CASCADE,
    source_node_id BIGINT NOT NULL REFERENCES agentcanvas.nodes(id) ON DELETE CASCADE,
    target_node_id BIGINT NOT NULL REFERENCES agentcanvas.nodes(id) ON DELETE CASCADE,
    label TEXT DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT NOW()
);
