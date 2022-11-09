A hashtable constitutes a hash function that converts incoming data into a unique hash value which is used as an index in the underlying array structure.  
This makes it extremely efficient for searching, as the same value always produces the same key.  
Often, hashtables have an aspect of collisstion handling, which aims at solving the problem encountered when different values produce the same hash value.  
These are:   
1. Open addressing:  
This adds an extra step when we try insert into a value that is already occupied and just stores it in the next address.  
It also adds a step when retrieving as we have to check whether the value exists in the positions after. This gets inefficient with more values as the "real" position can be unpredictable
2. Separate chaining:  
This mostly uses a linked list as the value inside an index that contains multiple values.  
this allows potentially multiple values to be contained in the same key, but in order to search we would need to sift through the entire linked list in order to find the value.  
This is still faster than manually searching an array, but the internal search within an index takes O(n) in the worst case.  
