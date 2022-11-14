package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

var database *sql.DB

var (
	 server = os.Getenv("MSSQL_DB_SERVER")
	 password = os.Getenv("MSSQL_DB_PASSWORD")
   	 user = os.Getenv("MSSQL_DB_USER")
   	 port = os.Getenv("MSSQL_DB_PORT")
   	//  db = os.Getenv("MSSQL_DB_DATABASE")
	// port = 1433
	// password = "iamnwani01"
	// user = "SA"
)

// type Database struct {
//    SqlDb *sql.DB
// }

// var dbContext = context.Background()
//
// var connectionString = fmt.Sprintf("user id=%s;password=%s;port=%d", user, password, port)

var dbContext = context.Background()

var connectionString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)