package main

import (
	"fmt"
)

//slices可以包含任何类型，包括其他的slices
func main() {
	//创建一个tic-tac-toe板
	//玩家默认修改几个值，输出修改结果

	//思路：建立一个二维的slice，修改具体的值，输出
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "x"
	board[1][1] = "x"
	board[2][2] = "x"

	fmt.Println(board)
	for i := 0; i < len(board); i++ {
		fmt.Printf()
	}
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Create a tic-tac-toe board.
// 	board := [][]string{
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 		[]string{"_", "_", "_"},
// 	}

// 	// The players take turns.
// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "O"
// 	board[0][2] = "X"

// 	for i := 0; i < len(board); i++ {
// 		fmt.Printf("%s\n", strings.Join(board[i], " "))
// 	}
// }
