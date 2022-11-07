#!/usr/bin/python
# -*- coding: utf-8 -*-

def reverse_integer(n: int) -> int:
    sign = 1
    if n < 0:
        sign = -1
        n = -n

    result = 0
    while n > 0:
        result = result * 10 + (n % 10)
        n //= 10

    return result * sign


def solve_problem(n: int) -> None:
    print(f"integer: {n}, reversed: {reverse_integer(n)}")


if __name__ == '__main__':
    solve_problem(1)
    solve_problem(123)
    solve_problem(-9864)
    solve_problem(0)
    solve_problem(-1)
