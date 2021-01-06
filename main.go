package main

import (
	"bufio"
	"cron-datablok/config"
	"cron-datablok/db"
	"cron-datablok/models"
	"log"
	"os"

	"github.com/jasonlvhit/gocron"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize main database
	db.Db = db.MariaDBInit()

	gocron.Every(1).Day().At("20:19").Do(DataBlok)

	<-gocron.Start()
}

func DataBlok() {
	log.Println("RUNNING CRON DATA BLOK")
	fileName := config.DataBlokPath
	log.Println("open file: ", fileName)
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	params := make(map[string]string)
	for _, eachline := range txtlines {
		params["name"] = eachline
		params["category_blok_id"] = config.BlokCategory
		log.Println("PARAMS: ", params)
		statusResponse, err := models.CreateBlok(params)
		if statusResponse != 200 {
			log.Println(err)
		}
	}

}
