package main

import (
	"fmt"
	"github.com/josephnormandev/pegs/game"
)

func main() {
	fmt.Println("Enter a Board: (15 characters long)")
	fmt.Println("...............")

	var stream string
	fmt.Scanln(&stream)

	var startBoard = game.CreateStreamBoard(stream)
	fmt.Println()
	fmt.Println("Results........")
	fmt.Print(startBoard.GetString())

	/*
	startBoard.GenerateMovesForwards()
	for i := 0; i < len(startBoard.MovesForward); i ++ {
		//fmt.Println(startBoard.MovesForward[i].GetString())
	}
	return*/

	var endBoards = runGames(startBoard)
	if len(endBoards) > 0 {
		var selectedWinner = endBoards[0]
		fmt.Println(selectedWinner.GetMovesString())
	}
}

func runGames(startBoard *game.Board) []*game.Board {
	var roundBoards = [][]*game.Board{
		{ startBoard },
	}
	var roundI int
	for roundI = 0; roundI < startBoard.GetPegsLeft() ; roundI ++ {
		if len(roundBoards[roundI]) == 0 {
			break
		}

		roundBoards = append(roundBoards, []*game.Board{}) // round + 1
		for boardI := 0; boardI < len(roundBoards[roundI]); boardI ++ {
			var currentBoard = roundBoards[roundI][boardI]
			currentBoard.SetRound(roundI)
			currentBoard.GenerateMovesForwards()
			var resultBoards = currentBoard.MovesForward

			for resultI := 0; resultI < len(resultBoards); resultI ++ {
				var resultBoard = resultBoards[resultI]
				roundBoards[roundI+1] = append(roundBoards[roundI+1], resultBoard)
			}
		}
	}
	return roundBoards[roundI - 1]
}