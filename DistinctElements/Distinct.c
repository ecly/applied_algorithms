#include <stdio.h>
#include <limits.h>
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
       res += (__builtin_popcount(A[bits-1-i]&x)%2)<<i;

    return res;
}

//as given
int f(int x){
    return ((x*0xbc164501) & 0x7fe00000) >> 21;
}

int max(int a, int b){
    if(a > b)
        return a;
    else
        return b;
}

float silly_sum(int* arr, int size){
    float sum = 0;
    for(int i = 0; i < size; i++){
            sum += pow(2,-arr[i]);
            //printf("added: %6f \n", pow(2,-arr[i]));
    }
    return sum;
}

int find_v(int* arr, int size){
    for(int i = 0; i < size; i++){
        if(arr[i] == 0)
            return i;
    }
    return 0;
}
    
int main(){
    int threshold, input;
    scanf("%d", &threshold);
    int m = 1024;
    int* M = malloc (sizeof (int) * m);;
    memset (M, 0, sizeof (int) * m);

    while(scanf("%d", &input) != EOF){
        int j = f(input);
        //printf("m[%d] is %d \n", j, M[j]);
        //printf("max M[%d], %d = %d \n",j, __builtin_clz(hash(input, 32)), max(M[j], __builtin_clz(hash(input, 32))));
        //printf("leading zeros: %d \n", __builtin_clz(hash(input, 32)));
        M[j] = max(M[j], __builtin_clz(hash(input, 32)));
        //printf("M[%d] is now %d \n", j, M[j]);
    }
    float Z = 1.0/silly_sum(M, m);
    printf("Z = %6f\n", Z);
    int V = find_v(M, m);
    printf("V = %d\n", V);
    float E = m*m*Z*0.7213/(1 + 1.079/m);
    printf("E = %6f\n", E);
    if (E < 2.5*m && V > 0)
        E = m * log(m/V);

    printf("Changed E = %6f\n", E);

    if(E > threshold)
        printf("above");
    else
        printf("below");

    return 0;
}
