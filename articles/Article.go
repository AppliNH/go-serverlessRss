package articles

// Article : defines an article from a rss feed
type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Link  string `json:"link"`
}
