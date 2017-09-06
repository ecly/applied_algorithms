#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>
int A[32] = {0x21ae4036, 0x32435171, 0xac3338cf, 0xea97b40c, 0x0e504b22,
    0x9ff9a4ef, 0x111d014d, 0x934f3787, 0x6cd079bf, 0x69db5c31,
    0xdf3c28ed, 0x40daf2ad, 0x82a5891c, 0x4659c7b0, 0x73dc0ca8,
    0xdad3aca2, 0x00c74c7e, 0x9a2521e2, 0xf38eb6aa, 0x64711ab6,
    0x5823150a, 0xd13a3a9a, 0x30a5aa04, 0x0fb9a1da, 0xef785119,
    0xc9f0b067, 0x1e7dde42, 0xdda4a7b2, 0x1a1c2640, 0x297c0633,
    0x744edb48, 0x19adce93};

int hash(int x, int bits){
    int res = 0;
    for(int i=0; i<bits; i++)
        res += (__builtin_popcount(A[bits-1-i]&x)&1)<<i;

    return res;
}

int f(int x){
    return ((x*0xbc164501) & 0x7fe00000) >> 21;
}

int max(int a, int b){
    return a > b ? a : b;
}

float silly_sum(int* arr, int size){
    float sum = 0;
    for(int i = 0; i < size; i++)
        sum += pow(2,-arr[i]);

    return sum;
}

//https://tinodidriksen.com/uploads/code/cpp/speed-string-to-int.cpp
int f_atoi(const char *p) {
    int x = 0;
    int neg = 0;
    if (*p == '-') {
        neg = 1;
        ++p;
    }
    while (*p >= '0' && *p <= '9') {
        x = (x*10) + (*p - '0');
        ++p;
    }
    if (neg) x = -x;
    return x;
}

int main(){
    char buf[30];
    fgets_unlocked(buf, sizeof(buf), stdin);
    int threshold = f_atoi(buf);
    int m = 1024;
    int V = m;
    int* M = (int*)calloc(m, sizeof(int));
    int input;
    while(fgets_unlocked(buf, sizeof(buf), stdin) != NULL){
        input = f_atoi(buf);
        int j = f(input);
        int val = M[j];
        if(!val) V--;
        M[j] = max(val, __builtin_ffs(hash(input, 32)));
    }

    float Z = 1.0/silly_sum(M, m);
    float E = m*m*Z*0.7213/(1 + 1.079/m);
    if (E < 2.5*m && V > 0)
        E = m * log(m/(float)V);

    E > threshold ? printf("above") : printf("below");

    return 0;
}
