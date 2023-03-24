package sender

import "github.com/google/uuid"

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

type PickDistributionList struct {
	Peer  int
	Count int
}
