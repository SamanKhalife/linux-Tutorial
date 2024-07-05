# free
The `free` command in Unix and Linux systems is used to display information about the system's memory usage, including total and used memory, free memory, swap usage, and buffer/cache usage.

### Basic Usage

Simply type `free` in your terminal:

```sh
free
```

### Output Format

The `free` command typically outputs information in the following format:

```
              total        used        free      shared  buff/cache   available
Mem:        1635412      467908      587084       44644      580420      922572
Swap:       2097148      192224     1904924
```

- **`total`**: Total amount of physical memory (RAM) in kilobytes.
- **`used`**: Used memory, in kilobytes.
- **`free`**: Free memory, in kilobytes.
- **`shared`**: Memory used (shared by processes), in kilobytes.
- **`buff/cache`**: Memory used as buffers/cache, in kilobytes.
- **`available`**: Estimate of how much memory is available for starting new applications, in kilobytes.

For swap:

- **`Swap`**: Total amount of swap memory (if available).
- **`used`**: Used swap memory.
- **`free`**: Free swap memory.

### Practical Use Cases

- **Monitoring Memory Usage**: Quickly check how much memory is being used and how much is free.
- **Identifying Memory Pressure**: Determine if the system is running low on available memory.
- **Checking Swap Usage**: Assess whether swap space is being utilized and if there might be performance implications.

### Summary

The `free` command provides essential information about system memory usage, including both physical memory (RAM) and swap space. It's a useful tool for monitoring system resources, diagnosing performance issues related to memory, and understanding the overall health of the system on Unix and Linux platforms.

# help

```
Usage:
 free [options]

Options:
 -b, --bytes         show output in bytes
     --kilo          show output in kilobytes
     --mega          show output in megabytes
     --giga          show output in gigabytes
     --tera          show output in terabytes
     --peta          show output in petabytes
 -k, --kibi          show output in kibibytes
 -m, --mebi          show output in mebibytes
 -g, --gibi          show output in gibibytes
     --tebi          show output in tebibytes
     --pebi          show output in pebibytes
 -h, --human         show human-readable output
     --si            use powers of 1000 not 1024
 -l, --lohi          show detailed low and high memory statistics
 -t, --total         show total for RAM + swap
 -s N, --seconds N   repeat printing every N seconds
 -c N, --count N     repeat printing N times, then exit
 -w, --wide          wide output

     --help     display this help and exit
 -V, --version  output version information and exit

For more details see free(1).
```
