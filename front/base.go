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
	Routes map[string] func(http.ResponseWriter,*http.Request)
}


func (base *BaseController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("dd")
}

func (base *BaseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println("serveing .....");
	uri := r.URL.Path;
	if action,ok:=base.Routes[uri];ok {
		action(w,r)
	} else {
		w.Write([]byte("Not Found!"));
		//base.C.(Controller).Get(w,r);
	}
}



