#include<stdlib.h>
#include<stdio.h>
#include<omp.h>

typedef unsigned int myentry;
typedef unsigned int myindex;

void multMatrices(int N, int M, int K, myentry *A, myentry *B, myentry *C) {
#pragma omp parallel for 
    for(myindex i = 0; i < N; i++){
        for(myindex j = 0; j < M; j++){
            myentry x = A[i*N+j];
            for(myindex k = 0; k < K; k++)
                C[k+i*N] += x * B[j*N+k];
        }
    }
}

int x = 0;
int nextPR() {
    x =  (x+234532)*((x>> 5 )+12234);
    return x & 16383;
}

long int hash(long int a, long int b) { return (a  | a<<27)*(b+2352351);}

int main(int argc, char **argv){
    if(argc != 3) {
        printf("Usage: mult N seed\n");
        exit(1);
    }

    omp_set_num_threads(4);
    myindex N = atoi(argv[1]);
    x = atoi(argv[2]);

    myentry *A = malloc( N*N*sizeof(myentry));
    myentry *B = malloc( N*N*sizeof(myentry));
    myentry *C = malloc( N*N*sizeof(myentry));

    if( A == NULL || B==NULL || C==NULL ) {
        printf("Could not allocate memory");
        exit(2);
    }

    for(int i=0; i< N; i++) {
        myindex row = i*N;
        for(int col=0; col<N; col++) {
            myindex index = row+col;
            A[index] = nextPR();
            B[index] = nextPR();
        }
    }
    multMatrices(N,N,N, A,B,C);

    int h = atoi(argv[2]);
    for(int k=0;k<3;k++)
        for(int i=0; i< N*N; i++) {
            //    printf("%f ", C[i]);
            h = hash(h, (long int) C[i]);
        }
    printf( "%d\n", h & 1023);
    return 0;
}
