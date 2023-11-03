# mpstat

The `mpstat` command in Linux is used to display information about the CPUs in your system. It can be used to see how busy the CPUs are, what processes are using the CPUs, and how much time the CPUs are spending in different states.

The syntax of the `mpstat` command is as follows:

```
mpstat [options]
```

The `options` argument specifies additional options for getting information about the CPUs. The most common options are as follows:

* `-P`: Display statistics for all CPUs.
* `-u`: Display user-mode statistics.
* `-i`: Display idle statistics.
* `-d`: Display wait-for-io statistics.

For example, the following command displays the statistics for all CPUs in your system:

```
mpstat
```

The `mpstat` command is a useful tool for monitoring the performance of your CPUs. It can be used to see if there are any CPUs that are overloaded or that are not being used efficiently.

Here are some additional things to keep in mind about the `mpstat` command:

* The `mpstat` command can be used to monitor the performance of both physical and virtual CPUs.
* The `mpstat` command can be used to monitor the performance of CPUs on both local and remote systems.
* The `mpstat` command can be used to monitor the performance of CPUs over a period of time.

It is important to be aware of these limitations when using the `mpstat` command, so that you do not get confused by the output or accidentally monitor CPUs that are not relevant to your system.

The `idle` column in the `mpstat` output shows the percentage of time that the CPU is in the idle state. A high idle percentage indicates that the CPU is not being used very much. A low idle percentage indicates that the CPU is being used heavily.

The `user` column in the `mpstat` output shows the percentage of time that the CPU is in user mode. User mode is the mode that the CPU is in when it is running user processes. A high user percentage indicates that the CPU is being used by user processes. A low user percentage indicates that the CPU is not being used by user processes.

The `system` column in the `mpstat` output shows the percentage of time that the CPU is in system mode. System mode is the mode that the CPU is in when it is running kernel processes. A high system percentage indicates that the CPU is being used by kernel processes. A low system percentage indicates that the CPU is not being used by kernel processes.

The `wait` column in the `mpstat` output shows the percentage of time that the CPU is waiting for I/O. A high wait percentage indicates that the CPU is waiting for data from disk or other devices. A low wait percentage indicates that the CPU is not waiting for I/O.




# help 

```

```
