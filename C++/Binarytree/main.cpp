#include <cstdio>
#include <algorithm>

/**
 * @note I didn't use std:fstream and std:queue here because
 * my C++ setup would build but not run any code with thosed in it.
 */

// The amount of integers to grow the buffer with.
#define BUFFER_GROWTH 1000

// This is the file to input with random values from.
#define INPUT_FILE "..\\..\\randomFile.txt"

// This is the file to output the sorted values to.
#define OUTPUT_FILE "..\\..\\sortedFile.txt"

/**
 * This reads all the values from the input file.
 * 
 * @param data A pointer to an array which will be replaced by the loaded data.
 * @param length The length of the data array.
 */
void readFile(int **data, int *length)
{
    int count = 0, value, i;
    int capacity = BUFFER_GROWTH;
    int *buf = NULL, *oldbuf;

    FILE *fid = fopen(INPUT_FILE, "r");
    if (fid)
    {
        buf = (int *)malloc(capacity * sizeof(int));
        while (!feof(fid))
        {
            fscanf(fid, "%d", &value);

            if (count + 1 >= capacity)
            {
                capacity += BUFFER_GROWTH;
                oldbuf = buf;
                buf = (int *)malloc(capacity * sizeof(int));
                for (i = 0; i < count; ++i)
                    buf[i] = oldbuf[i];
                free(oldbuf);
            }

            buf[count] = value;
            ++count;
        }
        fclose(fid);
    }

    *length = count;
    *data = buf;
}

// This is the binary tree node.
class Node
{
public:
    // This is the value for this node.
    const int value;

    // This is the value less than this value.
    Node *left;

    // This is the value greater than or equal to this value.
    Node *right;

    /**
     * This constructs a new Node.
     * 
     * @param value The value for this node.
     */
    Node(int value);

    // Deconstructor for this Node.
    ~Node();
};

Node::Node(int value) : value(value),
                        left(NULL),
                        right(NULL)
{
    /* Do Nothing */
}

Node::~Node()
{
    if (this->left != NULL)
        free(this->left);
    if (this->right != NULL)
        free(this->right);
}

/**
 * This inserts a value into the tree recursively.
 * 
 * @param value This is the value to insert.
 * @param n This is the current node to insert inside of.
 */
void insertValue(Node *n, Node *p)
{
    if (n->value > p->value)
    {
        if (n->left != NULL)
            insertValue(n->left, p);
        else
            n->left = p;
    }
    else
    {
        if (n->right != NULL)
            insertValue(n->right, p);
        else
            n->right = p;
    }
}

/**
 * This recursively gets all the values from the binary tree.
 * 
 * @param index The current index to start outputting to.
 * @param n The current node to dive into.
 * @param data The array to write the sorted values to.
 * @returns The index after all the parts of the tree were written.
 */
int outputValues(int index, Node *n, int *data)
{
    if (n->left != NULL)
        index = outputValues(index, n->left, data);

    data[index] = n->value;
    ++index;

    if (n->right != NULL)
        index = outputValues(index, n->right, data);
    return index;
}

/**
 * This writes the values to the output file.
 * 
 * @param data The sorted values to write to the output file.
 * @param length The number of values to write.
 */
void writeFile(int *data, int length)
{
    int i;
    FILE *fid = fopen(OUTPUT_FILE, "w");
    if (fid)
    {
        for (i = 0; i < length; ++i)
            fprintf(fid, "%d\n", data[i]);
        fclose(fid);
    }
}

/**
 * This is the entry point for a binary tree sort in C++.
 * 
 * This reads a file of numbers, sorts the numbers with
 * the built-in standard template library's sort algorithm,
 * then writes the sorted numbers to another file. 
 */
int main()
{
    int *data = NULL;
    int length = 0;
    readFile(&data, &length);

    Node *root = new Node(data[0]);
    for (int i = 1; i < length; ++i)
    {
        Node *n = new Node(data[i]);
        insertValue(root, n);
    }

    outputValues(0, root, data);

    writeFile(data, length);

    if (data)
        free(data);
    free(root);
    return 0;
}
