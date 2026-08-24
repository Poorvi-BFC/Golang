# Go: Learn, Practice, Push

A hands-on Go learning repository built around a simple loop:

> **Learn one concept. Practice it with code. Push what I built.**

The goal is steady improvement. Keep each lesson small, understandable, and easy to revisit.

## Learning Path

### 1. Learn

Study one topic at a time and explain it in your own words.

- Go installation, `go mod`, packages, and the toolchain
- Variables, constants, types, zero values, and type conversions
- Functions, multiple return values, and error handling
- Arrays, slices, maps, and structs
- Pointers and methods
- Interfaces and composition
- Goroutines, channels, and synchronization
- Testing, benchmarks, and table-driven tests
- JSON, files, HTTP clients, and servers
- Documentation, readability, and idiomatic Go

Useful commands to learn early:

```bash
go mod init example.com/learn-go
go run .
go test ./...
go fmt ./...
go vet ./...
```

### 2. Practice with small challenges

Do not move to the next topic until you can use the current one without copying the solution. Suggested exercises:

| Stage | Practice | Main concepts |
| --- | --- | --- |
| Beginner | Temperature and unit converter | Functions, input, types |
| Beginner | CLI calculator | Switches, errors, packages |
| Beginner | Word and character counter | Strings, maps, files |
| Intermediate | To-do list CLI | Structs, JSON, persistence |
| Intermediate | URL checker | HTTP, goroutines, channels |
| Intermediate | Log parser | Files, regular expressions, tests |
| Advanced | REST API for tasks | HTTP handlers, JSON, middleware |
| Advanced | Concurrent worker pool | Context, synchronization, cancellation |

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
