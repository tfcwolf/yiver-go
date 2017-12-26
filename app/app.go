package app
import (
	"net/http"
	"front"
	"log"
	"reflect"
	//"log"
)
var err error 
var mux *http.ServeMux;

func Init() {
	
	front.Init()
	log.Println("init.....");
	if err != nil {
		log.Println(err)
	}
    mux = http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir("public")))   
}

func NewHandler(c interface{}) *front.BaseController {
	base := &front.BaseController{}
	base.C = c
	elm := reflect.ValueOf(c).Elem().Field(0).Field(2).Interface();
	base.Routes = elm.(map[string]func(http.ResponseWriter,*http.Request))
	return    base
}

func Handle(pattern string,controller front.Controller) {
	log.Println("handle.....");
	mux.Handle(pattern,NewHandler(controller));
}

func Run() {
	
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
	    log.Fatal("ListenAndServe: ", err)
	}
}
