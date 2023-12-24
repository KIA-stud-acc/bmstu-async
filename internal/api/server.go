package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandomStatus() int {
	time.Sleep(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000) + 1
}

type Result struct {
	QuantityOfVotes int `json:"QuantityOfVotes"`
	Key             int
	Id              int
}

func PerformPUTRequest(url string, data Result) (*http.Response, error) {
	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Создание PUT-запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

func SendStatus(pk int, url string) {
	// Выполнение расчётов с randomStatus
	result := RandomStatus()

	// Отправка PUT-запроса к основному серверу
	data := Result{QuantityOfVotes: result, Key: 123456, Id: pk}
	_, err := PerformPUTRequest(url, data)
	if err != nil {
		fmt.Println("Error sending status:", err)
		return
	}

	fmt.Println("Status sent successfully for pk:", pk)
}
