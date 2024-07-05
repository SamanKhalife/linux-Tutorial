# pkill

The `pkill` command in Unix and Linux is used to send signals to processes based on their name and other attributes. It provides a way to terminate processes by matching their command name or other attributes, similar to `killall` but with more flexible pattern matching capabilities.

### Basic Usage

The basic syntax for the `pkill` command is:

```sh
pkill [options] pattern
```

- **`pattern`**: The pattern used to match against process names.
- **`options`**: Optional flags to modify the behavior of `pkill`.

### Examples

#### Kill a Single Process

To kill all instances of a process named `firefox`:

```sh
pkill firefox
```

- This command sends the default signal (`SIGTERM`) to all processes whose name matches `firefox`, asking them to terminate gracefully.

#### Forcefully Terminate Processes

To forcefully terminate all instances of a process named `java`:

```sh
pkill -9 java
```

- The `-9` option sends the `SIGKILL` signal, which forcefully terminates the `java` processes without allowing them to clean up resources.

#### Kill Processes Matching a Pattern

To kill all processes whose name matches a pattern using regular expressions:

```sh
pkill '^myapp.*'
```

- This command terminates all processes whose command name starts with `myapp`.

#### Verbose Mode

To display verbose output while killing processes:

```sh
pkill -v chrome
```

- The `-v` option (verbose mode) prints a message for each process killed.

### Options

#### `-u username`

- Limits the selection to processes belonging to the specified user.

#### `-g gid`

- Limits the selection to processes belonging to the specified group ID.

#### `-t terminal`

- Limits the selection to processes associated with the specified terminal.

#### `-x`

- Only match processes whose name exactly matches the pattern.

#### `-n`

- Selects only the newest (most recently started) process matching the criteria.

### Practical Use Cases

#### Terminating Problematic Processes

Use `pkill` to terminate misbehaving or stuck processes based on their name or other attributes.

#### Batch Processing

In scripts or automated tasks, `pkill` can be used to ensure specific processes are terminated before starting a new task or performing maintenance.

#### Managing Multiple Instances

Efficiently manage multiple instances of a process by using `pkill` with pattern matching to target specific instances.

### Caution

- **Be cautious when using `pkill`**, especially with the `-9` option (`SIGKILL`), as it forcefully terminates processes without allowing them to clean up resources. This can lead to data loss or corruption in some cases.
- **Ensure you specify the correct pattern** to avoid unintended termination of critical processes.

### Summary

The `pkill` command is a flexible tool for terminating processes based on their name or other attributes in Unix and Linux systems. It offers powerful pattern matching capabilities and various options for customizing process selection and signal handling. Understanding its usage and options ensures effective process management and system maintenance.
# help

```
Usage:
 pkill [options] <pattern>

Options:
 -<sig>, --signal <sig>    signal to send (either number or name)
 -q, --queue <value>       integer value to be sent with the signal
 -e, --echo                display what is killed
 -c, --count               count of matching processes
 -f, --full                use full process name to match
 -g, --pgroup <PGID,...>   match listed process group IDs
 -G, --group <GID,...>     match real group IDs
 -i, --ignore-case         match case insensitively
 -n, --newest              select most recently started
 -o, --oldest              select least recently started
 -O, --older <seconds>     select where older than seconds
 -P, --parent <PPID,...>   match only child processes of the given parent
 -s, --session <SID,...>   match session IDs
 -t, --terminal <tty,...>  match by controlling terminal
 -u, --euid <ID,...>       match by effective IDs
 -U, --uid <ID,...>        match by real IDs
 -x, --exact               match exactly with the command name
 -F, --pidfile <file>      read PIDs from file
 -L, --logpidfile          fail if PID file is not locked
 -r, --runstates <state>   match runstates [D,S,Z,...]
 --ns <PID>                match the processes that belong to the same
                           namespace as <pid>
 --nslist <ns,...>         list which namespaces will be considered for
                           the --ns option.
                           Available namespaces: ipc, mnt, net, pid, user, uts

 -h, --help     display this help and exit
 -V, --version  output version information and exit

```
