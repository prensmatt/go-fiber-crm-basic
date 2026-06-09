package main
import(
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/leads",GetLeads)
	app.Get("/leads/:id",GetLeads)
	app.Post("/leads",NewLead)
	app.Delete("/leads/:id",DeleteLead)
}

func main(){
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}