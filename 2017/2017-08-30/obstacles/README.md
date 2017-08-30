# Exercise

https://gist.github.com/juanvallejo/890fbd971988b52c6f83a185df09fbf2


## Given

a two-dimensional matrix `b`, a list of obstacle points `ob`, a list of
waypoints `waypoints`,  with the starting waypoint at `waypoints[0]` and the
ending point as `waypoints[N]`(where `N` is the last waypoint).

## Find

the shortest path while visiting all given `waypoints` *in order*.

## Definitions

```c
typedef struct point2d {int x, y}

struct matrix {
    int maxX;
    int maxY;
};

struct obstacle {
    point2d point;
    int padding; // obstacle padding to avoid. Adheres to following rules:
                 // 1. Diagonal movements are allowed
                 // 2. all obstacles are circles
                 // 3. this property is measured by euclidean distance
}

point2d[] findPath(matrix b, point2d[] waypoints, obstacle[] ob);
```

## Clarifications

- The matrix is the grid defined by `<0, 0>` (inclusive) and `<maxX, maxY>` (exclusive)
- `obstacle.padding` is the radius of the obstacle centered at `obstacle.point`.
