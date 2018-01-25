a, b = 0, 1

def fibonacci():
    global a, b
    while True:
        a, b = b, a + b
        yield a

def main():
    f = fibonacci()
    last_number = 0
    fib_sum = 0

    while last_number < 4e6:
        last_number = next(f)
        if last_number % 2 == 0:
            fib_sum += last_number

    print(fib_sum)

if __name__ == "__main__":
    main()

