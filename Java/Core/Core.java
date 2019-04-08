import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileReader;
import java.io.FileWriter;
import java.util.ArrayList;
import java.nio.file.Paths;
import java.util.Arrays;

/**
 * The Core class reads a file of numbers, core sorts the numbers,
 * then writes the sorted numbers to another file.
 */
public class Core {

    /** This is the file to input with random values from. */
    static private final String inputFile =
        Paths.get("..", "..", "randomFile.txt").toString();

    /** This is the file to output the sorted values to. */
    static private final String outputFile =
        Paths.get("..", "..", "sortedFile.txt").toString();

    /**
     * This is the entry point for a core sort in Go.
     * 
     * @param args not used.
     */
    public static void main(String[] args) {
        try {
            int[] data = readFile();
            int length = data.length;
            if (length <= 0)
                throw new Exception("Failed to read input file");
            
            Arrays.sort(data);

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
}
