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
export const genFibFunc = (max = Number.MAX_SAFE_INTEGER) => {
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

/**
 * @example
 * const fib = genFib(6) // { next: [Function: next] }
 * fib.next() // { value: 1, done: false }
 * fib.next() // { value: 1, done: false }
 * fib.next() // { value: 2, done: false }
 * fib.next() // { value: 3, done: false }
 * fib.next() // { value: 5, done: false }
 * fib.next() // { value: undefined, done: true }
 * @example
 * const fib = genFib(6) // { next: [Function: next] }
 * [...fib] // [1, 1, 2, 3, 5]
 */
export const genFibIterator = (max = Number.MAX_SAFE_INTEGER) => {
  // initialize default values in the scope
  let n1 = 0
  let n2 = 0

  // this time we return an iterator object (rather than a function)
  return {
    // this satisfies the Iterator protocol
    next: () => {
      // calculates the next value
      let nextVal = n2 === 0 ? 1 : n1 + n2

      // redefines n1 and n2 to match new values
      const prevVal = n2
      n2 = nextVal
      n1 = prevVal

      // if we reached the upper bound (iteration completed)
      // set done to true and nextVal to undefined
      let done = false
      if (nextVal >= max) {
        nextVal = undefined
        done = true
      }

      // return the iteration object as for the iteration protocol
      return { value: nextVal, done }
    },

    // this satisfies the Iterable protocol
    [Symbol.iterator] () {
      // returns `this` because the object itself is an iterator
      return this
    }
  }
}

export function* genFib (max = Number.MAX_SAFE_INTEGER) {
  // initialize default values in the scope
  let n1 = 0
  let n2 = 0

  // we can now pre-initialize nextVal to 1 as part of the this.state
  let nextVal = 1

  // loop until we exceed the max number
  while (nextVal <= max) {
    // yields the current value
    yield nextVal

    // shifts nextVal -> v2 and n2 -> n1
    const prevVal = n2
    n2 = nextVal
    n1 = prevVal

    // calculates the next value
    nextVal = n1 + n2
  }
}
