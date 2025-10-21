package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bjhermid/go-api-1/cmd/api"
	"github.com/bjhermid/go-api-1/config"
	"github.com/bjhermid/go-api-1/db"
	"github.com/go-sql-driver/mysql"
)

func main(){
	db,err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd : config.Envs.DBPassword,
		Addr: config.Envs.DBAddres,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}) 
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	
	addr := fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port)
	server := api.NewAPIServer(addr,db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected Succefully")
}