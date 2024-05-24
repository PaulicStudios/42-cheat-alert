package requests

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

var httpClient *http.Client

func GetTokenSetClient() {
	config := clientcredentials.Config{
		ClientID:     os.Getenv("API_CLIENT_ID"),
		ClientSecret: os.Getenv("API_CLIENT_SECRET"),
		TokenURL:     "https://api.intra.42.fr/oauth/token",
	}

	ctx := context.Background()
	_, err := config.Token(ctx)
	if err != nil {
		panic(err)
	}

	httpClient = config.Client(ctx)
}
