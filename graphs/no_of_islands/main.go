/*
Number of Islands

Problem:
	Given an m x n grid of '1' (land) and '0' (water), count the islands.
	An island is a group of '1's connected up, down, left or right
	(diagonal cells do NOT count as connected).

Idea:
	- Scan every cell of the grid.
	- When we find a '1' that is not visited yet, it is the start of a NEW island:
	  increase count, then run DFS from it.
	- DFS visits every land cell connected to it and marks them visited,
	  so the rest of that island is never counted again.
	- Final count = number of times we started a DFS.

Time:  O(m * n)  - each cell is visited at most once.
Space: O(m * n)  - visited matrix + recursion stack (worst case all land).
*/

package main

import "fmt"

/*
valid checks if cell (i, j) lies inside the grid.

Input:
	i, j - row and column of the cell to check
	m    - number of rows
	n    - number of columns

Returns true if 0 <= i < m and 0 <= j < n, else false.
*/
func valid(i, j, m, n int) bool {
	if i < 0 || j < 0 || i >= m || j >= n {
		return false
	}
	return true
}

/*
dfs marks every land cell connected to (i, j) as visited.

Input:
	grid    - the map of '1' (land) and '0' (water)
	visited - visited[r][c] is true if cell (r, c) is already explored
	i, j    - current cell (row, column), always a land cell
	n       - number of columns
	m       - number of rows

How it works:
	1. Mark (i, j) as visited.
	2. Look at the 4 neighbours using xaxis/yaxis:
	       k=0: (-1, 0) up
	       k=1: ( 0, 1) right
	       k=2: ( 1, 0) down
	       k=3: ( 0,-1) left
	3. If a neighbour is inside the grid, is land, and not visited,
	   call dfs on it (go deeper). When a cell has no new neighbour,
	   the recursion returns back to the previous cell (backtracking).
*/
func dfs(grid [][]byte, visited [][]bool, i, j, n, m int) {
	visited[i][j] = true

	xaxis := []int{-1, 0, 1, 0}
	yaxis := []int{0, 1, 0, -1}

	for k := range 4 {

		row := i + xaxis[k]
		col := j + yaxis[k]

		if valid(row, col, m, n) && grid[row][col] == '1' && !visited[row][col] {
			dfs(grid, visited, row, col, n, m)
		}
	}
}

/*
numIslands returns the number of islands in the grid.

Input:
	grid - m x n matrix of bytes, '1' = land, '0' = water

How it works:
	1. Create an m x n visited matrix (all false).
	2. Loop over every cell (i, j).
	3. If the cell is '1' and not visited, a new island is found:
	   count++ and call dfs to mark the whole island as visited.
	4. Return count.
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	count := 0

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				dfs(grid, visited, i, j, n, m)
			}
		}
	}

	return count
}

/*
main runs numIslands on a sample grid.

	1 1 0 0 0
	1 1 0 0 0
	0 0 1 0 0
	0 0 0 1 1

Islands: top-left block, the single cell at (2,2), and (3,3)-(3,4).
Expected output: 3
*/
func main() {
	fmt.Println("Number of islands: ", numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}
