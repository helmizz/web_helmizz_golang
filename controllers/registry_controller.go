package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(c echo.Context, view, pageConfig, redirectTrue, redirectFalse string) error

var (
	dbInstance      *sql.DB
	GetRegistry     = map[string]HandlerFunc{}
	PostRegistry    = map[string]HandlerFunc{}
	PutRegistry     = map[string]HandlerFunc{}
	DeleteRegistry  = map[string]HandlerFunc{}
)

// InitControllers sets the shared DB instance used by all controllers
func InitControllers(db *sql.DB) {
	dbInstance = db
}

// GetDB returns the global DB instance
func GetDB() *sql.DB {
	return dbInstance
}

// HandlerFunc

