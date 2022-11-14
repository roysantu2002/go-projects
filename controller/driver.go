package controller

import (
	// "context"
	"database/sql"
	"fmt"
	"os"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	 "github.com/joho/godotenv"
	 "strconv"
	 "reflect"
)

var database *sql.DB

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
	fmt.Printf("%s uses %s\n", os.Getenv("MSSQL_DB_SERVER"), os.Getenv("MSSQL_DB_USER"))
}

// use os package to get the env variable which is already set
func envVariable(key string) string {

  // return the env variable using os package
  return os.Getenv(key)
}


// var (
// 	 server = os.Getenv("MSSQL_DB_SERVER")
// 	 password = os.Getenv("MSSQL_DB_PASSWORD")
//    	 user = os.Getenv("MSSQL_DB_USER")
//    	//  port = os.Getenv("MSSQL_DB_PORT")
//    	 port = 1433
//    	//  db = os.Getenv("MSSQL_DB_DATABASE")
// 	// port = 1433
// 	// password = "iamnwani01"
// 	// user = "SA"
// )

// type Database struct {
//    SqlDb *sql.DB
// }

// // var dbContext = context.Background()
// //
// // var connectionString = fmt.Sprintf("user id=%s;password=%s;port=%d", user, password, port)
//
// var dbContext = context.Background()
//

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {

	server := envVariable("MSSQL_DB_SERVER")
	user := envVariable("MSSQL_DB_USER")
	password := envVariable("MSSQL_DB_PASSWORD")
	port := envVariable("MSSQL_DB_PORT")

	fmt.Printf("t1: %s\n", reflect.TypeOf(port))
	port_, err := strconv.Atoi(port)

	var connectionString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port_)


	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}