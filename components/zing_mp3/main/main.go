package main

import (
	"fmt"

	zingmp3 "example.com/musicafy_be/components/zing_mp3"
)

func main() {
	cookie := zingmp3.GetCookie()
	fmt.Println("Cookie string:", cookie)
}
