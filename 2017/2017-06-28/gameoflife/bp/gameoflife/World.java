package org.bbelovic.kata.gameoflife;

import java.util.HashSet;
import java.util.Set;
import java.util.function.Predicate;
import java.util.stream.Collectors;

public class World {

    private Set<Cell> state;

    public World() {
        this.state = new HashSet<>();
    }

    public Set<Cell> getState() {
        return state;
    }

    public void evolve() {
        if (state.isEmpty()) {
            return;
        } else {
            Set<Cell> nextState = new HashSet<>(state);
            ruleNo1And2(nextState);
            ruleNo4(nextState);
            state = nextState;

        }
    }

    private void ruleNo4(Set<Cell> state) {
        Set<Cell> allDeadNeighbours = getAllDeadNeighbours();
        Predicate<Cell> predicate = (cell) -> {
            Set<Cell> neighbours = cell.getAllNeighbours();
            neighbours.retainAll(state);
            return neighbours.size() == 3;

        };
        Set<Cell> resurrected = allDeadNeighbours.stream()
                .filter(cell -> predicate.test(cell))
                .collect(Collectors.toSet());
        state.addAll(resurrected);
    }

    private void ruleNo1And2(Set<Cell> nextState) {
        Predicate<Cell> predicate = (cell) -> {
            Set<Cell> neighbours = cell.getAllNeighbours();
            neighbours.retainAll(nextState);
            return neighbours.size() < 2 || neighbours.size() > 3;

        };
        Set<Cell> toBeKilled = this.state.stream()
                .filter(cell -> predicate.test(cell))
                .collect(Collectors.toSet());
        nextState.removeAll(toBeKilled);
    }

    public void addCell(Cell cell) {
        state.add(cell);
    }

    private Set<Cell> getAllDeadNeighbours() {
        Set<Cell> allNeighbours = new HashSet<>();
        for (Cell each: state) {
            allNeighbours.addAll(each.getAllNeighbours());
        }
        allNeighbours.removeAll(state);
        return allNeighbours;
    }
}
