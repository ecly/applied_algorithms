import java.util.Scanner;
import java.io.File;
public class NaiveMM{
    public static void main(String[] args){
        int n = Integer.parseInt(args[0]);
        int[][] a = matrixFromFile(args[1], n);
        int[][] b = matrixFromFile(args[2], n);
        int[][] res = multiplyMatrices(a,b);
        System.out.println(matrixToString(res));
    }

    private static int[][] multiplyMatrices(int[][] a, int[][] b){
        int[][] res = new int[a.length][a.length];
        for(int i = 0; i < a.length; i++)
            for(int j = 0; j < a.length; j++)
                for(int k = 0; k < a.length; k++)
                    res[i][j] += a[i][k]*b[k][j];

        return res;
    }

    private static String matrixToString(int[][] matrix){
        StringBuilder builder = new StringBuilder();
        for(int i = 0; i < matrix.length; i++){
            for(int j = 0; j < matrix.length; j++){
                builder.append(matrix[i][j]); 
                builder.append(" "); 
            }
        }

        return builder.toString();
    }

    private static int[][] matrixFromFile(String fileName, int n){
        int[][] matrix = new int[n][n];
        try {
            Scanner scanner = new Scanner(new File(fileName));
            // ignore first line
            scanner.nextLine();

            for(int m = 0; m < n*n; m++){
                String line = scanner.nextLine();
                String[] vals = line.split(",");
                int i = Integer.parseInt(vals[0]);
                int j = Integer.parseInt(vals[1]);
                int x = Integer.parseInt(vals[2]);
                matrix[i][j] = x;
            }
        } catch (Exception e){}
        return matrix;
    }
}
