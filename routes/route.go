package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	
	"app/controllers"
	"app/middlewares"
	"app/utils"
)

// RouteConfig represents the structure of the route configuration
type RouteConfig struct {
	Routes []struct {
		Path       string `json:"path"`
		Method     string `json:"method"`
		Handler    string `json:"handler"`
		View       string `json:"view"`
		PageConfig string `json:"pageconfig"`
		ResponReturn string   `json:"responreturn"`
		RedirectTrue string `json:"redirecttrue"`
		RedirectFalse string `json:"redirectfalse"`
		Middlewares   []string `json:"middlewares"`
	} `json:"routes"`
}

// InitRoutes initializes all routes from the route configuration file
func InitRoutes(e *echo.Echo, db *sql.DB) {
	// Load route configuration
	file, err := os.Open(filepath.Join("configs", "routeconfig.json"))
	if err != nil {
		log.Fatalf("Error opening route config file: %v", err)
	}
	defer file.Close()
	
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading route config file: %v", err)
	}
	
	var config RouteConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		log.Fatalf("Failed to parse route config: %v", err)
	}
	
	// Initialize template renderer
	renderer := utils.NewTemplateRenderer("views")
	e.Renderer = renderer
	
	for _, r := range config.Routes {
		method := r.Method
		handlerKey := r.Handler
		view := r.View
		pageConfig := r.PageConfig
		responReturn := r.ResponReturn
		redirectTrue := r.RedirectTrue
		redirectFalse := r.RedirectFalse

		var h controllers.HandlerFunc
		switch method {
		case "GET":
			h = controllers.GetRegistry[handlerKey]
		case "POST":
			h = controllers.PostRegistry[handlerKey]
		case "PUT":
			h = controllers.PutRegistry[handlerKey]
		case "DELETE":
			h = controllers.DeleteRegistry[handlerKey]
		}

		if h == nil {
			log.Printf("No handler found for %s %s", method, handlerKey)
			continue
		}

		// Convert handler
	routeFunc := func(c echo.Context) error {
		log.Printf("Accessed %s %s by %s", method, r.Path, handlerKey)
		return h(c, view, pageConfig, redirectTrue, redirectFalse)
	}

	// ==== Ambil middleware dinamis ====
	var mwChain []echo.MiddlewareFunc
	for _, m := range r.Middlewares {
		if strings.Contains(m, ":") {
			parts := strings.SplitN(m, ":", 2)
			mid := middlewares.GetWithParam(parts[0], parts[1])
			if mid != nil {
				mwChain = append(mwChain, mid)
			}
		} else {
			mid := middlewares.Get(m)
			if mid != nil {
				mwChain = append(mwChain, mid)
			}
		}
	}

		// Bungkus handler dengan middleware (manual chaining)
		finalHandler := routeFunc
		for i := len(mwChain) - 1; i >= 0; i-- {
			finalHandler = mwChain[i](finalHandler)
		}



			switch method {
		case "GET":
			e.GET(r.Path, finalHandler)
		case "POST":
			e.POST(r.Path, finalHandler)
		case "PUT":
			e.PUT(r.Path, finalHandler)
		case "DELETE":
			e.DELETE(r.Path, finalHandler)
		}
	}
		
}
