def main():

    numbers = []
    for i in range(1000):
        if i % 3 == 0 or i % 5 == 0:
            numbers.append(i)

    sum = 0
    for number in numbers:
        sum += number

    print(sum)

if __name__ == "__main__":
    main()
