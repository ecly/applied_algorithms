/**
 * Maintain a dynamic set under insertion, reporting the approximate number of distinct items inserted using a HyperLogLog counter.
 * @author Rasmus Pagh
 * @version 2016.11.21
 */

public class ApproxSet {

	private static int logm = 10;
	private static int m = 1<<logm;
	private byte[] M = new byte[m];

	public void add(Object x) {
		int xh = x.hashCode();
		if (xh!=0) {
			int i = LinearHash.f(xh,logm);
			byte val = (byte)ExpHash.hash(xh);
			if (val>M[i]) M[i]=val;
		}
	}
	
	public int sizeEstimate() {
		double wsum = 0;
		int zerosum = 0;
		for (int j=0; j<m; j++) {
			wsum += Math.pow(2.0,-M[j]);
			if (M[j]==0) zerosum++;
		}
		double Z = 1/wsum;
		double estimate = m*m*Z*0.7213/(1 + 1.079/m);
		if ((estimate < 2.5 * m) && (zerosum>0))
			estimate = m * Math.log((double)m/zerosum);
		return (int)estimate;
	}

    public byte[] getSet(){
        return M;
    }

    public void addSet(ApproxSet a) {
        byte[] Mx = a.getSet();

        for(int i = 0; i < m; i++){
            if (M[i] < Mx[i])
                M[i] = Mx[i];
        }
	}
}
