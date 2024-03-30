package main

type User struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Bookmarks []string `json:"bookmarks"`
	//CreatedAt Time     `json:"createdAt"`
}

type Bookmarks struct {
	UserID    string `json:"userId"`
	ArticleID string `json:"articleId"`
	//CreatedAt  Time   `json:"createdAt"`
}

type Article struct {
	URL        string `json:"url"`
	ThumbsUp   int    `json:"thumbsUp"`
	ThumbsDown int    `json:"thumbsDown"`
	Love       int    `json:"love"`
	//CreatedAt  Time   `json:"createdAt"`
}

type UserReaction struct {
	UserID     string `json:"userId"`
	ArticleID  string `json:"articleId"`
	ThumbsUp   bool   `json:"thumbsUp"`
	ThumbsDown bool   `json:"thumbsDown"`
	Love       bool   `json:"love"`
	//CreatedAt  Time   `json:"createdAt"`
}

func NewUser(id string) *User {
	return &User{
		ID: id,
	}
}

func NewArticle(url string) *Article {
	return &Article{
		URL: url,
	}
}

func NewUserReaction(userId, articleUrl string) *UserReaction {
	return &UserReaction{
		UserID:    userId,
		ArticleID: articleUrl,
		// Actual reaction
	}
}
