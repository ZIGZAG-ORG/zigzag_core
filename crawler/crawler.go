package crawler

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GoogleCrawler() {
	// 크롤러 생성
	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com", "news.google.com"),
	)

	// 검색어 설정 (TSLA 관련 뉴스 검색)
	searchQuery := "TSLA"
	url := fmt.Sprintf("https://www.google.com/search?q=%s&tbm=nws", searchQuery)

	// 크롤링할 요소 정의
	c.OnHTML("div.SoaBEf", func(e *colly.HTMLElement) {
		title := e.ChildText("div.mCBkyc")      // 기사 제목
		link := e.ChildAttr("a.WlydOe", "href") // 기사 링크

		// 구글 뉴스 링크는 "/url?q=" 이후 실제 URL이 포함됨
		if strings.Contains(link, "/url?q=") {
			link = strings.Split(link, "/url?q=")[1]
			link = strings.Split(link, "&")[0] // 불필요한 파라미터 제거
		}

		fmt.Printf("Title: %s\nLink: %s\n\n", title, link)
	})

	// 요청 실행
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("크롤링 중 오류 발생: %v", err)
	}
}
