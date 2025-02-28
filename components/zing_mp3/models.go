package zingmp3

type StreamingSongRes struct {
	File128 string `json:"128"`
	File320 string `json:"320"`
}

type resSuggestion struct {
	Err     int    `json:"err"`
	Message string `json:"msg"`
	Data    struct {
		Items []struct {
			Suggestions []Song `json:"suggestions"`
		} `json:"items"`
	} `json:"data"`
}

type Song struct {
	Type           int    `json:"type"`
	Title          string `json:"title"`
	ID             string `json:"id"`
	RadioPid       string `json:"radioPid"`
	HasVideo       bool   `json:"hasVideo"`
	Thumb          string `json:"thumb"`
	ThumbVideo     string `json:"thumbVideo"`
	Duration       int    `json:"duration"`
	Link           string `json:"link"`
	ModifiedTime   int64  `json:"modifiedTime"`
	LyricLink      string `json:"lyricLink"`
	LyricID        string `json:"lyricId"`
	DownloadTypes  string `json:"downloadTypes"`
	OrgMD5         string `json:"orgMD5"`
	UserID         int    `json:"userId"`
	EuID           string `json:"euId"`
	Privacy        int    `json:"privacy"`
	HLyricVersion  int64  `json:"hLyricVersion"`
	ReleaseTime    int64  `json:"releaseTime"`
	DownloadStatus int    `json:"downloadStatus"`
	Status         int    `json:"status"`
	PlayStatus     int    `json:"playStatus"`
	// Artists        []Artist `json:"artists"`
	// Genres         []Genre  `json:"genres"`
	DisSPlatform int    `json:"disSPlatform"`
	DisDPlatform int    `json:"disDPlatform"`
	BoolAtt      int    `json:"boolAtt"`
	Tracking     string `json:"tracking"`
}
