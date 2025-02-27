package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Структура для парсинга JSON
type Competition struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Response struct {
	Competitions []Competition `json:"competitions"`
}

func Football_get() {
	apiKey := os.Getenv("FOOTBALL_API_KEY") // Используем API-ключ из переменной окружения
	url := "https://api.football-data.org/v4/competitions"

	// Создаём HTTP-клиент
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Auth-Token", apiKey)

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, _ := ioutil.ReadAll(resp.Body)

	// Парсим JSON
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}

	// Выводим список лиг
	for _, comp := range response.Competitions {
		fmt.Printf("ID: %d, Лига: %s (%s)\n", comp.ID, comp.Name, comp.Code)
	}
}
