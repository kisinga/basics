# Solution

## Approach

- Create independent quickfind and quickunion data objects that both independently implement the `Percolation` interface.
- Create a 2D array of size `n` x `n` and initialize each object with blocked sites.