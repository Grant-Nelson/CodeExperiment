using System;
using System.Collections.Generic;
using System.IO;

namespace QuickSort
{
    /// <summary>
    /// This class reads a file of numbers, quick sorts the numbers,
    /// then writes the sorted numbers to another file.
    /// </summary>
    static public class EntryPoint
    {
        /// <summary>This is the file to input with random values from.</summary>
        static private string inputFile = Path.Combine("..", "..", "randomFile.txt");

        /// <summary>This is the file to output the sorted values to.</summary>
        static private string outputFile = Path.Combine("..", "..", "sortedFile.txt");

        /// <summary>This is the entry point for a quicksort in Go.</summary>
        static public void Main()
        {
            int[] data = ReadFile();
            int length = data.Length;
            if (length <= 0)
                throw new Exception("Failed to read input file");

            Quicksort(data, 0, length - 1);

            WriteFile(data);
        }

        /// <summary>This reads all the values from the input file.</summary>
        /// <returns>An unsorted list of values from the input file.</returns>
        static private int[] ReadFile()
        {
            using(StreamReader file = new StreamReader(inputFile))
            {
                List<int> data = new List<int>();
                string line = file.ReadLine();
                while (line != null)
                {
                    data.Add(int.Parse(line));
                    line = file.ReadLine();
                }
                return data.ToArray();
            }
        }

        /// <summary>
        /// This performs a quick sort in the low inclusive and high inclusive range.
        /// </summary>
        /// <param name="data">The data being quick sorted.</param>
        /// <param name="low">The low inclusive index for the range to sort.</param>
        /// <param name="high">The hight inclusive index for the range to sort.</param>
        static private void Quicksort(int[] data, int low, int high)
        {
            if (low < high)
            {
                int p = low, temp;
                for (int j = low, pivot = data[high]; j < high; j++)
                {
                    if (data[j] < pivot)
                    {
                        temp = data[p];
                        data[p] = data[j];
                        data[j] = temp;
                        p++;
                    }
                }
                temp = data[p];
                data[p] = data[high];
                data[high] = temp;

                Quicksort(data, low, p - 1);
                Quicksort(data, p + 1, high);
            }
        }

        /// <summary>This writes the values to the output file.</summary>
        /// <param name="data">The sorted values to write to the output file.</param>
        static private void WriteFile(int[] data)
        {
            using(StreamWriter file = new StreamWriter(outputFile))
            {
                foreach (int value in data)
                    file.WriteLine(value);
            }
        }
    }
}