package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ductruonghoc/MyFirstGo/main/model"

	"github.com/gin-gonic/gin"
)

// type Message struct {
// 	Content []model.Procedure `json:"content"`
// }

// func (msg *Message) SetContent(content []model.Procedure) {
// 	msg.Content = content;
// }

func Inventory(db *sql.DB) gin.HandlerFunc {
	//Handler func as lambda
	fn := func (context *gin.Context)  {
		rows, err := db.Query("SELECT id, procs_name, price, cap, inventory, name_dict FROM Procs");
		if err != nil {
			log.Fatal(err);
		}
		procedure := model.ProcedureModel(rows);
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.IndentedJSON(http.StatusOK, procedure);		
	} 
	return gin.HandlerFunc(fn);
}

func ImportExportSingleItem(db *sql.DB) gin.HandlerFunc {
	//Handler func as lambda
	fn := func (context *gin.Context)  {
		id := context.Param("id");
		value := context.Param("value");
		_, err := db.Query("UPDATE Procs SET inventory=$1 WHERE id=$2", value, id);
		if err != nil {
			log.Fatal(err);
		}
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.Status(http.StatusOK);	
	} 
	return gin.HandlerFunc(fn);
}