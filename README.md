# Advent of Code

## Overview
"Decided to dive back into Go, and what better way to learn than doing Advent of Code.


### Key Takeaways

1. **Stack / Heap**
   Variables declared in a function are often on the stack, while dynamically allocated ones (like slices grown beyond a certain size) go to the heap.
   ```go
   func example() {
       x := 42                // Likely stored on the stack
       largeSlice := make([]int, 1000000) // Allocated on the heap
       fmt.Println(x, largeSlice)
   }
   ```

2. **Pointer Dereferencing**
   Accessing the value behind a pointer.
   ```go
   func example() {
       x := 42
       ptr := &x            // Pointer to x
       fmt.Println(*ptr)    // Dereferences ptr to access x's value
   }
   ```

3. **Contiguous Memory Allocation**
   Slices store elements in sequential memory, allowing fast iteration.
   ```go
   func example() {
       nums := []int{1, 2, 3, 4}
       for i, n := range nums {
           fmt.Printf("Index: %d, Value: %d\n", i, n)
       }
   }
   ```

4. **Slices / Hash Map**
   Slices for ordered data, maps for key-value pairs.
   ```go
   func example() {
       slice := []int{10, 20, 30}
       fmt.Println("Slice:", slice)

       hashMap := map[string]int{"A": 1, "B": 2}
       fmt.Println("Hash Map:", hashMap)
   }
   ```

5. **Function with Multiple Return Values**
   Returning multiple values directly from a function.
   ```go
   func divide(a, b int) (int, error) {
       if b == 0 {
           return 0, fmt.Errorf("division by zero")
       }
       return a / b, nil
   }
   func example() {
       result, err := divide(10, 2)
       if err != nil {
           fmt.Println("Error:", err)
       } else {
           fmt.Println("Result:", result)
       }
   }
   ```

6. **Error Handling**
   Explicit handling of errors using the returned error value.
   ```go
   func readFile(fileName string) error {
       file, err := os.Open(fileName)
       if err != nil {
           return fmt.Errorf("failed to open file: %w", err)
       }
       defer file.Close()
       return nil
   }
   ```

7. **`make` and `new`**
   Allocating slices, maps, and channels with `make`, and pointers with `new`.
   ```go
   func example() {
       slice := make([]int, 5)          // Creates a slice with 5 elements
       hashMap := make(map[string]int) // Creates a map
       channel := make(chan int)       // Creates a channel

       ptr := new(int)  // Allocates a pointer to an int initialized to 0
       fmt.Println(slice, hashMap, channel, *ptr)
   }
   ```

8. **If Special Syntax**
   Declaring and checking variables inline in an `if` statement.
   ```go
   func example() {
       myMap := map[string]int{"A": 1, "B": 2}
       if value, exists := myMap["A"]; exists {
           fmt.Printf("Found: %d\n", value)
       } else {
           fmt.Println("Not found")
       }
   }
   ```
