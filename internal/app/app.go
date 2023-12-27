package app

import (
	"mydb/internal/database"
	"mydb/internal/database/compute"
	"mydb/internal/database/compute/analyzer"
	"mydb/internal/database/compute/parser"
	"mydb/internal/database/storage"

	"go.uber.org/zap"
)

type App struct {
	Database *database.Database
}

func New(logger *zap.Logger) *App {
	prs := parser.New(logger)
	anz := analyzer.New(logger)
	stg := storage.New(logger)
	cmp := compute.New(anz, prs, logger)
	db := database.New(cmp, stg, logger)

	return &App{
		Database: db,
	}
}
