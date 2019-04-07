import os.path

inputFile = os.path.join("..", "..", "randomFile.txt")
""" This is the file to input with random values from. """

outputFile = os.path.join("..", "..", "sortedFile.txt")
""" This is the file to output the sorted values to. """


def readFile():
    """ This reads all the values from the input file. """
    data = []
    f = open(inputFile, "r")
    for line in f:
        data.append(int(line))
    return data


def quicksort(data, low, high):
    """ This performs a quick sort in the low inclusive
    and high inclusive range. """
    if low < high:
        p = partition(data, low, high)
        quicksort(data, low, p-1)
        quicksort(data, p+1, high)


def partition(data, low, high):
    """ This performs a top down merge where it zippers together two
    parts from `a` into `b`. start is inclusive and stop is exclusive. """
    pivot = data[high]
    i = low
    for j in range(low, high):
        if data[j] < pivot:
            temp = data[i]
            data[i] = data[j]
            data[j] = temp
            i = i+1

    temp = data[i]
    data[i] = data[high]
    data[high] = temp
    return i


def writeFile(data):
    """ This writes the values to the output file. """
    f = open(outputFile, "w")
    for value in data:
        f.write("{}\n".format(value))
    f.close()


data = readFile()
length = len(data)

quicksort(data, 0, length-1)

writeFile(data)
exit(0)
