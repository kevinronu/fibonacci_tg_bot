package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevinronu/fibonacci-tg-bot/utils"
)

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handleTelegramWebHook)
}

func handleTelegramWebHook(ctx context.Context, update Update) {
	if update.Message.Text == "/start" {
		log.Println("Printed welcome")
		sendTextToTelegramChat(update.Message.Chat.Id, "Welcome! Please enter the position of the Fibonacci number you want to calculate:")
	} else {
		pos, err := utils.StringToInt(update.Message.Text)
		if err != nil {
			sendTextToTelegramChat(update.Message.Chat.Id, err.Error())
			return
		}

		result, err := utils.GetFibonacciNumberWithPosition(pos)
		if err != nil {
			sendTextToTelegramChat(update.Message.Chat.Id, err.Error())
			return
		}

		log.Printf("Calculated number for position %s: %d", update.Message.Text, result)
		sendTextToTelegramChat(update.Message.Chat.Id, fmt.Sprintf("The fibonacci number in position %d one is: %d", pos, result))
	}
}

func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return nil, err
	}

	return &update, nil
}

func sendTextToTelegramChat(chatId int, text string) error {
	tgBotToken := utils.GetEnv("TELEGRAM_BOT_TOKEN")
	telegramApiBaseUrl := "https://api.telegram.org/bot"
	telegramApiSendMessage := "/sendMessage"
	var telegramApiUrl string = telegramApiBaseUrl + tgBotToken + telegramApiSendMessage

	log.Printf("Sending %s to chat_id: %d", text, chatId)

	response, err := http.PostForm(
		telegramApiUrl,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return err
	}

	defer response.Body.Close()

	var bodyBytes, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return err
	}

	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return nil
}
