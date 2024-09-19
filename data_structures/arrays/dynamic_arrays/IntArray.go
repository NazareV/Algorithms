/*
Go already has a very efficient dynamic array implementation called slices.
Slices are dynamic arrays with automatic resizing built-in, making them suitable for many of the operations you're trying to implement in a custom way in Java.
This eliminates the need for manually resizing arrays and handling capacity.
*/

package main

import (
	"fmt"
	"sort"
)

// IntArray struct represents a dynamic array of integers
type IntArray struct {
	arr []int // The array is represented as a slice of integers
}

// NewIntArray creates and returns an empty dynamic array
func NewIntArray() *IntArray {
	return &IntArray{
		arr: []int{}, // Initialize an empty slice of integers
	}
}

// Size returns the number of elements in the dynamic array
func (ia *IntArray) Size() int {
	return len(ia.arr) // len() returns the length of the slice
}

// IsEmpty checks if the array is empty (i.e., has zero elements)
func (ia *IntArray) IsEmpty() bool {
	return len(ia.arr) == 0 // Check if the length of the array is zero
}

// Get returns the element at the specified index, with error handling for out-of-bounds access
func (ia *IntArray) Get(index int) (int, error) {
	if index < 0 || index >= len(ia.arr) { // Check if index is within valid bounds
		return 0, fmt.Errorf("index out of bounds") // Return an error if index is invalid
	}
	return ia.arr[index], nil // Return the element at the given index
}

// Set modifies the element at the specified index, with error handling for out-of-bounds access
func (ia *IntArray) Set(index int, elem int) error {
	if index < 0 || index >= len(ia.arr) { // Check if index is within valid bounds
		return fmt.Errorf("index out of bounds") // Return an error if index is invalid
	}
	ia.arr[index] = elem // Modify the element at the given index
	return nil
}

// Add appends a new element to the end of the array
func (ia *IntArray) Add(elem int) {
	ia.arr = append(ia.arr, elem) // Append the new element to the end of the slice
}

// RemoveAt removes the element at the specified index, shifting the elements to fill the gap
func (ia *IntArray) RemoveAt(index int) error {
	if index < 0 || index >= len(ia.arr) { // Check if index is within valid bounds
		return fmt.Errorf("index out of bounds") // Return an error if index is invalid
	}
	// Rebuild the array without the element at index by slicing around it
	ia.arr = append(ia.arr[:index], ia.arr[index+1:]...)
	return nil
}

// Remove removes the first occurrence of the specified element, if found
func (ia *IntArray) Remove(elem int) error {
	for i, e := range ia.arr {
		if e == elem { // Check if the current element matches the target element
			ia.RemoveAt(i) // Remove the element at the found index
			return nil     // Return success
		}
	}
	return fmt.Errorf("element not found") // Return an error if the element is not found
}

// Reverse reverses the order of elements in the array
func (ia *IntArray) Reverse() {
	// Use two indices, one starting from the beginning and one from the end
	for i, j := 0, len(ia.arr)-1; i < j; i, j = i+1, j-1 {
		ia.arr[i], ia.arr[j] = ia.arr[j], ia.arr[i] // Swap elements at i and j
	}
}

// BinarySearch performs a binary search for the specified key in the array
// (Assumes the array is sorted)
func (ia *IntArray) BinarySearch(key int) int {
	// Use sort.Search to find the position where the element should be
	index := sort.Search(len(ia.arr), func(i int) bool { return ia.arr[i] >= key })
	if index < len(ia.arr) && ia.arr[index] == key {
		return index
	}
	return -1
}

// Sort sorts the array in ascending order
func (ia *IntArray) Sort() {
	sort.Ints(ia.arr) // Use sort.Ints to sort the slice of integers in ascending order
}

// String returns a string representation of the array
func (ia *IntArray) String() string {
	return fmt.Sprintf("%v", ia.arr) // Format the slice as a string using fmt.Sprintf
}

func main() {
	ar := NewIntArray() // Create a new dynamic array

	ar.Add(3)  // Add element 3 to the array
	ar.Add(7)  // Add element 7 to the array
	ar.Add(6)  // Add element 6 to the array
	ar.Add(-2) // Add element -2 to the array

	ar.Sort() // Sort the array in ascending order: [-2, 3, 6, 7]

	// Print each element of the sorted array
	for i := 0; i < ar.Size(); i++ { // Iterate through the array using Size() to get the length
		val, _ := ar.Get(i) // Get the value at index i (ignore the error for simplicity)
		fmt.Println(val)    // Print the value
	}

	// Print the array using its string representation
	fmt.Println(ar) // Prints: [-2, 3, 6, 7]
}
