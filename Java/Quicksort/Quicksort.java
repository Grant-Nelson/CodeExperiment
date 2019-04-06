import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.util.ArrayList;

/**
 * The Quicksort class reads a file of numbers, quick sorts the numbers,
 * then writes the sorted numbers to another file.
 */
public class Quicksort {

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
            int[] data = readFile();
            int length = data.length;
            if (length <= 0)
                throw new Exception("Failed to read a file");

            quicksort(data, 0, length - 1);
            
            writeFile(data);
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
        ArrayList<Integer> values = new ArrayList<Integer>(100000);
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
     * This performs a quick sort in the low inclusive and high inclusive range.
     * 
     * @param data The data being quick sorted.
     * @param low  The low inclusive index for the range to sort.
     * @param high The hight inclusive index for the range to sort.
     */
    private static void quicksort(int[] data, int low, int high) {
        if (low < high) {
            int p = partition(data, low, high);
            quicksort(data, low, p - 1);
            quicksort(data, p + 1, high);
        }
    }

    /**
     * This shifts values lower than a pivot and returns the pivot index.
     * 
     * @param data The data being quick sorted.
     * @param low  The low inclusive index for the range to sort.
     * @param high The hight inclusive index for the range to sort.
     * @return The pivot index to perform the next split at.
     */
    private static int partition(int[] data, int low, int high) {
        int pivot = data[high];
        int i = low;
        for (int j = low; j < high; j++) {
            if (data[j] < pivot) {
                int temp = data[i];
                data[i] = data[j];
                data[j] = temp;
                i++;
            }
        }
        int temp = data[i];
        data[i] = data[high];
        data[high] = temp;
        return i;
    }

    /**
     * This writes the values to the output file.
     * 
     * @param data The sorted values to write to the output file.
     * @throws Exception
     */
    private static void writeFile(int[] data) throws Exception {
        BufferedWriter writer = new BufferedWriter(new FileWriter(outputFile));
        for (int value: data) {
            writer.write(value +"\n");
        }
        writer.close();
    }
}
