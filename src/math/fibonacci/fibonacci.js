/**
 * @example
 * const fib = genFib(6)
 * fib() // 1
 * fib() // 1
 * fib() // 2
 * fib() // 3
 * fib() // 5
 * fib() // null
 */
export const genFib = (max = Number.MAX_SAFE_INTEGER) => {
  // initialize default values in the scope
  let n1 = 0
  let n2 = 0

  // return an anounymous function that will return the next element
  // every time that it is called
  return () => {
    // calculates the next value
    const nextVal = n2 === 0 ? 1 : n1 + n2

    // redefines n1 and n2 to match new values
    const prevVal = n2
    n2 = nextVal
    n1 = prevVal

    // if we reached the upper bound return null (iteration completed)
    if (nextVal >= max) {
      return null
    }

    return nextVal
  }
}
