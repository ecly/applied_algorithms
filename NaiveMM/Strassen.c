// Inspired by this Java implementation: http://www.sanfoundry.com/java-program-strassen-algorithm/
// Hint of using NaiveMM for smaller from Riko Jacob
// Idea for NaiveMM optimization with 1 val of first matrix at a time from Michael Vesterli
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
            to[i*size+j] = from[offset+i*2*size+j];
        }
    }
}

int join(int* from, int* to, int colOff, int rowOff, int size){
    int offset = colOff*size*2+rowOff;
    for(int i = 0; i < size; i++){
        for(int j = 0; j < size; j++){
            to[offset+i*size*2+j] = from[i*size+j];
        }
    }
}

int* add(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for (int i = 0; i < n*n; i++){
        mat[i] = a[i] + b[i];
    }
    return mat;
}

int* sub(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for (int i = 0; i < n*n; i++){
        mat[i] = a[i] - b[i];
    }
    return mat;
}

int* multNaive(int* a, int* b, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            int x = a[i*n+j];
            for(int k = 0; k < n; k++){
                mat[k+i*n] += x * b[j*n+k];
            }
        }
    }

    return mat;
}

int* mult(int* m1, int* m2, int n){
    int *mat = (int *)malloc(n*n * sizeof(int));

    // Base cases
    if(n <= 32){
        return multNaive(m1, m2, n);
    } 
    else {
        int newN = n/2;
        // Create new sub matrices
        int* a = m1; 
        int* b = &m1[newN]; 
        int* c = &m1[newN*n]; 
        int* d = &m1[newN*n+newN]; 
        int* e = m2; 
        int* f = &m2[newN]; 
        int* g = &m2[newN*n]; 
        int* h = &m2[newN*n+newN]; 

        int *p1 = mult(a, sub(f, h, newN), newN);
        int *p2 = mult(add(a, b, newN), h, newN);
        int *p3 = mult(add(c, d, newN), e, newN);
        int *p4 = mult(d, sub(g, e, newN), newN);
        int *p5 = mult(add(a, d, newN), add(e, h, newN), newN);
        int *p6 = mult(sub(b, d, newN), add(g, h, newN), newN);
        int *p7 = mult(sub(a, c, newN), add(e, f, newN), newN);

        int *c11 = add(sub(add(p5, p4, newN), p2, newN), p6, newN);
        int *c12 = add(p1, p2, newN);
        int *c21 = add(p3, p4, newN);
        int *c22 = sub(sub(add(p5, p1, newN), p3, newN), p7, newN);

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
