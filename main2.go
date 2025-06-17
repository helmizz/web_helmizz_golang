package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"app/middlewares"
	"app/routes"
	"app/utils"
)

func init() {
	// Logging ke file
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)

	// Load config umum
	utils.LoadConfig("generalconfig")

	// Init audit logger (log login/logout)
	utils.InitAuditLogger()

	// Inisialisasi Redis client
	middlewares.InitRedisClient()
}

func main() {
	e := echo.New()

	// === MIDDLEWARE GLOBAL ===

	// Gunakan Redis session store
	e.Use(session.Middleware(middlewares.NewRedisSessionStore()))

	// CORS
	e.Use(middleware.CORS())

	// CSRF (optional)
	/*
	const CSRFTokenHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))
	*/

	// === INISIALISASI DATABASE ===
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	// === STATIC FILES ===
	e.Static("/assets", "statics/assets/")

	// === ROUTING DINAMIS ===
	routes.InitRoutes(e, db)

	// === KONFIG SERVER ===
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")

	log.Printf("🚀 Starting server at %s:%d", host, port)

	if err := e.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatal(err)
	}
}
