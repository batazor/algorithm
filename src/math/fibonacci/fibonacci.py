import unittest


def fibonacci(n: int) -> int:
    if n < 2:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)


if __name__ == '__main__':
    assert fibonacci(9) == 34, "Should be 34"
    unittest.main()
