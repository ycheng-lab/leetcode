#!/usr/bin/python3
# -*- coding: utf-8 -*-

def is_palindrome_number(n: int) -> bool:
    if n < 0:
        return False

    while n > 0:
        if n < 10:
            return True

        i = 1
        left = n // 10
        while left > 10:
            left //= 10
            i += 1
        
        if left != n % 10:
            return False
        
        for _ in range(i):
            left *= 10

        n -= left
        n //= 10

    return True

def solve_problem(n : int):
	print(f"{n} is palindrome number: {is_palindrome_number(n)}")


if __name__ == "__main__":
	solve_problem(2)
	solve_problem(28273642)
	solve_problem(123456654321)
	solve_problem(17361)




