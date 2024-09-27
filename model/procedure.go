package model

import (
	"database/sql"
	"log"
)

type Procedure struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Price uint32 `json:"price"`
	Cost uint32 `json:"cost"`
	Inventory uint16 `json:"inventory"`
	NameDict string `json:"name_dict"`
	Unit string `json:"unit"`
	RootID uint32 `json:"root_id"`
	RootIs uint16 `json:"root_is"`
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
			&procedure.NameDict,
			&procedure.Unit,
			&procedure.RootID,
			&procedure.RootIs);
		//Err
		if err != nil {
			log.Fatal(err);
		}
		//Put clone to struct
		procedures = append(procedures, *procedure);
	}

	return;
}