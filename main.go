package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ShawnRong/bento/models"

	"github.com/ShawnRong/bento/db"

	"github.com/ShawnRong/bento/config"
	"github.com/ShawnRong/bento/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	models.AutoMigrate()
	server.Init()
}
