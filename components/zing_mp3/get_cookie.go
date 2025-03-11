package zingmp3

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func GetCookie() string {
	// Tạo cookie jar để lưu trữ cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Lỗi khi tạo cookie jar:", err)
		os.Exit(1)
	}

	// Tạo HTTP client với cookie jar
	client := &http.Client{
		Jar: jar,
	}

	// URL cần truy cập
	urlString := "https://zingmp3.vn/"
	targetURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Lỗi khi phân tích URL:", err)
		os.Exit(1)
	}

	// Thực hiện request GET
	resp, err := client.Get(urlString)
	if err != nil {
		fmt.Println("Lỗi khi thực hiện request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Lấy tất cả cookie từ response
	cookies := jar.Cookies(targetURL)

	cookieString := ""
	// In ra các cookie
	fmt.Println("Cookies từ", urlString, ":")
	if len(cookies) == 0 {
		fmt.Println("Không tìm thấy cookie nào")
	} else {
		for i, cookie := range cookies {
			cookieString += fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value)
			fmt.Printf("%d. Tên: %s, Giá trị: %s\n", i+1, cookie.Name, cookie.Value)
			fmt.Printf("   Domain: %s, Path: %s, Secure: %t, HttpOnly: %t\n",
				cookie.Domain, cookie.Path, cookie.Secure, cookie.HttpOnly)
			fmt.Printf("   Expires: %s\n", cookie.Expires)
			fmt.Println("-----------------------------------")
		}
	}

	fmt.Println("Trạng thái HTTP:", resp.Status)
	return cookieString
}
