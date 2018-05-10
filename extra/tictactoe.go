package extra

import (
	"math/rand"
	"strconv"
	"fmt"
)

var board = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}

const X int = 11
const O int = 12

var moves int = 0
var myTurn bool = true

var winner int

var lastX, lastY int

func main() {
	fmt.Println("Let's play Tic Tac Toe!")
	fmt.Println("I'm X, you're O.")
	fmt.Println("Here's the board.")
	printBoard()
	fmt.Println("All positions marked with a number are empty, so input the number of the box you want to put O into when it is your turn.")


	for isGameOver() == false {
		myTurn = !myTurn
		if myTurn {
			fmt.Println("It's my turn.")
			computerPicksPosition()
		} else {
			fmt.Println("It's your turn.")
			userPicksPosition()
		}
		printBoard()
		moves++
	}

	gameOutcome()
}

func isGameOver() bool {
	winConditionSatisfied := false

	if board[0][lastY] == board[1][lastY] && board[1][lastY] == board[2][lastY]{ // Vertical
		winConditionSatisfied = true
	}
	if board[lastX][0] == board[lastX][1] && board[lastX][1] == board[lastX][2] { // Horizontal
		winConditionSatisfied = true
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] { // Diagonal 1
		winConditionSatisfied = true
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] { // Diagonal 2
		winConditionSatisfied = true
	}

	if winConditionSatisfied {
		if myTurn {
			winner = 11
		} else {
			winner = 12
		}
		return true
	}

	if moves == 9 {
		return true
	}
	return false
}


func gameOutcome() {
	switch winner {
	case 11:
		fmt.Println("I won! :D")
	case 12:
		fmt.Println("You won. Drats! :/")
	default:
		fmt.Println("I guess it was a draw. :(")
	}
}

func computerPicksPosition() {
	lastX = rand.Int() % 3
	lastY = rand.Int() % 3

	if isPositionTaken(lastX, lastY) {
		computerPicksPosition()
		return
	}
	board[lastX][lastY] = X

}

func userPicksPosition() {
	fmt.Print("Pick a position to move: ")
	var input = -1
	fmt.Scanf("%d", &input)

	if input < 1 && input > 9 {
		fmt.Print("That is an invalid position. Try again.")
		userPicksPosition()
		return
	}

	lastX = (input - 1) / 3
	lastY = (input - 1) % 3

	if isPositionTaken(lastX,lastY) {
		fmt.Print("That is an invalid move. Try again.")
		userPicksPosition()
		return
	}
	board[lastX][lastY] = O
}

func isPositionTaken(x, y int) bool {
	if getEntry(x,y) == "X" || getEntry(x,y) == "O" {
		return true
	}
	return false
}

func getEntry(x, y int) string {
	if board[x][y] == X {
		return "X"
	}
	if board[x][y] == O {
		return "O"
	}
	return strconv.Itoa(board[x][y])
}

func printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				fmt.Println(getEntry(i, j))
				if i < 2 {
					fmt.Println("--+---+--")
				}
			} else {
				fmt.Print(getEntry(i,j), " | ")
			}
		}
	}
}
