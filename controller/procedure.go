package controller

import (
	"database/sql"
	"encoding/json"
	"io"
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
	fn := func(context *gin.Context) {
		query := "select * from latest_procs_version()";
		rows, err := db.Query(query)
		if err != nil { //Query got error
			log.Println(err)
			context.Status(http.StatusInternalServerError);
			return;
		}
		procedure, err := model.ProcedureModel(rows);
		if err != nil { //Model err
			log.Println(err)
			context.Status(http.StatusBadRequest);
			return;
		}
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.IndentedJSON(http.StatusOK, procedure);
	}
	return gin.HandlerFunc(fn);
}

func ImportExportSingleItem(db *sql.DB) gin.HandlerFunc {
	//Handler func as lambda
	fn := func(context *gin.Context) {
		id := context.Param("id")
		value := context.Param("value")
		_, err := db.Query("UPDATE Procs_version SET inventory=$1 WHERE id=$2", value, id)
		if err != nil { //Query got error
			log.Println(err)
			context.Status(http.StatusBadRequest)
			return;
		}
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.Status(http.StatusOK)
	}
	return gin.HandlerFunc(fn)
}

type ioItem struct {
	ID       uint32 `json:"id"`
	Quantity int16  `json:"quantity`
}

func ImportExportItems(db *sql.DB) gin.HandlerFunc {
	//Handler func as lambda
	fn := func(context *gin.Context) {
		jsonData, err := io.ReadAll(context.Request.Body)
		if err != nil { //Cant reâd body
			log.Println(err)
			context.Status(http.StatusInternalServerError)
			return;
		}
		var arr []ioItem;
		err = json.Unmarshal([]byte(jsonData), &arr);
		if err != nil {
			log.Println(err)
			context.Status(http.StatusBadRequest)
			return;
		}
		for _, element := range arr {
			log.Println(element);
			id := element.ID;
			quantity := element.Quantity;
			_, err = db.Query("UPDATE Procs_version SET inventory=inventory+$1 WHERE id=$2", quantity, id);
			if err != nil {
				context.Status(http.StatusInternalServerError);
				return;
			}
		}
		// id := context.Param("id");
		// value := context.Param("value");
		// msg := new(Message);
		// msg.SetContent(procedure);
		//Intend struct to JSON
		context.Status(http.StatusOK);
	}
	return gin.HandlerFunc(fn)
}
