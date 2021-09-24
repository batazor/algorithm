import unittest
from typing import Dict, Generator
from functools import lru_cache


# fibonacci - use recursion method
def fibonacci(n: int) -> int:
    if n < 2:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)


# fibonacciWithMemoization - use memoization for processing
memo: Dict[int, int] = {0: 0, 1: 1}  # basic case


def fibonacciWithMemoization(n: int) -> int:
    if n not in memo:
        memo[n] = fibonacciWithMemoization(n - 1) + fibonacciWithMemoization(n - 2)  # memoization
    return memo[n]


@lru_cache(maxsize=None)
def fibonacciWithCache(n: int) -> int:
    if n < 2:
        return n
    return fibonacciWithCache(n - 1) + fibonacciWithCache(n - 2)


def fibonacciWithIterative(n: int) -> int:
    if n == 0: return n  # special case
    last: int = 0  # begin value
    next: int = 1  # end value

    for _ in range(1, n):
        last, next = next, last + next
    return next


# fibonacciGenerator - return generator
def fibonacciGenerator(n: int) -> Generator[int, None, None]:
    yield 0  # special case
    if n > 0: yield 1  # special case
    last: int = 0  # begin value
    next: int = 1  # end value

    for _ in range(1, n):
        last, next = next, last + next
        yield next  # main step generator


if __name__ == '__main__':
    assert fibonacci(9) == 34, "Should be 34"
    assert fibonacciWithMemoization(30) == 832040, "Should be 832040"
    assert fibonacciWithCache(30) == 832040, "Should be 832040"
    assert fibonacciWithIterative(30) == 832040, "Should be 832040"

    for i in fibonacciGenerator(50):
        print(i)

    unittest.main()
