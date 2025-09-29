from collections import deque
from grid.main import Grid


def make_grid():
    with open("input_test_simple.txt") as datafile:
        return Grid(datafile.read())


directions = [
    (0, 1),  # up
    (1, 0),  # right
    (0, -1),  # down
    (-1, 0),  # left
]

grid = make_grid()


def find_plot_perimiter(start):
    curr = start
    queue = deque([curr])

    val = curr.val
    perimeter = 0
    area = 0

    while len(queue):
        curr = queue.popleft()
        area += 1

        if not curr.visited:
            curr.visited = True

        for cx, cy in directions:
            next = grid.get_point(x=curr.x + cx, y=curr.y + cy)
            if not next or next.val != val:
                perimeter += 1
            elif next and next.visited:
                continue
            else:
                next.visited = True
                queue.append(next)

    return area, perimeter


# [(0, 1), (0, -1), (-1, 0), (1, 1), (1, -1), (2, 1), (2, -1), (3, 1), (4, 0), (3, -1)]
def is_adjacent(p1, p2) -> bool:
    x1, y1 = p1
    x2, y2 = p2
    if abs(x1 - x2) == 1 and abs(y1 - y2) == 1:
        return False
    return abs(x1 - x2) <= 1 and abs(y1 - y2) <= 1


def find_plot_sides(start):
    curr = start
    queue = deque([curr])

    val = curr.val
    sides = 0
    area = 0

    axises = []

    while len(queue):
        curr = queue.popleft()
        area += 1

        if not curr.visited:
            curr.visited = True

        for cx, cy in directions:
            next = grid.get_point(x=curr.x + cx, y=curr.y + cy)
            # axis = (cx, cy)
            if not next or next.val != val:
                axises.append((cx + curr.x, cy + curr.y))
                sides += 1
            elif next and next.visited:
                continue
            else:
                next.visited = True
                queue.append(next)

    result = []
    rejects = []
    print(axises)
    for ax in set(axises):
        include = True
        for ax2 in result:
            if is_adjacent(ax, ax2):
                rejects.append(ax)
                include = False
        for ax2 in rejects:
            if is_adjacent(ax, ax2):
                include = False

        if include:
            result.append(ax)

    print(start.val, result)
    return area, sides


def part_1():
    price = 0
    for y in range(len(grid.rows)):
        for x in range(len(grid.rows[0])):
            point = grid.rows[y][x]
            if point.visited:
                continue
            area, perimeter = find_plot_perimiter(point)
            price += area * perimeter

    print("total price:", price)


def part_2():
    price = 0
    for y in range(len(grid.rows)):
        for x in range(len(grid.rows[0])):
            point = grid.rows[y][x]
            if point.visited:
                continue
            area, sides = find_plot_sides(point)
            # print("plant", point.val, "area", area, "sides", sides)
            price += area * sides

    print("total price:", price)


part_2()
