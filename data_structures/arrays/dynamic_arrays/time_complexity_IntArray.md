# IntArray Implementation

This project contains an implementation of a dynamic array for integers in Go. The dynamic array supports various operations such as adding, removing, sorting, and accessing elements.

## Time Complexity Analysis

Let's analyze the total time complexity of the code, focusing on the operations inside the `main()` function. Below is a breakdown of each operation:

### Operations in `main()`

1. **Creating an Instance**: `NewIntArray()`

   - **Complexity**: O(1)
   - This operation creates an empty dynamic array.

2. **Adding Elements**: `ar.Add(3)`, `ar.Add(7)`, `ar.Add(6)`, `ar.Add(-2)`

   - **Complexity**: O(1) (amortized for each)
   - Appending elements to a slice takes amortized O(1) time. For 4 additions, the total complexity is:
     - O(1) + O(1) + O(1) + O(1) = O(4) ≈ O(1) (constants are dropped in Big-O notation).

3. **Sorting the Array**: `ar.Sort()`

   - **Complexity**: O(n log n)
   - Sorting the array takes O(n log n), where `n` is the number of elements. For `n = 4`, this translates to:
     - O(4 log 4) = O(4 \* 2) = O(8) ≈ O(n log n).

4. **Iterating Through Elements**: `for i := 0; i < ar.Size(); i++ { ar.Get(i) }`

   - **Complexity**: O(n)
   - The loop runs `n` times, with each iteration performing a constant time O(1) operation (`Get(i)`). Thus, the total time for the loop is:
     - O(n) = O(4) = O(n).

5. **Printing the Array**: `fmt.Println(ar)`
   - **Complexity**: O(n)
   - Converting the array to a string requires iterating over all `n` elements, resulting in:
     - O(n).

### Total Time Complexity

Combining all the steps gives us:

1. **O(1)** for `NewIntArray()`
2. **O(1)** for adding elements
3. **O(n log n)** for sorting
4. **O(n)** for iterating and accessing elements
5. **O(n)** for printing the array

The largest term dominates the overall complexity, which is **O(n log n)** due to the sorting operation.

### Final Total Time Complexity: **O(n log n)**

This is the time complexity for the entire code in `main()`, where `n` is the number of elements in the array (in this case, `n = 4`).
