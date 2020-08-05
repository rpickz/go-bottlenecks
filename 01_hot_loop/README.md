# Problem: Hot Loop/Hot Spot

## Summary

Hot spots/hot loops are a common performance problem in certain types of programs.

A hot spot or hot loop is an area of a program in which the CPU spends the majority of its time.  

It can be caused by a number of things - including long-running loops, poor use of buffering causing frequent I/O bottlenecks, and more.

This particular example demonstrates two issues - and these are resolved one at a time, leading to a performance boost of more than 10x.

Benchmarks are used within these examples to demonstrate the performance gains acquired, and profiling tools are used to diagnose the symptoms of this issue.

Resources:
[Wikipedia entry](https://en.wikipedia.org/wiki/Hot_spot_(computer_programming))

## Instructions

Start by analsing the first directory `01_problem` - this demonstrates the unoptimised program.

Use the following to build the program:

```bash
cd 01_problem
go build -o problem .
```

Once the program has been built, use the following to run the program (with time information):

```shell script
time ./problem
```

You will be provided some time statistics, including a "Real" value - which is the amount of time taken to run the program according to the "wall clock" (or if you were to time this using a stop-watch).
