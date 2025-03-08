package request_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	service "zigzag-core/service/just"

	"gorm.io/gorm"
)

const (
	RealDomain = "https://openapi.koreainvestment.com:9443"
)

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	AppKey    string `json:"appkey"`
	SecretKey string `json:"secretkey"`
}

type TokenResponse struct {
	ApprovalKey string `json:"approval_key"`
}

func GetKoreaInvestmentToken(appKey, secretKey string, db *gorm.DB) (*TokenResponse, error) {
	fmt.Println("진입중..")
	apikeys, err := service.GetAPIKeyByUserPK(db, "7")
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	// API URL 설정
	baseURL := RealDomain // 실전 투자 선택 시 변경

	url := baseURL + "/oauth2/Approval"

	// 요청 데이터 생성
	payload := TokenRequest{
		GrantType: "client_credentials",
		AppKey:    apikeys.AppKey,
		SecretKey: apikeys.AppSecret,
	}

	fmt.Println(payload.GrantType)
	fmt.Println(payload.AppKey)
	fmt.Println(payload.SecretKey)

	// JSON 변환
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("JSON 변환 실패: %v", err)
	}

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 생성 실패: %v", err)
	}

	// HTTP 헤더 설정
	req.Header.Set("Content-Type", "application/json")

	// HTTP 클라이언트 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP 요청 실패: %v", err)
	}
	defer resp.Body.Close()

	// 응답 바디 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 실패: %v", err)
	}

	// JSON 응답 파싱
	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return nil, fmt.Errorf("응답 JSON 파싱 실패: %v", err)
	}

	// 응답 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 요청 실패: %s, 응답: %s", resp.Status, string(body))
	}

	// ✅ 성공적으로 토큰 반환
	return &tokenResp, nil
}
