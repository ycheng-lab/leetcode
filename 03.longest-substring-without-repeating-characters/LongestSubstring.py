
def calc_longest_substr(s : str) -> tuple[str, int]:
    locations = [-1 for x in range(255)]

    index, max_length, left = 0, 0, 0
    for i, c in enumerate(s):       
        if locations[ord(c)] >= left:
            left = locations[ord(c)] + 1
        elif i - left + 1 > max_length:
            max_length = i - left + 1
            index = i
        locations[ord(c)] = i
    
    return s[index - max_length + 1 : index + 1], max_length


def solve_problem(s:str) -> None:
    substr, length = calc_longest_substr(s)
    print(f'string: {s}, longest substr: {substr}, length: {length}')


if __name__ == '__main__':
    solve_problem("abcabcbb")
    solve_problem("bbbbb")
    solve_problem("pwwkew")

