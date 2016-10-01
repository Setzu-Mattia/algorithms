#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define TRUE 1
#define FALSE 0

/* Pick the best distance for a given node. */
int pick(int* distance, int n, short int* found);

int raggiungi(int N, int M, int A[], int B[], int* start, int* end) {
	return -1;
}

void dijkstra(int vertex, int* weights, int* distance, int n, short int* found) {
	int i, u, w;

	for (i = 0; i < n; i++) {
		found[i] = FALSE;
		distance[i] = weights[i];
	}

	found[vertex] = TRUE;
	distance[vertex] = 0;

	for (i = 0; i < n-2; i++) {
		u = pick(distance, n, found);
		found[u] = TRUE;

		for (w = 0; w < n; w++){
			if(!found[w])
				if (distance[u] + weights[u][w] < distance[w])
					distance[w] = distance[u] + weights[u][w];
		}
	}
}

int pick(int* distance, int n, short int* found) {
	int i, min, min_pos;

	min = INT_MAX;
	min_pos = -1;

	for (i = 0; i < n; i++)
		if(distance[i] < min && !found[i]) {
			min = distance[i];
			min_pos = i;
		}

	return min_pos;
}
