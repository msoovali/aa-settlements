package googledrive

import (
	"fmt"

	"google.golang.org/api/drive/v3"
)

type search struct {
	fileName string
	parentId string
	fileType mimeType
	call     *drive.FilesListCall
}

func (g *googleDriveGateway) initSearch(fileName, parentId string, fileType mimeType) search {
	return search{
		fileName: fileName,
		parentId: parentId,
		fileType: fileType,
		call:     g.service.Files.List(),
	}
}

func (s search) find() (*drive.File, error) {
	r, err := s.call.Q(fmt.Sprintf("name='%s' and mimeType='%s' and '%s' in parents", s.fileName, s.fileType, s.parentId)).Do()
	if err != nil {
		return nil, err
	}
	if len(r.Files) > 0 {
		return r.Files[0], nil
	}
	return nil, nil
}
