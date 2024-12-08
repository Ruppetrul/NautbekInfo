package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.telegram.org/bot7952827286:AAG7EAV4s1jvuHmop7vPejvxWSZzbYbyJxs/sendMessage"
const chatId = int64(923984580)

type Message struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(text string) error {
	message := Message{
		ChatID: chatId,
		Text:   text,
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать данные в JSON: %v", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return fmt.Errorf("не удалось отправить запрос: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка от Telegram API: %v", resp.Status)
	}

	return nil
}
