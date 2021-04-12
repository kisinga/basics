A graph is a tree-like structure where nodes are called vertices and the relationships between the nodes is called an edge.  
The relationship is part of the data itself, therefore you can have directed, weighted and cyclic representations.  
The vertices can be represented in one of two major ways:  
1. Adjacency list  
2. Adjacency matrix  

Adjacency list  
Here the edges are stored in a list within the vertices. 


Adjacency matrix  
Here the edges are stored in a 2D-Array within the vertices.  
Every vertex is present in every matrix, therefore adding new vertices has a complexity of O(v*v) where v is the number of vertices.  
This is because new copies of each matrix must be created in the process.   
Searching is however very fast, at O(1) since it's very easy to check data at an index in an array  
Removing is also tidious at O(v*v)