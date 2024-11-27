Here is the `README.md` content you requested:

```markdown
# **Go Concurrent API Fetcher**

This project demonstrates how to use Go's concurrency features, such as **GoRoutines**, **Channels**, and **Mutexes**, to perform concurrent API calls efficiently.

## **Table of Contents**
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [How to Run](#how-to-run)
- [Example](#example)
- [Concurrency Concepts](#concurrency-concepts)

## **Introduction**

In this project, we explore how Go handles concurrency with its powerful tools:
- **GoRoutines** for running tasks concurrently
- **Channels** for safe communication between GoRoutines
- **Mutexes** to manage concurrent access to shared resources.

The goal of this project is to fetch data from multiple APIs concurrently using Go's concurrency features.

## **Prerequisites**

To run this project, you need to have the following installed on your machine:

- [Go](https://golang.org/dl/) version 1.16 or later
- A terminal (Linux, macOS, or Windows with WSL)

### **Main Files:**
- **main.go**: The entry point of the program where the GoRoutines are managed, and the WaitGroup synchronizes their completion.
- **fetcher.go**: Contains the logic for fetching data from an API and using channels to pass data.

## **How to Run**

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/Aekshant/go-concurrent-api-fetcher.git
   ```

2. Change to the project directory:
   ```bash
   cd go-concurrent-api-fetcher
   ```

3. Run the project:
   ```bash
   go run main.go
   ```

   The program will fetch data concurrently from a list of APIs and display the results.

## **Example**

Here’s an example of how the program works:

```bash
$ go run main.go
Fetching data concurrently from APIs...
API 1: Fetched data successfully
API 2: Fetched data successfully
API 3: Fetched data successfully
API 4: Fetched data successfully
All workers are done!
```

In this example, the program makes concurrent HTTP requests to different API endpoints and processes the responses asynchronously.

## **Concurrency Concepts**

### **GoRoutines**
GoRoutines allow concurrent execution of functions in Go. Each GoRoutine runs independently, so we can perform tasks like API calls concurrently, speeding up the execution time.

### **Channels**
Channels are used to communicate between GoRoutines. In this project, we use a channel to pass data back from the GoRoutines to the main function after each API call finishes.

### **Mutexes**
Mutexes are used to ensure safe access to shared resources. In this project, a mutex is used to protect any shared data, ensuring no race conditions occur when multiple GoRoutines are accessing the same resource.

Happy_Coding:)
