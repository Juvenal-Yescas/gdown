package process

import (
	// "github.com/Juvenal-Yescas/gdown/internal/helpers/auth"
	"github.com/Juvenal-Yescas/gdown/pkg/gdriveapi"
	"google.golang.org/api/drive/v3"
)

func CreateCopyInDrive(clientGDrive *drive.Service, idFile string, outputName string) (*drive.File, *drive.File, error) {
	idFolder, err := gdriveapi.CreateFolder(clientGDrive, "gdown")
	if err != nil {
		return nil, nil, err
	}

	_, err = gdriveapi.MakeSharedFolder(clientGDrive, idFolder.Id)
	if err != nil {
		return idFolder, nil, err
	}
	arrayFolder := []string{idFolder.Id}

	idFileCopied, err := gdriveapi.CreateACopy(clientGDrive, idFile, outputName, arrayFolder)
	if err != nil {
		return idFolder, nil, err
	}

	return idFolder, idFileCopied, nil
}

func CleanCopyInDrive(clientGDrive *drive.Service, idFolder string) error {
	err := gdriveapi.Delete(clientGDrive, idFolder)
	return err
}

// func SkipQuotaExceeded(clientWeb *http.Client, idFile string, outputName string) (string, error) {
// 	client, err := auth.CreateClient()
// 	if err != nil {
// 		return "", err
// 	}

// 	log.Debug("Try download again")
// 	output, err := StartDownload(clientWeb, idFileCopied.Id, outputName)
// 	if err != nil {
// 		return "", err
// 	}

// 	err = gdriveapi.Delete(client, idFolder.Id)
// 	if err != nil {
// 		return "", err
// 	}
// 	return output, err
// }
