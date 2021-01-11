package urlshort

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

var dbFile = "mapDB"
var debugLog *log.Logger

func init() {
	//log.Printf("DEBUG: mapDataLevelDB()): Init()\n\n")

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	//debugLog.Printf("Debug logger is set\n")

	err := populateMapData()
	if err != nil {
		log.Fatalf("Init(): Error Populating %v DB file: %v\n", dbFile, err)
	}
}

// populateMapData Populates the DB with the URL Mapping data.
// "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
func populateMapData() error {

	//debugLog.Println("populateMapData(): Start Populating the Database")
	var err error
	var db *leveldb.DB

	db, err = leveldb.OpenFile(dbFile, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	batch := new(leveldb.Batch)
	batch.Put([]byte("/urlshort-godoc"), []byte("https://godoc.org/github.com/gophercises/urlshort"))
	batch.Put([]byte("/yaml-godoc"), []byte("https://godoc.org/gopkg.in/yaml.v2"))
	err = db.Write(batch, nil)

	//time.Sleep(5 * time.Second)
	//debugLog.Println("populateMapData(): Finish Poulating the Database")
	return err
}

// PrintMapData prints the map data in the DB file.
func PrintMapData() {

	//debugLog.Println("PrintMapData(): Start Printing the DB")
	var err error
	var db *leveldb.DB

	db, err = leveldb.OpenFile(dbFile, nil)
	if err != nil {
		log.Fatalf("PrintMapData(): Error opening %v DB file: %v\n", dbFile, err)
	}
	defer db.Close()

	//debugLog.Printf("PrintMapData(): %v DB file is open\n", dbFile)
	fmt.Printf("%v File Contents: \n", dbFile)

	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%s: %s\n", key, value)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Fatalf("PrintMapData(): Error Iterating over %v DB file: %v\n", dbFile, err)
	}
	fmt.Println("")

	//debugLog.Println("PrintMapData(): Finish Printing the DB")
}

// GetURLRedirect returns the mapped value URL we want to redirect the key URL to
func GetURLRedirect(key string) (string, bool) {

	//debugLog.Printf("GetURLRedirect(): Start getting the Mapped URL for %v\n", key)
	var err error
	var db *leveldb.DB
	var value string
	valueFound := false

	db, err = leveldb.OpenFile(dbFile, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	var data []byte
	data, err = db.Get([]byte(key), nil)
	if err != nil {
		switch err {
		case leveldb.ErrNotFound: // The key is not in the database
			//debugLog.Printf("GetURLRedirect(): key: %v, Not found in the DB\n", key)
			return value, valueFound
		default:
			log.Fatalf(err.Error())
		}
	}
	value = string(data)
	valueFound = true
	//debugLog.Printf("GetURLRedirect(): %v: %v\n", key, value)

	//debugLog.Printf("GetURLRedirect(): Finish getting the Mapped URL for %v\n", key)
	return value, valueFound
}
