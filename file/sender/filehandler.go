package sender

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"

	"github.com/google/uuid"
)

type folderName string
type FileName string

var copyofPeerlist = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var peerdistrlist []PickDistributionList = nil
var previouschoices []int = nil

const rootFolder folderName = "core"
const mapfilefolder folderName = "core/mapfiles"
const piecefolder folderName = "core/piecefolders"
const sendfolder folderName = "core/send"
const subFolder4 folderName = "core/recieve"

func (fn folderName) MakeFolder() error {
	_, err := os.Stat(string(fn))
	if os.IsNotExist(err) {
		fmt.Println("Creating the folder " + string(fn))
		err := os.Mkdir(string(fn), 0755)
		if err != nil {
			fmt.Println("[ERROR] during creating the folder" + string(fn))
			return err
		}
		fmt.Println("[SUCCESS] - in creating the folder [" + string(fn) + "]")
		return nil
	} else {
		fmt.Println("Folder " + string(fn) + " already exits")
		return nil
	}
}

func InitFolders() {
	rootFolder.MakeFolder()
	mapfilefolder.MakeFolder()
	piecefolder.MakeFolder()
	sendfolder.MakeFolder()
	subFolder4.MakeFolder()
}

func (fn *FileName) HandleFile() {
	file, err := os.Open(string(sendfolder) + "/" + string(*fn))
	if err != nil {
		fmt.Println("[ERROR] - during opening the file :[" + string(*fn) + "]")
		return
	} else {
		fileName := strings.Split(string(*fn), ".")[0]
		fileType := strings.Split(string(*fn), ".")[1]
		fileInfo := ComposeFileInfo(fileName, fileType)
		fmt.Printf("File Name :[%s]\nFile Type :[%s]\nFile Size : [%d]\nFile Pieces: [%d]\nFile uuid : [%v]", fileInfo.FileName, fileInfo.FileType, fileInfo.FileSize, fileInfo.FilePieces, fileInfo.UniqueID)
		fileInfo.SplitAndSave(file)
	}
}

func GetFileSize(fileName string) int {
	file, err := os.Stat(string(sendfolder) + "/" + fileName)
	if err != nil {
		fmt.Println("[ERROR] - during determining the file size")
		return 0
	} else {
		return int(file.Size())
	}
}

func ComposeFileInfo(fileName string, fileType string) *FileInfo {
	uuID := uuid.New()
	return &FileInfo{
		FileName:   fileName,
		FileType:   fileType,
		FileSize:   GetFileSize(fileName + "." + fileType),
		FilePieces: 0,
		UniqueID:   uuID,
		Pieces:     nil,
	}
}

func ComposePieceInfo(pieceName string, pieceSize int) *PieceInfo {
	return &PieceInfo{
		PieceName: pieceName,
		PieceSize: pieceSize,
		Sources:   nil,
	}
}

func (fI *FileInfo) preparePieceDistributionList() {
	peerdistrlist = nil
	// count := int(math.Ceil(float64(fI.FilePieces/len(copyofPeerlist)))) * 2
	count := int(math.Ceil((float64(fI.FilePieces*3))/float64(len(copyofPeerlist)))) + 2
	fmt.Println("Count = ", count)
	for _, peer := range copyofPeerlist {
		peerdistrlist = append(peerdistrlist, PickDistributionList{
			Peer:  peer,
			Count: count,
		})
	}
}

func (fI *FileInfo) AppendPiecesMapList(piece *PieceInfo) {
	fI.Pieces = append(fI.Pieces, *piece)
	// fmt.Println(fI.Pieces)
}

func (pI *PieceInfo) AddSources() {
	previouschoices = nil
	for i := 0; i < 3; i++ {
		// fmt.Println("Calling choice from add sources")
		choice := choosePeer()
		// fmt.Println(choice)
		// fmt.Println(peerdistrlist)
		previouschoices = append(previouschoices, choice)
		// fmt.Println(previouschoices)
		// fmt.Println(choice, previouschoices)
		pI.Sources = append(pI.Sources, choice)
	}
	// fmt.Println(pI.Sources)
	// fmt.Println(peerdistrlist)
}

func choosePeer() int {
	choice := rand.Intn(len(peerdistrlist))
	// fmt.Println("Choice is ", choice+1)
	for i, peer := range peerdistrlist {

		if i == choice {
			if peer.Count == 0 {
				// time.Sleep(1 * time.Second)
				return choosePeer()
			}
			if previouschoices != nil {
				for _, prvpeer := range previouschoices {
					if peer.Peer == prvpeer {
						// fmt.Println(peer.Peer, prvpeer, "previously chosen")
						return choosePeer()
					}
				}
			}
			peerdistrlist[i].Count = peerdistrlist[i].Count - 1
			// fmt.Println("Reduced the count of ", peer.Peer)
			return peer.Peer
		}
	}
	fmt.Println("[ERROR] - outside loop of choose peer")
	return 0
}

func (fI *FileInfo) Save() {
	fileName := fmt.Sprint(string(mapfilefolder) + "/" + fmt.Sprint(fI.FileName, "_", fI.FileType, "_", fI.UniqueID, ".txt"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] - during creating the map file for [" + fI.FileName + "]")
		return
	}
	fileInfoBytes, err := json.Marshal(fI)
	if err != nil {
		fmt.Println("[ERROR] - during marshalling the map file for [" + fI.FileName + "]")
	}
	file.Write(fileInfoBytes)
	file.Close()
	fmt.Println("[SUCCESS] - in saving the map file for [" + fI.FileName + "]")

}
