package rss_article

import (
	"fmt"
	"go-serverlessRss/articles"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

// GetRssArticle : retrieve an article by its index from any rss feed - CLI only
func GetRssArticle(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)
	rssURI := pathParams["rss"]
	index, err := strconv.Atoi(pathParams["item"])

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rssURI, err = url.PathUnescape(rssURI)

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

	article := listOfArticles[index]

	w.Write([]byte(fmt.Sprintf(`
	| %s |
	 
	%s
	
	Read more at : %s
	
	`, article.Title, article.Desc, article.Link)))

}
