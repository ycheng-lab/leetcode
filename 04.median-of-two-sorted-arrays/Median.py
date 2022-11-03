

def find_median(arr : list[int]) -> float:
    k = len(arr) // 2
    if len(arr) % 2 == 1:
        return arr[k]
    return (arr[k-1] + arr[k]) / 2


def merge(one: list[int], two: list[int]) -> list[int]:
    i, j = 0, 0
    whole = []
    for _ in range(len(one) + len(two)):
        if j == len(two) or (i < len(one) and j < len(two) and one[i] <= two[j]):
            whole.append(one[i])
            i += 1
            continue
        if i == len(one) or (i < len(one) and j < len(two) and one[i] >= two[j]):
            whole.append(two[j])
            j += 1
    return whole

def calc_median(one:list[int], two:list[int]) -> float:
    return find_median(merge(one, two))

if __name__ == '__main__':
    print(calc_median([], [7, 8, 9, 10]))