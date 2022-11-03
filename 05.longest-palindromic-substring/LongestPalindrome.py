
def is_palindrome(s: str) -> bool:
    if len(s) <= 1:
        return True
    i, j = 0, len(s) - 1
    while i < j:
        if s[i] != s[j]:
            return False
        else:
            i += 1
            j -= 1
    return True


def get_longest_palindrome(s: str) -> str:
    palindrome = ""
    max_length = 0
    for i, j in zip(range(len(s)), reversed(range(len(s)))):
        if is_palindrome(s[i:j + 1]):
            if j - i + 1 > max_length:
                max_length = j - i + 1
                palindrome = s[i:j + 1]
    return palindrome


def solve_problem(s: str) -> None:
    print(f"string: {s}, max palindrome: {get_longest_palindrome(s)}")


if __name__ == '__main__':
    solve_problem("babad")
    solve_problem("cbbd")
