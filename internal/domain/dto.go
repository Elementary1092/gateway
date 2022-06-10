package domain

import services "github.com/elem1092/gateway/pkg/client/grpc/CRUDService"

type GetPostInfo struct {
	Id     int32  `json:"id"`
	IdType string `json:"id_type"`
}

type DeleteDTO struct {
	Id int32 `json:"id"`
}

type ErrorDTO struct {
	Error string `json:"error"`
	Code  int    `json:"error_code"`
}

type StatusDTO struct {
	Status string `json:"status"`
}

type UpdatePostDTO struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (dto *UpdatePostDTO) formGRPCRequest() *services.UpdatePostDTO {
	return &services.UpdatePostDTO{
		Id: dto.Id,
		Content: &services.ContentDTO{
			Title: dto.Title,
			Body:  dto.Body,
		},
	}
}

type SavePostDTO struct {
	UserId int32  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (dto *SavePostDTO) formGRPCRequest() *services.SavePostDTO {
	return &services.SavePostDTO{
		UserId: dto.UserId,
		Content: &services.ContentDTO{
			Title: dto.Title,
			Body:  dto.Body,
		},
	}
}
