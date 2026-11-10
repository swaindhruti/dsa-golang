package main

import "fmt"

/*
 * DFS - DEPTH FIRST SEARCH (on an Adjacency List)
 *
 * WHAT IT IS:
 *   A way to visit every node in a graph.
 *   "Depth first" = go as DEEP as possible down one path first,
 *   and only come back (backtrack) when there is nowhere new to go.
 *
 * INPUT:
 *   adjList = adjacency list, adjList[a] = list of nodes connected to a
 *   start   = the node to start from
 *
 * EXAMPLE (start = 0):
 *
 *     adjList:            graph:
 *     0: [2 3 1]              0
 *     1: [0]                / | \
 *     2: [0 4]             2  3  1
 *     3: [0]               |
 *     4: [2]               4
 *
 *   Order: 0 -> 2 -> 4 -> (back) -> 3 -> (back) -> 1
 *   Result: [0 2 4 3 1]
 *
 * STEPS:
 *   1. Mark the current node as visited and add it to the result.
 *   2. Go through its neighbors one by one.
 *   3. If a neighbor is NOT visited, call DFS on it (go deeper).
 *   4. If it IS visited, skip it (otherwise we loop forever).
 *
 * WHY visited IS []bool:
 *   Nodes are 0 to n-1, so visited[node] tells us in O(1) if we've been there.
 *   (slices.Contains would be O(n) for every check, too slow.)
 *
 * COMPLEXITY (V = number of vertices, E = number of edges):
 *   Time  : O(V + E) -> each node visited once, each edge checked once (twice if undirected)
 *   Space : O(V)     -> visited array + recursion stack (can go V deep)
 *
 * WHEN TO USE:
 *   Find paths, check if nodes are connected, count components, detect cycles,
 *   explore every possibility (backtracking).
 */
func dfs(adjList [][]int, start int) []int {
	visited := make([]bool, len(adjList)) // all false at start
	result := []int{}                     // empty, we append as we visit

	// recursive helper inside the function (closure)
	// declared first with var so it can call itself
	var visit func(node int)
	visit = func(node int) {
		visited[node] = true
		result = append(result, node)

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				visit(neighbor) // go deeper
			}
		}
	}

	visit(start)

	fmt.Println("DFS:", result)
	return result
}

//bfs

func main() {
	adjList := [][]int{
		{2, 3, 1}, // 0
		{0},       // 1
		{0, 4},    // 2
		{0},       // 3
		{2},       // 4
	}

	dfs(adjList, 0)
}
