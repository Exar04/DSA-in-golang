package main

import (
	"fmt"
)

func isSafe(maze [][]int, x int, y int, len int)bool{
	if x < len && y < len && maze[x][y] == 1 {
		return true	
	}
	return false
}


func ratMaze(maze [][]int, x int, y int, solution *[5][5]int) bool{

	if x == len(solution) - 1 && y == len(solution)-1{
		solution[x][y] = 1
		return true
	}

	if isSafe(maze, x, y, len(solution)) {
		solution[x][y] = 1
		if ratMaze(maze, x+1, y, solution) {
			return true
		}
		if ratMaze(maze, x, y+1, solution) {
			return true
		}
		solution[x][y] = 0
		return false
	}

	return false
}












// Failed solution 

// func solveMaze(maze [][]int) [5][5]int{
// 	currentPositionX := 0;
// 	currentPositionY := 0;
// 	var solution [5][5]int
// 	checkPositionAndMove(maze, &solution, currentPositionX, currentPositionY);

// 	return solution
// }

// func checkPositionAndMove(maze [][]int, solution *[5][5]int, x int,y int){
// 	visited := 1
// 	// Init condition
// 	// refSol := *solution


// 	if solution[4][4] == 1 {
// 		fmt.Println("hell")
// 		return
// 	}
// 	solution[x][y] = 1

// 	// if maze[0][0] == 1 && solution[0][0] == 0 {
// 	// 	solution[0][0]=1
// 	// 	fmt.Println("init")
// 	// }
// 	// later conditions
// 	if x+1 < 5 && y < 5 && maze[x + 1][y] == 1 && solution[x + 1][y] != visited{
// 		fmt.Println("1")
// 		solution[x+1][y] = visited
// 		checkPositionAndMove(maze, solution, x + 1, y)
// 	}
	
// 	if x < 5 && y+1 < 5 && x >= 0 && y >= 0 && maze[x][y + 1] == 1 && solution[x][y + 1] != visited  {
// 		fmt.Println("2")
// 		solution[x][y+1] = visited
// 		checkPositionAndMove(maze, solution, x, y + 1)
// 	}
		
// 	if x-1 < 5 && y < 5 && x-1 >= 0 && y >= 0 && maze[x - 1][y] == 1 && solution[x - 1][y] != visited{
// 		fmt.Println("3")
// 		solution[x - 1][y] = visited
// 		checkPositionAndMove(maze, solution, x - 1, y)
// 	}
		
// 	if x < 5 && y-1 < 5 && x >= 0 && y-1 >= 0 && maze[x][y - 1] == 1 && solution[x][y - 1] != visited{
// 		fmt.Println("4")
// 		solution[x][y - 1] = visited
// 		checkPositionAndMove(maze, solution, x, y - 1)
// 	}

// 	if maze[x][y+1] ==0 && maze[x + 1][y] ==0 && maze[x - 1][y]==0 && maze[x][y - 1]==0{
// 		solution[x][y] = 2
// 	}
// 	return

// }

func main(){
	maze := [][]int{{1,0,1,1,0},
					{1,1,0,0,0},
					{0,1,1,1,1},
					{0,1,0,1,0},
					{1,1,0,1,1}}
	
	var solution [5][5] int
	ratMaze(maze, 0,0,&solution)	

	for i := 0; i < 5; i++ {
		fmt.Println(solution[i])
	}

}

func niso(arr *[4]int, x int){
	if x < 4 {
		arr[x]= x+1
		niso(arr,x + 1)
	}
}