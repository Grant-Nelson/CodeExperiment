import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.util.ArrayList;

/**
 * The Mergesort class reads a file of numbers, merge sorts the numbers,
 * then writes the sorted numbers to another file.
 */
public class Mergesort {

    /** This is the file to input with random values from. */
    static private final String inputFile = "../../randomFile.txt";

    /** This is the file to output the sorted values to. */
    static private final String outputFile = "../../sortedFile.txt";

    /**
     * This is the entry point for a quicksort in Go.
     * 
     * @param args not used.
     */
    public static void main(String[] args) {
        try {
            int[] randomData = readFile();
            int length = randomData.length;
            if (length <= 0)
                throw new Exception("Failed to read a file");
            
            int[] sortedData = randomData.clone();
            split(randomData, sortedData, 0, length);

            writeFile(sortedData);
        } catch (Exception e) {
            System.out.println(e);
            System.exit(1);
        }
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
     * This performs a top down merge sort by splitting the current level into 2
     * parts to sort, then merging the two parts.
     * 
     * @param a     The source array for merging from.
     * @param b     The copy of the array for merging to.
     * @param start The inclusive index to start merging at.
     * @param stop  The inclusive index to stop merging at.
     */
    private static void split(int[] a, int[] b, int start, int stop) {
        if (stop-start < 2) return;
        int mid = (stop + start) / 2;
        split(b, a, start, mid);
        split(b, a, mid, stop);
        merge(a, b, start, mid, stop);
    }

    /**
     * merge performs a top down merge where it zippers together
     * two parts from `a` into `b`.
     * 
     * @param a     The source array for merging from.
     * @param b     The copy of the array for merging to.
     * @param start The inclusive index to start merging at.
     * @param mid   The index halfway between the start and stop.
     * @param stop  The inclusive index to stop merging at.
     */    
    private static void merge(int[] a, int[] b, int start, int mid, int stop) {
        for (int i = start, j = mid, k = start; k < stop; k++) {
            if ((i < mid) && ((j >= stop) || (a[i] <= a[j]))) {
                b[k] = a[i];
                i++;
            } else {
                b[k] = a[j];
                j++;
            }
        }
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
}
