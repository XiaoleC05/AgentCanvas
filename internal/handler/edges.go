package handler

import (
	"net/http"
	"strconv"

	"github.com/XiaoleC05/AgentCanvas/internal/db"
	"github.com/XiaoleC05/AgentCanvas/internal/model"
	"github.com/gin-gonic/gin"
)

type EdgeHandler struct {
	repo        *db.EdgeRepository
	projectRepo *db.ProjectRepository
	nodeRepo    *db.NodeRepository
}

func NewEdgeHandler() *EdgeHandler {
	return &EdgeHandler{
		repo:        db.NewEdgeRepository(),
		projectRepo: db.NewProjectRepository(),
		nodeRepo:    db.NewNodeRepository(),
	}
}

// Create POST /api/projects/:id/edges — 创建边（需要项目归属验证 + 节点合法性校验）
func (h *EdgeHandler) Create(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	project, err := h.projectRepo.GetByID(c.Request.Context(), projectID, userID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if project == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	var req model.CreateEdgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceExists, err := h.nodeRepo.ExistsInProject(c.Request.Context(), req.SourceNodeID, projectID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if !sourceExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "source node not in project"})
		return
	}

	targetExists, err := h.nodeRepo.ExistsInProject(c.Request.Context(), req.TargetNodeID, projectID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if !targetExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "target node not in project"})
		return
	}

	edge, err := h.repo.Create(c.Request.Context(), projectID, req)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"edge": edge})
}

// Delete DELETE /api/edges/:id — 通过边反查项目做权限校验
func (h *EdgeHandler) Delete(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	edgeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid edge id"})
		return
	}

	projectID, err := h.repo.GetProjectID(c.Request.Context(), edgeID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if projectID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "edge not found"})
		return
	}

	project, err := h.projectRepo.GetByID(c.Request.Context(), projectID, userID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if project == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	if err := h.repo.Delete(c.Request.Context(), edgeID, projectID); err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
