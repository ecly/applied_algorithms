/**
 * Compute array where n[i] is number of ordered vertex pairs (v1, v2) at distance i
 * Takes undirected graph input from stdin with one edge per line: "nodeID1 nodeID2"
 * @author Rasmus Pagh
 * @version 2016.10.22
 * Usage: java AllBFS < graph.txt
 */

import java.util.*;

public class AllBFS {

	private static Map<Integer,Set<Integer>> graph = new HashMap<Integer,Set<Integer>>();
	private static int[] n;

	private static void readGraph() {
		Scanner input = new Scanner(System.in);
		while (input.hasNextLine()) {
			String line = input.nextLine();
			Integer a = Integer.parseInt(line.split(" ")[0]);
			Integer b = Integer.parseInt(line.split(" ")[1]);
			if (graph.get(a) == null) graph.put(a,new HashSet<Integer>());
			graph.get(a).add(b);
			if (graph.get(b) == null) graph.put(b,new HashSet<Integer>());
			graph.get(b).add(a);
		}
	}

	private static void BFS(Integer start) {
		Queue<Integer> queue = new ArrayDeque<Integer>();
		Set<Integer> visited = new HashSet<Integer>();
		
		visited.add(start);
		queue.add(start);
		int d = 0;
		while (queue.peek()!=null) {
			int newNodes = queue.size();
			n[d++] += newNodes;
			for (int i=0; i<newNodes; i++) {
				Integer v = queue.poll();
				for (Integer w : graph.get(v)) {
					if (!visited.contains(w)) {
						visited.add(w);
						queue.add(w);
					}
				}
			}
		}
	}
	
	public static void main(String[] args) {
		readGraph();
		n = new int[graph.keySet().size()];

		for (Integer v : graph.keySet()) { BFS(v); }

		for (int i=0; i<n.length; i++) {
			if (n[i]>0) System.out.println("n["+i+"] = "+n[i]);
		}
		
	}
}
