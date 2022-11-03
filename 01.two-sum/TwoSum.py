
def two_sum(arr: list[int], sum: int) -> tuple[int, int]:
    num_index = {}
    for i, value in enumerate(arr):
        if sum - value in num_index:
            return num_index[sum - value], i
        num_index[value] = i
    return -1, -1


def solve_problem(arr: list[int], sum: int) -> None:
    i, j = two_sum(arr, sum)
    print(f"array: {arr}, sum: {sum}, indexes: [{i}, {j}], values: ({arr[i]}, {arr[j]})")


if __name__ == '__main__':
    solve_problem([2, 7, 11, 15], 9)
    solve_problem([2, 7, 11, 15], 26)
