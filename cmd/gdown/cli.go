package main

import (
	"fmt"
	"github.com/Juvenal-Yescas/gdown"
	"github.com/Juvenal-Yescas/gdown/internal/helpers/auth"
	"github.com/Juvenal-Yescas/gdown/internal/process"
	"github.com/Juvenal-Yescas/gdown/internal/utils"
	"os"
)

func cliDownload(url string) {
	output, err := gdown.Download(url)
	if err != nil {
		if utils.CaseInsensitiveContains(err.Error(), "exceeded") {
			fmt.Println("The file has exceeded the download limit, try download with your account")
			output, err := downloadWithAccount(url)
			if err != nil {
				fmt.Println("Error : ", err)
				os.Exit(1)
			}
			fmt.Println("File downloaded : ", output)
			os.Exit(0)
		}
		fmt.Println("Error : ", err)
		os.Exit(1)

	}
	fmt.Println("File downloaded: ", output)
	os.Exit(0)
}

func downloadWithAccount(url string) (string, error) {
	idFileUrl, err := process.GetIdFromUrl(url)
	if err != nil {
		return "", err
	}

	clientHttp := utils.CreateClientHttp()
	outputName, err := process.GetNameOutput(clientHttp, idFileUrl)
	if err != nil {
		return "", err
	}

	clientGDrive, err := auth.CreateClientApiDrive()
	if err != nil {
		return "", err
	}
	infoFolderDrive, infoFileDrive, err := process.CreateCopyInDrive(clientGDrive, idFileUrl, outputName)
	if err != nil {
		return "", err
	}
	fmt.Println("Downloading ... ", outputName)
	output, err := gdown.DownloadOutput("https://drive.google.com/uc?id="+infoFileDrive.Id, outputName)
	if err != nil {
		return "", err
	}

	err = process.CleanCopyInDrive(clientGDrive, infoFolderDrive.Id)
	if err != nil {
		return "", err
	}

	return output, nil
}
