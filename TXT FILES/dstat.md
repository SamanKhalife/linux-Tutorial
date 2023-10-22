# dstat


dstat is a system monitoring tool that can be used to collect and display real-time statistics about the performance of a Linux system. It can be used to monitor CPU usage, memory usage, disk I/O, network traffic, and other metrics.

dstat is a powerful tool that can be used to troubleshoot performance problems and to identify bottlenecks. It is also a valuable tool for system administrators who need to monitor the performance of their systems.

Here are some of the most commonly used dstat statistics:

cpu: This statistic shows the CPU usage of each CPU core.

mem: This statistic shows the memory usage of the system.

disk: This statistic shows the disk I/O activity of the system.

net: This statistic shows the network traffic activity of the system.


# help 

```
dstat [options]

Display system statistics.

Options:

-c, --count   Number of samples to collect (default 1).
-d, --delay   Delay between samples (default 1 second).
-h, --help     Display this help message.
-i, --interval   Delay between updates to the screen (default 1 second).
-l, --list     Show a list of available statistics.
-m, --mode     Output mode (default 'default').
-n, --noupdate   Do not update the screen.
-t, --top       Show only the top-N statistics.
-V, --version   Display the version number.

Statistics:

    cpu   CPU usage.
    mem   Memory usage.
    disk   Disk I/O.
    net   Network traffic.
    irq   Interrupts.
    ctxt   Context switches.
    swap   Swap usage.
    iowait   Time spent waiting for I/O.
    sys   System time.
    user   User time.
    nice   Nice time.
    idle   Idle time.

Examples:

    dstat
    dstat -c 10
    dstat -l
    dstat -m 'cpu,mem,disk,net'
```
