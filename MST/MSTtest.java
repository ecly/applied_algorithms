public class MSTtest{

    static int mstToInt(int[] mst, int mstsize){
        int total = 0;
        for(int i = 0; i < mstsize; i++)
        {
            total += hashRand(mst[i]);
        }
        return total;
    }

    static int hashRand(int inIndex){
        final int b = 0x5f375a86;//bunch of random bits
        for(int i = 0; i < 8; i++)
        {
            inIndex = (inIndex + 1)*( (inIndex >> 1)^b);
        }
        return inIndex;
    }

    public static void main(String[] args){
        int[] mst = new int[]{-60078, -78884, 14222};
        System.out.println(mstToInt(mst, mst.length));
    }
}
