package game

import (
	"log"
	"os"
	"strconv"
)

type Board struct {
	Round int
	Pieces [15]*Piece
	From *Board
	MovesForward []*Board
}

func CreateForwardBoard(from *Board, pieces [15]*Piece) *Board {
	var board = &Board{
		From: from,
		Pieces: pieces,
	}
	for i := 0; i < len(board.Pieces); i ++ {
		board.Pieces[i].SetNeighbors(board.Pieces)
	}
	return board
}

func CreateStreamBoard(stream string) *Board {
	if len(stream) != 15 {
		log.Fatal("String must be 15 characters")
		os.Exit(1)
		return nil
	}

	var board = &Board{
		Round: 0,
	}

	for i := 0; i < len(stream); i ++ {
		board.Pieces[i] = CreatePiece(i, stream[i] == '1')
	}

	for i := 0; i < len(board.Pieces); i ++ {
		board.Pieces[i].SetNeighbors(board.Pieces)
	}

	return board
}

func (board *Board) SetRound(round int) {
	board.Round = round
}

func (board *Board) GetMovesString() string {
	if board != nil {
		var fromBoard = board.From
		if fromBoard != nil {
			return fromBoard.GetMovesString() + board.GetString()
		}
	}
	return ""
}

func (board *Board) GenerateMovesForwards() {
	for pieceI := 0; pieceI < len(board.Pieces); pieceI ++ {
		if !board.Pieces[pieceI].Pegged {
			var hole = board.Pieces[pieceI]
			for direction := 0; direction < len(hole.Neighbors); direction ++ {
				var neighbor = hole.Neighbors[direction]
				if neighbor != nil && neighbor.Pegged {
					var neighborsNeighbor = neighbor.Neighbors[direction]
					if neighborsNeighbor != nil && neighborsNeighbor.Pegged {
						var newPieces [15]*Piece
						for newPieceI := 0; newPieceI < len(newPieces); newPieceI ++ {
							var toCopyPiece = board.Pieces[newPieceI]
							if toCopyPiece == neighborsNeighbor {
								newPieces[newPieceI] = CreatePiece(newPieceI, false)
							} else if toCopyPiece == neighbor {
								newPieces[newPieceI] = CreatePiece(newPieceI, false)
							} else if toCopyPiece == hole {
								newPieces[newPieceI] = CreatePiece(newPieceI, true)
							} else {
								newPieces[newPieceI] = CreatePiece(newPieceI, toCopyPiece.Pegged)
							}
						}
						var newBoard = CreateForwardBoard(board, newPieces)
						board.MovesForward = append(board.MovesForward, newBoard)
					}
				}
			}
		}
	}
}

func (board *Board) GetPegsLeft() int {
	var count = 0
	for pieceI := 0; pieceI < len(board.Pieces); pieceI ++ {
		if board.Pieces[pieceI].Pegged {
			count ++
		}
	}
	return count
}

func (board *Board) GetString() string {
	var i = 0
	var output = "Round: " + strconv.Itoa(board.Round) + "\n"
	var pegsPerLevel = 1
	for levelI := 0; levelI < 5; levelI ++ {
		for levelSpacingI := 0; levelSpacingI < 5 - pegsPerLevel; levelSpacingI ++ {
			output += " "
		}

		for levelPegI := 0; levelPegI < pegsPerLevel; levelPegI ++ {
			if board.Pieces[i].Pegged {
				output += "*"
			} else {
				output += "-"
			}
			output += " "
			i ++
		}
		output += "\n"
		pegsPerLevel ++
	}
	return output + "\n"
}