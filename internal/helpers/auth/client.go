// refs https://developers.google.com/drive/v3/web/quickstart/go
package auth

import (
	"github.com/Juvenal-Yescas/gdown/internal/utils"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateClientApiDrive() (*drive.Service, error) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	client := getClient(config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
		return nil, err
	}
	return srv, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.

	tokFile := utils.GetDefaultConfigDir() // tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
