// intended solution: naive matrix multiplication in C
// itu, course APALG
// by Riko Jacob
// first created Fall 16
 
#include<stdlib.h>
#include<stdio.h>
#include<omp.h>
 
typedef double myfloat;
typedef unsigned int myindex;
   
void MxMnaive(int N, int M, int K, myfloat *A, myfloat *B, myfloat *C) {
    for(int i = 0; i < N; i++){
        for(int j = 0; j < M; j++){
            int x = A[i*N+j];
            for(int k = 0; k < K; k++)
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
 
  if( argc != 3 ) {
    printf("Usage: mult N seed\n");
    exit(1);
  }
 
  omp_set_num_threads(1);
   
  myindex N = atoi(argv[1]);
  x = atoi(argv[2]);
   
  myfloat *A = malloc( N*N*sizeof(myfloat));
  myfloat *B = malloc( N*N*sizeof(myfloat));
  myfloat *C = malloc( N*N*sizeof(myfloat));
 
  int i,j;
 
  if( A == NULL || B==NULL || C==NULL ) {
    printf("Could not allocate memory");
    exit(2);
  }
 
  for(i=0; i< N; i++) {
    for(j=0; j<N; j++) {
      A[i*N+j] = nextPR();
      B[i*N+j] = nextPR();
      C[i*N+j] = 0;
    }
  }
  MxMnaive(N,N,N, A,B,C);
 
  int h = atoi(argv[2]);
  for(int k=0;k<3;k++)
    for(int i=0; i< N*N; i++) {
      //    printf("%f ", C[i]);
      h = hash(h, (long int) C[i]);
    }
  printf( "%d\n", h & 1023);
  return 0;
}
