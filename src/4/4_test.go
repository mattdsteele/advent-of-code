package main

import (
	"testing"

	util "github.com/mattdsteele/advent-of-code"
	tst "github.com/mattdsteele/advent-of-code/testing"
)

var data string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

var lines []string = util.SliceAtLine(data)

func TestParse(t *testing.T) {
	game := parse(lines)
	tst.Equals(t, 7, game.draws[0])
	tst.Equals(t, 1, game.draws[len(game.draws)-1])
	tst.Equals(t, 3, len(game.boards))
	board := game.boards[0]
	tst.Equals(t, 5, len(board.rows))
	tst.Equals(t, 22, board.rows[0].cols[0].value)
	tst.Equals(t, false, board.rows[0].cols[0].selected)
}

func TestMark(t *testing.T) {
	game := parse(lines)
	board := game.boards[0]
	board.mark(23)
	tst.Equals(t, true, board.rows[1].cols[2].selected)
	tst.Equals(t, false, board.rows[1].cols[3].selected)
}

func TestBingo(t *testing.T) {
	game := parse(lines)
	board := game.boards[0]
	board.mark(0)
	board.mark(11)
	board.mark(17)
	board.mark(13)
	board.mark(15)
	tst.Equals(t, false, board.solved())
	board.mark(22)
	tst.Equals(t, true, board.solved())
}

func TestBingoVertical(t *testing.T) {
	game := parse(lines)
	board := game.boards[0]
	// 22 13 17 11  0
	//  8  2 23  4 24
	// 21  9 14 16  7
	//  6 10  3 18  5
	//  1 12 20 15 19
	board.mark(11)
	board.mark(4)
	board.mark(16)
	board.mark(18)
	tst.Equals(t, false, board.solved())
	board.mark(15)
	tst.Equals(t, true, board.solved())
}
