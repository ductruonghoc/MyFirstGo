package model

import (
	"database/sql"
	"log"
)

type Procedure struct {
	ID uint16 `json:"id"`
	Name string `json:"name"`
	Price int16 `json:"price"`
	Cost int16 `json:"cost"`
	Inventory int16 `json:"inventory"`
	NameDict string `json:"name_dict"`
};
//Scan to modeling procedure
func ProcedureModel(rows *sql.Rows) (procedures []Procedure){
	//run when loop done, free resources
	defer rows.Close();
	//Each rows
	for rows.Next() {
		//Clone struct
		procedure := new(Procedure);
		//Structure rows by clone
		err := rows.Scan(
			&procedure.ID, 
			&procedure.Name, 
			&procedure.Price, 
			&procedure.Cost, 
			&procedure.Inventory,
			&procedure.NameDict);
		//Err
		if err != nil {
			log.Fatal(err);
		}
		//Put clone to struct
		procedures = append(procedures, *procedure);
	}

	return;
}