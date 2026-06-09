package lead
import(
	"github.com/prensmatt/go-fiber-crm-basic/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber/v2"
)

type Lead struct{
	gorm.Model
	Name 			string
	Company 	string
	Email 		string
	Phone 		int
}

func GetLeads(c *fiber.Ctx){

}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){

}

func DeleteLead(c *fiber.Ctx){

}