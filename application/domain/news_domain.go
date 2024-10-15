package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults string
	Articles     []Article
}

type Article struct {
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
