package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/XiaoleC05/AgentCanvas/internal/config"
	"github.com/XiaoleC05/AgentCanvas/internal/db"
	"github.com/XiaoleC05/AgentCanvas/internal/handler"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsOrigins() []string {
	if v := os.Getenv("CORS_ALLOWED_ORIGINS"); v != "" {
		return strings.Split(v, ",")
	}
	return []string{"http://localhost:5173"}
}

func main() {
	config.Load()

	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	runMigrations()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-User-Id", "X-Username", "X-Role"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	projectHandler := handler.NewProjectHandler()
	nodeHandler := handler.NewNodeHandler()
	edgeHandler := handler.NewEdgeHandler()

	r.GET("/api/health", handler.Health)

	api := r.Group("/api")
	api.Use(handler.AuthMiddleware())
	{
		api.GET("/projects", projectHandler.List)
		api.POST("/projects", projectHandler.Create)
		api.GET("/projects/:id", projectHandler.Get)
		api.PUT("/projects/:id", projectHandler.Update)
		api.DELETE("/projects/:id", projectHandler.Delete)

		api.POST("/projects/:id/nodes", nodeHandler.Create)
		api.POST("/projects/:id/edges", edgeHandler.Create)

		api.PUT("/nodes/:id", nodeHandler.Update)
		api.DELETE("/nodes/:id", nodeHandler.Delete)

		api.DELETE("/edges/:id", edgeHandler.Delete)
	}

	srv := &http.Server{
		Addr:    ":" + config.Cfg.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func runMigrations() {
	ctx := context.Background()

	initSQL := `
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
	`

	_, err := db.Pool.Exec(ctx, initSQL)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed")
}
