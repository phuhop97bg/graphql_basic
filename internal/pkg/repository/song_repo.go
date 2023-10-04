package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"
	"fmt"
)

type songRepository struct {
	DB db.Database
}

func (u *songRepository) FirstSong(condition entity.Song) (entity.Song, error) {
	song := entity.Song{}
	err := u.DB.First(&song, condition)
	return song, err
}
func (u *songRepository) FindSongList(condition entity.Song) (song []entity.Song, err error) {
	err = u.DB.Find(&song, condition)
	return
}
func (u *songRepository) FindSongListWithPagination(condition entity.Song, pageNum int, pageSize int) (songs []entity.Song, err error) {
	offset := (pageNum - 1) * pageSize
	userClause := ""
	if condition.UserID > 0 {
		userClause = fmt.Sprintf(" AND user_id = %d \n", condition.UserID)
	}
	sql := "SELECT * FROM songs where  deleted_at IS NULL\n" +
		"AND title LIKE '%" + condition.Title + "%' AND \n" +
		" singer LIKE '%" + condition.Singer + "%'\n" +
		userClause +
		" ORDER BY view DESC\n" +
		" LIMIT " + fmt.Sprintf("%d,%d", offset, pageSize) + " ;"
	err = u.DB.RawQuery(&songs, sql)
	return
}
func (u *songRepository) CreateSong(song entity.Song) (entity.Song, error) {
	err := u.DB.Create(&song)
	return song, err
}
func (u *songRepository) DeleteSong(song entity.Song) error {
	err := u.DB.Delete(&song)
	return err
}
func (u *songRepository) UpdateSong(song, oldsong entity.Song) (entity.Song, error) {
	err := u.DB.Update(entity.Song{}, &oldsong, &song)
	return song, err
}

func NewSongRepository(db db.Database) *songRepository {
	return &songRepository{
		DB: db,
	}
}
