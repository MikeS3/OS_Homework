#!/bin/bash
#ChatGPT helped with the implementation of variables and time formatting in bash

echo "Starting Implementation with Go routines..."
START1=$(date +%s.%N)
go run HW0_1.go #File name of implementation with concurrency
END1=$(date +%s.%N)
TIME1=$(echo "$END1 - $START1" | bc)

echo "Starting Python Threads Implementation ..."
START2=$(date +%s.%N)
python3 HW1_2.py #File name of implementation with concurrency
END2=$(date +%s.%N)
TIME2=$(echo "$END2 - $START2" | bc)

echo "Goroutine Implementation: $TIME1 seconds"
echo "Python Thread Implementation: $TIME2 seconds"
