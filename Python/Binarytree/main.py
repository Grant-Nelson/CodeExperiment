import os.path

# This is the file to input with random values from.
inputFile = os.path.join("..", "..", "randomFile.txt")

# This is the file to output the sorted values to.
outputFile = os.path.join("..", "..", "sortedFile.txt")

# This is the binary tree node.
class Node:
    def __init__(self, value, left, right):
        self.value = value
        self.left = left
        self.right = right

# This reads all the values from the input file.
def readFile():
    data = []
    f = open(inputFile, "r")
    for line in f:
        data.append(int(line))
    return data

# This inserts a value into the tree recursively.
def insertValue(value, n):
    if n.value > value:
        if n.left is None:
            n.left = Node(value, None, None)
        else:
            insertValue(value, n.left)
    else:
        if n.right is None:
            n.right = Node(value, None, None)
        else:
            insertValue(value, n.right)

# This recursively gets all the values from the binary tree.
def outputValues(index, n, data):
    if n.left is not None:
        index = outputValues(index, n.left, data)

    data[index] = n.value
    index += 1

    if n.right is not None:
        index = outputValues(index, n.right, data)
    return index

# This writes the values to the output file.
def writeFile(data):
    f = open(outputFile, "w")
    for value in data:
        f.write("{}\n".format(value))
    f.close()

data = readFile()

root = Node(data[0], None, None)
for value in data[1:]:
    insertValue(value, root)

outputValues(0, root, data)

writeFile(data)
exit(0)
