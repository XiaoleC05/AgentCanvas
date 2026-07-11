package handler

import (
	"net/http"
	"strconv"

	"github.com/XiaoleC05/AgentCanvas/internal/db"
	"github.com/XiaoleC05/AgentCanvas/internal/model"
	"github.com/gin-gonic/gin"
)

type NodeHandler struct {
	repo        *db.NodeRepository
	projectRepo *db.ProjectRepository
}

func NewNodeHandler() *NodeHandler {
	return &NodeHandler{
		repo:        db.NewNodeRepository(),
		projectRepo: db.NewProjectRepository(),
	}
}

// Create POST /api/projects/:id/nodes — 创建节点（需要项目归属验证）
func (h *NodeHandler) Create(c *gin.Context) {
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

	var req model.CreateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node, err := h.repo.Create(c.Request.Context(), projectID, req)
	if err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"node": node})
}

// Update PUT /api/nodes/:id — 更新节点位置/配置（通过节点反查项目做权限校验）
func (h *NodeHandler) Update(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	nodeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid node id"})
		return
	}

	projectID, err := h.repo.GetProjectID(c.Request.Context(), nodeID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if projectID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "node not found"})
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

	var req model.UpdateNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node, err := h.repo.Update(c.Request.Context(), nodeID, projectID, req)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "node not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"node": node})
}

// Delete DELETE /api/nodes/:id
func (h *NodeHandler) Delete(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}

	nodeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid node id"})
		return
	}

	projectID, err := h.repo.GetProjectID(c.Request.Context(), nodeID)
	if err != nil {
		respondInternalError(c, err)
		return
	}
	if projectID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "node not found"})
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

	if err := h.repo.Delete(c.Request.Context(), nodeID, projectID); err != nil {
		respondInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
