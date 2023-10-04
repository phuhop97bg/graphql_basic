package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"
	"fmt"
)

type commentRepository struct {
	DB db.Database
}

func (u *commentRepository) FindCommentList(condition entity.Comment) (comments []entity.Comment, err error) {
	err = u.DB.Find(&comments, condition)
	return
}
func (u *commentRepository) FindCommentListWithPagination(condition entity.Comment, pageNum int, pageSize int) (comments []entity.Comment, err error) {
	offset := (pageNum - 1) * pageSize
	sql := "SELECT * FROM comment \n" +
		" WHERE song_id =" + fmt.Sprint(condition.SongID) + "\n" +
		" ORDER BY created_at DESC\n" +
		" LIMIT " + fmt.Sprintf("%d,%d", offset, pageSize) + " ;"

	err = u.DB.RawQuery(&comments, sql)
	return
}
func (u *commentRepository) CreateComment(comment entity.Comment) (entity.Comment, error) {
	err := u.DB.Create(&comment)
	return comment, err
}
func NewCommentRepository(db db.Database) *commentRepository {
	return &commentRepository{
		DB: db,
	}
}
