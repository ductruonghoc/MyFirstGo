package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
);

type Database struct {
	DB *sql.DB
};

func (d *Database) DBConnection() {
	connStr := os.Getenv("DBCONNECTION"); //Get credentials
	// Set the maximum number of idle connections in the pool
	idleConn := 50
	// Set the maximum number of connections in the pool
	maxConnections := 90
	// Set the maximum amount of time a connection can be reused
	maxConnLifetime := 2 * time.Minute
	db, err := sql.Open("postgres", connStr); //init
	if err != nil { 
		log.Fatal(err);
	}
	if db == nil { //nil db
		log.Fatal("Dont have db");
	}
	db.SetMaxOpenConns(maxConnections);
	db.SetMaxIdleConns(idleConn);
	db.SetConnMaxLifetime(maxConnLifetime);
	d.DB = db; //update
}