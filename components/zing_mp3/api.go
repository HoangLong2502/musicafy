package zingmp3

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SuggestionSong(query string, num int) {
	baseURL := "https://ac.zingmp3.vn/v1/web/ac-suggestions"

	ct := time.Now().Unix()
	// Xây dựng query params động
	params := url.Values{}
	params.Add("ctime", string(ct))
	params.Add("version", "1.13.0")

	n := strings.ReplaceAll(params.Encode(), "&", "")
	sig := GenerateHash("/v1/web/ac-suggestions", n)

	params.Add("num", fmt.Sprintf("%d", num))
	params.Add("query", query)
	params.Add("language", "vi")
	params.Add("sig", sig)
	params.Add("apiKey", "X5BM3w8N7MKozC0B85o4KMlzLZKhV00y")
	// Tạo URL đầy đủ
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Gửi request
	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

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
		return
	}

	// In kết quả
	fmt.Println(string(body))
}
