package imdb

import (
	"api/internal/helpers"
	"api/pkg/logging"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"

	http "github.com/bogdanfinn/fhttp"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

var (
	// Default Headers
	headers = http.Header{
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"},
		"Accept-Encoding":           {"gzip, deflate, br, zstd"},
		"Accept-Language":           {"en-GB,en;q=0.7"},
		"Connection":                {"keep-alive"},
		"Host":                      {"www.imdb.com"},
		"sec-ch-ua":                 {`"Brave";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {`"macOS"`},
		"Sec-Fetch-Dest":            {"document"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-Site":            {"none"},
		"Sec-Fetch-User":            {"?1"},
		"Sec-GPC":                   {"1"},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"},
	}

	// NotFoundURL is the URL to return when no image is found
	NotFoundURL = "https://ih1.redbubble.net/image.4905811447.8675/flat,750x,075,f-pad,750x1000,f8f8f8.jpg"
)

func GetImageURL(imdbID string) string {
	// Construct the Request for the Query Search

	request, err := http.NewRequest("GET", "https://www.imdb.com/find/?q="+url.QueryEscape(imdbID), nil)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create request for IMDb image")
		return ""
	}
	request.Header = headers
	client, err := SetupClient()

	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create HTTP client for IMDb image")
		return ""

	}
	err = Authenticate(client)
	if err != nil {
		return ""
	}
	response, err := client.Do(request)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to fetch IMDb image")
		return ""
	}
	if response.StatusCode != http.StatusOK {
		logging.Logger.Error().Msgf("Failed to fetch IMDb image, status code: %d", response.StatusCode)
		return NotFoundURL
	}
	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)
	bstring := string(b)
	helpers.SaveBodyToFile(response.Body)
	return ParseImageURL(bstring)

}

func Authenticate(client tls_client.HttpClient) error {
	// Request to Set Cookies For Authentication
	request, err := http.NewRequest("GET", "https://www.imdb.com/", nil)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create request for IMDb authentication")
	}
	request.Header = headers
	response, err := client.Do(request)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to authenticate with IMDb")
	}
	if response.StatusCode != http.StatusOK {
		logging.Logger.Error().Msgf("Failed to authenticate with IMDb, status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	return nil
}

func SetupClient() (tls_client.HttpClient, error) {

	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_133),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create HTTP client")
		return nil, err
	}

	return client, nil

}

func ParseImageURL(htmlContent string) string {
	const baseMarker = `ipc-image" loading="lazy" src="https://m.media-amazon.com/images/M/`

	// Check if the HTML content contains the image URL pattern
	if !strings.Contains(htmlContent, baseMarker) {
		logging.Logger.Error().Msg("Image URL not found in HTML content")
		return NotFoundURL
	}

	// Extract the raw partial URL
	rawPart := strings.Split(htmlContent, baseMarker)[1]
	rawURL := "https://m.media-amazon.com/images/M/" + strings.Split(rawPart, `"`)[0]

	// Now clean and replace modifiers (e.g., _V1_QL75_UX50_CR..._.jpg)
	re := regexp.MustCompile(`_V1[^.]+`)
	highQualityModifier := "_V1_UX1280"

	// Replace or insert new quality suffix
	if re.MatchString(rawURL) {
		rawURL = re.ReplaceAllString(rawURL, highQualityModifier)
	} else {
		// If no modifier, append one before .jpg
		rawURL = strings.Replace(rawURL, ".jpg", highQualityModifier+".jpg", 1)
	}

	return rawURL
}
