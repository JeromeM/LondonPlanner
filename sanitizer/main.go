package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"encoding/xml"

	"github.com/JeromeM/LondonPLanner/sanitizer/database"
	"github.com/JeromeM/LondonPLanner/sanitizer/helper"
	"github.com/JeromeM/LondonPLanner/sanitizer/types"

	_ "github.com/mattn/go-sqlite3"
	"github.com/schollz/progressbar/v3"
)

const dataDir string = "data"
const zipFile string = "journey-planner-timetables.zip"

var (
	db *sql.DB
)

type XMLFile struct {
	File fs.FileInfo
}

func main() {
	// Create SQLite database
	db = database.CreateDatabase()

	// Unzip File
	_, err := os.Stat(filepath.Join(dataDir, zipFile))
	if os.IsNotExist(err) {
		helper.GFatalLn("Error! Zipfile doesn't exists.")
	}
	helper.UnzipSource(filepath.Join(dataDir, zipFile), dataDir)

	// Searching for files concerning Underground Train
	helper.GInfoLn("Searching for Underground Train files and updating database ...")
	var file XMLFile
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		helper.GFatalLn("Cant't read dir %s", dataDir)
	}
	bar := progressbar.NewOptions(len(files),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionOnCompletion(func() {
			helper.GBlank()
			helper.GInfoLn("Finished")
		}),
	)
	for _, file.File = range files {
		file.undergroundTrain()
		bar.Add(1)
	}

	// Remove XML files when finished
	xmlFiles, err := filepath.Glob(filepath.Join(dataDir, "*.xml"))
	if err != nil {
		helper.GFatalLn("Can't find XML files")
	}
	for _, f := range xmlFiles {
		if err := os.Remove(f); err != nil {
			helper.GFatalLn("Can't remove file %s : %s", f, err.Error())
		}
	}

	helper.GBlank()
}

func (xmlFile *XMLFile) undergroundTrain() {

	var (
		trans        types.TransXChange
		file         fs.FileInfo = xmlFile.File
		fullFilename string      = filepath.Join(dataDir, file.Name())
	)

	if IsExist("<VehicleTypeCode>UT</VehicleTypeCode>", fullFilename) {

		// We will read the XML and keep necessary informations for the search
		xmlFile, err := os.Open(fullFilename)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(xmlFile)
		xml.Unmarshal(byteValue, &trans)

		lineName := trans.Services.Service[0].Line_name
		for i := 0; i < len(trans.StopPoints.StopPoint); i++ {
			database.AddStation(db, trans.StopPoints.StopPoint[i].Ref, trans.StopPoints.StopPoint[i].Name, lineName)
		}

		defer xmlFile.Close()
	}
}

func IsExist(str, filepath string) bool {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		helper.GFatalLn("Can't read file %s : ", err.Error())
	}

	isExist, err := regexp.Match(str, b)
	if err != nil {
		helper.GFatalLn("Can't search if file %s exists : ", err.Error())
	}
	return isExist
}
