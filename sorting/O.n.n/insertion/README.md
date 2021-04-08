in this algo, for each element, we check if the order is correct until the current element.  
Since the first element is in order, we start from second element and check
if the order is maintained. If not, then we swap them. So, on any given element,
we check if the current element is greater than previous element. If it is not,
we keep swapping elements until our current element is greater than previous element.
https://medium.com/javarevisited/sorting-algorithms-slowest-to-fastest-a9f0e30937b9
it has a complexity of O(n^2)