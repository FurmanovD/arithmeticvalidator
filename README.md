# Arithmetic Expression Validator

This Go project implements two distinct approaches to validating arithmetic expressions consisting of:
- Decimal numbers (positive and negative)
- Operators: `+`, `-`
- Parentheses `(` and `)`
- Optional whitespace (ignored)

The validation ensures correct structure, balanced parentheses, and proper operator/operand sequences. It **does not evaluate** the expression, only validates its syntax.

---

## 📁 Project Structure

├── main.go # Entry point (example usage)
├── validator/
│ ├── token.go # Tokenizer logic for recursive parser
│ ├── parser.go # Recursive descent validation logic
│ ├── validatorlinear.go # Linear state machine-based validation (faster)
│ └── validator_test.go # Unit tests and benchmarks
├── go.mod
├── Makefile
└── README.md


---

## ✅ Features

- Supports real numbers with optional decimal parts
- Handles deeply nested parentheses
- Ignores whitespace
- Detects:
  - Unbalanced parentheses
  - Consecutive operators
  - Trailing or leading syntax errors

---

## 🚫 Limitations

- Only supports `+` and `-` (no `*`, `/`, etc.)
- No evaluation or AST construction
- No support for functions or variables
- Unary `-` supported only at valid positions (start or after `(`)

---

## 🛠️ Approaches

### 1. Recursive Descent Parser (`ValidateExpression`)
- Implements a recursive parser based on EBNF grammar
- Tokenizes input, then validates via grammar traversal
- Easier to extend in the future (e.g., add new operators)

### 2. Linear State Machine (`ValidateLinear`)
- Scans the input once using an explicit FSM with states:
  - `stateExpStart`: start of expression or after `(`
  - `stateNumberIntPart`: integer digits
  - `stateNumberDecPart`: after `.`
- Tracks parenthesis balance inline
- Significantly faster and alloc-free for valid input

---

## 🧪 Benchmark Results

| Benchmark                                 | Ops/sec | Time/op    | Allocations |
|------------------------------------------|---------|------------|-------------|
| `ValidateExpression`                     | 104k    | 13.4 µs    | 11 allocs   |
| `ValidateExpression_LargeExpression`     | 15.6k   | 68.0 µs    | 13 allocs   |
| `ValidateExpression_DeepNestedExpression`| 5.3k    | 197.9 µs   | 16 allocs   |
| `ValidateLinear`                         | **229k**| **5.2 µs** | **4 allocs**|
| `ValidateLinear_LargeExpression`         | 43.2k   | 24.1 µs    | 0 allocs    |
| `ValidateLinear_DeepNestedExpression`    | 20.6k   | 59.3 µs    | 0 allocs    |

> Benchmarks were run with `go test -bench . -benchmem`.

---

## ✅ Pros and Cons

| Approach           | Pros                                               | Cons                                   |
|--------------------|----------------------------------------------------|----------------------------------------|
| **Recursive Parser** | Grammar-like clarity, easier extensibility         | Slower, more allocations               |
| **Linear Validator** | Extremely fast and memory-efficient                | Harder to maintain for complex syntax  |

---

## 🧑‍💻 Usage

Run validation using either method:

```go
ok, err := validator.ValidateExpression("(1.1 + (-2.3))")
ok2, err2 := validator.ValidateLinear("1 - (2.5 + 3.1)")
```

## 🔧 Commands
```bash
make build        # Build the binary
make test         # Run unit tests
make benchmark    # Run benchmarks
make lint         # Run linter (requires golangci-lint)
make run          # Run main.go
```

