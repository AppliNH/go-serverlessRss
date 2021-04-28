package main

import (
	"go-serverlessRss/routes/rss"
	rss_article "go-serverlessRss/routes/rss/article"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pathReq := r.RequestURI

		if strings.HasPrefix(pathReq, "/api/cli/rss/") {

			s1 := strings.TrimLeft(pathReq, "/api/cli/rss/")

			if strings.Contains(s1, "/item/") {

				i := strings.Index(s1, "/item/")
				rssURI := url.PathEscape(s1[:i])

				r.URL.Path = "/api/cli/rss/" + rssURI + s1[i:]

			} else {
				rssURI := url.PathEscape(s1)
				r.URL.Path = "/api/cli/rss/" + rssURI
			}

		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/cli").Subrouter()
	apiJSON := r.PathPrefix("/api/json").Subrouter()

	api.SkipClean(true)
	api.HandleFunc("/rss/{rss}", rss.GetRssArticles).Methods("GET")
	api.HandleFunc("/rss/{rss}/item/{item}", rss_article.GetRssArticle).Methods("GET")

	apiJSON.SkipClean(true)
	apiJSON.HandleFunc("/rss", rss.GetRssArticles).Methods("POST")

	log.Println("goRss is running on port " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), rewriter(r)))
}
