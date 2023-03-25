package reciever

import (
	"github.com/google/uuid"
)

type folderName string

const rootFolder folderName = "core"
const mapfilefolder folderName = "core/mapfiles"
const piecefolder folderName = "core/piecefolders"
const sendfolder folderName = "core/send"
const recieverfolder folderName = "core/recieve"

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
