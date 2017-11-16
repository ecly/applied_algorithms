// intended solution: naive matrix multiplication in C
// itu, course APALG
// by Riko Jacob
// first created Fall 16
 
#include<stdlib.h>
#include<stdio.h>
#include<omp.h>
 
typedef double myfloat;
typedef unsigned int myindex;
 
// always row, colum
// row major
myindex rm(int i, int m, int N, int M) {return i*M + m; }
// column major
myindex cm(int i, int m, int N, int M) {return m*N + i; }
   
void MxMnaive(int N, int M, int K, myfloat *A, myfloat *B, myfloat *C) {
  // computes C += AB, C: N x M, A: N x K, B: K x M, all in row major layout
  //int i,j,k;
  for(int i=0; i< N; i++)
    for(int j=0; j<M; j++)
      for(int k=0; k<K; k++) 
        C[rm(i,j, N,M)] += A[rm(i,k, N,K)] * B[rm(k,j, K,M)];  
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
      A[rm(i,j,N,N)] = nextPR();
      B[rm(i,j,N,N)] = nextPR();
      C[rm(i,j,N,N)] = 0;
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
