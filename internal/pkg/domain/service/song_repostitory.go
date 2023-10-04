package service

import "backend-food/internal/pkg/domain/domain_model/entity"

type SongRepository interface {
	FirstSong(entity.Song) (entity.Song, error)
	FindSongList(entity.Song) (user []entity.Song, err error)
	CreateSong(Song entity.Song) (entity.Song, error)
	DeleteSong(Song entity.Song) error
	UpdateSong(newSong, oldSong entity.Song) (entity.Song, error)
	FindSongListWithPagination(condition entity.Song, pageNum int, pageSize int) ([]entity.Song, error)
}
