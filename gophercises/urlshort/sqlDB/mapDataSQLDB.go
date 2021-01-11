// The following mysql commands must be run in the terminal before running this program.
// Start mysql Server
//		$ mysql.server start
// Create the DB
// Switch to the DB
// 		mysql>use mapDB;
// Create the Tables
//		mysql> CREATE TABLE URLMap ( url_original varchar(100) NOT NULL, url_redirect varchar(100) NOT NULL );
// Drop the Tables as needed
// 		mysql> DROP TABLE URLMap;
// Stop mysql Server
//		$ mysql.server stop

package urlshort

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//need to initilize the driver here.
	_ "github.com/go-sql-driver/mysql"
)

//Database Variables
var dbHost = "localhost"
var dbDriver = "mysql"
var dbTable = "URLMap"
var dbName = "mapDB"
var dbPort = "3306"

// MYSQL CREDS
var dbUser string = "root"
var dbPass string = ""

var debugLog *log.Logger

type urlPathMap struct {
	URLOriginal string `yaml:"url_original" json:"url_original"`
	URLRedirect string `yaml:"url_redirect" json:"url_redirect"`
}

func init() {
	log.Printf("DEBUG: mapDataSQLDB()): Init()\n\n")

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog.Printf("Debug logger is set\n")

	//Setting up the database
	//makeMySQLDatabase(dbUser, dbPass, dbName)

	//Populating the database
	populateMapData()
}

func makeMySQLDatabase(username string, password string, dbName string) {
	debugLog.Println("makeMySQLDatabase(): Making new MySQL Database with name '" + dbName + "' on server: " + dbHost)
	mysqlString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/"
	db, err := sql.Open(dbDriver, mysqlString)
	if err != nil {
		debugLog.Printf("makeMySQLDatabase(): Unable to open DB: %v\n", mysqlString)
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		debugLog.Println("makeMySQLDatabase(): Unable to create database: '" + dbName + "' on server: " + dbHost)
		panic(err)
	}

	// _, err = db.Exec("GRANT ALL PRIVILEGES ON " + dbName + ".* To '" + username + "'@'%' IDENTIFIED BY '" + password + "'")
	// if err != nil {
	// 	panic(err)
	// }
}

// populateMapData Populates the DB with the URL Mapping data.
// "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
func populateMapData() {

	var err error
	var db *sql.DB

	dbFile := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	debugLog.Printf("populateMapData(): Start Populating the Database %v", dbFile)
	db, err = sql.Open(dbDriver, dbFile)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	debugLog.Printf("populateMapData():  %v DB file is open\n", dbFile)

	// Add the following map to the DB: "/urlshort-godoc", "https://godoc.org/github.com/gophercises/urlshort"
	addURLMap(db, "/urlshort-godoc", "https://godoc.org/github.com/gophercises/urlshort")

	// Add the follwong map to the DB: "/yaml-godoc", "https://godoc.org/gopkg.in/yaml.v2"
	addURLMap(db, "/yaml-godoc", "https://godoc.org/gopkg.in/yaml.v2")

	// For Debug Purpose Only
	PrintMapData()

}

// PrintMapData prints the map data in the DB file.
func PrintMapData() {

	debugLog.Println("PrintMapData(): Start Printing the DB")
	dbFile := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open(dbDriver, dbFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	debugLog.Printf("PrintMapData(): %v DB file is open\n", dbFile)

	res, err := db.Query("SELECT * FROM " + dbTable)
	defer res.Close()
	if err != nil {
		log.Fatalf("PrintMapData(): Error selecting all the rows from %v DB file - %v DB table: %v\n", dbFile, dbTable, err)
	}

	for res.Next() {
		var urlMap urlPathMap
		err := res.Scan(&urlMap.URLOriginal, &urlMap.URLRedirect)
		if err != nil {
			log.Fatalf("PrintMapData(): Error Iterating over %v DB file: %v\n", dbFile, err)
		}
		fmt.Printf("%+v\n", urlMap)
	}
	debugLog.Println("PrintMapData(): Finish Printing the DB")
}

// GetURLRedirect returns the mapped value URL we want to redirect the key URL to
func GetURLRedirect(key string) (string, bool) {

	debugLog.Printf("GetURLRedirect(): Start getting the Mapped URL for %v\n", key)
	var err error
	var db *sql.DB
	var sqlQuery string
	var value string
	valueFound := false

	dbFile := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err = sql.Open(dbDriver, dbFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	debugLog.Printf("GetURLRedirect(): %v DB file is open\n", dbFile)

	var urlMap urlPathMap

	//Query Format:   SELECT * from URLMap WHERE url_original = '/urlshort-godoc';
	sqlQuery = "SELECT * from " + dbTable + " WHERE url_original = '" + key + "'"
	debugLog.Printf("GetURLRedirect(): Query(%v)\n", sqlQuery)
	tmpRow := db.QueryRow(sqlQuery)
	err = tmpRow.Scan(&urlMap.URLOriginal, &urlMap.URLRedirect)

	if err != nil {
		switch err {
		case sql.ErrNoRows: // The key is not in the database
			debugLog.Printf("GetURLRedirect(): key: %v, Not found in the DB\n", key)
			return value, valueFound
		default:
			log.Fatalf("ERROR: GetURLRedirect(): key: %v resulted in Error: %v\n", key, err)
		}
	}
	value = urlMap.URLRedirect
	valueFound = true
	debugLog.Printf("GetURLRedirect(): %v: %v\n", key, value)

	debugLog.Printf("GetURLRedirect(): Finish getting the Mapped URL for %v:%v\n", key, value)
	return value, valueFound
}

func addURLMap(db *sql.DB, urlOriginal string, urlRedirect string) {

	var lastID int64
	var err error
	var sqlQuery string
	var res sql.Result

	sqlQuery = "INSERT INTO " + dbTable + "(url_original, url_redirect)" + " VALUES " + "('" + urlOriginal + "', '" + urlRedirect + "')"
	debugLog.Printf("populateMapData(): %v\n", sqlQuery)
	res, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	lastID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("populateMapData(): The last inserted row id: %d\n", lastID)

}
