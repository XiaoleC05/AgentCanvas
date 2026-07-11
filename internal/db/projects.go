package db

import (
	"context"

	"github.com/XiaoleC05/AgentCanvas/internal/model"
	"github.com/jackc/pgx/v5"
)

type ProjectRepository struct{}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{}
}

func (r *ProjectRepository) Create(ctx context.Context, userID int64, req model.CreateProjectRequest) (*model.Project, error) {
	query := `
		INSERT INTO agentcanvas.projects (user_id, name)
		VALUES ($1, $2)
		RETURNING id, user_id, name, created_at, updated_at
	`

	var p model.Project
	err := Pool.QueryRow(ctx, query, userID, req.Name).Scan(
		&p.ID, &p.UserID, &p.Name, &p.CreatedAt, &p.UpdatedAt,
	)

	return &p, err
}

func (r *ProjectRepository) List(ctx context.Context, userID int64) ([]*model.Project, error) {
	query := `
		SELECT id, user_id, name, created_at, updated_at
		FROM agentcanvas.projects
		WHERE user_id = $1
		ORDER BY updated_at DESC
	`

	rows, err := Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		var p model.Project
		if err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	return projects, nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id, userID int64) (*model.Project, error) {
	query := `
		SELECT id, user_id, name, created_at, updated_at
		FROM agentcanvas.projects
		WHERE id = $1 AND user_id = $2
	`

	var p model.Project
	err := Pool.QueryRow(ctx, query, id, userID).Scan(
		&p.ID, &p.UserID, &p.Name, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &p, err
}

func (r *ProjectRepository) GetDetail(ctx context.Context, id, userID int64) (*model.ProjectDetail, error) {
	p, err := r.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, nil
	}

	nodeRepo := NewNodeRepository()
	edgeRepo := NewEdgeRepository()

	nodes, err := nodeRepo.ListByProject(ctx, id)
	if err != nil {
		return nil, err
	}

	edges, err := edgeRepo.ListByProject(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.ProjectDetail{
		Project: *p,
		Nodes:   nodes,
		Edges:   edges,
	}, nil
}

func (r *ProjectRepository) Update(ctx context.Context, id, userID int64, req model.UpdateProjectRequest) (*model.Project, error) {
	query := `
		UPDATE agentcanvas.projects
		SET name = COALESCE(NULLIF($3, ''), name),
		    updated_at = NOW()
		WHERE id = $1 AND user_id = $2
		RETURNING id, user_id, name, created_at, updated_at
	`

	var p model.Project
	err := Pool.QueryRow(ctx, query, id, userID, req.Name).Scan(
		&p.ID, &p.UserID, &p.Name, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &p, err
}

func (r *ProjectRepository) Delete(ctx context.Context, id, userID int64) (int64, error) {
	query := `DELETE FROM agentcanvas.projects WHERE id = $1 AND user_id = $2`
	result, err := Pool.Exec(ctx, query, id, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
