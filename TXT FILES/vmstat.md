# vmstat

The vmstat command in Linux is a system monitoring tool that can be used to display information about the virtual memory system.

For example, the following command will display the virtual memory statistics every 5 seconds for 10 iterations:

`vmstat 5 10 `

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
