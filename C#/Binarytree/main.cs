using System;
using System.Collections.Generic;
using System.IO;

namespace Binarytree
{
    /// <summary>
    /// This class reads a file of numbers, binary tree sort the numbers,
    /// then writes the sorted numbers to another file.
    /// </summary>
    static public class EntryPoint
    {
        /// <summary>This is the file to input with random values from.</summary>
        static private string inputFile = Path.Combine("..", "..", "randomFile.txt");

        /// <summary>This is the file to output the sorted values to.</summary>
        static private string outputFile = Path.Combine("..", "..", "sortedFile.txt");

        /// <summary>This is the binary tree node.</summary>
        private class Node
        {
            /// <summary>This is the value for this node.</summary>
            public int value;

            /// <summary>This is the value less than this value.</summary>
            public Node left;

            /// <summary>This is the value greater than or equal to this value.</summary>
            public Node right;

            /// <summary>This constructs a new Node.</summary>
            /// <param name="value">The value for this node.</param>
            public Node(int value)
            {
                this.value = value;
                this.left = null;
                this.right = null;
            }
        }

        /// <summary>This is the entry point for a binary tree sort in C#.</summary>
        static public void Main()
        {
            int[] data = ReadFile();
            int length = data.Length;
            if (length <= 0)
                throw new Exception("Failed to read input file");

            Node root = new Node(data[0]);
            for (int i = 1; i < length; ++i)
                InsertValue(data[i], root);

            OutputValues(0, root, data);

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

        /// <summary>This inserts a value into the tree recursively.</summary>
        /// <param name="value">This is the value to insert.</param>
        /// <param name="n">This is the current node to insert inside of.</param>
        static private void InsertValue(int value, Node n)
        {
            if (n.value > value)
            {
                if (n.left != null) InsertValue(value, n.left);
                else n.left = new Node(value);
            }
            else
            {
                if (n.right != null) InsertValue(value, n.right);
                else n.right = new Node(value);
            }
        }

        /// <summary>This recursively gets all the values from the binary tree.</summary>
        /// <param name="index">The current index to start outputting to.</param>
        /// <param name="n">The current node to dive into.</param>
        /// <param name="data">The array to write the sorted values to.</param>
        /// <returns>The index after all the parts of the tree were written.</returns>
        static private int OutputValues(int index, Node n, int[] data)
        {
            if (n.left != null) index = OutputValues(index, n.left, data);

            data[index] = n.value;
            ++index;

            if (n.right != null) index = OutputValues(index, n.right, data);
            return index;
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