package game

type Piece struct {
	Index int
	Pegged bool
	Neighbors [6]*Piece
}

func CreatePiece(index int, pegged bool) *Piece {
	return &Piece{
		Index: index,
		Pegged: pegged,
	}
}

func (piece *Piece) SetNeighbors(everybody [15]*Piece) {
	var indexes = [15][6]int{
		{ -1, -1, -1, -1, 2, 1 }, // 0
		{ -1, -1, 0, 2, 4, 3}, // 1
		{ 1, 0, -1, -1, 5, 4}, // 2
		{ -1, -1, 1, 4, 7, 6 }, // 3
		{ 3, 1, 2, 5, 8, 7 }, // 4
		{ 4, 2, -1, -1, 9, 8 }, // 5
		{ -1, -1, 3, 7, 11, 10 }, // 6
		{ 6, 3, 4, 8, 12, 11 }, // 7
		{ 7, 4, 5, 9, 13, 12 }, // 8
		{ 8, 5, -1, -1, 14, 13 }, // 9
		{ -1, -1, 6, 11, -1, -1 }, // 10
		{ 10, 6, 7, 12, -1, -1 }, // 11
		{ 11, 7, 8, 13, -1, -1 }, // 12
		{ 12, 8, 9, 14, -1, -1 }, // 13
		{ 13, 9, -1, -1, -1, -1 }, // 14
	}

	var neighborIndexes = indexes[piece.Index]
	for neighborI := 0; neighborI < len(neighborIndexes); neighborI ++ {
		if neighborIndexes[neighborI] != -1 {
			piece.Neighbors[neighborI] = everybody[neighborIndexes[neighborI]]
		} else {
			piece.Neighbors[neighborI] = nil
		}
	}
}