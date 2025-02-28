package songsbiz

import (
	"example.com/musicafy_be/components/appctx"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

type SearchSuggestionReq struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
}

type searchSuggestionStore interface {
	FindSong(id string) (songmodels.Songs, error)
	CreateSong(data songmodels.Songs) (*songmodels.Songs, error)
}

type searchSuggestionBiz struct {
	store searchSuggestionStore
}

func NewSearchSuggestionBiz(store searchSuggestionStore) searchSuggestionBiz {
	return searchSuggestionBiz{
		store: store,
	}
}

func (biz *searchSuggestionBiz) SearchSuggestion(appctx appctx.AppContext, params SearchSuggestionReq) ([]songmodels.Songs, error) {
	zingApi := appctx.GetZingmp3Api()
	zingSongs := zingApi.SuggestionSong(params.Search, params.Limit)
	var songs []songmodels.Songs
	for _, e := range zingSongs {
		// song, err := biz.store.FindSong(e.ID)
		// if err != nil {
		// 	if err.Error() == "record not found" {
		// 		song_ins, err := biz.store.CreateSong(songmodels.Songs{
		// 			MaskId:    e.ID,
		// 			Title:     e.Title,
		// 			Thumbnail: &e.Thumb,
		// 			Duration:  e.Duration,
		// 		})
		// 		if err != nil {
		// 			return songs, err
		// 		}
		// 		song = *song_ins
		// 	} else {
		// 		return songs, err
		// 	}

		// }
		songs = append(songs, songmodels.Songs{
			MaskId:    e.ID,
			Title:     e.Title,
			Thumbnail: &e.Thumb,
			Duration:  e.Duration,
		})
	}

	return songs, nil
}
