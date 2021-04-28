package rss

import (
	"encoding/json"
	"fmt"
	"go-serverlessRss/articles"
	"log"
	"net/http"
	"net/url"
	"strings"

	. "go-serverlessRss/newsfeed"

	"github.com/gorilla/mux"
)

// GetRssArticles : Retrieve and send back articles in JSON or CLI format
func GetRssArticles(w http.ResponseWriter, r *http.Request) {

	var rssURI string
	var err error
	var mode string
	log.Println(r.URL.Path)
	if strings.Contains(r.URL.Path, "/json/") {
		mode = "json"
		w.Header().Set("Content-Type", "application/json")

		var newsFeed NewsFeed
		err = json.NewDecoder(r.Body).Decode(&newsFeed)
		rssURI = newsFeed.Uri

	} else if strings.Contains(r.URL.Path, "/cli/") {
		mode = "cli"
		pathParams := mux.Vars(r)
		rssURI = pathParams["rss"]

		rssURI, err = url.PathUnescape(rssURI)

	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listOfArticles, err := articles.GenArticlesFromRssFeed(rssURI)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mode == "json" {
		json.NewEncoder(w).Encode(listOfArticles)
	} else if mode == "cli" {

		for i, v := range listOfArticles {

			w.Write([]byte(fmt.Sprintf(`
			| %d --> %s |
			
			`, i, v.Title)))
		}

		w.Write([]byte(fmt.Sprintf(`
	
			Query the same uri + /item/[articleNumber] to read an article
	
			`)))
	}

}
