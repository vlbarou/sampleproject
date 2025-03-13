package main

import (
	"github.com/vlbarou/sampleproject/config"
	"github.com/vlbarou/sampleproject/internal/database"
	"log"

	"github.com/vlbarou/sampleproject/internal/server"
)

// docker run --name mariadb-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 -d mariadb:latest
// barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o main cmd/api/main.go
// barousis@barousis-HP-ZBook-Fury-16-G9-Mobile-Workstation-PC:~/Documents/projects/NCMT/trainings/my_context/myrepos/sampleproject$ go build -o migrate cmd/migrate/migrate.go

// curl -X POST http://localhost:8080/users -H "Content-Type: application/json"      -d @data.json
// curl http://localhost:8080/users
func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	defer database.Close(db)

	srv := server.NewServer(cfg, db)
	log.Fatal(srv.Start())
}
