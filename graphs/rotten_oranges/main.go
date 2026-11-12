/*
Rotting Oranges

Problem:
	Given an m x n grid where each cell is:
		0 - empty cell
		1 - fresh orange
		2 - rotten orange
	Every minute, a rotten orange rots its fresh neighbours (up, down, left, right).
	Return the minimum minutes until no fresh orange is left.
	If some fresh orange can never rot, return -1.

Idea (multi-source BFS):
	- All rotten oranges start rotting at the SAME time, so put all of them
	  in the queue at the start (many sources, not one).
	- Process the queue level by level. One level = one minute.
	- Each fresh neighbour of a rotten orange becomes rotten and joins the queue
	  for the next minute. Decrease freshCount every time an orange rots.
	- At the end, if freshCount > 0, some oranges were unreachable -> -1.

Why BFS and not DFS?
	BFS spreads out one layer at a time, exactly like the rot spreads
	one step per minute, so the number of levels = minutes needed.

Time:  O(m * n)  - each cell enters the queue at most once.
Space: O(m * n)  - visited matrix + queue.
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
	return i >= 0 && i < m && j >= 0 && j < n
}

/*
bfs spreads the rot from all rotten oranges and returns the minutes taken.

Input:
	grid       - the m x n grid of 0 / 1 / 2
	visited    - visited[r][c] is true if cell (r, c) is already rotten
	queue      - all rotten oranges at minute 0, as {row, col}
	freshCount - pointer to the number of fresh oranges; reduced as they rot
	m          - number of rows
	n          - number of columns

How it works:
	1. While the queue is not empty, take the current level (size := len(queue)).
	   Every orange in this level rotted in the same minute.
	2. For each orange, look at the 4 neighbours using xaxis/yaxis:
	       k=0: (-1, 0) up
	       k=1: ( 0, 1) right
	       k=2: ( 1, 0) down
	       k=3: ( 0,-1) left
	3. If a neighbour is inside the grid, fresh (1), and not visited,
	   mark it visited, add it to the queue, and freshCount--.
	4. If at least one orange rotted in this level, minutes++.
	   (The last level rots nothing, so it does not add a minute.)
*/
func bfs(grid [][]int, visited [][]bool, queue [][]int, freshCount *int, m, n int) int {
	xaxis := []int{-1, 0, 1, 0}
	yaxis := []int{0, 1, 0, -1}

	minutes := 0

	for len(queue) > 0 {
		size := len(queue)
		rotted := false

		for range size {
			cell := queue[0]
			queue = queue[1:]

			for k := range 4 {
				row := cell[0] + xaxis[k]
				col := cell[1] + yaxis[k]

				if valid(row, col, m, n) && grid[row][col] == 1 && !visited[row][col] {
					visited[row][col] = true
					queue = append(queue, []int{row, col})
					*freshCount--
					rotted = true
				}
			}
		}

		if rotted {
			minutes++
		}
	}

	return minutes
}

/*
orangesRotting returns the minimum minutes until no fresh orange is left,
or -1 if that is impossible.

Input:
	grid - m x n matrix, 0 = empty, 1 = fresh, 2 = rotten

How it works:
	1. Create an m x n visited matrix (all false).
	2. Loop over every cell:
	   - rotten (2): add to the queue and mark visited (BFS starting points).
	   - fresh  (1): freshCount++.
	3. If there are no fresh oranges, return 0.
	4. Run bfs to spread the rot and get the minutes.
	5. If freshCount > 0, some orange was never reached -> return -1.
	   Otherwise return the minutes.
*/
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := [][]int{}
	freshCount := 0

	for i := range m {
		for j := range n {
			switch grid[i][j] {
			case 2:
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			case 1:
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	minutes := bfs(grid, visited, queue, &freshCount, m, n)

	if freshCount > 0 {
		return -1
	}

	return minutes
}

/*
main runs orangesRotting on sample grids.

Example 1:
	2 1 1
	1 1 0
	0 1 1
	minute 1: (0,1), (1,0) rot
	minute 2: (0,2), (1,1) rot
	minute 3: (2,1) rots
	minute 4: (2,2) rots
	Expected output: 4

Example 2:
	2 1 1
	0 1 1
	1 0 1
	(2,0) has no neighbour that can rot it.
	Expected output: -1

Example 3:
	0 2
	No fresh oranges.
	Expected output: 0
*/
func main() {
	fmt.Println("Minutes: ", orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))

	fmt.Println("Minutes: ", orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}))

	fmt.Println("Minutes: ", orangesRotting([][]int{
		{0, 2},
	}))
}
