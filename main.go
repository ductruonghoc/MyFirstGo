package main

import (
	"github.com/ductruonghoc/MyFirstGo/main/database"
	"github.com/ductruonghoc/MyFirstGo/main/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//DB
	db := new(database.Database); //init
	db.DBConnection(); //connect
	//API
	app := gin.Default();
	app.Use(cors.Default());
	procedureRoute := new(routes.Procedure);
	routes := app.Group("/api")
	{
		procedureRoute.SetRoot(routes.Group("/procedure"))
		{
			procedureRoute.SetDB(db);
			procedureRoute.Routes();
		}
	}

	app.Run(":8080");
}