package songstorage

import (
	"example.com/musicafy_be/common"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

func (s *store) CreateSong(data songmodels.Songs) (*songmodels.Songs, error) {
	results := s.db.Table(songmodels.Songs{}.TableName()).Create(&data)
	if results.Error != nil {
		return nil, common.ErrDB(results.Error)
	}
	return &data, nil
}
