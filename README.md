# Go Learning Repository

A comprehensive, hands-on Go learning project documenting a complete journey from fundamentals to advanced concurrency patterns.

> **Learn → Practice → Commit**

This repository tracks steady progress through Go concepts, with each topic studied, implemented, and committed to version control.

---

## 📚 Learning Progression

### Phase 1: Fundamentals (24-27 Aug)
**Branch:** `praD2`, `praD4`  
**Topics covered:**
- Variables, constants, types, and zero values
- Functions and multiple return values
- Switches and basic control flow
- Loops and iteration patterns
- Maps and basic data structures
- Introduction to pointers

**Key files:**
- `24AUG/` - Hello world, variables, functions, arrays
- `25AUG/` - Functions, loops, maps, structs, switches
- `27AUG/` - Pointers fundamentals

---

### Phase 2: Object-Oriented Concepts (01 Sept)
**Branch:** `struct`  
**Topics covered:**
- Structs and composition
- Methods and receivers
- Interfaces and polymorphism
- Error handling and custom errors
- Interface composition

**Key files:**
- `01SEP/struct/` - Struct definition and usage
- `01SEP/interface/` - Single and multiple interfaces
- `01SEP/Errors/` - Error types and handling

---

### Phase 3: Collections Deep Dive (02 Sept)
**Branch:** `2sep`  
**Topics covered:**
- Advanced loops and iteration
- Slices - creation, modification, and operations
- Array manipulation
- Modulo operations
- Loop assignments and patterns

**Key files:**
- `02sep/loops/` - Loop patterns and best practices
- `02sep/slices/` - Slice fundamentals and assignments
- `02sep/loop-test/` - Loop testing exercises

---

### Phase 4: Advanced Data Structures & Pointers (03 Sept)
**Branch:** `3sep`  
**Topics covered:**
- Maps - creation, iteration, manipulation
- Pointer arithmetic and dereferencing
- Advanced function patterns
- Type assertions and switches
- Pointer to structs

**Key files:**
- `03sep/maps/` - Map operations and assignments
- `03sep/ptr/` - Pointer concepts and exercises
- `03sep/advF/` - Advanced function patterns

---

### Phase 5: Concurrency & Generics (05 Sept)
**Branch:** `5sep`  
**Topics covered:**
- Generics - type parameters and constraints
- Generic interfaces with parametric constraints
- Concurrency primitives
- Mutexes and thread-safe code
- Synchronization patterns

**Key files:**
- `05sep/Generics/` - Generic types, functions, and interfaces
- `05sep/Mutexes/` - Thread-safe counters and synchronization

---

## 🗂️ Project Structure

```
Golang/
├── 24AUG/              # Week 1 - Fundamentals
├── 25AUG/              # Week 1 continued
├── 27AUG/              # Week 1.5 - Pointers intro
├── 01SEP/              # Week 2 - Structs, Interfaces, Errors
├── 02SEP/              # Week 2.5 - Loops, Slices, Arrays
├── 03SEP/              # Week 3 - Maps, Pointers (adv), Functions
├── 05SEP/              # Week 4 - Generics, Mutexes, Concurrency
├── GoProverbs.md       # Key Go principles and idioms
└── README.md           # This file
```

---

## 🌿 Git Branches

| Branch           | Focus                           | Status     |
| ---------------- | ------------------------------- | ---------- |
| `main`           | Stable, completed learning path | ✅ Active   |
| `5sep`           | Generics and concurrency        | ✅ Complete |
| `3sep`           | Maps and advanced pointers      | ✅ Complete |
| `2sep`           | Slices and loops                | ✅ Complete |
| `struct`         | Interfaces and error handling   | ✅ Complete |
| `praD2`, `praD4` | Practice exercises              | ✅ Complete |

---

## 📋 Key Concepts Mastered

### Data Types & Structures
- ✅ Variables, constants, type conversion
- ✅ Arrays, slices, maps
- ✅ Structs and composition
- ✅ Pointers and dereferencing

### Functions & Methods
- ✅ Multiple return values
- ✅ Error handling patterns
- ✅ Receivers and methods
- ✅ Variadic functions
- ✅ Advanced function patterns

### Object-Oriented Programming
- ✅ Interfaces and polymorphism
- ✅ Interface composition
- ✅ Type assertions and switches

### Advanced Topics
- ✅ Generics with type constraints
- ✅ Generic interfaces
- ✅ Concurrency primitives
- ✅ Mutex-based synchronization
- ✅ Thread-safe code patterns

---

## 🚀 Quick Start

Navigate to any topic folder and run:

```bash
# Run a specific example
go run <topic>/main.go

# Run all tests in a folder
go test ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...
```

---

## 📖 Resources

- [Go Proverbs](./GoProverbs.md) - Essential Go design principles
- Official Go docs: https://golang.org/doc/
- Effective Go: https://golang.org/doc/effective_go

---

## 📝 Notes

Each topic includes:
- **Concept files** (`main.go`) - Implementation of the concept
- **Assignment files** (`.md` files) - Problem statements and explanations
- **Practice exercises** - Apply concepts to real problems

---

**Status:** ✅ Learning path completed across 5 phases covering fundamentals through advanced concurrency patterns.

For every exercise, try to include:

- A short problem statement
- A clean `main` or package entry point
- Input validation and useful errors
- At least three tests, including one edge case
- A short README describing how to run it

### 3. Push your progress

Use Git as a learning journal. A small, working commit is better than a large unfinished upload.

```bash
git status
git add .
git commit -m "feat: add URL checker exercise"
git push origin main
```

Recommended commit prefixes:

- `learn:` for notes or examples
- `practice:` for exercises
- `feat:` for project functionality
- `test:` for tests
- `docs:` for documentation
- `fix:` for corrections

## Weekly Routine

Repeat this cycle for each topic:

1. **Learn**: Read the official documentation and write a tiny example.
2. **Practice**: Solve one exercise without following a tutorial line by line.
3. **Review**: Run formatting, tests, and `go vet`.
4. **Push**: Commit with a focused message and write down what you learned.
5. **Reflect**: Record one question or improvement for the next session.

## Definition of Done

Before pushing a lesson or exercise, check:

- [ ] The code runs with `go run .` or has a documented command.
- [ ] `go fmt ./...` has been run.
- [ ] `go test ./...` passes.
- [ ] Errors are handled instead of silently ignored.
- [ ] The change has a focused commit message.
- [ ] The README or notes explain the key idea in your own words.



Keep each lesson small and independent. Prefer several focused folders over one increasingly large program.

## Best Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go documentation](https://go.dev/doc/)
- [Go Blog](https://go.dev/blog/)



**Keep learning. Keep practicing. Keep pushing.**
