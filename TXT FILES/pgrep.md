# pgrep

The `pgrep` command in Unix and Linux is used to search for processes based on their names and other attributes and then prints the process IDs (PIDs) of matching processes. It allows for flexible pattern matching and is often used in scripts or command-line operations to identify and work with specific processes.

### Basic Usage

The basic syntax for the `pgrep` command is:

```sh
pgrep [options] pattern
```

- **`pattern`**: The pattern used to match against process names.
- **`options`**: Optional flags to modify the behavior of `pgrep`.

### Examples

#### Find PID of a Process

To find the PID of a process named `firefox`:

```sh
pgrep firefox
```

- This command prints the PID of all processes whose name matches `firefox`.

#### Find PIDs of Processes Matching a Pattern

To find the PIDs of all processes whose command name matches a pattern:

```sh
pgrep '^myapp.*'
```

- This command finds and prints the PIDs of all processes whose command name starts with `myapp`.

#### Verbose Mode

To display verbose output while finding processes:

```sh
pgrep -l chrome
```

- The `-l` option (list) prints the PID and the process name for each matching process.

### Options

#### `-u username`

- Limits the selection to processes belonging to the specified user.

#### `-g gid`

- Limits the selection to processes belonging to the specified group ID.

#### `-x`

- Only match processes whose name exactly matches the pattern.

#### `-n`

- Selects only the newest (most recently started) process matching the criteria.

#### `-a`

- Displays the process name as well as the PID.

### Practical Use Cases

#### Automating Process Management

In scripts or automated tasks, `pgrep` can be used to identify and manipulate processes based on their names or other attributes.

#### Monitoring and Logging

Use `pgrep` to monitor the existence or activity of specific processes, feeding the output into other commands or scripts for further processing.

#### System Administration

`pgrep` is useful for system administrators to quickly identify and manage processes on servers or workstations.

### Caution

- **Be cautious when using `pgrep`**, as it can affect running processes. Ensure you understand the implications of terminating or manipulating processes based on their PIDs.

### Summary

The `pgrep` command is a powerful tool for identifying and managing processes based on their names and other attributes in Unix and Linux systems. It provides flexible pattern matching and options for selecting processes, making it useful for automation, monitoring, and administrative tasks. Understanding its usage and options enables effective process management and system monitoring.

# help 

```
Usage:
 pgrep [options] <pattern>

Options:
 -d, --delimiter <string>  specify output delimiter
 -l, --list-name           list PID and process name
 -a, --list-full           list PID and full command line
 -v, --inverse             negates the matching
 -w, --lightweight         list all TID
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

For more details see pgrep(1).

```
