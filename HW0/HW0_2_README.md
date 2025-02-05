# HW0_2 Simple Data Structure

A simple stack implementation in Go that demonstrates basic stack operations—**Push**, **Pop**, and **Top (Peek)**—using a fixed-size array.

## Overview

This project implements a basic stack data structure in Go with the following functionalities:

- **Push:** Adds an element to the top of the stack. Returns `false` if the stack is full.
- **Pop:** Removes and returns the top element of the stack. Returns a boolean to indicate success or failure (e.g., when the stack is empty).
- **Top (Peek):** Returns the top element without removing it from the stack.

The stack is implemented using a fixed-size array of 100 integers. The `top` variable tracks the current position in the stack.

## How It Works

- **Push:** The function checks if the stack has reached its capacity. If not, it places the new element at the current `top` index and increments `top`.
- **Pop:** The function checks if the stack is empty. If not, it decrements `top` and returns the element that was at the top.
- **Top:** The function returns the element at `top - 1` without modifying the stack.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) install Go 

### Cloning the Repository

Clone the repository using the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW0
```
### Running the Code

To run the code enter this command

```
go run HW0_2.go
```

### Results

- Program shows multiple push and pop commands successfully being called and displaying the accurate top variable
- Program creates a successful usable stack in GO

