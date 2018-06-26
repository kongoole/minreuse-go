package route

import (
	"net/http"
	"github.com/kongoole/minreuse-go/controller"
)

func Register() {
	http.HandleFunc("/", controller.Home{}.Index)
	http.HandleFunc("/blog", controller.Blog{}.Index)
	http.HandleFunc("/article", controller.Blog{}.Article)
	http.HandleFunc("/tag", controller.Blog{}.TagArticles)
	http.HandleFunc("/blog/search", controller.Blog{}.Search)
	http.HandleFunc("/public/", serveResource)
}

// serveResource returns static files
func serveResource(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}
