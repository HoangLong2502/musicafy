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
	FindSong(id string) (*songmodels.Songs, error)
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
	songs := make([]songmodels.Songs, len(zingSongs))
	for i, e := range zingSongs {
		songs[i] = e.ToModelDb()
	}
	return songs, nil
}
