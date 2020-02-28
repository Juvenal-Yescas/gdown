package process

import (
	"errors"
	"github.com/Juvenal-Yescas/gdown/internal/helpers/webscraping"
	"github.com/Juvenal-Yescas/gdown/internal/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetUrlConfirmation(clientHttp *http.Client, idFile string) (string, error) {

	urlGoogleDrive := "https://drive.google.com/uc?id=" + idFile

	// Generate cookie
	resp, err := clientHttp.Get(urlGoogleDrive)
	if err != nil {
		log.Error("Client get : %v", err)
		return "", err
	}

	defer resp.Body.Close()

	var urlDirect string

	if len(resp.Cookies()) == 0 {

		data, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(data)

		if utils.CaseInsensitiveContains(bodyString, "exceeded") {
			log.Info("Google Drive - Quota exceeded : Try download using your account...")
			log.Debug(urlGoogleDrive)
			return "", errors.New("This file is limited: Quota exceeded")
		}
		// small file
		return urlGoogleDrive, nil
	}

	// big file
	codeConfirmation := resp.Cookies()[0].Value
	urlDirect = urlGoogleDrive + "&confirm=" + codeConfirmation
	return urlDirect, nil
}

func StartDownload(clientHttp *http.Client, urlDirect string, outputName string) (string, error) {
	log.Info("Downloading ... ", outputName)

	fileOutput, err := utils.CreateFile(outputName)
	if err != nil {
		return "", err
	}

	resp, err := clientHttp.Get(urlDirect)
	if err != nil {
		log.Error("Client get : %v", err)
		return "", err
	}

	defer resp.Body.Close()
	_, err = utils.WriteToOutput(fileOutput, resp.Body)
	if err != nil {
		log.Error("Write data error : %v", err)
		return "", err
	}
	return outputName, nil
}

func GetIdFromUrl(url string) (string, error) {
	if utils.CaseInsensitiveContains(url, "google.com") {
		if utils.CaseInsensitiveContains(url, "id") {
			arrayUrl := strings.Split(url, "=")
			if utils.CaseInsensitiveContains(arrayUrl[1], "export") {
				arrayUrl := strings.Split(arrayUrl[1], "&")
				return arrayUrl[0], nil
			}
			return arrayUrl[1], nil
		} else if utils.CaseInsensitiveContains(url, "docs") {
			log.Error("Documents not implemented")
			return "", errors.New("Documents not implemented")
		} else {
			arrayUrl := strings.Split(url, "/")
			return arrayUrl[5], nil
		}
	}
	return url, nil
}

func GetNameOutput(client *http.Client, idFile string) (string, error) {
	urlOpen := "https://drive.google.com/open?id=" + idFile

	tittleWebsite := webscraping.GetTittle(client, urlOpen)

	extraTittle := regexp.MustCompile(" - Google Drive")
	outputName := extraTittle.ReplaceAllString(tittleWebsite, "")
	log.Debug("Correct name output: ", outputName)

	if outputName == "" {
		return "", errors.New("Tittle dont found")
	}
	return outputName, nil
}
