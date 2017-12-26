package main
import (
	"app"
	"front"
	//"log"
)


func main() {
	app.Init();
	index := front.NewIndexHandler()
    app.Handle("/index", index)
    app.Handle("/index/",index )
    //app.Handle("/test",&front.TextHandler{})
    
    app.Run()

}
