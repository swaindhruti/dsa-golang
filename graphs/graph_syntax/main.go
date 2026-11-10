package main

import "fmt"

/*
 * ADJACENCY MATRIX (Undirected + Unweighted Graph)
 *
 * WHAT IT IS:
 *   A table (grid) of size n x n.
 *   Rows = vertices, columns = vertices.
 *   matrix[a][b] = 1  -> a and b are connected
 *   matrix[a][b] = 0  -> not connected
 *
 * INPUT:
 *   n     = how many vertices (named 0, 1, 2 ... n-1)
 *   edges = list of connections, each one is [from, to]
 *           e.g. [[0,1], [1,2], [1,3]]
 *
 * EXAMPLE (n = 4, edges = [[0,1], [1,2], [1,3]]):
 *
 *          0  1  2  3
 *     0: [ 0  1  0  0 ]   -> 0 is connected to 1
 *     1: [ 1  0  1  1 ]   -> 1 is connected to 0, 2, 3
 *     2: [ 0  1  0  0 ]   -> 2 is connected to 1
 *     3: [ 0  1  0  0 ]   -> 3 is connected to 1
 *
 * STEPS:
 *   1. Make an n x n grid filled with 0.
 *   2. For every edge [a, b], set matrix[a][b] = 1 AND matrix[b][a] = 1
 *      (undirected = the connection works both ways).
 *
 * WHY EACH ROW HAS SIZE n:
 *   Each vertex can connect to any of the n vertices,
 *   so every row needs one spot for each vertex.
 *
 * COMPLEXITY (V = number of vertices, E = number of edges):
 *   Build              : O(V^2 + E)  -> make the full grid, then fill the edges
 *   Space              : O(V^2)      -> the grid always has V x V cells
 *   Is a connected to b?: O(1)       -> just look at matrix[a][b]
 *   All neighbors of a : O(V)        -> have to check the whole row a
 *
 * WHEN TO USE:
 *   Graph is small or has LOTS of edges, or you often ask "are a and b connected?"
 */
func adjacencyMatrix(n int, edges [][]int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n) // n columns, one per vertex
	}

	for _, e := range edges {
		src, dest := e[0], e[1]
		matrix[src][dest] = 1
		matrix[dest][src] = 1
	}

	// print matrix
	fmt.Println("Adjacency Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// print edges
	for i := range matrix {
		for j := i + 1; j < n; j++ { // j > i so each edge prints once and to ignore the diagonals
			if matrix[i][j] == 1 {
				fmt.Printf("%d is connected to %d\n", i, j)
			}
		}
	}

	return matrix
}

/*
 * ADJACENCY LIST (Undirected + Unweighted Graph)
 *
 * WHAT IT IS:
 *   An array of size n. Each index is a vertex.
 *   At each index we keep a list of the vertices it is connected to.
 *   Only real connections are stored, no 0s like in the matrix.
 *
 * INPUT: same as the matrix (n, edges)
 *
 * EXAMPLE (n = 4, edges = [[0,1], [1,2], [1,3]]):
 *
 *     0: [1]         -> 0 is connected to 1
 *     1: [0 2 3]     -> 1 is connected to 0, 2, 3
 *     2: [1]         -> 2 is connected to 1
 *     3: [1]         -> 3 is connected to 1
 *
 * STEPS:
 *   1. Make an array of n empty lists.
 *   2. For every edge [a, b], add b to a's list AND a to b's list
 *      (undirected = the connection works both ways).
 *
 * COMPLEXITY (V = number of vertices, E = number of edges):
 *   degree(a) = how many neighbors a has
 *   Build              : O(V + E)       -> one list per vertex, one add per edge side
 *   Space              : O(V + E)       -> only stores real connections
 *   Is a connected to b?: O(degree(a))  -> have to search through a's list
 *   All neighbors of a : O(degree(a))   -> they're already in adjList[a]
 *
 * WHEN TO USE:
 *   Most of the time! Best for big graphs with FEW edges, and for BFS / DFS.
 */
func adjacencyList(n int, edges [][]int) [][]int {
	adjList := make([][]int, n)

	for _, e := range edges {
		src, dest := e[0], e[1]
		adjList[src] = append(adjList[src], dest)
		adjList[dest] = append(adjList[dest], src)
	}

	// print list
	fmt.Println("Adjacency List:")
	for i, neighbors := range adjList {
		fmt.Printf("%d: %v\n", i, neighbors)
	}

	return adjList
}

func main() {
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}

	adjacencyMatrix(n, edges)
	fmt.Println()
	adjacencyList(n, edges)
}
