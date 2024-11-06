 **pbits** is designed for bit masking and data packing in Go, enabling efficient memory management and storage of numbers in a specified number of bits. It can be used for:

- **Memory Optimization**: By using defined masks (e.g., Mask5, Mask10), you can limit the range of numbers you want to store, saving memory. For example, Mask10 allows storing numbers from 0 to 1023 using only 10 bits, instead of the standard 64 bits.

- **Packing Data** in a Specific Number of Bits: The Pack and Unpack functions enable storing and retrieving data using defined masks. This is useful when you want to store multiple values in a single 64-bit word, occupying exactly the space needed.

- **Manipulating and Storing Large Bitwise Data**: The package enables storing numerical values in a compact form, useful when serializing data for disk storage, especially when maximizing space efficiency is essential.

- **Low-level Applications**: It is ideal for systems handling large datasets, network operations, binary protocols, or hardware-level tasks where bitwise efficiency is critical.

This package can be applied in systems requiring optimal storage or transmission of large numbers of values, such as databases, index files, or data transmission systems where data is sent in binary form.

**pbits** was created to make bit manipulation more transparent and to simplify error detection in the code. By providing a structured set of masks and utility functions, it helps developers work with bits more intuitively and reduces the chances of mistakes in bit-level operations, enhancing code readability and maintainability.

### Example

Suppose we have two numbers:

- a = 1939, which we want to store using 11 bits (i.e., using Mask11).
- b = 512, which we want to store using 10 bits (i.e., using Mask10).

- Using Pack and Unpack functions, we can store these numbers in a single 64-bit integer and retrieve them.


```go
package main

import (
    "fmt"
    "github.com/malumar/pbits"
)

func main() {
    // Define the numbers to pack
    a := uint64(1939) // Number to pack with Mask11
    b := uint64(512)  // Number to pack with Mask10

    // Step 1: Pack the numbers into a single variable (bigendian format)
    packed := pbits.Pack(a, pbits.NoMask, pbits.Mask11) | 
		pbits.Pack(b, pbits.Mask11, pbits.Mask10)
    fmt.Printf("Packed value: %d\n", packed)

    // Step 2: Unpack the numbers
    unpackedA := pbits.Unpack(packed, pbits.NoMask, pbits.Mask11)
    unpackedB := pbits.Unpack(packed, pbits.Mask11, pbits.Mask10)
    
    fmt.Printf("Unpacked values: a = %d, b = %d\n", unpackedA, unpackedB)
}
```

More examples You find in pbits_test.go