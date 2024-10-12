# Concurrent, Thread-Safe Queue in Go

This repository contains the implementation of a **concurrent, thread-safe queue** in Go. The queue is implemented using a `sync.Mutex` to ensure thread safety when multiple goroutines access or modify the queue concurrently.

## Key Concepts

- **Concurrency in Go**: Go provides goroutines for lightweight concurrent programming. This queue implementation ensures safe concurrent access using goroutines.
- **Thread Safety**: We use a `sync.Mutex` to protect the critical sections of the code where the queue is modified (enqueue, dequeue, and checking the size).
- **Instance-specific Locking**: Each `ConcQueue` instance has its own mutex, ensuring that only one goroutine can modify the queue at a time, without affecting other instances.
# thread-safe-queue
