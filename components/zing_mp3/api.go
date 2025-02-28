package zingmp3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

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

func (api *ZingMp3API) SuggestionSong(query string, num int) []Song {
	baseURL := fmt.Sprintf("%s/v1/web/ac-suggestions", api.urlAc)
	// Xây dựng query params động
	params := url.Values{}
	params.Add("num", fmt.Sprintf("%d", num))
	params.Add("query", query)
	params.Add("language", "vi")
	// Tạo URL đầy đủ
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		panic(err)
	}

	var res resSuggestion
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}
	if res.Err != 0 {
		panic(res.Message)
	}
	fmt.Println(string(body))
	return res.Data.Items[1].Suggestions
}

func (api *ZingMp3API) StreamFileSong(id string) *StreamingSongRes {
	baseURL := "https://zingmp3.vn/api/v2/song/get/streaming"

	ct := time.Now().Unix()
	// Xây dựng query params động
	params := url.Values{}
	ct_params := fmt.Sprintf("%d", ct)
	params.Add("ctime", ct_params)
	params.Add("id", id)
	params.Add("version", "1.13.0")
	fmt.Println("===== ctime:", ct_params)
	// n := strings.ReplaceAll(params.Encode(), "&", "")
	n := fmt.Sprintf("ctime=%did=%sversion=%s", ct, id, "1.13.0")
	sig := GenerateHash("/api/v2/song/get/streaming", n)

	fmt.Println("===== Generated Hash:", sig)

	params.Add("sig", sig)
	params.Add("apiKey", "X5BM3w8N7MKozC0B85o4KMlzLZKhV00y")

	// Tạo URL đầy đủ
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Gửi request
	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Add("Cookie", "zmp3_rqid=MHwyMjIdUngMjUyLjE4LjM4fG51WeBGx8MTmUsIC0MDmUsICwNzgyMTM3OQ")

	// Thực hiện request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Đọc response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	var raw map[string]json.RawMessage
	err = json.Unmarshal(body, &raw)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	// In kết quả
	fmt.Println(string(body))
	var data StreamingSongRes
	err = json.Unmarshal(raw["data"], &data)
	if err != nil {
		log.Fatal().Msg("Can't decode response zing api")
	}

	return &data
}
