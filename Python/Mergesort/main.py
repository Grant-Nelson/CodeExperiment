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


def split(b, a, start, stop):
    """ This performs a top down merge sort by splitting the
    current level into 2 parts to sort, then merging the two parts.
    start is inclusive and stop is exclusive. """
    if stop-start < 2:
        return
    mid = int((stop + start)/2)
    split(a, b, start, mid)
    split(a, b, mid, stop)
    merge(b, a, start, mid, stop)


def merge(a, b, start, mid, stop):
    """ This performs a top down merge where it zippers together two
    parts from `a` into `b`. start is inclusive and stop is exclusive. """
    i = start
    j = mid
    for k in range(start, stop):
        if (i < mid) and ((j >= stop) or (a[i] <= a[j])):
            b[k] = a[i]
            i = i+1
        else:
            b[k] = a[j]
            j = j+1


def writeFile(data):
    """ This writes the values to the output file. """
    f = open(outputFile, "w")
    for value in data:
        f.write("{}\n".format(value))
    f.close()


randomData = readFile()
length = len(randomData)

sortedData = randomData.copy()
split(randomData, sortedData, 0, length)

writeFile(sortedData)
exit(0)
