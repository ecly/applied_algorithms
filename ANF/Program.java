import java.util.Scanner;

public class Program{
    public static void main(String args[]){
        Scanner scanner = new Scanner(System.in);
        String first = scanner.nextLine();
        String second = scanner.nextLine();

        Integer[] firstInts = readInts(first);
        Integer[] secondInts = readInts(second);

        ApproxSet firstSet = createSet(firstInts);
        ApproxSet secondSet = createSet(secondInts);

        int initialSize = firstSet.sizeEstimate();
        firstSet.addSet(secondSet);
        if (1.2 * initialSize < firstSet.sizeEstimate())
            System.out.println("almost disjoint");
        else
            System.out.println("almost same");
    }

    private static Integer[] readInts(String line){
        String[] split = line.split(" "); 
        Integer[] integers = new Integer[split.length];
        for(int i = 0; i < split.length; i++)
            integers[i] = new Integer(split[i]);

        return integers;
    }

    private static ApproxSet createSet(Integer[] ints){
        ApproxSet set = new ApproxSet();
        for(Integer i: ints)
            set.add(i);
        return set;
    }
}
