import os.path

inputFile = os.path.join("..", "..", "randomFile.txt")
""" This is the file to input with random values from. """

outputFile = os.path.join("..", "..", "sortedFile.txt")
""" This is the file to output the sorted values to. """


def split(a, b, start, stop):
    """ This performs a top down merge sort by splitting the
    current level into 2 parts to sort, then merging the two parts.
    start is inclusive and stop is exclusive. """
    if stop-start < 2:
        return

    mid = int((stop + start)/2)
    split(b, a, start, mid)
    split(b, a, mid, stop)

    i = start
    j = mid
    for k in range(start, stop):
        if (i < mid) and ((j >= stop) or (a[i] <= a[j])):
            b[k] = a[i]
            i = i+1
        else:
            b[k] = a[j]
            j = j+1


def sort(data):
    """ This sorts the given data. """
    length = len(data)
    sortedData = data.copy()
    split(data, sortedData, 0, length)
    for i in range(0, length):
        data[i] = sortedData[i]


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
