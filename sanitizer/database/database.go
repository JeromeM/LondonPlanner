package database

import (
	"database/sql"
	"errors"
	"io/fs"
	"os"

	"github.com/JeromeM/LondonPLanner/sanitizer/helper"
)

const databaseFile string = "database.db"

type Db struct {
	*sql.DB
}

func CreateDatabase() *sql.DB {
	helper.GInfoLn("Creating database in file %s", databaseFile)
	if _, err := os.Stat(databaseFile); !errors.Is(err, fs.ErrNotExist) {
		helper.GWarningLn("File %s already exists.", databaseFile)
	} else {
		file, err := os.Create(databaseFile)
		checkErr(err)
		file.Close()
	}

	database, err := sql.Open("sqlite3", databaseFile)
	checkErr(err)

	CreateTables(database)

	return database
}

func CreateTables(database *sql.DB) {
	/////// STATIONS
	drop := `DROP TABLE IF EXISTS stations;`
	_, err := database.Exec(drop)
	checkErr(err)
	helper.GInfoLn("Creating table stations...")
	stationTable := `CREATE TABLE stations (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Reference" TEXT,
		"Name" TEXT
		);`
	_, err = database.Exec(stationTable)
	checkErr(err)

	/////// LINES
	drop = `DROP TABLE IF EXISTS lines;`
	_, err = database.Exec(drop)
	checkErr(err)
	helper.GInfoLn("Creating table lines...")
	lineTable := `CREATE TABLE lines (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Reference" TEXT,
		"Name" TEXT
		);`
	_, err = database.Exec(lineTable)
	checkErr(err)
}

func AddStation(db *sql.DB, ref string, name string) {
	record := `INSERT INTO stations(Reference, Name) VALUES (?, ?)`
	query, err := db.Prepare(record)
	checkErr(err)

	_, err = query.Exec(ref, name)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		helper.GFatalLn(err.Error())
	}
}
