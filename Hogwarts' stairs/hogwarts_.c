#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

/* Nodes colors */
int* visited;

/* Create a new weighted edge in the dijkstra-compatible graph. */
int build_edge(int edge_index, int past_weight, int* weights, int* A, int* B, int* start, int* end, int M, int N);

/* Return 0 if a given edge exists at time "time", 1 otherwise. */
int exists(int edge_index, int time, int* stairs_appear, int* stairs_disappear);

/* Returns weight for edge at edge_index at time "time", -1 if not existing. */
int weight(int edge_index, int time);

/* Get neighbors for node "node". */
int* neighbors(int node, int* A, int* B, int M, int N);

/* Visit neighbors. */
void visit_neighbors(int node, int weight, int* weights, int* A, int* B, int* start, int* end, int M, int N);


int raggiungi(int N, int M, int A[], int B[], int* inizio, int* fine) {
	int* ancestor_neighbors = (int*) malloc(N * sizeof(int));
	int* weights = (int*) malloc(M * sizeof(int));
	// Startup values
	visited = (int*) malloc(N * sizeof(int));
	visited[0] = 1;

	int i = 0;
	for (i = 1; i < N; i++) visited[i] = 0;

	ancestor_neighbors = neighbors(0, A, B, M, N);

	if (ancestor_neighbors == NULL) return -1;
	visited[0] = 1;

	// Build neighborhood
	for (int i, cur_edge = 0; ancestor_neighbors[i] != -1; i++) {
		// Find start time
		for (cur_edge = 0; cur_edge < M && A[cur_edge] == 0 && B[cur_edge] == ancestor_neighbors[i]; cur_edge++);

		printf("Raggiungi building edge\n");
		int new_weight = build_edge(cur_edge, inizio[cur_edge], weights, A, B, inizio, fine, M, N);
		weights[cur_edge] = new_weight;
		printf("Edge built, visiting ancestor neighborhood %d at time %d\n", ancestor_neighbors[i], new_weight);
		visit_neighbors(ancestor_neighbors[i], new_weight, weights, A, B, inizio, fine, M, N);
		printf("done\n");
	}

	for (int i = 0; i < M; i++) {
		printf("Arc %d -> %d, w: %d\n", A[i], B[i], weights[i]);
	}
	return -1;
}


int* neighbors(int node, int* A, int* B, int M, int N) {
	printf("Neighbors for node %d...\n", node);
	int* node_neighbors = (int*) malloc(N * sizeof(int));
	int j = -1;

	for (int i = 0; i < M; i++) {
		if (A[i] == node) {
			printf("j: %d\n", j);
			j++;
			printf("%d -> %d\n", node, B[i]);
			printf("j: %d\n", j);
			node_neighbors[j] = B[i];
		}
	}

	printf("-j: %d\n", j);
	node_neighbors[++j] = -1;
	printf("-j: %d\n", j);
	printf("Done with the neighbors.\n");

	// Realloc size
	node_neighbors = (int*) realloc(node_neighbors, j);
	printf("Reallocated %d neighbors for %d\n", j, node);

	return node_neighbors;
}


int build_edge(int edge_index, int past_weight, int* weights, int* A, int* B, int* start, int* end, int M, int N) {
	printf("Building edge %d -> %d at time %d\n", A[edge_index], B[edge_index], past_weight);

	// Edge already disappeared
	if (past_weight >= end[edge_index]) {
		printf("Time already gone, +infinite weight: %d -> %d: %d\n", A[edge_index], B[edge_index], INT_MAX);
		weights[edge_index] = INT_MAX;
		return INT_MAX;
	}

	// Edge is available with a waiting time
	if (past_weight < start[edge_index]) {
		printf("Need to wait...%d -> %d: %d\n", A[edge_index], B[edge_index], start[edge_index] + 1);
		weights[edge_index] = start[edge_index] + 1;
		return start[edge_index] + 1;
	}

	// Edge is available immediatly
	if (past_weight >= start[edge_index] && past_weight < end[edge_index])  {
		printf("Can go right away...%d -> %d: %d\n", A[edge_index], B[edge_index], past_weight + 1);
		weights[edge_index] = past_weight + 1;
		return past_weight + 1;
	}

	return INT_MAX;
}

void visit_neighbors(int node, int weight, int* weights, int* A, int* B, int* start, int* end, int M, int N) {
	printf("Visiting neighbors for %d - %d\n", node, visited[node]);
	if (visited[node]) {
		printf("Node %d visited, return\n", node);
		return;
	}

	printf("putting visited[%d] to 1\n", node);
	visited[node] = 1;
	int* node_neighbors = neighbors(node, A, B, M, N);
	int num_neighbors = sizeof(node_neighbors) / sizeof(int);
	printf("Node %d has %d neighbors\n", node, num_neighbors);
	int* neighbors_edges = (int*) malloc(num_neighbors * sizeof(int));
	printf("Node edges\n");

	for (int i = 0; i < N; i++) {
		int j = 0;

		printf("%d --- %p: vis: %d\n", node, node_neighbors, visited[node]);
		// Find edges indexes
		if (node_neighbors == NULL || node_neighbors[0] == -1) {
			printf("no neighbors for %d\n", node);
			return;
		}

		for (i = 0; j < num_neighbors && node_neighbors[j] != -1; i++) {
			if (A[i] == node) {
				printf("node_edge, vis: %d\n", visited[node]);
				neighbors_edges[j] = i;
				j++;
			}
		}

		printf("node_edges done\n");
		//for (i = 0; node_neighbors[i] != -1 && !visited[node_neighbors[i]]; i++) {
		for (i = 0; i < N || node_neighbors == NULL || node_neighbors[i] == -1; i++) {
			printf("Visit neighbors: %d | visited; %d\n", num_neighbors, visited[node_neighbors[i]]);
			printf("Edge for %d: %d -> %d\n", node, node, node_neighbors[i]);
			int new_weight = build_edge(neighbors_edges[i], weight, weights, A, B, start, end, M, N);
			visit_neighbors(node_neighbors[i], new_weight, weights, A, B, start, end, M, N);
		}
	}
}

//
// void dijkstra(){
//
// }
