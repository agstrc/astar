package astar

import "github.com/agstrc/heuristic-search/pqueue"

// Node refers to a type capable of returning its own neighbours and its associated
// traversal cost.
//
// Any implementations must also be comparable as when searching for a given path, the
// goal node will be compared to the traversed nodes. Furthermore, the nodes are also
// used as map keys. Therefore, when calling FindPath, these details must be taken into
// consideration before asserting that the type which implements this is comparable.
type Node[T any] interface {
	comparable

	Neighbors() []T
	Cost() int
}

// Heuristic is a heuristic function used in the A* algorithm implementation. Its return
// will be used when defining a a node's priority during the search. The higher the
// returned value, the more likely the node is to be traversed next.
type Heuristic[T any] func(from T, to T) int

// FindPath implements the A* algorithm to find a path from start to goal. The returned
// slice is the computed path. If no path is found, a nil slice is returned
func FindPath[N Node[N]](start N, goal N, heuristic Heuristic[N]) ([]N, int) {
	var frontier pqueue.PriorityQueue[N]
	frontier.Push(start, 0)

	costTo := map[N]int{start: 0}
	cameFrom := map[N]*N{
		start: nil,
	}

	for !frontier.Empty() {
		currentNode := frontier.Pop()

		if currentNode == goal {
			// build a slice, starting from cameFrom[goal], that specifies the reverse
			// path (from goal to start)
			path := []N{goal}
			previousNode := cameFrom[goal]
			for {
				if previousNode == nil {
					break
				}
				path = append(path, *previousNode)
				previousNode = cameFrom[*previousNode]
			}
			// reverse the slice, therefore making the slice point from start to goal
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path, costTo[goal]
		}

		for _, next := range currentNode.Neighbors() {
			costToNext := costTo[currentNode] + next.Cost()
			previousCostToNext, isNextVisited := costTo[next]

			if !isNextVisited || costToNext < previousCostToNext {
				costTo[next] = costToNext
				// as "cost" represents the traversal cost, the higher the cost, the
				// lower its priority should be. Therefore, it is turned into a negative
				// so the higher costs have a lower priority
				priority := (costToNext * (-1)) + heuristic(next, goal)
				frontier.Push(next, priority)
				cameFrom[next] = &currentNode
			}
		}
	}

	return nil, 0
}
