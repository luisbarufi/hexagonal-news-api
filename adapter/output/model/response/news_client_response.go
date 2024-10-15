package response

type NewsClientResponse struct {
	Status       string
	TotalResults string
	Articles     []ArticleResponse
}

type ArticleResponse struct {
	Source      ArticleSourceResponse
	Id          string
	Name        string
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponse struct {
	ID   *string
	Name string
}
