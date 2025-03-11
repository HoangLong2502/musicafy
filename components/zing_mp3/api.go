package zingmp3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"example.com/musicafy_be/common"
	"github.com/rs/zerolog/log"
)

type ZingMp3API struct {
	urlAc   string
	url     string
	version string
	apiKey  string
}

func NewZingMp3Api(urlAc string, url string, version string, apiKey string) ZingMp3API {
	return ZingMp3API{
		urlAc:   urlAc,
		url:     url,
		version: version,
		apiKey:  apiKey,
	}
}

// makeRequest thực hiện HTTP request với các tham số chung
func (api *ZingMp3API) makeRequest(method, url string, params url.Values, needCookie bool) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, common.NewCustomError(err, "Error creating request", "lỗi tạo request", "makeRequest")
	}

	if needCookie {
		cookie := GetCookie()
		req.Header.Add("cookie", cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, common.NewCustomError(err, "Error sending request", "lỗi gửi request", "makeRequest")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, common.NewCustomError(err, "Error reading response", "lỗi đọc response", "makeRequest")
	}

	return body, nil
}

// buildRequestURL tạo URL với các tham số
func (api *ZingMp3API) buildRequestURL(doamin string, path string, params url.Values) string {
	return fmt.Sprintf("%s%s?%s", doamin, path, params.Encode())
}

func (api *ZingMp3API) SuggestionSong(query string, num int) []SongSearch {
	params := url.Values{}
	params.Add("num", fmt.Sprintf("%d", num))
	params.Add("query", query)
	params.Add("language", "vi")

	fullURL := api.buildRequestURL(api.urlAc, "/v1/web/ac-suggestions", params)
	body, err := api.makeRequest("GET", fullURL, params, false)
	if err != nil {
		log.Error().Err(err).Msg("Error in SuggestionSong")
		return nil
	}

	var res resSuggestion
	if err := json.Unmarshal(body, &res); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JSON in SuggestionSong")
		return nil
	}

	if res.Err != 0 {
		log.Error().Msg(res.Message)
		return nil
	}

	return res.Data.Items[1].Suggestions
}

func (api *ZingMp3API) StreamFileSong(id string) *StreamingSongRes {
	path := "/api/v2/song/get/streaming"
	ct := time.Now().Unix()

	params := url.Values{}
	params.Add("ctime", fmt.Sprintf("%d", ct))
	params.Add("id", id)
	params.Add("version", "1.13.3")

	// Tạo signature
	n := strings.ReplaceAll(params.Encode(), "&", "")
	sig := GenerateHash(path, n)
	params.Add("sig", sig)
	params.Add("apiKey", api.apiKey)

	fullURL := api.buildRequestURL(api.url, path, params)
	body, err := api.makeRequest("GET", fullURL, params, true)
	if err != nil {
		log.Error().Err(err).Msg("Error in StreamFileSong")
		return nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JSON in StreamFileSong")
		return nil
	}

	var errCode int
	if err := json.Unmarshal(raw["err"], &errCode); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling error code")
		return nil
	}

	if errCode != 0 {
		return nil
	}

	var data StreamingSongRes
	if err := json.Unmarshal(raw["data"], &data); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling streaming data")
		return nil
	}

	return &data
}

func (api *ZingMp3API) DetailSong(id string) (*SongDetail, error) {
	path := "/api/v2/song/get/info"
	ct := time.Now().Unix()

	params := url.Values{}
	params.Add("ctime", fmt.Sprintf("%d", ct))
	params.Add("id", id)
	params.Add("version", "1.13.3")
	params.Add("sig", GenerateHash(path, params.Encode()))
	params.Add("apiKey", api.apiKey)

	fullURL := api.buildRequestURL(api.url, path, params)
	body, err := api.makeRequest("GET", fullURL, params, true)
	if err != nil {
		return nil, err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, common.NewCustomError(err, "Error unmarshalling JSON", "lỗi unmarshalling JSON", "DetailSong")
	}

	var song SongDetail
	if err := json.Unmarshal(raw["data"], &song); err != nil {
		return nil, common.NewCustomError(err, "Error unmarshalling song detail", "lỗi unmarshalling chi tiết bài hát", "DetailSong")
	}

	return &song, nil
}
