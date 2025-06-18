# Dashboard App Documentation

## Overview

Dashboard App is a flexible web application built with Go using the Echo framework. It follows the MVC (Model-View-Controller) pattern with additional service and utilities layers. The application is designed to be highly configurable, with most changes possible via JSON configuration files without requiring code changes.

## Features

- **Configurable**: Most functionality can be adjusted via JSON configuration files
- **Multiple Database Support**: Compatible with MySQL, PostgreSQL, and MS SQL Server
- **Hot Reload Templates**: In development mode, HTML templates are reloaded on each request
- **Middleware Support**: Authentication, logging, and more built-in
- **MVC Architecture**: Clean separation of concerns with Models, Views, Controllers

## Project Structure

- **configs/**: Configuration files in JSON format
- **controllers/**: Controller logic
- **middlewares/**: Middleware functions
- **models/**: Data models
- **routes/**: Route definitions
- **services/**: Business logic services
- **utils/**: Utility functions
- **views/**: HTML templates
- **docs/**: Documentation
- **logs/**: Log files

## Configuration

The application uses three types of configuration files:

1. **generalconfig.json**: Main application configuration
2. **routeconfig.json**: Route definitions and handlers
3. **pageXconfig.json**: Page-specific configuration (variables, queries, etc.)

### General Configuration

```json
{
  "app_name": "Dashboard App",
  "is_production": false,
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "database": {
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "user": "root",
    "password": "root",
    "name": "bi_db"
  }
}
```

### Route Configuration

```json
{
  "routes": [
    {
      "path": "/",
      "method": "GET",
      "handler": "generalcontroller",
      "view": "page1.html",
      "page1config": "page1config.json"
    }
  ]
}
```

### Page Configuration

```json
{
  "data": [
    { "type": 1, "variable_name": "title", "variable_value": "Welcome" },
    { "type": 6, "variable_name": "users", "variable_value": "SELECT * FROM users LIMIT 10" }
  ]
}
```

## Variable Types

- **Type 1**: Direct value assignment
- **Type 2**: Form value
- **Type 3**: Query parameter
- **Type 4**: Path parameter
- **Type 5**: Form file
- **Type 6**: SQL query with variable replacement

## Running the Application

1. Install dependencies: `go mod download`
2. Run the application: `go run main.go`
3. Access the application at `http://localhost:8080`

## Making Changes

To make changes to the application:

1. Edit the appropriate configuration file in the `configs/` directory
2. Edit the corresponding HTML template in the `views/` directory
3. If in development mode (`is_production = false`), changes to templates will be reflected without restarting the server

## Database Support

The application supports multiple database backends:

- MySQL
- PostgreSQL
- Microsoft SQL Server

Configure the database connection in `generalconfig.json`.
