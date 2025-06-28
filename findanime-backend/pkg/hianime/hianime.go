package hianime

import (
	"api/internal/helpers"
	"api/pkg/logging"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
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

func GetImageURL(animeID string) string {
	// Construct the Request for the Query Search

	request, err := http.NewRequest("GET", fmt.Sprintf("https://myanimelist.net/search/prefix.json?type=all&keyword=%s&v=1", url.QueryEscape(animeID)), nil)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create request for MyAnimeList image")
		return ""
	}
	request.Header = headers
	client, err := SetupClient()

	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create HTTP client for MyAnimeList image")
		return ""

	}
	err = Authenticate(client)
	if err != nil {
		return ""
	}
	response, err := client.Do(request)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to fetch MyAnimeList image")
		return ""
	}
	if response.StatusCode != http.StatusOK {
		logging.Logger.Error().Msgf("Failed to fetch MyAnimeList image, status code: %d", response.StatusCode)
		return NotFoundURL
	}
	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)
	helpers.SaveBodyToFile(response.Body)
	url, err := GetFirstImageURL(b)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to parse image URL from MyAnimeList response")
		return NotFoundURL
	}
	return url

}

func Authenticate(client tls_client.HttpClient) error {
	// Request to Set Cookies For Authentication
	request, err := http.NewRequest("GET", "https://myanimelist.net/", nil)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to create request for MyAnimeList authentication")
	}
	request.Header = headers
	response, err := client.Do(request)
	if err != nil {
		logging.Logger.Error().Err(err).Msg("Failed to authenticate with MyAnimeList")
	}
	if response.StatusCode != http.StatusOK {
		logging.Logger.Error().Msgf("Failed to authenticate with MyAnimeList, status code: %d", response.StatusCode)
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

func GetFirstImageURL(jsonData []byte) (string, error) {
	// Return first image url from the json response as the correct image banner for the anime query
	var resp ApiResponse
	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		return "", err
	}

	if len(resp.Categories) == 0 || len(resp.Categories[0].Items) == 0 {
		return "", fmt.Errorf("no categories or items found")
	}

	var imageURL string

	if strings.Contains(resp.Categories[0].Items[0].ImageURL, "r/116x180/") {
		imageURL = strings.Replace(resp.Categories[0].Items[0].ImageURL, `r/116x180/`, "", 1)
	} else {
		imageURL = resp.Categories[0].Items[0].ImageURL
	}
	return imageURL, nil
}
