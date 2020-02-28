package gdriveapi

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
)

func CreateFolder(client *drive.Service, nameFolder string) (*drive.File, error) {

	file := &drive.File{Name: nameFolder, MimeType: "application/vnd.google-apps.folder"}

	resp, err := client.Files.Create(file).Do()
	if err != nil {
		log.Error("An error occurred: %v\n", err)
		return nil, err
	}
	log.Debug("Id folder: %+v", resp.Id)
	return resp, nil
}

func CreateACopy(client *drive.Service, fileId string, nameOutput string, folderId []string) (*drive.File, error) {

	file := &drive.File{Name: nameOutput, Parents: folderId}

	resp, err := client.Files.Copy(fileId, file).Do()
	if err != nil {
		log.Debug("An error occurred: %v\n", err)
		log.Error(err)
		return nil, err
	}

	log.Debug("Id file copied: %+v", resp.Id)
	return resp, err
}

func Delete(client *drive.Service, fileIde string) error {
	error := client.Files.Delete(fileIde).Do()
	if error != nil {
		log.Error("An error occurred: %v\n", error)
	}
	return error
}
