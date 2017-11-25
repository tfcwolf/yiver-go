package front
import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
)



type TextHandler struct {
	BaseController
}

func (th *TextHandler) Get(w http.ResponseWriter, r *http.Request){
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
        rows.Scan(&post.Id,&post.Post_content)

        posts =append(posts,post)
    }


     // b, err := json.Marshal(posts)
     // if err != nil {
     //    log.Println(err)
     // }
     json.NewEncoder(w).Encode(posts);
}
