#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

typedef struct {
	int path_length;
	int time;
} Path;


/* Copy a path. */
Path* copy_Path(Path* path);

/* Walk from node a to node b, returns the cost for said path, time added. */
Path* walk(int a, int b, int time, int A[], int B[], int inizio[], int fine[], int M, int N, Path* walk_path);

/* Explore neighbors nodes. */
void walk_neighborhood(int start_node, Path* history, int A[], int B[], int inizio[], int fine[], int M, int N);

/* The optimal path */
Path* optimal_path;

/* Nodes colors */
int* colors;


int raggiungi(int N, int M, int A[], int B[], int inizio[], int fine[]) {
	int i = 0;
	colors = (int*) malloc(N * sizeof(int));

	Path* opt_path = (Path*) malloc(sizeof(Path));
	opt_path->path_length = INT_MAX;
	opt_path->time = 0;

	optimal_path = opt_path;
	for (i = 0; i < N; i++) colors[i] = 0;

	colors[0] = 1;
	walk_neighborhood(0, opt_path, A, B, inizio, fine, M, N);

	return optimal_path->time == INT_MAX ? -1 : optimal_path->time;
}


void walk_neighborhood(int start_node, Path* history, int A[], int B[], int inizio[], int fine[], int M, int N) {
	Path* path = copy_Path(history);
	int* neighbors = (int*) malloc(N * sizeof(int));
	int i, j = 0;

	for (i = 0, j = 0; i < M; i++) {
		if (A[i] == start_node) {
			neighbors[j] = B[i];
			j++;
		}
	}

	for (i = j;i < N; i++) neighbors[i] = -1;

	free(history);
	for (i = 0; i < j; i++) {
		//printf("On edge %d - %d\n", start_node, neighbors[i]);
		Path* current = walk(start_node, neighbors[i], path->time, A, B, inizio, fine, M, N, path);
		Path* new_path = copy_Path(current);

		colors[start_node] = 1;

		// Reached N - 1
		if (neighbors[i] == (N - 1)) {
			//printf("Reached with time: %d\n", current->time);
			// First optimal time
			if (optimal_path->time == INT_MAX) {
				optimal_path = current;
				//printf("Reached first optimal path in %d.\n", optimal_path->time);
			}
			// New optimal time
			if (current->time < optimal_path->time) {
				optimal_path = current;
				//printf("Reached new optimal path in %d.\n", optimal_path->time);
			}

			// No new optimal time
			if (current->time >= optimal_path->time) {
				optimal_path = current;
				//printf("No new optimal path, current: %d\n", optimal_path->time);
			}

		}

		//free(current);
		//printf("Walking neighbor of %d\n", neighbors[i]);
		walk_neighborhood(neighbors[i], new_path, A, B, inizio, fine, M, N);
	}
}


Path* walk(int a, int b, int time, int A[], int B[], int inizio[], int fine[], int M, int N, Path* walk_path) {
	Path* path = copy_Path(walk_path);
	int stair = 0;

	// Find the stair
	for (stair = 0; A[stair] != a || B[stair] != b; stair++);

	//printf("For edge %d - %d, time %d\n", a, b, time);
	// Deadlock: stair already disappeared
	if (time > fine[stair]) {
		//printf("Deadlock on %d %d\n", a, b);
		path->time = INT_MAX;
		path->path_length = INT_MAX;
	} // Stair available
	else {
		path->path_length = path->path_length + 1;
		// /printf("Summing time for traversal %d -> %d, %d + 1 from %d\n", a, b, path->time, walk_path->time);
		path->time = (inizio[stair] <= path->time) ? path->time + 1 : inizio[stair] + 1;
	}

	return path;
}


Path* copy_Path(Path* path) {
	Path* new_path = (Path*) malloc(sizeof(Path));

	new_path->path_length = path->path_length;
	new_path->time = path->time;

	return new_path;
}
