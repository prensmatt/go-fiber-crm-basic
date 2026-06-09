package main
import(
	"github.com/prensmatt/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func setupRoutes(app *fiber.App){
	app.Get("/leads",GetLeads)
	app.Get("/leads/:id",GetLeads)
	app.Post("/leads",NewLead)
	app.Delete("/leads/:id",DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3","leads.db")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}