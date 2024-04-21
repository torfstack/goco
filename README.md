## goco

`goco` is a library to construct quantum circuits in Go. It is inspired by Qiskit, a quantum computing library in Python.
Quantum circuits can be constructed with an unlimited number of qubits and gates. The library also supports the simulation of quantum circuits.
As the simulation is based on the statevector representation, the number of qubits should be limited to 20 or less to avoid memory issues.

### Installation

```bash
go get github.com/torfstack/goco
```

### Example

```go
package main

import (
    "fmt"
    "github.com/torfstack/goco"
)

func main() {
    // Create a quantum circuit with 2 qubits
    qs := goco.NewSystem(2) // 2 qubits

    // Apply a Hadamard gate to the first qubit
    qs.H(0)

    // Apply a CNOT gate to the first and second qubits
    qs.CNOT(0, 1)

    // Simulate the circuit
    b := goco.NewLinearAlgebraBackend(qs)
    result := b.Simulate()

    // Print the result: [0.5, 0, 0, 0.5]
    fmt.Println(result)
}
```