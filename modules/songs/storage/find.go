package songstorage

import songmodels "example.com/musicafy_be/modules/songs/models"

func (s *store) FindSong(id string) (songmodels.Songs, error) {
	var songs songmodels.Songs

	results := s.db.Table(songmodels.Songs{}.TableName()).Where("mask_id = ?", id).First(&songs)
	if results.Error != nil {
		return songs, results.Error
	}

	return songs, nil
}
