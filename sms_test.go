package esms

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func setup(t *testing.T) *Client {
	apiKey := os.Getenv("ESMS_API_KEY")
	secretKey := os.Getenv("ESMS_SECRET_KEY")
	client, err := NewClient(apiKey, secretKey)
	if err != nil {
		t.Fatal(err)
	}
	client.logger = log.Println
	return client
}

func TestGetBalance(t *testing.T) {
	client := setup(t)

	result, _, err := client.Sms.GetBalance()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.CodeResponse != CodeSuccess {
		t.Fatalf("expected %s, got %s", CodeSuccess, result.CodeResponse)
	}
	if result.UserID == 0 {
		t.Fatalf("UserID should not be empty")
	}
}

func TestSendMultipleMessage_Sandbox(t *testing.T) {
	client := setup(t)

	result, _, err := client.Sms.SendMultipleMessage(SendMultipleMessageInput{
		Phone:     "0123456789",
		Content:   "Cam on quy khach da su dung dich vu cua chung toi. Chuc quy khach mot ngay tot lanh!",
		SmsType:   SmsTypeBrandName,
		BrandName: "Baotrixemay",
		IsUnicode: "0",
		RequestId: fmt.Sprintf("request_%d", time.Now().UnixNano()),
		Sandbox:   Sandbox,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.CodeResult != CodeSuccess {
		t.Fatalf("CodeResult: expected %s, got %s", CodeSuccess, result.CodeResult)
	}
	if result.SmsID == "" {
		t.Fatalf("SMSID should not be empty")
	}
}
