package articles

type Articles struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Llink string `json:"link"`
}
