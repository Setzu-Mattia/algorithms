#include <stdio.h>
#include <stdlib.h>
#include <tgmath.h>


int flips = 0;


/* Sort with a spatula. */
long paletta_sort(int N, int V[]);

/* Is Marco really a good hamburger flipper? */
int are_you_loving_it(int* even, int* odd, int even_size, int odd_size);

void merge_sort(int *A,int n);

void merge_sort(int *A,int n);

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

    merge_sort(even, even_size);
    merge_sort(odd, odd_size);

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

void merge(int *A,int *L,int leftCount,int *R,int rightCount) {
	int i,j,k;

	// i - to mark the index of left aubarray (L)
	// j - to mark the index of right sub-raay (R)
	// k - to mark the index of merged subarray (A)
	i = 0; j = 0; k =0;

	while(i<leftCount && j< rightCount) {
		if(L[i]  < R[j]) A[k++] = L[i++];
		else {
            flips++;
            A[k++] = R[j++];
        }
	}
	while(i < leftCount) A[k++] = L[i++];
	while(j < rightCount) A[k++] = R[j++];
}

// Recursive function to sort an array of integers.
void merge_sort(int *A,int n) {
	int mid,i, *L, *R;
	if(n < 2) return; // base condition. If the array has less than two element, do nothing.

	mid = n/2;  // find the mid index.

	// create left and right subarrays
	// mid V (from index 0 till mid-1) should be part of left sub-array
	// and (n-mid) V (from mid to n-1) will be part of right sub-array
	L = (int*)malloc(mid*sizeof(int));
	R = (int*)malloc((n- mid)*sizeof(int));

	for(i = 0;i<mid;i++) L[i] = A[i]; // creating left subarray
	for(i = mid;i<n;i++) R[i-mid] = A[i]; // creating right subarray

	merge_sort(L,mid);  // sorting the left subarray
	merge_sort(R,n-mid);  // sorting the right subarray
	merge(A,L,mid,R,n-mid);  // Merging L and R into A as sorted list.
        free(L);
        free(R);
}
