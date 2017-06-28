package org.bbelovic.kata.gameoflife;

import org.junit.Assert;
import org.junit.Test;

import java.util.HashSet;
import java.util.Set;

/**
 * Created by bbelovic on 6/28/17.
 */
public class GameOfLifeTest {
    @Test
    public void
    dead_world_stays_dead() {
        World world = new World();
        world.evolve();
        Set<Cell> want = new HashSet<>();
        Set<Cell> next = world.getState();
        Assert.assertTrue(next.equals(want));
    }

    @Test
    public void
    dead_cell_should_resurrect_when_it_has_exactly_three_live_neighbours() {
        World world = new World();
        world.addCell(new Cell(1, 1));
        world.addCell(new Cell(1, 2));
        world.addCell(new Cell(2, 1));

        world.evolve();
        Set<Cell> want = new HashSet<>();
        want.add(new Cell(1, 1));
        want.add(new Cell(1, 2));
        want.add(new Cell(2, 1));
        want.add(new Cell(2, 2));
        Set<Cell> next = world.getState();
        Assert.assertEquals(want, next);

    }

    @Test
    public void
    live_cell_should_become_dead_when_it_has_less_than_two_alive_neighbours() {
        World world = new World();
        world.addCell(new Cell(1, 1));
        world.addCell(new Cell(1, 2));

        world.evolve();
        Set<Cell> want = new HashSet<>();
        Set<Cell> next = world.getState();
        Assert.assertEquals(want, next);

    }

    @Test
    public void
    live_cell_should_become_dead_when_it_has_more_than_three_alive_neighbours() {
        World world = new World();
        world.addCell(new Cell(1, 1));
        world.addCell(new Cell(1, 2));
        world.addCell(new Cell(2, 1));
        world.addCell(new Cell(2, 2));
        world.addCell(new Cell(3, 1));

        world.evolve();
        Set<Cell> want = new HashSet<>();
        want.add(new Cell(1, 1));
        want.add(new Cell(1, 2));
        want.add(new Cell(3, 1));
        Set<Cell> next = world.getState();
        Assert.assertEquals(want, next);

    }
}
