package geminai

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent"
	apiKey       = "AIzaSyAY6Jj71TlYdlCimYuDdvtgr-GjFOBjup0" // 여기에 실제 API 키를 입력하세요.
)

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type RequestBody struct {
	Contents []Content `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []Part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func GeminiRequest() {
	// 환경 변수에서 API 키와 URL 가져오기
	apiKey := apiKey
	geminiAPIURL := geminiAPIURL

	if apiKey == "" || geminiAPIURL == "" {
		log.Fatal("❌ API 키 또는 URL이 설정되지 않았습니다.")
	}

	// 보낼 프롬프트 설정
	prompt := "Explain how black holes work."

	// Resty 클라이언트 생성
	client := resty.New()

	// 요청 본문 구성
	requestBody := RequestBody{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}

	// Gemini API 호출
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("key", apiKey).
		SetBody(requestBody).
		Post(geminiAPIURL)

	// 오류 처리
	if err != nil {
		log.Fatalf("❌ Gemini API 호출 실패: %v", err)
	}

	// 응답 상태 코드 확인
	if resp.StatusCode() != 200 {
		log.Fatalf("❌ Gemini API 응답 실패: 상태 코드 %d, 응답: %s", resp.StatusCode(), resp.String())
	}

	// 응답 JSON 파싱
	var geminiResponse GeminiResponse
	err = json.Unmarshal(resp.Body(), &geminiResponse)
	if err != nil {
		log.Fatalf("❌ Gemini 응답 JSON 파싱 실패: %v", err)
	}

	// 응답 출력
	fmt.Println("=======================================")
	fmt.Println(" 입력 프롬프트:")
	fmt.Println(prompt)
	fmt.Println("=======================================")
	fmt.Println(" Gemini 응답:")
	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		fmt.Println(geminiResponse.Candidates[0].Content.Parts[0].Text)
	} else {
		fmt.Println("응답 내용을 찾을 수 없습니다.")
	}
	fmt.Println("=======================================")
}
