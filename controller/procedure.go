package controller

import (
	"database/sql"
	"log"
	"github.com/ductruonghoc/MyFirstGo/main/model"
	"net/http"

	"github.com/gin-gonic/gin"
);

// type Message struct {
// 	Content []model.Procedure `json:"content"`
// }

// func (msg *Message) SetContent(content []model.Procedure) {
// 	msg.Content = content;
// }

func Inventory(db *sql.DB) gin.HandlerFunc {
	//Handler func as lambda
	fn := func (context *gin.Context)  {
		rows, err := db.Query("SELECT id, procs_name, price, cap, inventory FROM Procs");
		if err != nil {
			log.Fatal(err);
		}
		procedure := model.ProcedureModel(rows);
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.IndentedJSON(http.StatusOK,procedure);		
	} 
	return gin.HandlerFunc(fn);
}