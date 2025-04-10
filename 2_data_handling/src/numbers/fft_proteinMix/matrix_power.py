# https://www.hackerrank.com/challenges/pmix/proble
import numpy as np
import math

def pmix(s, k):

    # IDEA1: bit 0 and 1 can be independent systems
    n=len(s)
    bit0=np.zeros((n,), dtype=np.byte)
    bit1=np.zeros((n,), dtype=np.byte)

    # so prepare initial bit states separately
    i=0
    for c in s:
        ascii=ord(c)-65
        bit0[i]=ascii&1
        bit1[i]=(ascii&10)>>1
        i=i+1

    # IDEA2: prepare XOR operation
    U=np.zeros((n,n), dtype=np.byte)
    for i in range(n-1):
        U[i,i] = 1
        U[i,i+1] = 1
    
    # connect last element with the first
    U[n-1,n-1]=1
    U[n-1,0] =1

    U1=np.linalg.matrix_power(U, k) # apply it k times
    U2=np.mod(U1, [2])                 # apply modulo 2

    # calculate the final state vector applying the U^k
    v0=np.dot(U2, bit0)
    v1=np.dot(U2, bit1)

    # convert back into 'A'-'D' form
    ret=np.zeros((n), dtype=np.byte)
    for i in range(n):
        ascii=(v1[i] << 1)|v0[i]
        ret[i]=ascii+65
    
    return ''.join(chr(i) for i in ret)

def pmix1(n, k):
    U=np.zeros((n,n), dtype=np.int64)
    for i in range(n-1):
        U[i,i] = 1
        U[i,i+1] = 1
    
    # connect last element with the first
    U[n-1,n-1]=1
    U[n-1,0] =1

    # Pascal triangle
    Uk=np.linalg.matrix_power(U, k)
    print("U^k=", Uk)

    w, v = np.linalg.eig(U)

    print("w=", w)
    print("v=", v)

                               
if __name__ == '__main__':
    # f = open("input13.txt", "r")
    # nk=f.readline().split()
    # #nk = input().split()
    # n = int(nk[0])
    # k = int(nk[1])
    # s = f.readline()
    # #s = input()

    #result=pmix1(5, 7)
    #print(result + '\n')
    U=np.array([
        [1,1,0,0],
        [0,1,1,0],
        [0,0,1,1],
        [1,0,0,1]])

    # input
    vec = np.array([0,0,0,1])
    y=np.dot(U, vec)
    print(y)

    
    angle=2 * np.pi/ 4
    w= complex(math.cos(angle), math.sin(angle))

    Fn=np.array([
        [1, 1, 1, 1],
        [1, pow(w, 1), pow(w, 2), pow(w, 3)],
        [1, pow(w, 2), pow(w, 4), pow(w, 6)],
        [1, pow(w, 3), pow(w, 6), pow(w, 9)]])

    invFn=np.linalg.inv(Fn)

    vec1=np.dot(Fn, vec)

    # diag(U)
    
    diagF=np.array([
        [2, 0, 0,  0],
        [0, 1+pow(w,1), 0,  0],
        [0, 0, 1+pow(w,2), 0],
        [0, 0,  0, 1+pow(w, 3)]
    ])
    vec2=np.dot(diagF, vec1)

    #  fft
    vec3=np.dot(invFn, vec2)
    print(vec3)
