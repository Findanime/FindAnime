package helpers

import (
	"api/pkg/logging"
	"api/pkg/mongodb"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Response(c fiber.Ctx, status int, data ...any) error {
	message := ""

	if len(data) > 0 {
		message = fmt.Sprintf("%v", data[0])
	} else {
		message = fmt.Sprintf("%v: %v", status, http.StatusText(status))
	}

	return c.Status(status).JSON(fiber.Map{
		"code":    status,
		"message": message,
	})
}

func FormatResponse(raw string) (string, error) {
	// Step 1: Parse the outer JSON
	var outer map[string]string
	if err := json.Unmarshal([]byte(raw), &outer); err != nil {
		return "", fmt.Errorf("failed to parse outer JSON: %w", err)
	}

	// Step 2: Extract and unmarshal the inner JSON string
	messageStr := outer["message"]
	var inner any
	if err := json.Unmarshal([]byte(messageStr), &inner); err != nil {
		return "", fmt.Errorf("failed to parse inner JSON: %w", err)
	}

	// Step 3: Pretty-print the inner JSON
	prettyJSON, err := json.MarshalIndent(inner, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal pretty JSON: %w", err)
	}

	return string(prettyJSON), nil
}

func SaveBodyToFile(body io.Reader) error {
	file, err := os.Create("debug.html")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	return err
}

func LocalizeQuery(anime string, data string) error {

	var recommendations []map[string]interface{}

	err := json.Unmarshal([]byte(data), &recommendations)
	if err != nil {
		logging.Logger.Error().Msgf("Failed to unmarshal recommendations: %v", err)
		return err
	}

	payload := bson.M{
		"anime": anime,
		"data":  recommendations,
	}
	_, err = mongodb.DB.InsertOne(payload, mongodb.DB.Collections.Recommendations)
	if err != nil {
		logging.Logger.Error().Msgf("Failed to insert recommendations into database: %v", err)
	}

	return nil
}
