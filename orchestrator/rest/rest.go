package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/jonas747/dshardorchestrator/orchestrator"
)

type RESTAPI struct {
	orchestrator *orchestrator.Orchestrator
	listenAddr   string

	g *gin.Engine
}

func NewRESTAPI(orchestrator *orchestrator.Orchestrator, listenAddr string) *RESTAPI {
	return &RESTAPI{
		orchestrator: orchestrator,
		listenAddr:   listenAddr,
		g:            gin.Default(),
	}
}

func (ra *RESTAPI) Run() error {
	ra.setupRoutes()
	err := ra.g.Run(ra.listenAddr)
	return err
}

func (ra *RESTAPI) setupRoutes() {
	ra.g.GET("/status", ra.handleGETStatus)
	ra.g.POST("/startnode", ra.handlePOSTStartNode)
	ra.g.POST("/shutdownnode", ra.handlePOSTShutdownNode)
	ra.g.POST("/migrateshard", ra.handlePOSTMigrateShard)
	ra.g.POST("/migratenode", ra.handlePOSTMigrateNode)
	ra.g.POST("/fullmigration", ra.handlePOSTFullMigration)
	ra.g.POST("/stopshard", ra.handlePOSTStopShard)
}
