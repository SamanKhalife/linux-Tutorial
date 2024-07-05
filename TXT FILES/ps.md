# ps 

The `ps` command in Unix and Linux is used to list information about currently running processes on a system. It provides a snapshot of processes running at the time the command is executed, offering various options to customize the output.

### Basic Usage

Simply type `ps` in your terminal:

```sh
ps
```

### Output Format

The `ps` command typically displays a minimal set of information about processes, including:

- **PID**: Process ID.
- **TTY**: Terminal associated with the process.
- **TIME**: CPU time consumed by the process.
- **CMD**: The command that started the process.

### Examples

#### List All Processes

To list all processes running under the current user's session:

```sh
ps aux
```

- **`a`**: Shows processes from all users.
- **`u`**: Displays detailed information, including user-oriented format.

#### Display Process Tree

To display a process tree, showing hierarchical relationships:

```sh
ps axjf
```

- **`j`**: Shows a process tree (parent/child relationship).
- **`f`**: Displays full listing.

### Options

The `ps` command offers various options to customize output:

- **`-e`**: Lists all processes, regardless of terminal.
- **`-f`**: Displays full-format listing.
- **`-l`**: Long format (extra details).
- **`-u username`**: Lists processes for a specific user.
- **`-p pid`**: Shows processes with specific PIDs.

### Practical Use Cases

- **Monitoring Processes**: View running processes to understand system activity.
- **Identifying Resource Usage**: Determine CPU and memory usage per process.
- **Troubleshooting**: Diagnose issues by identifying misbehaving or stuck processes.

### Summary

The `ps` command is a fundamental tool for viewing information about processes on Unix and Linux systems. It provides essential insights into system activity, process management, and resource utilization. Understanding its usage and options helps in effective system administration, performance monitoring, and troubleshooting.
# help 

```
Usage:
 ps [options]

 Try 'ps --help <simple|list|output|threads|misc|all>'
  or 'ps --help <s|l|o|t|m|a>'
 for additional help text.

For more details see ps(1).
```

