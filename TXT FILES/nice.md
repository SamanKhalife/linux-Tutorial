# nice 

In Unix and Linux systems, the `nice` command is used to run a program with modified scheduling priority. It allows you to adjust the priority of a process so that it runs with a higher or lower priority level in relation to other processes on the system.

### Basic Usage

The basic syntax for `nice` is:

```sh
nice [options] [command [arguments...]]
```

- **`command`**: The command you want to execute with modified priority.
- **`arguments`**: Optional arguments passed to the command.

### Examples

#### Increase Priority (Lower Nice Value)

To run a command with a higher priority (lower nice value), you typically need root (superuser) privileges:

```sh
sudo nice -n -10 command
```

- **`-n -10`**: Sets the nice value to -10, which increases the priority of `command`.

#### Decrease Priority (Higher Nice Value)

To run a command with a lower priority (higher nice value), you do not need root privileges:

```sh
nice -n 10 command
```

- **`-n 10`**: Sets the nice value to 10, which decreases the priority of `command`.

### Nice Values

- Nice values range from -20 (highest priority) to 19 (lowest priority).
- Higher nice values mean lower priority, and lower nice values mean higher priority.
- Regular users typically can only increase the nice value (lower priority), while root can both increase and decrease it.

### Practical Use Cases

- **Background Processes**: Set lower priority for background tasks to avoid impacting foreground tasks.
- **Resource Management**: Adjust process priority to balance system resources more effectively.
- **Batch Jobs**: Schedule CPU-intensive tasks to run at off-peak times with lower priority.

### Tips

- **`renice` Command**: Use `renice` to change the priority of existing processes.
- **Monitoring**: Use `top` or `ps` commands to monitor the CPU usage and priority of processes.

### Summary

The `nice` command is essential for adjusting the scheduling priority of processes in Unix and Linux systems, allowing users to manage system resources efficiently. It provides flexibility in task management by modifying process priorities based on workload and system requirements. Understanding how to use `nice` effectively enhances performance tuning and system administration tasks on Unix-based platforms.


# help 

```

```
