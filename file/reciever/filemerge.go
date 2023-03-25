package reciever

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetFileSize(filePath string) int {
	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("[ERROR] - during determining the file size")
		return 0
	} else {
		return int(file.Size())
	}
}

func (fbi FileBasicInfo) RetreiveFullInfo() *FileInfo {
	filepath := string(mapfilefolder) + "/" + fbi.FileName + "_" + fbi.FileType + "_" + fbi.UniqueID.String() + ".txt"
	bufferSize := GetFileSize(filepath)
	buffer := make([]byte, bufferSize)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("[ERROR] - during trying to open the file map list")
		fmt.Println("[PATH : " + filepath + "]")
		return nil
	}
	file.Read(buffer)
	fileInfo := &FileInfo{}
	err = json.Unmarshal(buffer, fileInfo)
	if err != nil {
		fmt.Println("[ERROR] - during unmarhsalling the file map list")
		return nil
	}
	fmt.Println("====================================================")
	fmt.Println(fileInfo.FileName)
	fmt.Println(fileInfo.FileType)
	fmt.Println(fileInfo.FileSize)
	fmt.Println(fileInfo.FilePieces)
	fmt.Println(fileInfo.UniqueID)
	fmt.Println("----------------------------------------------------")
	for _, piece := range fileInfo.Pieces {
		fmt.Println(piece)
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println("====================================================")
	return fileInfo
}

func (p *PieceInfo) GetPieceBytes(fI *FileInfo) []byte {
	piecePath := string(piecefolder) + "/" + fI.FileName + "_" + fI.FileType + "_" + fI.UniqueID.String() + "/" + p.PieceName
	file, err := os.Open(piecePath)
	if err != nil {
		fmt.Println("[ERROR] - during opening the piece ", p.PieceName)
		fmt.Println("[PATH:" + piecePath + "]")
		return nil
	}
	pieceBuffer := make([]byte, p.PieceSize)
	file.Read(pieceBuffer)
	return pieceBuffer
}

func (fI *FileInfo) Merge() {
	destinationFilePath := string(recieverfolder) + "/" + string(fI.FileName) + "." + string(fI.FileType)
	file, err := os.Create(destinationFilePath)
	if err != nil {
		fmt.Println("[ERROR] - during creating the destination file")
		fmt.Println("[PATH :" + destinationFilePath + "]")
		return
	}
	for i := 0; i < fI.FilePieces; i++ {
		file.Write(fI.Pieces[i].GetPieceBytes(fI))
	}
	fmt.Println("[SUCCESS] - in reassembling the file")
	file.Close()
}
