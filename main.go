package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

type Articles struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Llink string `json:"link"`
}

func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pathReq := r.RequestURI

		if strings.Contains(pathReq, `/item/`) {
			if strings.HasPrefix(pathReq, "/api/v1/rss/") || strings.HasPrefix(pathReq, "/api/json/rss/") {
				var s1 string
				var prefix string

				if strings.HasPrefix(pathReq, "/api/v1/rss/") {
					s1 = strings.TrimLeft(pathReq, "/api/v1/rss/")
					prefix = "/api/v1/rss/"

				}
				if strings.HasPrefix(pathReq, "/api/json/rss/") {
					s1 = strings.TrimLeft(pathReq, "/api/json/rss/")
					prefix = "/api/json/rss/"

				}

				i := strings.Index(s1, "/item/")
				uri := s1[:i]
				s2 := s1[i:]

				pe := url.PathEscape(uri)

				r.URL.Path = prefix + pe + s2

				r.URL.RawQuery = ""
			}
		}
		if !strings.Contains(pathReq, `/item/`) {
			if (strings.HasPrefix(pathReq, "/api/v1/rss/") || strings.HasPrefix(pathReq, "/api/json/rss/")) && !strings.HasSuffix(pathReq, "/item/") {

				var pe string
				var prefix string

				if strings.HasPrefix(pathReq, "/api/v1/rss/") {
					pe = url.PathEscape(strings.TrimLeft(pathReq, "/api/v1/rss/"))
					prefix = "/api/v1/rss/"

				}
				if strings.HasPrefix(pathReq, "/api/json/rss/") {
					pe = url.PathEscape(strings.TrimLeft(pathReq, "/api/json/rss/"))
					prefix = "/api/json/rss/"

				}

				r.URL.Path = prefix + pe
				r.URL.RawQuery = ""
			}
		}

		h.ServeHTTP(w, r)
	})
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

func rssArticlesJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rssArticlesJSON")

	pathParams := mux.Vars(r)
	rssUri := pathParams["rss"]

	rss, err := url.PathUnescape(rssUri)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	fmt.Println(err)

	w.Header().Set("Content-Type", "application/json")
	var listOfArticles []Articles

	for i := 0; i < len(feed.Items); i++ {

		article := Articles{i, feed.Items[i].Title, feed.Items[i].Description, feed.Items[i].Link}

		listOfArticles = append(listOfArticles, article)
	}

	json.NewEncoder(w).Encode(listOfArticles)
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	apiJSON := r.PathPrefix("/api/json").Subrouter()

	api.SkipClean(true)
	api.HandleFunc("/rss/{rss}", rssArticles)
	api.HandleFunc("/rss/{rss}/item/{item}", rssArticle)

	apiJSON.SkipClean(true)
	apiJSON.HandleFunc("/rss/{rss}", rssArticlesJSON)

	log.Fatal(http.ListenAndServe(":8080", Rewriter(r)))
}
