package tictactoe

type TicTacToe struct {
	board [3][3]string
	turn  string
}

func NewGame() *TicTacToe {
	return &TicTacToe{
		board: [3][3]string{},
		turn:  "X",
	}
}

func (t *TicTacToe) MakeMove(x, y int) bool {
	if t.board[x][y] == "" {
		t.board[x][y] = t.turn
		t.switchTurn()
		return true
	}
	return false
}

func (t *TicTacToe) switchTurn() {
	if t.turn == "X" {
		t.turn = "O"
	} else {
		t.turn = "X"
	}
}

func (t *TicTacToe) CheckWinner() string {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if t.board[i][0] != "" && t.board[i][0] == t.board[i][1] && t.board[i][1] == t.board[i][2] {
			return t.board[i][0]
		}
		if t.board[0][i] != "" && t.board[0][i] == t.board[1][i] && t.board[1][i] == t.board[2][i] {
			return t.board[0][i]
		}
	}
	// Check diagonals
	if t.board[0][0] != "" && t.board[0][0] == t.board[1][1] && t.board[1][1] == t.board[2][2] {
		return t.board[0][0]
	}
	if t.board[0][2] != "" && t.board[0][2] == t.board[1][1] && t.board[1][1] == t.board[2][0] {
		return t.board[0][2]
	}
	return ""
}

func (t *TicTacToe) IsDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.board[i][j] == "" {
				return false
			}
		}
	}
	return t.CheckWinner() == ""
}

func (t *TicTacToe) Board() [3][3]string {
	return t.board
}
