package gdown

import (
	"github.com/Juvenal-Yescas/gdown/internal/process"
	"github.com/Juvenal-Yescas/gdown/internal/utils"
)

func Download(url string) (string, error) {
	idFile, err := process.GetIdFromUrl(url)
	if err != nil {
		return "", err
	}

	clientHttp := utils.CreateClientHttp()
	outputName, err := process.GetNameOutput(clientHttp, idFile)
	if err != nil {
		return "", err
	}

	urlDirect, err := process.GetUrlConfirmation(clientHttp, idFile)
	if err != nil {
		return "", err
	}

	downloaded, err := process.StartDownload(clientHttp, urlDirect, outputName)
	if err != nil {
		return "", err
	}
	return downloaded, nil
}

func DownloadOutput(url string, outputName string) (string, error) {
	idFile, err := process.GetIdFromUrl(url)
	if err != nil {
		return "", err
	}

	clientHttp := utils.CreateClientHttp()
	urlDirect, err := process.GetUrlConfirmation(clientHttp, idFile)
	if err != nil {
		return "", err
	}

	downloaded, err := process.StartDownload(clientHttp, urlDirect, outputName)
	if err != nil {
		return "", err
	}
	return downloaded, nil
}
