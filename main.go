package main

import (
	"bufio"
	"cron-datablock/config"
	"cron-datablock/db"
	"cron-datablock/models"
	"log"
	"os"

	"github.com/jasonlvhit/gocron"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize main database
	db.Db = db.MariaDBInit()

	gocron.Every(1).Day().At("20:45").Do(DataBlock)

	<-gocron.Start()
}

func DataBlock() {
	log.Println("RUNNING CRON DATA BLOCK")
	fileName := config.DataBlockPath
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
		params["domain"] = eachline
		params["category_block_id"] = config.BlockCategory
		statusResponse, err := models.CreateBlock(params)
		if statusResponse != 200 {
			log.Println(err)
		}
	}

}
