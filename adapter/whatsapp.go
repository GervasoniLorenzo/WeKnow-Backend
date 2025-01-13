package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"weKnow/model"
)

type KnownWhatsApp struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

func (a KnownAdapter) SendWhatsApp(recipient string, message string) error {
	token := "EAARN7ok63p8BO7UOW7qiT0CmMZAmaPPxEMK6ZAVMZCcOXZBx78Ut3yZBc3KQRZA6i42OnpYlC9gdcPvlF7DB8soFrFyWN4fxymJp4rM8ZComv52NclBii6p9BZA8vFczGEVexSR7wdzNJbvzf9JSP6jgHLLXvZBkDodbRBRGykuR7RVwmmxFAhETdiU6hdYPTCxJyYg82oLYOeZBtiiKXNZBZALP3S30nvjA"
	phoneNumberID := "445724098626095"
	// , recipient
	url := fmt.Sprintf("https://graph.facebook.com/v17.0/%s/messages", phoneNumberID)

	// Crea il payload del messaggio
	msg := model.WhatsAppMessage{
		MessagingProduct: "whatsapp",
		To:               recipient,
		Type:             "text",
	}
	msg.Text.Body = message

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Crea la richiesta HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Invia la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verifica se la richiesta è andata a buon fine
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	fmt.Println("Message sent successfully!")
	return nil
}
