package resource

import "github.com/go-restful/app/model"

type PostResource struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
}

func NewPostResource(post *model.Post) *PostResource {
	return &PostResource{
		Title:    post.Title,
		Category: post.Category,
		Content:  post.Content,
		UserID:   post.UserID,
	}
}

func NewPostsResource(users *[]model.Post) []PostResource {
	var usersResponse []PostResource
	for _, user := range *users {
		usersResponse = append(usersResponse, *NewPostResource(&user))
	}

	return usersResponse
}
