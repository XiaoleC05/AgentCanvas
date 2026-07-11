package model

import (
	"time"
)

type Project struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectDetail struct {
	Project
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

type CreateProjectRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateProjectRequest struct {
	Name string `json:"name"`
}

type Node struct {
	ID        int64     `json:"id"`
	ProjectID int64     `json:"project_id"`
	Type      string    `json:"type"`
	Label     string    `json:"label"`
	Config    string    `json:"config"`
	PositionX float64   `json:"position_x"`
	PositionY float64   `json:"position_y"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNodeRequest struct {
	Type      string  `json:"type" binding:"required"`
	Label     string  `json:"label"`
	Config    string  `json:"config"`
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
}

type UpdateNodeRequest struct {
	Label     string   `json:"label"`
	Config    string   `json:"config"`
	PositionX *float64 `json:"position_x"`
	PositionY *float64 `json:"position_y"`
}

type Edge struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"project_id"`
	SourceNodeID int64     `json:"source_node_id"`
	TargetNodeID int64     `json:"target_node_id"`
	Label        string    `json:"label"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateEdgeRequest struct {
	SourceNodeID int64  `json:"source_node_id" binding:"required"`
	TargetNodeID int64  `json:"target_node_id" binding:"required"`
	Label        string `json:"label"`
}
