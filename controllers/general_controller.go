package controllers

import (
	"database/sql"
	_ "log"
	"github.com/labstack/echo/v4"
	"app/utils"
	"net/http"
)

/*
// GeneralController handles all general requests
type GeneralController struct {
	db          *sql.DB
}

// NewGeneralController creates a new general controller
func newGeneralController(db *sql.DB) *GeneralController {
	return &GeneralController{
		db: db,
	}
}

*/


type GeneralController struct {
	db *sql.DB
}

func newGeneralController() *GeneralController {
	return &GeneralController{db: GetDB()}
}





func init() {
	controller := newGeneralController()
	GetRegistry["generalcontroller"] = controller.HandleRequestget
	PostRegistry["generalcontroller"] = controller.HandleRequestpost
	PutRegistry["generalcontroller"] = controller.HandleRequestput
	DeleteRegistry["generalcontroller"] = controller.HandleRequestdelete
}


// HandleRequestget processes requests based on the page config
func (c *GeneralController) HandleRequestget(ctx echo.Context, viewName, pageConfigFile,redirectTrue,redirectFalse string) error {

	// Load page configuration and process variables
	data := utils.LoadConfigPage(pageConfigFile, ctx)
	data["csrf"] = ctx.Get("csrf")
	// Add any additional data needed for the view
	 data["appName"] = "Dashboard App"	
	// Render the template with the data
	return ctx.Render(200, viewName, data)
}

// HandleRequestpost processes requests based on the page config
func (c *GeneralController) HandleRequestpost(ctx echo.Context, viewName, pageConfigFile,redirectTrue,redirectFalse string) error {
	// Load page configuration and process variables
	 data := utils.LoadConfigPage(pageConfigFile, ctx)
	
	// return ctx.String(http.StatusOK,"berhasil" )
	return ctx.JSON(http.StatusOK, data)	
}

// HandleRequestput processes requests based on the page config
func (c *GeneralController) HandleRequestput(ctx echo.Context, viewName, pageConfigFile,redirectTrue,redirectFalse string) error {
	// Load page configuration and process variables
	data := utils.LoadConfigPage(pageConfigFile, ctx)
	
	// Add any additional data needed for the view
	data["appName"] = "Dashboard App"
	
	// Render the template with the data
	return ctx.JSON(http.StatusOK, data)
}

// HandleRequestdelete processes requests based on the page config
func (c *GeneralController) HandleRequestdelete(ctx echo.Context, viewName, pageConfigFile,redirectTrue,redirectFalse string) error {
	// Load page configuration and process variables
	data := utils.LoadConfigPage(pageConfigFile, ctx)
	
	// Add any additional data needed for the view
	data["appName"] = "Dashboard App"
	
	// Render the template with the data
	return ctx.JSON(http.StatusOK, data)
}
