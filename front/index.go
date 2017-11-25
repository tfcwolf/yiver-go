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

func (ih *IndexHanlder) Create(w http.ResponseWriter,r *http.Request) {

}


