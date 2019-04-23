using System;
using System.Collections.Generic;
using System.IO;

namespace Core
{
    /// <summary>
    /// This class reads a file of numbers, sorts the numbers with
    /// the built-in sort algorithm, then writes the sorted numbers to another file.
    /// </summary>
    static public class EntryPoint
    {
        /// <summary>This is the file to input with random values from.</summary>
        static private string inputFile = Path.Combine("..", "..", "randomFile.txt");

        /// <summary>This is the file to output the sorted values to.</summary>
        static private string outputFile = Path.Combine("..", "..", "sortedFile.txt");

        /// <summary>This sorts the given data.</summary>
        /// <param name="data">The data to sort.</param>
        static private void Sort(int[] data)
        {
            Array.Sort(data);
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

        /// <summary>This is the entry point for a core sort in C#.</summary>
        static public void Main()
        {
            int[] data = ReadFile();
            Sort(data);
            WriteFile(data);
        }
    }
}