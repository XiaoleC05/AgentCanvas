package db

import (
	"context"

	"github.com/XiaoleC05/AgentCanvas/internal/model"
	"github.com/jackc/pgx/v5"
)

type NodeRepository struct{}

func NewNodeRepository() *NodeRepository {
	return &NodeRepository{}
}

func (r *NodeRepository) Create(ctx context.Context, projectID int64, req model.CreateNodeRequest) (*model.Node, error) {
	query := `
		INSERT INTO agentcanvas.nodes (project_id, type, label, config, position_x, position_y)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, project_id, type, label, config, position_x, position_y, created_at
	`

	label := req.Label
	config := req.Config
	if config == "" {
		config = "{}"
	}

	var n model.Node
	err := Pool.QueryRow(ctx, query,
		projectID, req.Type, label, config, req.PositionX, req.PositionY,
	).Scan(&n.ID, &n.ProjectID, &n.Type, &n.Label, &n.Config, &n.PositionX, &n.PositionY, &n.CreatedAt)

	return &n, err
}

func (r *NodeRepository) ListByProject(ctx context.Context, projectID int64) ([]*model.Node, error) {
	query := `
		SELECT id, project_id, type, label, config, position_x, position_y, created_at
		FROM agentcanvas.nodes
		WHERE project_id = $1
		ORDER BY id
	`

	rows, err := Pool.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []*model.Node
	for rows.Next() {
		var n model.Node
		if err := rows.Scan(&n.ID, &n.ProjectID, &n.Type, &n.Label, &n.Config, &n.PositionX, &n.PositionY, &n.CreatedAt); err != nil {
			return nil, err
		}
		nodes = append(nodes, &n)
	}

	return nodes, nil
}

func (r *NodeRepository) Update(ctx context.Context, id int64, projectID int64, req model.UpdateNodeRequest) (*model.Node, error) {
	query := `
		UPDATE agentcanvas.nodes
		SET label = COALESCE(NULLIF($3, ''), label),
		    config = COALESCE(NULLIF($4, ''), config),
		    position_x = COALESCE($5, position_x),
		    position_y = COALESCE($6, position_y)
		WHERE id = $1 AND project_id = $2
		RETURNING id, project_id, type, label, config, position_x, position_y, created_at
	`

	var posX, posY interface{}
	if req.PositionX != nil {
		posX = *req.PositionX
	}
	if req.PositionY != nil {
		posY = *req.PositionY
	}

	var n model.Node
	err := Pool.QueryRow(ctx, query,
		id, projectID, req.Label, req.Config, posX, posY,
	).Scan(&n.ID, &n.ProjectID, &n.Type, &n.Label, &n.Config, &n.PositionX, &n.PositionY, &n.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &n, err
}

func (r *NodeRepository) Delete(ctx context.Context, id int64, projectID int64) error {
	query := `DELETE FROM agentcanvas.nodes WHERE id = $1 AND project_id = $2`
	_, err := Pool.Exec(ctx, query, id, projectID)
	return err
}

func (r *NodeRepository) ExistsInProject(ctx context.Context, nodeID, projectID int64) (bool, error) {
	var exists bool
	err := Pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM agentcanvas.nodes WHERE id = $1 AND project_id = $2)`,
		nodeID, projectID,
	).Scan(&exists)
	return exists, err
}

// GetProjectID 通过节点 ID 反查所属项目 ID，用于权限校验
func (r *NodeRepository) GetProjectID(ctx context.Context, nodeID int64) (int64, error) {
	var projectID int64
	err := Pool.QueryRow(ctx,
		`SELECT project_id FROM agentcanvas.nodes WHERE id = $1`,
		nodeID,
	).Scan(&projectID)

	if err == pgx.ErrNoRows {
		return 0, nil
	}

	return projectID, err
}
