package db

import (
	"context"

	"github.com/XiaoleC05/AgentCanvas/internal/model"
	"github.com/jackc/pgx/v5"
)

type EdgeRepository struct{}

func NewEdgeRepository() *EdgeRepository {
	return &EdgeRepository{}
}

func (r *EdgeRepository) Create(ctx context.Context, projectID int64, req model.CreateEdgeRequest) (*model.Edge, error) {
	query := `
		INSERT INTO agentcanvas.edges (project_id, source_node_id, target_node_id, label)
		VALUES ($1, $2, $3, $4)
		RETURNING id, project_id, source_node_id, target_node_id, label, created_at
	`

	var e model.Edge
	err := Pool.QueryRow(ctx, query,
		projectID, req.SourceNodeID, req.TargetNodeID, req.Label,
	).Scan(&e.ID, &e.ProjectID, &e.SourceNodeID, &e.TargetNodeID, &e.Label, &e.CreatedAt)

	return &e, err
}

func (r *EdgeRepository) ListByProject(ctx context.Context, projectID int64) ([]*model.Edge, error) {
	query := `
		SELECT id, project_id, source_node_id, target_node_id, label, created_at
		FROM agentcanvas.edges
		WHERE project_id = $1
		ORDER BY id
	`

	rows, err := Pool.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var edges []*model.Edge
	for rows.Next() {
		var e model.Edge
		if err := rows.Scan(&e.ID, &e.ProjectID, &e.SourceNodeID, &e.TargetNodeID, &e.Label, &e.CreatedAt); err != nil {
			return nil, err
		}
		edges = append(edges, &e)
	}

	return edges, nil
}

func (r *EdgeRepository) Delete(ctx context.Context, id int64, projectID int64) error {
	query := `DELETE FROM agentcanvas.edges WHERE id = $1 AND project_id = $2`
	_, err := Pool.Exec(ctx, query, id, projectID)
	return err
}

func (r *EdgeRepository) ExistsInProject(ctx context.Context, edgeID, projectID int64) (bool, error) {
	var exists bool
	err := Pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM agentcanvas.edges WHERE id = $1 AND project_id = $2)`,
		edgeID, projectID,
	).Scan(&exists)
	return exists, err
}

// GetProjectID 通过边 ID 反查所属项目 ID，用于权限校验
func (r *EdgeRepository) GetProjectID(ctx context.Context, edgeID int64) (int64, error) {
	var projectID int64
	err := Pool.QueryRow(ctx,
		`SELECT project_id FROM agentcanvas.edges WHERE id = $1`,
		edgeID,
	).Scan(&projectID)

	if err == pgx.ErrNoRows {
		return 0, nil
	}

	return projectID, err
}
