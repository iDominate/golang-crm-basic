package lead

import (
	"github.com/gofiber/fiber"
	"github.com/iDominate/golang-crm-basic/database"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id", "1")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(&lead); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		return
	}
	db.Create(lead)
	c.Status(fiber.StatusCreated).JSON(lead)

}

func DeleteLead(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		c.Status(fiber.StatusNotFound).JSON("No lead found")
		return
	}
	db.Delete(id)
	c.Status(fiber.StatusAccepted).JSON("Success")
}
