package request_api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	RealWebSocketURL = "ws://ops.koreainvestment.com:21000/tryitout/H0GSCNI0" // 실전 투자
)

// ✅ 웹소켓 요청 데이터 구조체
type WebSocketRequest struct {
	Header struct {
		ApprovalKey string `json:"approval_key"`
		CustType    string `json:"custtype"`
		TrType      string `json:"tr_type"`
		ContentType string `json:"content-type"`
	} `json:"header"`
	Body struct {
		Input struct {
			TrID  string `json:"tr_id"`
			TrKey string `json:"tr_key"`
		} `json:"input"`
	} `json:"body"`
}

// ✅ 웹소켓 응답 데이터 구조체
type WebSocketResponse struct {
	Header struct {
		TrID    string `json:"tr_id"`
		TrKey   string `json:"tr_key"`
		Encrypt string `json:"encrypt"`
	} `json:"header"`
	Body struct {
		RtCd   string `json:"rt_cd"`
		MsgCd  string `json:"msg_cd"`
		Msg    string `json:"msg1"`
		Output struct {
			IV  string `json:"iv"`
			Key string `json:"key"`
		} `json:"output"`
	} `json:"body"`
}

func ConnectToTradingWebSocket(approvalKey string) error {

	webSocketURL := RealWebSocketURL

	// 📌 웹소켓 연결
	header := http.Header{}
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, header)
	if err != nil {
		return fmt.Errorf("웹소켓 연결 실패: %v", err)
	}
	defer conn.Close()

	fmt.Println("✅ 웹소켓 연결 성공!")

	// 📌 구독 요청 데이터 생성
	request := WebSocketRequest{
		Header: struct {
			ApprovalKey string `json:"approval_key"`
			CustType    string `json:"custtype"`
			TrType      string `json:"tr_type"`
			ContentType string `json:"content-type"`
		}{
			ApprovalKey: approvalKey,
			CustType:    "P", // 개인 계정
			TrType:      "1", // 1: 등록
			ContentType: "utf-8",
		},
		Body: struct {
			Input struct {
				TrID  string `json:"tr_id"`
				TrKey string `json:"tr_key"`
			} `json:"input"`
		}{
			Input: struct {
				TrID  string `json:"tr_id"`
				TrKey string `json:"tr_key"`
			}{
				TrID:  "H0GSCNI0", // ✅ 실시간 해외주식 체결통보 (실전)
				TrKey: "@2512782",
			},
		},
	}

	// 📌 JSON 변환 후 웹소켓으로 요청 전송
	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("JSON 변환 실패: %v", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		return fmt.Errorf("구독 요청 실패: %v", err)
	}
	fmt.Println("📩 구독 요청 전송 완료!")

	// 📌 실시간 데이터 수신
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("📌 웹소켓 수신 오류:", err)
			log.Println("🚨 연결 종료 감지. 재연결 시도 중...")
			time.Sleep(5 * time.Second) // 재연결 전 대기
			ConnectToTradingWebSocket(approvalKey)
		}

		// 📌 JSON 응답 파싱
		var response WebSocketResponse
		err = json.Unmarshal(message, &response)
		if err != nil {
			fmt.Println("📊 실시간 체결 데이터 (TEXT):", string(message))
			continue
		}

		// ✅ 성공적인 응답 확인
		if response.Body.RtCd == "0" {
			fmt.Println("🔹 실시간 체결통보 구독 성공:", response.Body)
		} else {
			fmt.Println("🔹 실시간 체결 전:", response.Body)
		}
	}
}
