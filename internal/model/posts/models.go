package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashtags []string `json:"post_hashtags"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserId       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashtags string    `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID           int64    `json:"id"`
		UserId       int64    `json:"user_id"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashtags []string `json:"post_hashtags"`
		IsLiked      bool     `json:"is_liked"`
	}

	Pagination struct {
		Limit  int64 `json:"limit"`
		Offset int64 `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail Post      `json:"post_detail"`
		LikeCount  int       `json:"like_count"`
		Comments   []Comment `json:"comment"`
	}

	Comment struct {
		Id             int64  `json:"id"`
		UserId         int64  `json:"user_id"`
		Username       string `json:"username"`
		CommentContent string `json:"comment_content"`
	}
)
