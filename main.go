package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pathReq := r.RequestURI

		if strings.Contains(pathReq, `/item/`) {
			if strings.HasPrefix(pathReq, "/api/v1/rss/") {

				s1 := strings.TrimLeft(pathReq, "/api/v1/rss/")
				i := strings.Index(s1, "/item/")
				uri := s1[:i]
				s2 := s1[i:]
				fmt.Println("s1 : ", s1)
				fmt.Println("s2 : ", s2)
				fmt.Println("uri : ", uri)

				pe := url.PathEscape(uri)

				r.URL.Path = "/api/v1/rss/" + pe + s2

				r.URL.RawQuery = ""
			}
		}
		if !strings.Contains(pathReq, `/item/`) {
			if strings.HasPrefix(pathReq, "/api/v1/rss/") && !strings.HasSuffix(pathReq, "/item/") {

				pe := url.PathEscape(strings.TrimLeft(pathReq, "/api/v1/rss/"))
				r.URL.Path = "/api/v1/rss/" + pe
				r.URL.RawQuery = ""
			}
		}

		h.ServeHTTP(w, r)
	})
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func rssArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rssArticles")

	pathParams := mux.Vars(r)
	rssUri := pathParams["rss"]

	rss, err := url.PathUnescape(rssUri)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	fmt.Println(err)

	for i := 0; i < len(feed.Items); i++ {
		articleTitle := feed.Items[i].Title
		w.Write([]byte(fmt.Sprintf(`
		| %d --> %s |
		
		`, i, articleTitle)))

	}
	w.Write([]byte(fmt.Sprintf(`

		Please query the same uri + /item/[articleNumber] to read an article

		`)))
}

func rssArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rssArticle")

	pathParams := mux.Vars(r)
	rssUri := pathParams["rss"]
	item, erro := strconv.Atoi(pathParams["item"])

	fmt.Println(erro)
	rss, err := url.PathUnescape(rssUri)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	fmt.Println(err)

	article := feed.Items[item]

	w.Write([]byte(fmt.Sprintf(`
	| %s |
	 
	%s
	
	Read more at : %s
	
	`, article.Title, article.Description, article.Link)))

}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.SkipClean(true)
	api.HandleFunc("/rss/{rss}", rssArticles)
	api.HandleFunc("/rss/{rss}/item/{item}", rssArticle)
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", Rewriter(r)))
}
