package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/iDominate/golang-crm-basic/database"
	"github.com/iDominate/golang-crm-basic/lead"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}
func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	InitDB()
	setupRoutes(app)
	fmt.Println("Listening on port: 8000")
	log.Fatal(app.Listen(":8000"))
	defer database.DBConn.Close()
}

func InitDB() {
	db, err := gorm.Open("sqlite3", "leads.db")

	database.DBConn = db

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	fmt.Println("Connected Successfully to DB...")
	database.DBConn.AutoMigrate(&lead.Lead{})
}
