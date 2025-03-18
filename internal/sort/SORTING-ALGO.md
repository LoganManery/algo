Here are the major sorting algorithms:

1. **Comparison-based algorithms**:
   - **Bubble Sort**: Repeatedly steps through the list, compares adjacent elements and swaps them if they're in the wrong order.
   - **Selection Sort**: Repeatedly finds the minimum element from the unsorted portion and puts it at the beginning.
   - **Insertion Sort**: Builds the sorted array one item at a time by taking elements from the unsorted portion and inserting them at the correct position.
   - **Merge Sort**: Divides the array into halves, sorts each half, then merges them back together.
   - **Quick Sort**: Picks a "pivot" element and partitions the array around it, then recursively sorts the sub-arrays.
   - **Heap Sort**: Converts the array into a heap data structure, then repeatedly extracts the maximum element.
   - **Shell Sort**: Improvement of insertion sort that allows the exchange of items that are far apart.
   - **Tim Sort**: Hybrid of merge sort and insertion sort, designed to perform well on many kinds of real-world data.

2. **Non-comparison based algorithms**:
   - **Counting Sort**: Works by counting objects having distinct key values, then doing arithmetic to calculate positions.
   - **Radix Sort**: Sorts integers by processing individual digits, starting from least significant digit to most significant.
   - **Bucket Sort**: Distributes elements into buckets, then sorts these buckets individually.

3. **Specialized algorithms**:
   - **Topological Sort**: For directed graphs, arranges vertices in order where for every edge (u,v), vertex u comes before v.
   - **External Sort**: For data that doesn't fit into memory, involves sorting in chunks and merging.

Each algorithm has different time and space complexity characteristics, making them suitable for different scenarios depending on factors like data size, memory constraints, and whether the data is already partially sorted.

Would you like me to explain any of these algorithms in more detail?