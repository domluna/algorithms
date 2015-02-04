Graph Representations
---------------------

O(E) - Edges
O(V) - Vertices

### Edge List

Have a list of all the edges in the graph. If there is a weight associated with
the edge add that property to the structure to represent edges.

O(E) Memory

Query time is O(E) as well since we are naively not storing edges in
any particular order.

Question: Can we store the edges ways that will improve our query time?

### Adjacency Matrix

Have a matrix and 0s and 1s. If A[i, j] = 1, then that means there is an edge
from vertices i and j. If A[i, j] = 0 there is no edge.

O(V^2) Memory

If the matrix is sparse then we can reduce our memory costs.

Query time is O(1). A Matrix is just a 2D array. If you want to find out which
vertices are connected a vertex then that's O(V).

### Adjacency List

Every vertex points to a list of the other vertices it's connected to. This can still
be O(V^2) if all vertices are connected to all other vertices. Query time is O(d), d
is the degree of vertex i. Note: d can be V-1.

Typically V^2 is too negative of an upper bound O(E) is more representative. E, of course, can be V^2 but this is very unlikely.

