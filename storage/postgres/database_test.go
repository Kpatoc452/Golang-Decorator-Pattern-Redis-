package postgres

import (
	"fmt"
	"testing"

	"log"

	"github.com/Kpatoc452/Dceorator/config"
)


var descriptions = []string{
	"hello",
	"world",
	"text",
	"leetcode",
	"math",
}

func TestAddInfo(t *testing.T){
	cfg ,err := config.New()
	if err != nil{
		log.Fatal(err)
	}
	db, err := New(cfg)
	if err != nil{
		log.Fatal(err)
	}
	for _, line := range descriptions{
		err = db.AddInfo(line)
		if err != nil{
			log.Fatal(err)
		}
	}

	items, err := db.GetAllInfo()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items{
		fmt.Println(item)
	}
}