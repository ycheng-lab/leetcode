#!/usr/bin/python
# -*- coding: utf-8 -*-

MIN_INT32 = -0x80000000
MAX_INT32 = 0x7fffffff


def atoi(s: str) -> int:
    i = 0
    for _, c in enumerate(s):
        if c == " ":
            i += 1
        else:
            break
    s = s[i:]
    sign = 1
    if s.startswith("-"):
        sign = -1
        s = s[1:]
    elif s.startswith("+"):
        sign = 1
        s = s[1:]

    result = 0
    for _, c in enumerate(s):
        if c.isdigit():
            result = result * 10 + (int(c))
        else:
            break
        if result * sign >= MAX_INT32:
            return MAX_INT32
        if result * sign <= MIN_INT32:
            return MIN_INT32

    return result * sign


def solve_problem(s: str) -> None:
    print(f"string: {s}, integer: {atoi(s)}")


if __name__ == '__main__':
    solve_problem("")
    solve_problem("42")
    solve_problem("   -42")
    solve_problem("4193 with words")
    solve_problem("words and 987")
    solve_problem("-91283472332")
