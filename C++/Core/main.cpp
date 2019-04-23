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
 * This sorts the given data.
 * 
 * @param data The array to sort.
 */
void sort(int *data, int length)
{
    std::sort(data, data + length);
}

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
 * This is the entry point for a core sort in C++.
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

    sort(data, length);

    writeFile(data, length);
    if (data) free(data);
    return 0;
}
