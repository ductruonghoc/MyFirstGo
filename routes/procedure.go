package routes

import (

	"github.com/gin-gonic/gin"

	"main/controller"
	"main/database"
)

type Procedure struct {
	Root *gin.RouterGroup
	DB *database.Database
}

func (proc *Procedure) SetRoot(root *gin.RouterGroup) {
	proc.Root = root;
}

func (proc *Procedure) SetDB(db *database.Database) {
	proc.DB = db;
}


func (proc *Procedure) Routes() {
	db := proc.DB.DB
	proc.Root.GET("/inventory", controller.Inventory(db));
}