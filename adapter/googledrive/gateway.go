package googledrive

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

type googleDriveGateway struct {
	service      *drive.Service
	rootFolderId string
}

func NewDriveGateway(service *drive.Service, rootFolderId string) *googleDriveGateway {
	return &googleDriveGateway{
		service:      service,
		rootFolderId: rootFolderId,
	}
}

func (g *googleDriveGateway) CopySpreadsheet(target string, destination string) (string, error) {
	file, err := g.initSearch(target, g.rootFolderId, Spreadsheet).find()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve spreadsheet: %v", err)
	}
	if file == nil {
		return "", fmt.Errorf("spreadsheet not found")
	}

	return g.CopySpreadsheetFromId(file.Id, destination)
}

func (g *googleDriveGateway) CopySpreadsheetFromId(targetId string, destinationName string) (string, error) {
	newFile, err := g.service.Files.Copy(targetId, &drive.File{
		Name: destinationName,
	}).Do()
	if err != nil {
		return "", fmt.Errorf("copy error: %v", err)
	}

	return newFile.Id, nil
}

func (g *googleDriveGateway) CreateFolder(name string, parentFolderId string) (string, error) {
	existingFolder, err := g.initSearch(name, parentFolderId, Folder).find()
	if err != nil {
		return "", fmt.Errorf("failed to determine if folder %s already exists: %v", name, err)
	}
	if existingFolder != nil {
		log.Printf("Folder %s already existis in parent %s. Skipping creation.", name, parentFolderId)
		return existingFolder.Id, nil
	}
	log.Printf("Creating folder %s to parent folder %s", name, parentFolderId)
	folder := &drive.File{
		Name:     name,
		MimeType: string(Folder),
		Parents:  []string{parentFolderId},
	}
	createdFolder, err := g.service.Files.Create(folder).Do()
	if err != nil {
		return "", err
	}

	return createdFolder.Id, nil
}

func (g *googleDriveGateway) MoveSpreadsheetToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error) {
	file, err := g.initSearch(fileName, sourceFolderId, Spreadsheet).find()
	if err != nil {
		return "", nil
	}

	return g.moveToFolder(file, targetFolderId)
}

func (g *googleDriveGateway) MoveFolderToAnotherFolder(fileName, sourceFolderId, targetFolderId string) (string, error) {
	file, err := g.initSearch(fileName, sourceFolderId, Folder).find()
	if err != nil {
		return "", nil
	}

	return g.moveToFolder(file, targetFolderId)
}

func (g *googleDriveGateway) moveToFolder(file *drive.File, targetFolderId string) (string, error) {
	file.Parents = []string{targetFolderId}
	file, err := g.service.Files.Update(file.Id, file).Do()
	if err != nil {
		return "", err
	}

	return file.Id, nil
}
