/**
 * Compute array where n[i] is number of ordered vertex pairs (v1, v2) at distance i
 * Takes undirected graph input from stdin with one edge per line: "nodeID1 nodeID2"
 * @author Rasmus Pagh
 * @version 2016.10.22
 * Usage: java AllBFS < graph.txt
 */

/**
 * Converted to HyperANF by Emil Lynegaard
 * 2017
 */

import java.util.*;

public class ANF {

	private static Map<Integer,ApproxSet> counters = new HashMap<Integer,ApproxSet>();
    private static Map<Integer,Set<Integer>> graph = new HashMap<Integer,Set<Integer>>();
	private static int[] n;

	private static void readGraph() {
		Scanner input = new Scanner(System.in);
		while (input.hasNextLine()) {
			String line = input.nextLine();
			Integer a = Integer.parseInt(line.split(" ")[0]);
			Integer b = Integer.parseInt(line.split(" ")[1]);

            if(!graph.containsKey(a)){
                graph.put(a, new HashSet<Integer>());
                ApproxSet counter = new ApproxSet();
                counter.add(a);
                counters.put(a, counter);
            }
            graph.get(a).add(b);

            if(!graph.containsKey(b)){
                graph.put(b, new HashSet<Integer>());
                ApproxSet counter = new ApproxSet();
                counter.add(b);
                counters.put(b, counter);
            }
            graph.get(b).add(a);
        }
	}

    private static int ANF(){
        int d = 0;
        int medianLocation = graph.size() * graph.size() / 2;
        while(true){
            int reach = 0;
            Map<Integer,ApproxSet> m = new HashMap<Integer,ApproxSet>();

            for(Map.Entry<Integer, Set<Integer>> entry : graph.entrySet()){
                ApproxSet mv = new ApproxSet();
                mv.addSet(counters.get(entry.getKey()));
                for(Integer neighbor: entry.getValue())
                    mv.addSet(counters.get(neighbor));

                reach += mv.sizeEstimate();
                m.put(entry.getKey(), mv);
            }
            d++;
            if (reach >= medianLocation) break;
            counters = m;
        }
        return d;
    }

	public static void main(String[] args) {
		readGraph();
        System.out.println(ANF());
	}
}
