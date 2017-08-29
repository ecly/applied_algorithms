#include <stdio.h>
#include <stdlib.h>

int* matrixFromFile(FILE *fp, int n){
    int *mat = (int *)malloc(n * n * sizeof(int));

    //ignore first comment line
    fscanf(fp, "%*[^\n]\n", NULL);

    int i;
    int j;
    int lineCount = n*n;
    int characterCount = 5;//number of chars on line
    for(i = 0; i < lineCount; i++){
            int col,row,val;
            fscanf(fp, "%d,%d,%d", &col, &row, &val);
            mat[col*n+row] = val;
            fgetc(fp);//newline
        }
    return mat;
}

int multiplyMatrices(int* a, int* b, int n){
    int i,j,k;
    for(i = 0; i < n; i++){
        for(j = 0; j < n; j++){
            int val = 0;
            for(k = 0; k < n; k++){
                val = val + a[i*n+k]*b[k*n+j];
            }
           printf("%d ", val);
        }
    }
    return 1;
}

int main(int argc, char *argv[]){
    int n = atoi(argv[1]);
    FILE *fp1;
    FILE *fp2;
    fp1 = fopen(argv[2], "r");
    fp2 = fopen(argv[3], "r");
    int* m1 = matrixFromFile(fp1, n);
    int* m2 = matrixFromFile(fp2, n);
    multiplyMatrices(m1, m2, n);
    return 0;
}
