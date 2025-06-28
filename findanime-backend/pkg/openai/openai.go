package openai

import (
	"api/internal/config"
	"api/internal/helpers"
	"api/pkg/imdb"
	"api/pkg/logging"
	"api/pkg/myanimelist"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var (
	client openai.Client
)

func init() {
	client = openai.NewClient(
		option.WithAPIKey(config.Configuration.OpenAIKey),
	)
}
func AskGPT(anime string) (string, error) {
	prompt := fmt.Sprintf("Return exactly 12 anime similar to %s in a clean JSON array format. Each object must contain: title, rating (as a number), banner (with value of empty string), and description (1–2 sentences). Output only the raw JSON array — no explanations, no code blocks, no extra text. Ensure you do not query for banner url and just set it as empty string to save time and tokens", anime)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelChatgpt4oLatest,
	})
	if err != nil {
		logging.Logger.Error().Msg(err.Error())
		return "", err

	}
	return chatCompletion.Choices[0].Message.Content, nil
}

func AskGPTIMDB(anime string) (string, error) {
	jsonStr, err := AskGPT(anime)
	if err != nil {
		logging.Logger.Error().Msg(err.Error())
		return "", err
	}

	var animeList []Anime
	if err := json.Unmarshal([]byte(jsonStr), &animeList); err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect shared data

	// Updating banners concurrently with mutex protection
	for i := range animeList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Locking for safe access to shared resources
			mu.Lock()
			animeList[i].Banner = imdb.GetImageURL(animeList[i].Title)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	// Marshal modified data to JSON
	editedJSON, err := json.MarshalIndent(animeList, "", "  ")
	if err != nil {
		return "", err
	}
	go helpers.LocalizeQuery(anime, string(editedJSON))
	return string(editedJSON), nil
}

func AskAnimeList(anime string) (string, error) {
	jsonStr, err := AskGPT(anime)
	if err != nil {
		logging.Logger.Error().Msg(err.Error())
		return "", err
	}

	var animeList []Anime
	if err := json.Unmarshal([]byte(jsonStr), &animeList); err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect shared data

	// Updating banners concurrently with mutex protection
	for i := range animeList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Locking for safe access to shared resources
			mu.Lock()
			animeList[i].Banner = myanimelist.GetImageURL(animeList[i].Title)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	// Marshal modified data to JSON
	editedJSON, err := json.MarshalIndent(animeList, "", "  ")
	if err != nil {
		return "", err
	}
	go helpers.LocalizeQuery(anime, string(editedJSON))
	return string(editedJSON), nil
}
