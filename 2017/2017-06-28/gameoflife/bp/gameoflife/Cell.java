package org.bbelovic.kata.gameoflife;

import java.util.*;
import java.util.stream.Collectors;

import static java.util.Arrays.asList;

/**
 * Created by bbelovic on 6/28/17.
 */
public class Cell {
    private int x;
    private int y;

    public Cell(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public Set<Cell> getAllNeighbours() {
        List<int[]> neighbourCoordinates = asList(new int [] {1, 0},
                new int [] {0, 1},
                new int [] {1, 1},
                new int [] {1, -1},
                new int [] {0, -1},
                new int [] {-1, -1},
                new int [] {-1, 0},
                new int [] {-1, 1}
                );
        return neighbourCoordinates.stream()
                .map(coordinates -> new Cell(x + coordinates[0], y + coordinates[1]))
                .collect(Collectors.toSet());
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Cell cell = (Cell) o;
        return x == cell.x &&
                y == cell.y;
    }

    @Override
    public int hashCode() {
        return Objects.hash(x, y);
    }

    public String toString() {
        return String.format("Cell[x=%d, y=%d]", x, y);
    }
}
