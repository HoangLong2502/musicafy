package songsbiz

import (
	"errors"

	"example.com/musicafy_be/components/appctx"
)

type StreamingSongRes struct {
	File128 string `json:"128"`
	File320 string `json:"320"`
}

type StreamingSongStore interface {
	// FindSong(id string) (*songmodels.Songs, error)
}

type StreamingSongBiz struct {
	store StreamingSongStore
}

func NewStreamingSongBiz(store StreamingSongStore) *StreamingSongBiz {
	return &StreamingSongBiz{store: store}
}

func (biz *StreamingSongBiz) StreamingSong(appContext appctx.AppContext, id string) (*StreamingSongRes, error) {
	zingMp3Api := appContext.GetZingmp3Api()
	res := zingMp3Api.StreamFileSong(id)
	if res == nil {
		return nil, errors.New("không lấy được file streaming")
	}
	return &StreamingSongRes{
		File128: res.File128,
		File320: res.File320,
	}, nil
}
