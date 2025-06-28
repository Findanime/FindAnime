package turnstile

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/goccy/go-json"
)

var InvalidToken = errors.New("invalid token")

func ValidateTurnstileToken(token, secret string) error {
	params := url.Values{
		"response": {token},
		"secret":   {secret},
	}

	req, err := http.NewRequest("POST", "https://challenges.cloudflare.com/turnstile/v0/siteverify", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var reply TurnstileValidationResponse
	if err = json.Unmarshal(body, &reply); err != nil {
		return err
	} else if !reply.Success {
		return InvalidToken
	}

	return nil
}
