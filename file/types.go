package file

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
)

type folderName string

const rootFolder folderName = "core"
const mapfilefolder folderName = "core/mapfiles"
const piecefolder folderName = "core/piecefolders"
const sendfolder folderName = "core/send"
const subFolder4 folderName = "core/recieve"

type PieceInfo struct {
	PieceName string
	PieceType string
	PieceSize int
	Sources   []int
}

type FileInfo struct {
	FileName   string
	FileType   string
	FileSize   int
	FilePieces int
	UniqueID   uuid.UUID
	Pieces     []PieceInfo
}

type FileBasicInfo struct {
	FileName string
	FileType string
	UniqueID uuid.UUID
}

func parseUUID(uuidStr string) uuid.UUID {
	uuidObj, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println("[ERROR] - during parsing a uuid from the string")
		fmt.Println("[" + uuidStr + "]")
		return uuid.Nil
	}
	return uuidObj
}

func ListAvailableFiles() []FileBasicInfo {
	files, err := ioutil.ReadDir(string(mapfilefolder))
	if err != nil {
		fmt.Println("[ERROR] - during reading files available in the node")
		return nil
	}
	var availableFileList []FileBasicInfo = nil
	for _, file := range files {
		filebasicinfo := FileBasicInfo{
			FileName: strings.Split(file.Name(), "_")[0],
			FileType: strings.Split(file.Name(), "_")[1],
			UniqueID: parseUUID(strings.Split(strings.Split(file.Name(), "_")[2], ".")[0]),
		}
		availableFileList = append(availableFileList, filebasicinfo)

	}
	return availableFileList

}
