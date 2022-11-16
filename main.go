package main

import (
		// "context"
	// "bufio"
	// "database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"test.com/packages/controller"
	// "os"
	  "log"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

type User struct {
	// ID        int             `db:"id"`
	// FirstName sql.NullString  `db:"first_name"`
	// loginName  string          `db:"loginname"`
	// Balance   sql.NullFloat64 `db:"balance"`
	serverName string `json:"serverName"`
	loginName  string
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// type LoginName struct {
//
//     Name       string
// }

type dbname struct {
		server string `json:"serverName"`
        name map[int]string `json:"name"`
    }


var (

	loginname string
	counter int
)



var connectionError error
// var database *sql.DB

func PrettyString(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}

func main() {
	// var err error

	// var slice []string // Create an empty slice of type int.
	 // initialize a slice literal
	// newSlice := []string
	  // Creating nil slice
    var newSlice []string
	userDetails := map[int]string{}

	db := controller.ConnectDB()
	// fmt.Println(db)

	//  var version string
	// Select for multiple values
	// users := []User{}
    rows, err2 := db.Query("select loginname from master.sys.syslogins")

    if err2 != nil {
        log.Fatal(err2)
    }






	for rows.Next() {

		// our query
		// user := User{}

		// dbname := dbname{}
		// if err := rows.Scan(&dbname.name); err != nil {
		if err := rows.Scan(&loginname); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

		newSlice =append(newSlice, loginname)
		//  usr := {'loginName':   loginname, serverName: "test1"}
		// newSlice = append(newSlice, usr)
		userDetails[counter+1] = loginname
		// userDetails["b"] = "test1"
    	//  usr.loginName = loginname
    	//  usr.serverName = "Server1"



		// users := []User{
        // {loginName: loginname, serverName: "Roger Roe"},}


// 		json_data2, err := json.Marshal(usr)
//
// 		if err != nil {
//
// 			log.Fatal(err)
// 		}
//
// 		fmt.Println(string(json_data2))

		// dbname.name = loginname
		// dbname.server = "teston"
		// fmt.Printf("found loginname: %+v\n", dbname)
		// newSlice= append(newSlice, {loginname, "testname"})
		counter = counter + 1
	}

	// dbname.server = "Test1"
	// dbname.name = userDetails

	// jsonStr, err := json.Marshal(userDetails)
    // if err != nil {
    //     fmt.Printf("Error: %s", err.Error())
    // } else {
    //     fmt.Println(string(jsonStr))
    // }

// 	a1 := dbname{"Test1", userDetails}
//
// 	fmt.Println(a1)

	// res, err := PrettyString(`{userDetails}`)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println(res)


	file, _ := json.MarshalIndent(userDetails, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}


	// s := User{}
  	// err2 = structScan(rows, &s)
    // defer res.Close()
//
//     if err != nil {
//         log.Fatal(err)
//     }

// 	fmt.Println(users)
//     var data dbname
//     for res.Next() {
// 		  err := res.Scan(&loginname)
//         if err != nil {
//             log.Fatal(err)
//         }
//
//         // fmt.Printf("", loginname)
// 		data := dbname{loginname, loginname}
		// data.server = "test"
		// data.name = loginname
		// slice = append(slice, loginname)
		// data = DBLoginname{server : "test1", loginname: loginname}


		// fmt.Printf(res['loginname'])
		// var alb newSlice
        // if err := res.Scan(&loginname); err != nil {
		// 	fmt.Printf(loginname)
        // }
        // newSlicem = append(albums, album)


//
	// fmt.Println(newSlice)
//
//         var loginName newSlice
//         err := res.Scan(&loginName.loginname)
//
//         if err != nil {
//             log.Fatal(err)
//         }

        // fmt.Printf("%v\n", loginName)
		// add an element to the slice
    	// newSlice = append(res, "e")
		// fmt.Println(s)

// 		file, _ := json.MarshalIndent(users, "", " ")
//
// 		_ = ioutil.WriteFile("test.json", file, 0644)
    // }


// 	j, err := json.Marshal(slice)
//     if err != nil {
//         fmt.Printf("Error: %s", err.Error())
//     } else {
//         fmt.Println(string(j))
//     }
//


	// fmt.Println(slice)
//
// 	query := `select loginname from master.sys.syslogins`
//     ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancelfunc()
//     stmt, err := db.PrepareContext(ctx, query)
//     if err != nil {
//         log.Printf("Error %s when preparing SQL statement", err)
//         return 0, err
//     }

	// database = controller.database
	// connectionString = controller.connectionString

// 	database, connectionError = sql.Open("mssql", controller.connectionString); if connectionError != nil {
// 		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
// 	}
//
// 	err = database.PingContext(dbContext); if err != nil {
// 		fmt.Printf("Error checking db connection: %v", err)
// 	}
//
// 	fmt.Println(err)
//
// 	fmt.Println("-> Welcome to Reminders Console App, built using Golang and Microsoft SQL Server")
// 	fmt.Println("-> Select a numeric option; \n [1] Create a new Reminder \n [2] Get a reminder \n [3] Delete a reminder")
// //
// 	consoleReader := bufio.NewScanner(os.Stdin)
// 	consoleReader.Scan()
// 	userChoice := consoleReader.Text()
//
// 	switch userChoice {
// 	case "1":
// 		var (
// 			titleInput,
// 			descriptionInput,
// 			aliasInput string
// 		)
// 		fmt.Println("You are about to create a new reminder. Please provide the following details:")
//
// 		fmt.Println("-> What is the title of your reminder?")
// 		consoleReader.Scan()
// 		titleInput = consoleReader.Text()
//
// 		fmt.Println("-> What is the description of your reminder?")
// 		consoleReader.Scan()
// 		descriptionInput = consoleReader.Text()
//
// 		fmt.Println("-> What is an alias of your reminder? [ An alias will be used to retrieve your reminder ]")
// 		consoleReader.Scan()
// 		aliasInput = consoleReader.Text()
//
// 		err := createReminder(titleInput, descriptionInput, aliasInput); if err != nil {
// 			return
// 		}
// 		break
//
// 	case "2":
// 		fmt.Println("-> Please provide an alias for your reminder:")
// 		consoleReader.Scan()
// 		aliasInput := consoleReader.Text()
//
// 		getErr, data := retrieveReminder(aliasInput); if getErr != nil {
// 			fmt.Println(getErr)
// 	  }
//
// 	  fmt.Println(data)
// 		break
//
// 	case "3":
// 		fmt.Println("-> Please provide the alias for the reminder you want to delete:")
//
// 		consoleReader.Scan()
// 		deleteAlias := consoleReader.Text()
//
// 		getErr := deleteReminder(deleteAlias); if getErr != nil {
// 			fmt.Println(getErr)
// 		}
// 		break
//
// 		default:
// 		 fmt.Printf("-> Option: %v is not a valid numeric option. Try 1 , 2 , 3", userChoice)
// 	}
