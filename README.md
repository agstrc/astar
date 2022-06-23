# astar

An implementation of the astar algorithm for a college assignment. It features a simple
graphical interface made with [Ebitengine](https://ebiten.org/).

## Project goals

First and foremost, this project aims to implement all that is required by the
assignment. In a nutshell, the assignment requires that the agent performs the following
actions:

- Visit all dungeons.
- Collect an item inside each visited dungeon.
- Return back to the starting point.
- Head to the Lost Woods' gate.

In order to perform these actions, the best path must be dynamically computed through use
of the A\* algorithm. The act of moving through the map has a cost. Each terrain type has
its own traversal cost, and the algorithm must take that into consideration when
computing the best path.  
Furthermore, these actions and the cost of the traversal must be visible, as in, *some*
kind of interface had to be implemented. With this in mind, the traversal cost is
displayed in the top left corner of the screen.  

Further details of the assignment may be read [here](./assignment.md).

## Screenshots

![Main map](https://i.imgur.com/9tDISYw.png)
![Dungeon](https://i.imgur.com/TXMW4o1.png)

## Running the project

Simply follow [Ebitengine's installation guide](https://ebiten.org/documents/install.html)
and compile the file.

```sh
go run main.go
```

The map may also be edited by passing a single JSON file as an argument. The JSON schema
follows the file found at `game/plan/default_plan.json`.
