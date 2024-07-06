# iostat
The `iostat` command in Linux is used to monitor system input/output (I/O) statistics for CPU, disks, and partitions. It provides detailed information about disk utilization, I/O rates (reads and writes), and CPU utilization. Here’s a detailed explanation of how to use `iostat` and what information it provides:

### Usage of `iostat`

#### Basic Usage

To use `iostat`, open a terminal and type:

```bash
iostat
```

By default, `iostat` displays CPU and disk I/O statistics since the last reboot.

#### Options and Output

`iostat` provides output with different columns representing various system and device attributes:

1. **Columns**
   - `%user`: Percentage of CPU utilization that occurred while executing at the user level (application).
   - `%nice`: Percentage of CPU utilization that occurred while executing at the user level with nice priority.
   - `%system`: Percentage of CPU utilization that occurred while executing at the system level (kernel).
   - `%iowait`: Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.
   - `%idle`: Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.

   For disk devices:
   - `Device`: Device name.
   - `tps`: Number of transfers per second that were issued to the device.
   - `kB_read/s`: Kilobytes read from the device per second.
   - `kB_wrtn/s`: Kilobytes written to the device per second.
   - `kB_read`: Total kilobytes read from the device.
   - `kB_wrtn`: Total kilobytes written to the device.

   Example output:
   ```
   Linux 5.4.0-97-generic (hostname)  07/06/24  _x86_64_    (1 CPU)

   avg-cpu:  %user   %nice %system %iowait  %steal   %idle
             5.04    0.00    1.63    0.36    0.00   92.97

   Device             tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
   sda               0.36         1.44         0.00      13152          0
   ```

2. **Options**
   - `-c <count>`: Display statistics <count> times, then exit.
   - `-d`: Display device utilization report.
   - `-p <device>`: Display statistics for specific devices (e.g., `-p sda`).

   Example:
   ```bash
   iostat -c 5    # Display statistics 5 times and exit
   iostat -d      # Display device utilization report
   iostat -p sda  # Display statistics for device 'sda'
   ```

### Use Cases

- **Performance Monitoring:** `iostat` helps in monitoring system performance, especially CPU and disk I/O activities.
  
- **Resource Utilization:** Useful for identifying disk bottlenecks and CPU utilization issues.
  
- **Capacity Planning:** Assists in capacity planning by providing insights into disk read/write rates and overall system utilization.

### Conclusion

`iostat` is a powerful command-line tool for monitoring CPU and disk I/O statistics on Linux systems. It provides essential information about system performance, disk utilization, and CPU efficiency. By understanding its output and options, administrators and users can effectively monitor system activities, diagnose performance issues, and optimize system performance.

# help 

```
iostat [options] [interval] [count]

Display I/O statistics.

Options:

-k, --kilobytes   Display I/O statistics in kilobytes.
-m, --megabytes   Display I/O statistics in megabytes.
-g, --gigabytes   Display I/O statistics in gigabytes.
-d, --disk=[device]   Only display statistics for the specified device.
-c, --cpu   Display CPU utilization statistics.
-h, --help           Show this help message.
```

##  breakdown

```
-k, --kilobytes: This option displays I/O statistics in kilobytes.
-m, --megabytes: This option displays I/O statistics in megabytes.
-g, --gigabytes: This option displays I/O statistics in gigabytes.
-d, --disk=[device]: This option only displays statistics for the specified device.
-c, --cpu: This option displays CPU utilization statistics.
-h, --help: This option shows this help message.
```
