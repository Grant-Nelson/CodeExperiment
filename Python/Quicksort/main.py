import os.path

inputFile = os.path.join("..", "..", "randomFile.txt")
""" This is the file to input with random values from. """

outputFile = os.path.join("..", "..", "sortedFile.txt")
""" This is the file to output the sorted values to. """


def sort(data):
    """ This sorts the given data. """
    quicksort(data, 0, len(data)-1)


def quicksort(data, low, high):
    """ This performs a quick sort in the low inclusive
    and high inclusive range. """
    if low < high:
        pivot = data[high]
        p = low
        for j in range(low, high):
            if data[j] < pivot:
                temp = data[p]
                data[p] = data[j]
                data[j] = temp
                p = p+1

        temp = data[p]
        data[p] = data[high]
        data[high] = temp

        quicksort(data, low, p-1)
        quicksort(data, p+1, high)


def readFile():
    """ This reads all the values from the input file. """
    data = []
    f = open(inputFile, "r")
    for line in f:
        data.append(int(line))
    return data


def writeFile(data):
    """ This writes the values to the output file. """
    f = open(outputFile, "w")
    for value in data:
        f.write("{}\n".format(value))
    f.close()


data = readFile()
sort(data)
writeFile(data)
exit(0)
