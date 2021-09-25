from enum import Enum
from typing import List, NamedTuple, Callable, Optional
import random
from math import sqrt

import importlib.util

spec = importlib.util.spec_from_file_location("generic_search", "../../search/common/generic_search.py")
foo = importlib.util.module_from_spec(spec)
spec.loader.exec_module(foo)


# from generic_search import dfs, bfs, node_to_path, aster, Node


class Cell(str, Enum):
    EMPTY = " "
    BLOCKED = "X"
    START = "S"
    GOAL = "G"
    PATH = "*"


class MazeLocation(NamedTuple):
    row: int
    column: int


class Maze:
    def __init__(self, rows: int = 10, columns: int = 10, sparseness: float = 0.2,
                 start: MazeLocation = MazeLocation(0, 0), goal: MazeLocation = MazeLocation(9, 9)) -> None:
        # initialize base variable
        self._rows: int = rows
        self._columns: int = columns
        self.start: MazeLocation = start
        self.goal: MazeLocation = goal

        # set empty cells
        self._grid: List[List[Cell]] = [[Cell.EMPTY for c in range(columns)] for r in range(rows)]

        # set blocked cells
        self._randomly_fill(rows, columns, sparseness)

        # set begin and end position in labyrinth
        self._grid[start.row][start.column] = Cell.START
        self._grid[goal.row][goal.column] = Cell.GOAL

    def _randomly_fill(self, rows: int, columns: int, sparseness: float):
        for row in range(rows):
            for column in range(columns):
                if random.uniform(0, 1.0) < sparseness:
                    self._grid[row][column] = Cell.BLOCKED

    def __str__(self) -> str:
        output: str = ""
        for row in self._grid:
            output += "".join([c.value for c in row]) + "\n"
        return output


if __name__ == "__main__":
    maze: Maze = Maze()
    print(Maze)
