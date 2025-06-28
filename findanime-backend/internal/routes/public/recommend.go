package routes

import (
	"api/internal/helpers"
	"api/pkg/mongodb"
	"api/pkg/openai"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Recommend(c fiber.Ctx) error {
	var recommendations string
	anime := strings.ToLower(c.Query("anime"))
	if anime == "" {
		return helpers.Response(c, fiber.StatusBadRequest, "Anime query parameter is required")
	}

	payload := bson.M{"anime": anime}
	query, err := mongodb.DB.FindOne(payload, mongodb.DB.Collections.Recommendations)
	if err == nil {
		data, ok := query["data"].(primitive.A)
		if !ok {
			return helpers.Response(c, fiber.StatusInternalServerError, "Failed to parse recommendations")
		}
		recjson, err := json.Marshal(data)
		if err != nil {
			return helpers.Response(c, fiber.StatusInternalServerError, "Failed to marshal recommendations")
		}
		recommendations = string(recjson)
		if recommendations != "" {
			return helpers.Response(c, fiber.StatusOK, recommendations)
		}
	}
	recommendations, err = openai.AskAnimeList(anime)
	if err != nil {
		return helpers.Response(c, fiber.StatusInternalServerError, "Failed to get recommendations")
	}

	return helpers.Response(c, fiber.StatusOK, recommendations)
}
