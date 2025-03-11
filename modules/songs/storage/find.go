package songstorage

import songmodels "example.com/musicafy_be/modules/songs/models"

func (s *store) FindSong(id string) (*songmodels.Songs, error) {
	var songs songmodels.Songs

	results := s.db.Table(songmodels.Songs{}.TableName()).Where("mask_id = ?", id).First(&songs)
	if results.Error != nil {
		return nil, results.Error
	}

	return &songs, nil
}

func (s *store) FindAlbum(id string) (*songmodels.Albums, error) {
	var albums songmodels.Albums
	results := s.db.Table(songmodels.Albums{}.TableName()).Where("mask_id = ?", id).First(&albums)
	return &albums, results.Error
}
