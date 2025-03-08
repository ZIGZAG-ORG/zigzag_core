package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"zigzag-core/request_api"
	"zigzag-core/route"

	"github.com/LabStars/selpo-common/crypto"
	"github.com/LabStars/selpo-common/db"
)

func main() {
	// ✅ 1️⃣ DB 설정 파일 읽기
	dbConfig, err := db.ReadJSONFromFile("./config/dbSettings.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ✅ 2️⃣ 사용자 입력으로 암호화된 DB 패스워드 복호화
	key, err := db.ReadUserInput()
	if err != nil {
		fmt.Println("Failed to read from console:", err)
		return
	}

	dbPassword, err := crypto.Decrypt(dbConfig.Password, key)
	if err != nil {
		fmt.Printf("Decryption failed: %v", err)
		return
	}

	// ✅ 3️⃣ DB 연결 (main에서 한 번만 실행)
	db, err := db.ConnectDatabase(&dbConfig, dbPassword)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// ✅ 4️⃣ API 서버 먼저 실행 (8080)
	go func() {
		fmt.Println("🚀 API 서버 실행 중... (포트 8080)")
		route.StartServer(db, 8080)
	}()

	// ✅ 6️⃣ 한국투자증권 API 토큰 요청 (API 서버 실행 후)
	tokenResp, err := request_api.GetKoreaInvestmentToken("", "", db)
	if err != nil {
		log.Fatalf("❌ 토큰 요청 실패: %v", err)
	}

	// ✅ 5️⃣ 웹소켓 서버 실행 (9090)
	go func() {
		request_api.ConnectToTradingWebSocket(tokenResp.ApprovalKey)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("\n🛑 서버 종료")
}
