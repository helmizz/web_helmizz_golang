package main

import (
	"fmt"
	"log"
	"os"
	_ "path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"	

//	"app/middlewares"	
	
	"app/routes"
	"app/utils"
)

func init() {
	// Setup logging
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
	
	// Load general configuration
	utils.LoadConfig("generalconfig")
	
	// Init audit logger (log login/logout)
	utils.InitAuditLogger()

	// Inisialisasi Redis client
//	middlewares.InitRedisClient()
}

func main() {
	// Create a new Echo instance
	e := echo.New()
	
//	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key"))))
	 
/*	 
	const CSRFTokenHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"


	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))
*/	
		// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	 e.Use(middleware.CORS())
	
	// Initialize database
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	
	// Static files
	e.Static("/assets", "statics/assets/")
	
	// Initialize routes
	routes.InitRoutes(e, db)
	
	// Get server configuration
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
	
	log.Printf("==========================")
	log.Printf("Starting server at %s", port)
	log.Printf("==========================")
	
	
	if err := e.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatal(err)
	}
	
	// Start server
	// e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
