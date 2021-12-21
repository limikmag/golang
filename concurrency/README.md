1. Concurrency primitives

What is a process?
 - instance of  running program
 - process provides environment for program to execute

 OS allocates memory for:
 - code (compiled machine instructions)
 - data (global data)
 - heap (dynamic memory allocation)
 - stack (local variables of functions)

 What are threads?
 - smallest unit of execution that CPU accepts
 - process has at least one thread - main
 - process can have multiple threads
 - threads share the same address space


 ```bash
 ulimit -a
core file size          (blocks, -c) 0
data seg size           (kbytes, -d) unlimited
file size               (blocks, -f) unlimited
max locked memory       (kbytes, -l) unlimited
max memory size         (kbytes, -m) unlimited
open files                      (-n) 10496
pipe size            (512 bytes, -p) 1
stack size              (kbytes, -s) 8192 (thread memory size)
cpu time               (seconds, -t) unlimited
max user processes              (-u) 2784
virtual memory          (kbytes, -v) unlimited
```

2. Deep dive into concurrency primitives
3. Concurrency patterns
4. Context package