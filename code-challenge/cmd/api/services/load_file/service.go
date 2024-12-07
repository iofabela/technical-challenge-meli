package load_file

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type service struct {
	repo *Repository
}

type Service interface {
	LoadFile(ctx *gin.Context, fileName *multipart.FileHeader) (*multipart.FileHeader, error)
}

func NewService(repo *Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) LoadFile(ctx *gin.Context, file *multipart.FileHeader) (*multipart.FileHeader, error) {
	return s.repo.LoadFile(ctx, file)
}
