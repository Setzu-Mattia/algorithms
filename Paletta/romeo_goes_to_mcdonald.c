#include <stdio.h>
#include <stdlib.h>
#include <tgmath.h>


int flips = 0;


/* Sort with a spatula. */
long paletta_sort(int N, int V[]);

/* Is Marco really a good hamburger flipper? */
int are_you_loving_it(int* even, int* odd, int even_size, int odd_size);

void flip_order(int *list,int n);

long paletta_sort(int N, int* V) {
    int even_size = ceil(N / 2) + 1;
    int odd_size = floor(N / 2);
    int* even = (int*) malloc(even_size * sizeof(int));
    int* odd = (int*) malloc(odd_size * sizeof(int));

    // Arrange V
    int e, o = 0;

    for (int i = 0; i < N; i++) {
        if (i % 2 == 0) {
            even[e] = V[i];
            e++;
        }
        else {
            odd[o] = V[i];
            o++;
        }
    }

    // 1-element vector
    if (N == 1) return 0;
    // 2-V vector ordered
    if (N == 2 && V[0] > V[1]) return 0;
    // 2-V vector not ordered
    if (N == 2 && V[0] < V[1]) return -1;
    // 3-V vector not ordered-able
    if (N == 3 && (V[1] < V[0] || V[1] > V[2])) return -1;
    // 3-V vector order-able
    if (N == 3 && V[0] <= V[2]
        && V[1] >= V[0] && V[1] <= V[2]) return 0;

    flip_order(even, even_size);
    flip_order(odd, odd_size);

    return are_you_loving_it(even, odd, even_size, odd_size) ? flips : -1;
}

int are_you_loving_it(int* even, int* odd, int even_size, int odd_size) {
    int fail = 0;
    int e = 0;
    int o = 0;

    do {
        // Reached evens end.
        if (e == even_size - 2 && o == odd_size - 1) {
            return even[e] <= odd[o] && odd[o] <= even[e + 1];
        }

        if (even[e] <= odd[o] && odd[o] <= even[e + 1]) {
            e++;
            o++;
        } else {
            return 0;
        }
    } while(!fail);

    return fail;
}

void flip_order(int *list,int n) {
    int i = 0, j = 1;
    int direction = 1;

    do {
        switch(direction) {
        case 1:
            if (list[i] > list[j]) {
                flips = flips + (j - i);
                if (i >= 0) {
                    j = i;
                    i--;
                }
            } else {
                i++;
                j++;
            }
        break;
        case -1:
            if (list[i] > list[j]) {
                flips = flips + (j - 1);
                i = 0;
                j = 1;
            } else {
                i++;
                j++;
            }
        break;
        }
    } while(j < n);
}