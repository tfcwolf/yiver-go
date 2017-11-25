package front 
import (
	"net/http"
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func Init() {
	Db,err = sql.Open("mysql", "root:root@/dyb?charset=utf8");
}

type Controller interface {
	Get(w http.ResponseWriter, r *http.Request)
}
type BaseController  struct {
	C interface{}
	Db *sql.DB
}


func (base *BaseController) Get(w http.ResponseWriter, r *http.Request) {

}

func (base *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("serveing .....");\\\
	\\\\
	base.C.(Controller).Get(w,r);
}



