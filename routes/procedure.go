package routes

import (

	"github.com/gin-gonic/gin"

	"github.com/ductruonghoc/MyFirstGo/main/controller"
	"github.com/ductruonghoc/MyFirstGo/main/database"
	"github.com/gin-contrib/cors"
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
	proc.Root.GET("/inventory", cors.Default() ,controller.Inventory(db));
}