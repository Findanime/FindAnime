package config

type RedisCredentials struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Config struct {
	Port                 int              `json:"port"`
	TurnstileSecret      string           `json:"turnstile-secret"`
	NowPaymentsMode      string           `json:"nowpayments-mode"`
	NowPaymentsSecretIPN string           `json:"nowpayments-secret-ipn"`
	NowPaymentsSecret    string           `json:"nowpayments-secret"`
	DatabaseURI          string           `json:"database-uri"`
	ProxiesGgAPIKey      string           `json:"proxiesgg-api"`
	DiscordWebhook       string           `json:"discord-webhook"`
	Redis                RedisCredentials `json:"redis-credentials"`
	OpenAIKey            string           `json:"openai-go"`
}
