using System;
using System.Collections.Generic;
using System.IO;

namespace Mergesort
{
    /// <summary>
    /// This class reads a file of numbers, merge sorts the numbers,
    /// then writes the sorted numbers to another file.
    /// </summary>
    static public class EntryPoint
    {
        /// <summary>This is the file to input with random values from.</summary>
        static private string inputFile = Path.Combine("..", "..", "randomFile.txt");

        /// <summary>This is the file to output the sorted values to.</summary>
        static private string outputFile = Path.Combine("..", "..", "sortedFile.txt");

        /// <summary>
        /// This performs a top down merge sort by splitting the current level into 2
        /// parts to sort, then merging the two parts.
        /// </summary>
        /// <param name="a">The source array for merging from.</param>
        /// <param name="b">The copy of the array for merging to.</param>
        /// <param name="start">The inclusive index to start merging at.</param>
        /// <param name="stop">The exclusive index to stop merging at.</param>
        static private void Split(int[] a, int[] b, int start, int stop)
        {
            if (stop - start < 2) return;

            int mid = (stop + start) / 2;
            Split(b, a, start, mid);
            Split(b, a, mid, stop);

            for (int i = start, j = mid, k = start; k < stop; ++k)
            {
                if ((i < mid) && ((j >= stop) || (a[i] <= a[j])))
                {
                    b[k] = a[i];
                    ++i;
                }
                else
                {
                    b[k] = a[j];
                    ++j;
                }
            }
        }

        /// <summary>This sorts the given data.</summary>
        /// <param name="data">The data to sort.</param>
        static private void Sort(int[] data)
        {
            int length = data.Length;
            int[] sortedData = new int[length];
            data.CopyTo(sortedData, 0);
            Split(data, sortedData, 0, length);
            sortedData.CopyTo(data, 0);
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

        /// <summary>This is the entry point for a merge sort in C#.</summary>
        static public void Main()
        {
            int[] data = ReadFile();
            Sort(data);
            WriteFile(data);
        }
    }
}