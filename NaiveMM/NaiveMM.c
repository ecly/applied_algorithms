#include <stdio.h>
#include <stdlib.h>

int* matrixFromFile(FILE *fp, int n){
    int *mat = (int *)malloc(n * n * sizeof(int));

    //ignore first comment line
    fscanf(fp, "%*[^\n]\n", NULL);
    for(int i = 0; i < n*n; i++){
        int val;
        fscanf(fp, "%*d,%*d,%d", &val);
        mat[i] = val;
    }
    return mat;
}

int multiplyMatrices(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            int x = a[i*n+j];
            for(int k = 0; k < n; k++){
                mat[k+i*n] += x * b[j*n+k];
            }
        }
    }
    for(int f = 0; f < n*n; f++)
        printf("%d ", *mat++);

    return 0;
}

int main(int argc, char *argv[]){
    int n = atoi(argv[1]);
    FILE *fp1 = fopen(argv[2], "r");
    FILE *fp2 = fopen(argv[3], "r");
    int* m1 = matrixFromFile(fp1, n);
    int* m2 = matrixFromFile(fp2, n);
    multiplyMatrices(m1, m2, n);
    return 0;
}
