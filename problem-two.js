// Let's do generators in JS too since we did it in Python

function* fibonacci() {
  let current = a = 1
  let b = 2

  while (true) {
    current = b

    yield a

    b = a + current
    a = current
  }
}

const main = () => {
  const f = fibonacci()
  let lastNumber = 0
  let fibSum = 0

  while (lastNumber < 4e6) {
    lastNumber = f.next().value
    fibSum += lastNumber % 2 === 0 ? lastNumber : 0
  }

  console.log(fibSum)
  return fibSum
}

require.main === module && main()

