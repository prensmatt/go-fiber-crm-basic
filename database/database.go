package database
import(
	"github.com/jinzhu/gorm"
	_ "github.com/glebarez/go-sqlite"
)

var(
	DBConn *gorm.DB
)