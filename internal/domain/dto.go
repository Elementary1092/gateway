package domain

type ErrorDTO struct {
	Error string `json:"error"`
}

type StatusDTO struct {
	Status string `json:"status"`
}

type UpdatePostDTO struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SavePostDTO struct {
	UserId int32  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
