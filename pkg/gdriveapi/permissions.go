package gdriveapi

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
)

func MakeSharedFolder(client *drive.Service, idFolder string) (*drive.Permission, error) {
	file := &drive.Permission{Role: "reader", Type: "anyone"}
	resp, err := client.Permissions.Create(idFolder, file).Do()
	if err != nil {
		log.Debug("An error occurred: %v\n", err)
		return nil, err
	}
	log.Debug("Type shared folder: %+v", resp.Id)
	return resp, nil
}
