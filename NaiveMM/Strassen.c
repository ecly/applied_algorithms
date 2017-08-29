// More or less a C conversion of http://www.sanfoundry.com/java-program-strassen-algorithm/
#include <stdio.h>
#include <stdlib.h>

int* matrixFromFile(FILE *fp, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));

    //ignore first comment line
    fscanf(fp, "%*[^\n]\n", NULL);
    for(int i = 0; i < n*n; i++){
        int val;
        fscanf(fp, "%*d,%*d,%d", &val);
        mat[i] = val;
    }
    return mat;
}

int split(int* from, int* to, int colOff, int rowOff, int size){
    int offset = colOff*size*2+rowOff;
    for(int i = 0; i < size; i++){
        for(int j = 0; j < size; j++){
            int loc = i*size+j;
            to[loc] = from[offset+loc];
            //printf("offset: %d\n", offset);
            //printf("loc: %d\n", loc);
            //printf("to[loc]: %d\n\n", to[loc]);
        }
    }
}

int join(int* from, int* to, int colOff, int rowOff, int size){
    int offset = colOff*size*2+rowOff;
    for(int i = 0; i < size; i++){
        for(int j = 0; j < size; j++){
            int loc = i*size+j;
            to[loc] = from[offset+loc];
        }
    }
}

int* add(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for (int i = 0; i < n; i++){
        for (int j = 0; j < n; j++){
            int loc = i*n+j;
            mat[loc] = a[loc] + b[loc];
        }
    }
    return mat;
}

int* sub(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for (int i = 0; i < n; i++){
        for (int j = 0; j < n; j++){
            int loc = i*n+j;
            mat[loc] = a[loc] - b[loc];
        }
    }
    return mat;
}

int* mult(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));

    // Base cases
    if(n == 1){
        mat[0] = a[0] * b[0];
    } 
    else {
        // Create new sub matrices
        int newN = n/2;
        int matrixSize = newN*newN;;
        int *a11 = (int *)malloc(matrixSize * sizeof(int));
        int *a12 = (int *)malloc(matrixSize * sizeof(int));
        int *a21 = (int *)malloc(matrixSize * sizeof(int));
        int *a22 = (int *)malloc(matrixSize * sizeof(int));

        int *b11 = (int *)malloc(matrixSize * sizeof(int));
        int *b12 = (int *)malloc(matrixSize * sizeof(int));
        int *b21 = (int *)malloc(matrixSize * sizeof(int));
        int *b22 = (int *)malloc(matrixSize * sizeof(int));

        // Fill sub matrices
        split(a, a11, 0, 0, newN);
        //printf("a11: %d\n", *a11);
        split(a, a12, 0, newN, newN);
        //printf("a12: %d\n", *a12);
        split(a, a21, newN, 0, newN);
        //printf("a21: %d\n", *a21);
        split(a, a22, newN, newN, newN);
        //printf("a22: %d\n", *a22);
        split(b, b11, 0, 0, newN);
        //printf("b11: %d\n", *b11);
        split(b, b12, 0, newN, newN);
        //printf("b12: %d\n", *b12);
        split(b, b21, newN, 0, newN);
        //printf("b21: %d\n", *b21);
        split(b, b22, newN, newN, newN);
        //printf("b22: %d\n", *b22);

        int *p1 = mult(add(a11, a22, newN), add(b11, b22, newN), newN);
        printf("p22: %d\n", *p1);
        int *p2 = mult(add(a21, a22, newN), b11, newN);
        int *p3 = mult(a11, sub(b12, b22, newN), newN);
        int *p4 = mult(a22, sub(b21, b11, newN), newN);
        int *p5 = mult(add(a11, a12, newN), b22, newN);
        int *p6 = mult(sub(a21, a11, newN), add(b11, b12, newN), newN);
        int *p7 = mult(sub(a12, a22, newN), add(b21, b22, newN), newN);

        int *c11 = add(sub(add(p1, p4, newN), p5, newN), p7, newN);
        int *c12 = add(p3, p5, newN);
        int *c21 = add(p2, p4, newN);
        int *c22 = add(sub(add(p1, p3, newN), p2, newN), p6, newN);

        join(c11, mat, 0, 0, newN);
        join(c12, mat, 0, newN, newN);
        join(c21, mat, newN, 0, newN);
        join(c22, mat, newN, newN, newN);
    }
    return mat;
}

int main(int argc, char *argv[]){
    int n = atoi(argv[1]);
    FILE *fp1 = fopen(argv[2], "r");
    FILE *fp2 = fopen(argv[3], "r");
    int* m1 = matrixFromFile(fp1, n);
    int* m2 = matrixFromFile(fp2, n);
    int* mat = mult(m1, m2, n);
    for(int f = 0; f < n*n; f++)
        printf("%d ", *mat++);
    return 0;
}
