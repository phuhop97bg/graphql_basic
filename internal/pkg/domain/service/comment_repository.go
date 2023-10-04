package service

import "backend-food/internal/pkg/domain/domain_model/entity"

type CommentRepository interface {
	FindCommentList(entity.Comment) ([]entity.Comment, error)
	CreateComment(entity.Comment) (entity.Comment, error)
	FindCommentListWithPagination(condition entity.Comment, pageNum int, pageSize int) ([]entity.Comment, error)
}
