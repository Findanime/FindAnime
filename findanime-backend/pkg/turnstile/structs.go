package turnstile

import "time"

type TurnstileValidationResponse struct {
	Success     bool          `json:"success"`
	ChallengeTs time.Time     `json:"challenge_ts"`
	Hostname    string        `json:"hostname"`
	ErrorCodes  []interface{} `json:"error-codes"`
	Action      string        `json:"action"`
	Cdata       string        `json:"cdata"`
}
