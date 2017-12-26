package front
import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
)

type IndexHandler struct {
    BaseController
}

type Posts struct {
    Id int
    Post_content string 
    Post_title string
}

type Navs struct {
    Id int
    Cid int 
    Label string
    Href string
}

func NewIndexHandler() *IndexHandler{
    obj := &IndexHandler{}
    obj.init()
    return obj
}

func (ih *IndexHandler) init() {

    ih.Routes = make(map[string]func(http.ResponseWriter,*http.Request))
    ih.Routes["/index/show"] = ih.Show   
}

func (ih *IndexHandler) Show(w http.ResponseWriter, r *http.Request) {
    rows,err:= Db.Query("select id,cid,href,label from dyb_nav limit 5")
    if err != nil {
        log.Println(err)
        return 
    }

    defer rows.Close()
    var navs [] Navs;
    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    for rows.Next() {
        nav := Navs{}
        rows.Scan(&nav.Id,&nav.Cid,&nav.Label,&nav.Href)

        navs = append(navs,nav)
    }

     json.NewEncoder(w).Encode(navs);
}

func (ih *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
    
    rows,err:= Db.Query("select id,post_title from dyb_posts limit 5")
    if err != nil {
        log.Println(err)
        return 
    }

    defer rows.Close()
    var posts [] Posts;
    fmt.Println("");
    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    for rows.Next() {
        post := Posts{}
        rows.Scan(&post.Id,&post.Post_title)

        posts = append(posts,post)
    }
     // b, err := json.Marshal(posts)
     // if err != nil {
     //    log.Println(err)
     // }
     json.NewEncoder(w).Encode(posts);
     
}

func (ih *IndexHandler) Create(w http.ResponseWriter,r *http.Request) {

}


