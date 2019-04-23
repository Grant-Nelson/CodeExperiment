import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.util.ArrayList;
import java.nio.file.Paths;

/**
 * The Binarytree class reads a file of numbers, binary tree sort the numbers,
 * then writes the sorted numbers to another file.
 */
public class Binarytree {

    /** This is the file to input with random values from. */
    static private final String inputFile =
        Paths.get("..", "..", "randomFile.txt").toString();

    /** This is the file to output the sorted values to. */
    static private final String outputFile =
        Paths.get("..", "..", "sortedFile.txt").toString();

    /** This is the binary tree node. */
    private static class Node {

        /** This is the value for this node. */
        public int value;

        /** This is the value less than this value. */
        public Node left;

        /** This is the value greater than or equal to this value. */
        public Node right;

        /**
         * This constructs a new Node.
         * 
         * @param value The value for this node.
         */
        public Node(int value) {
            this.value = value;
        }
    }

    /**
     * This inserts a value into the tree recursively.
     * 
     * @param value This is the value to insert.
     * @param n     This is the current node to insert inside of.
     */    
    private static void insertValue(int value, Node n) {
        if (n.value > value) {
            if (n.left != null)
                insertValue(value, n.left);
            else n.left = new Node(value);
        } else {
            if (n.right != null)
                insertValue(value, n.right);
            else n.right = new Node(value);
        }
    }

    /**
     * This recursively gets all the values from the binary tree.
     *
     * @param index The current index to start outputting to.
     * @param n     The current node to dive into.
     * @param data  The array to write the sorted values to.
     * @return The index after all the parts of the tree were written.
     */
    private static int outputValues(int index, Node n, int[] data) {
       if (n.left != null)
            index = outputValues(index, n.left, data);

        data[index] = n.value;
        index++;

        if (n.right != null)
            index = outputValues(index, n.right, data);
        return index;
    }

    /**
     * This sorts the given data.
     *
     * @param data  The data to sort.
     */
    private static void sort(int[] data) {
        int length = data.length;
        Node root = new Node(data[0]);
        for (int i = 1; i < length; i++)
            insertValue(data[i], root);

        outputValues(0, root, data);
    }

    /**
     * This reads all the values from the input file.
     * 
     * @return An unsorted list of values from the input file.
     * @throws Exception
     */
    private static int[] readFile() throws Exception {
        ArrayList<Integer> values = new ArrayList<Integer>();
        BufferedReader reader = new BufferedReader(new FileReader(inputFile));
        String line = reader.readLine();
        while (line != null) {
            values.add(Integer.parseInt(line));
            line = reader.readLine();
        }
        reader.close();

        int length = values.size();
        int[] data = new int[length];
        for (int i = 0; i < length; i++)
            data[i] = values.get(i).intValue();
        return data;
    }

    /**
     * This writes the values to the output file.
     * 
     * @param data The sorted values to write to the output file.
     * @throws Exception
     */
    private static void writeFile(int[] data) throws Exception {
        BufferedWriter writer = new BufferedWriter(new FileWriter(outputFile));
        for (int value : data) {
            writer.write(value + "\n");
        }
        writer.close();
    }

    /**
     * This is the entry point for a binary tree in Go.
     * 
     * @param args not used.
     */
    public static void main(String[] args) {
        try {
            int[] data = readFile();
            sort(data);
            writeFile(data);
        } catch (Exception e) {
            System.out.println(e);
            System.exit(1);
        }
    }
}
