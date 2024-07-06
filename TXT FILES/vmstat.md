# vmstat

The `vmstat` command in Linux provides a comprehensive snapshot of system-wide performance metrics, focusing on virtual memory statistics. It reports information about processes, memory, paging, block IO, traps, and CPU activity. Here’s a detailed explanation of how to use `vmstat` and what information it provides:

### Usage of `vmstat`

#### Basic Usage

To use `vmstat`, open a terminal and type:

```bash
vmstat
```

By default, `vmstat` displays a one-line summary of system statistics since the last reboot. It provides information on processes, memory, paging, block IO, traps, and CPU activity.

#### Options and Output

`vmstat` typically provides output in columns representing different metrics:

1. **Procs**
   - `r`: The number of processes waiting for CPU time (runnable).
   - `b`: The number of processes in uninterruptible sleep (blocked).

   Example output:
   ```
   procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
   0  0      0 244712  67768 510780    0    0     0     0    4   26  0  0 99  0  0
   ```

2. **Memory**
   - `swpd`: Amount of virtual memory used (in kilobytes).
   - `free`: Amount of idle memory (in kilobytes).
   - `buff`: Amount of memory used as buffers (in kilobytes).
   - `cache`: Amount of memory used as cache (in kilobytes).

3. **Swap**
   - `si`: Amount of memory swapped in from disk (in kilobytes per second).
   - `so`: Amount of memory swapped out to disk (in kilobytes per second).

4. **IO**
   - `bi`: Blocks received from a block device (blocks per second).
   - `bo`: Blocks sent to a block device (blocks per second).

5. **System**
   - `in`: Number of interrupts per second, including the clock.
   - `cs`: Number of context switches per second.

6. **CPU**
   - `us`: Percentage of CPU time spent running user-level processes.
   - `sy`: Percentage of CPU time spent running kernel-level processes.
   - `id`: Percentage of CPU time spent idle.
   - `wa`: Percentage of CPU time spent waiting for IO.
   - `st`: Percentage of CPU time stolen from a virtual machine.

#### Additional Options

- **Interval:** You can specify the interval in seconds for which `vmstat` displays statistics. For example, to refresh every 2 seconds:

  ```bash
  vmstat 2
  ```

- **Count:** Specify a number of samples to display before exiting. For example, to display 5 samples at 1-second intervals:

  ```bash
  vmstat 1 5
  ```

- **Detailed Memory Statistics:** Use `-s` option to display detailed memory statistics (sizes, counts, and limits).

  ```bash
  vmstat -s
  ```

### Use Cases

- **Performance Monitoring:** `vmstat` is useful for monitoring system performance, identifying memory and CPU bottlenecks, and understanding system resource usage patterns over time.

- **Troubleshooting:** Helps in diagnosing system performance issues such as excessive CPU usage, high memory usage, or excessive IO operations.

- **Capacity Planning:** Provides insights into system resource utilization, aiding decisions on hardware upgrades or optimizations.

### Conclusion

`vmstat` is a powerful command-line tool for monitoring system-wide performance metrics related to memory, CPU, and IO activities on Linux systems. By understanding its output and options, administrators and users can effectively monitor and manage system resources, optimize performance, and troubleshoot issues.

# help 

```
vmstat [options]

Display virtual memory statistics.

Options:

-c, --clear   Clear the statistics after each report.
-d, --delay=DELAY   Delay between reports in seconds.
-f, --full   Display all statistics.
-m, --megabytes   Display memory statistics in megabytes.
-s, --seconds=SECONDS   Number of reports to generate.

Examples:

    vmstat
    vmstat 5 10
    vmstat -c
    vmstat -m
```



## breakdown

```
-c, --clear: This option clears the statistics after each report.
-d, --delay=DELAY: This option sets the delay between reports in seconds.
-f, --full: This option displays all statistics.
-m, --megabytes: This option displays memory statistics in megabytes.
-s, --seconds=SECONDS: This option sets the number of reports to generate.
```
