package main
import (
	"app"
	"front"
	//"log"
)


func main() {
	app.Init();
    app.Handle("/list", &front.IndexHandler{})
    app.Handle("/test",&front.TextHandler{})
    
    app.Run()

}
