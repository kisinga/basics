Quick Sort is also known as Partition Sort.   
This sorting algorithm is faster than the previous algorithms because this algorithm uses the concept of Divide and Conquer  
  
First, we decide a pivot element. We then find the correct index for this pivot position and then divide the array into two subarrays.  
One subarray will contain elements that are lesser than our pivot and other subarray will contain elements that are greater than our pivot.
(unsorted)  
We then recursively call these two subarrays and the process goes on until we can’t further divide the array.  
  
But how do we divide the subarray?  
We assume the last element of our array to be our pivot. We then traverse the entire array using two pointer technique. The elements encountered by the left pointer should be lower than the pivot and elements encountered by the right pointer should be greater than pivot. If not, we swap the elements at the left and right pointers so that for a particular position in the array, the elements to the left are lower while higher at the right. We then insert our pivot at this position.

Best-case: O(nlogn)- First, we are dividing the array into two subarrays recursively which will cost a time complexity of O(logn). For each function call, we are calling the partition function which costs O(n) time complexity. Hence the total time complexity is O(nlogn).  
  
Worst-case: O(n²)- When the array is sorted in descending order or all the elements are the same in the array, the time complexity jumps to O(n²) since the subarrays are highly unbalanced.

Links:  
https://www.youtube.com/watch?v=COk73cpQbFQ