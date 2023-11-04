# vnstat

The vnstat command in Linux is a network traffic monitor. It can be used to track network traffic usage over time.

For example, the following command will show the network traffic usage for the past 24 hours:

`vnstat -i eth0 -l 24`

# help 

```
vnstat [options]

Network traffic monitor.

Options:

-h, --help           Show this help message.
-i, --interface=IFACE   Monitor the specified network interface.
-l, --line-graph      Show the traffic usage in a line graph.
-t, --time=TIME   Show the traffic usage for the specified time period.
-d, --daily   Show the daily traffic usage.
-w, --weekly   Show the weekly traffic usage.
-m, --monthly   Show the monthly traffic usage.
-y, --yearly   Show the yearly traffic usage.
-r, --reset   Reset the traffic counters.

Examples:

    vnstat -i eth0
    vnstat -i eth0 -l 24
    vnstat -i eth0 -t 1w
```



## breakdown

```
-h, --help: This option shows this help message.
-i, --interface=IFACE: This option tells vnstat to monitor the specified network interface.
-l, --line-graph: This option tells vnstat to show the traffic usage in a line graph.
-t, --time=TIME: This option tells vnstat to show the traffic usage for the specified time period. The time period can be specified in days, weeks, months, or years.
-d, --daily: This option tells vnstat to show the daily traffic usage.
-w, --weekly: This option tells vnstat to show the weekly traffic usage.
-m, --monthly: This option tells vnstat to show the monthly traffic usage.
-y, --yearly: This option tells vnstat to show the yearly traffic usage.
-r, --reset: This option resets the traffic counters.
-f, --full: This option tells vnstat to show all available traffic data.
-s, --summary: This option tells vnstat to show a summary of the traffic usage.
-c, --csv: This option tells vnstat to output the traffic usage in CSV format.
-b, --bytes: This option tells vnstat to display the traffic usage in bytes.
-k, --kilobytes: This option tells vnstat to display the traffic usage in kilobytes.
-m, --megabytes: This option tells vnstat to display the traffic usage in megabytes.
-g, --gigabytes: This option tells vnstat to display the traffic usage in gigabytes.
```
