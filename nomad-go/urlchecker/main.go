package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("request error")

func main() {
	// Go에서 map은 초기화를 하지 않으면 nil이 된다는데 이건 다른 동적타입 언어와 차이점
	// map 말고 다른 것도 그런가
	// results := make(map[string]string)
	results := map[string]string{}

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.naver.com",
		"https://www.academy.nomadcoders.co",
	}

	for _, url := range urls {
		result := "OK"
		start := time.Now()
		err := hitURL(url)
		fmt.Println("Get에 걸린 시간: ", time.Since(start))
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

// net 패키지 자체에 대해 따로 정리를 해보자
// 우선 HTTP 메소드를 사용하는 방법은 net/http 패키지의 함수를 사용한다.
// 사용법을 지금 단정짓기엔 모르는게 많고 아직 이해하기 힘든 공식 문서의 내용이 많으므로 우선 서버 구성하는
// 실습을 해본 다음 감이 잡히면 그때 정리를 해야겠다.
func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
