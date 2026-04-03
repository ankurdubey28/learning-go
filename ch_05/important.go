package main

import "fmt"

func main() {
	// ---------------- MAP EXAMPLE ----------------

	mp := map[int]int{
		1: 1,
	}
	// mp is a map (reference type).
	// The variable 'mp' holds a reference to a runtime-managed hash table.
	// When mp is passed to a function, the reference is copied,
	// but both caller and callee point to the SAME underlying map.

	// ---------------- SLICE EXAMPLE ----------------

	sl := make([]int, 0, 3)
	// sl is a slice created with:
	// len = 0  → no elements are logically visible
	// cap = 3  → backing array has space for 3 elements
	//
	// Slice is a small descriptor:
	//   (pointer to backing array, len, cap)
	//
	// IMPORTANT:
	// Passing a slice to a function COPIES this descriptor.
	// The backing array may be shared, but the header is not.

	// Modify map and slice
	modMap(mp)
	modSlice(sl)

	// Print results
	fmt.Println(mp) // map[1:1 2:3 3:4]
	fmt.Println(sl) // []

	// Even though append wrote '5' into the backing array,
	// the caller's slice length is still 0,
	// so the element is out of bounds and invisible.
}

// ---------------- FUNCTIONS ----------------

func modMap(mp map[int]int) {
	// mp is a copy of the map header,
	// but it still references the SAME hash table as the caller.

	mp[2] = 3
	mp[3] = 4
	// These operations mutate the shared map data directly.
	// Therefore, the caller sees the changes.
}

func modSlice(sl []int) {
	// sl is a COPY of the slice header (ptr, len, cap).

	sl = append(sl, 5)
	// append performs two actions:
	//
	// 1) Writes '5' into the backing array at index len (0).
	//    Since cap > len, no new array is allocated.
	//
	// 2) Returns a NEW slice header with len = 1.
	//
	// The new header is assigned ONLY to the local variable 'sl'.
	// The caller's slice header remains unchanged (len = 0).

	// Result:
	// backing array: [5 _ _]
	// caller slice:  len = 0 → prints []
}

// ---------------- KEY TAKEAWAYS ----------------
//
// 1. Maps are reference types:
//    - Passing a map shares the same underlying data.
//    - Mutations are always visible to the caller.
//
// 2. Slices share backing arrays, NOT slice headers:
//    - Element mutation (sl[i]) is visible if i < caller's len.
//    - append changes the slice header (len/cap), so it is NOT visible
//      unless the slice is returned or passed as *[]T.
//
// 3. Capacity allows writing, but ONLY length allows seeing.
