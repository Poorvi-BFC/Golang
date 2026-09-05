# Go Proverbs

These are the Go proverbs shared by Rob Pike and the Go community. They represent core principles for writing effective Go code.

## The Proverbs

1. **Don't communicate by sharing memory, share memory by communicating.**
   - Use channels to pass data between goroutines instead of shared variables.

2. **Concurrency is not parallelism.**
   - Concurrency is about managing multiple tasks; parallelism is about executing them simultaneously.

3. **Channels orchestrate; mutexes serialize.**
   - Use channels for coordination, mutexes for synchronizing shared data.

4. **The bigger the interface, the weaker the abstraction.**
   - Keep interfaces small and focused for better composability.

5. **Make the zero value useful.**
   - Design types so their default (zero) values are valid and functional.

6. **interface{} says nothing.**
   - Empty interfaces are too generic; prefer explicit type constraints.

7. **Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.**
   - Use consistent formatting over personal preference; follow the standard.

8. **A little copying is better than a little dependency.**
   - Sometimes duplicating small amounts of code is better than adding a dependency.

9. **Syscall must always be guarded with build tags.**
   - Use build tags like `// +build` when working with system calls.

10. **Cgo must always be guarded with build tags.**
    - Use build tags when using C interop to ensure platform compatibility.

11. **Cgo is not Go.**
    - C interoperability adds complexity; use it only when necessary.

12. **With the unsafe package there are no guarantees.**
    - The `unsafe` package bypasses Go's safety features; use with extreme caution.

13. **Clear is better than clever.**
    - Write readable, understandable code over clever, compact code.

14. **Reflection is never clear.**
    - Reflection can be powerful but makes code harder to understand; avoid when possible.

15. **Errors are values.**
    - Treat errors like any other value; handle them explicitly.

16. **Don't just check errors, handle them gracefully.**
    - Checking for errors is not enough; provide meaningful error handling.

17. **Design the architecture, name the components, document the details.**
    - Focus on high-level design, clear naming, and thorough documentation.

18. **Documentation is for users.**
    - Write documentation that helps users understand and use your code.

19. **Don't panic.**
    - Use panic sparingly; it should be for truly exceptional situations.