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
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).SendString(err.Error())
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx){

}