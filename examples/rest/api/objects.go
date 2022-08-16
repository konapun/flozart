package api

type CreateAccountRequest struct {
	Username string
	Email    string
	Password string
}

type CreateAccountResponse struct{}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

type CreatePostRequest struct {
	Title   string
	Content string
}

type CreatePostResponse struct {
	PostId string
	Url    string
}

type DeleteThreadRequest struct {
	Url string
}

type DeleteThreadResponse struct{}

type CreateCommentRequest struct {
	Parent  string
	Content string
}

type CreateCommentResponse struct {
	CommentId string
}

type EditCommentRequest struct {
	CommentId string
	Content   string
}

type EditCommentResponse struct{}

type DeleteCommentRequest struct {
	CommentId string
}

type DeleteCommentResponse struct{}
