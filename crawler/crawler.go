package crawler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GoogleCrawler() {
	// Chrome 실행 옵션 추가
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath("/usr/bin/google-chrome"), // Chrome 실행 경로 설정
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.DisableGPU,
		chromedp.Headless, // 헤드리스 모드 활성화
		chromedp.Flag("disable-software-rasterizer", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-dev-shm-usage", true), // Docker 환경에서 메모리 문제 방지
	)

	// Chrome 실행 컨텍스트 생성
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 뉴스 제목과 URL을 저장할 변수
	var titles, urls []string

	// 크롤링 실행
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.google.com/search?lr=lang_en&cr=countryUS&sca_esv=4a0a8c718760d881&tbs=qdr:d,lr:lang_1en,ctr:countryUS&sxsrf=AHTn8zpR9dSRYHtsevgbXPR9jYxQ5SwT8g:1741345786008&q=tesla&tbm=nws&source=lnms`),
		chromedp.Sleep(2*time.Second), // 페이지 로딩 대기

		// 뉴스 제목 가져오기 (CSS 선택자 사용)
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".n0jPhd.ynAwRc.MBeuO.nDgy9d")).map(e => e.innerText)`, &titles),

		// 뉴스 URL 가져오기 (a 태그의 href 속성 크롤링)
		chromedp.Evaluate(`Array.from(document.querySelectorAll("a.WlydOe")).map(e => e.href)`, &urls),
	)

	if err != nil {
		log.Fatal(err)
	}

	// 뉴스 기사 제목 및 URL 출력
	fmt.Println("🔗 크롤링된 뉴스 기사 목록:")
	for i := range titles {
		// 유효한 링크만 출력
		if i < len(urls) && strings.HasPrefix(urls[i], "http") {
			fmt.Printf("📌 %s\n🔗 %s\n\n", titles[i], urls[i])
		}
	}
}
