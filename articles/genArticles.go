package articles

import "go-serverlessRss/services"

// GenArticlesFromRssFeed : Generate list of articles from any rssFeed
func GenArticlesFromRssFeed(rssURI string) ([]Article, error) {

	items, err := services.RetrieveItemsFromRss(rssURI)

	if err != nil {
		return nil, err
	}

	var listOfArticles []Article
	for i, v := range items {
		article := Article{ID: i, Title: v.Title, Desc: v.Description, Link: v.Link}

		listOfArticles = append(listOfArticles, article)
	}

	return listOfArticles, nil
}
