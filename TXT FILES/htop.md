# htop
`htop` is an interactive process viewer for Unix systems, similar to the `top` command but with an improved user interface and more features. It allows users to monitor system processes and their resource usage in real-time. Here's a detailed explanation of `htop`, including how to use it and what features it provides:

### Usage of `htop`

#### Basic Usage

To use `htop`, open a terminal and simply type:

```bash
htop
```

By default, `htop` displays an interactive process list with various columns showing detailed information about each process, such as CPU and memory usage.

#### Key Features and Interface

`htop` provides an enhanced interface compared to `top`, with the following notable features:

1. **Interactive Commands**
   - **Navigation**: Use arrow keys or mouse to navigate through the process list.
   - **Process Management**: Press `F9` to kill a process, `F7` to renice a process, `F5` to sort processes, etc.
   - **Filtering**: Press `F4` to search/filter processes by name.
   - **Detailed Process Information**: Displays detailed information about each process, including command line arguments and environment variables.
   - **Color-coded Display**: Different colors are used to highlight processes based on their resource usage.

2. **Columns**
   - `PID`: Process ID
   - `USER`: User running the process
   - `%CPU`: Percentage of CPU time used by the process
   - `%MEM`: Percentage of memory used by the process
   - `COMMAND`: Command name or command line used to start the process
   - `TIME+`: Total CPU time used by the process (in seconds)
   - `Threads`: Number of threads used by the process

   Example:
   ```
   PID   USER     PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
   1234  user     20   0  123456  7890   4567  R   10.0   2.0   0:05.12 htop
   ```

3. **Options**
   - `-u <username>`: Display processes for a specific user.
   - `-p <PID>`: Monitor specific process IDs.
   - `-d <delay>`: Set the delay between updates (in seconds).
   - `-C`: Sort by the command name rather than by CPU usage.

   Example:
   ```bash
   htop -u username    # Show processes for user 'username'
   htop -p 1234,5678   # Monitor processes with PID 1234 and 5678
   ```

### Use Cases

- **Real-time Monitoring**: `htop` is ideal for real-time monitoring of system processes and resource usage.
  
- **Process Management**: Helps in managing processes interactively, such as killing or renicing processes based on their resource consumption.
  
- **Performance Analysis**: Useful for troubleshooting performance issues and identifying processes that are consuming excessive CPU or memory.

### Conclusion

`htop` is a feature-rich and user-friendly alternative to `top`, offering enhanced visualization and interactive capabilities for monitoring system processes on Unix-like systems. It provides detailed insights into CPU, memory, and process activity, making it a valuable tool for system administrators, developers, and users concerned with system performance and resource management. By leveraging its interactive commands and customizable options, users can effectively monitor, manage, and optimize system resources.
help
```
Usage: head [OPTION]... [FILE]...
Print the first 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[-]NUM       print the first NUM bytes of each file;
                             with the leading '-', print all but the last
                             NUM bytes of each file
  -n, --lines=[-]NUM       print the first NUM lines instead of the first 10;
                             with the leading '-', print all but the last
                             NUM lines of each file
  -q, --quiet, --silent    never print headers giving file names
  -v, --verbose            always print headers giving file names
  -z, --zero-terminated    line delimiter is NUL, not newline
      --help     display this help and exit
      --version  output version information and exit

NUM may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.
Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Report any translation bugs to <https://translationproject.org/team/>
Full documentation <https://www.gnu.org/software/coreutils/head>
or available locally via: info '(coreutils) head invocation'
root@SAMANLEARN:~# ^C
root@SAMANLEARN:~# man head
root@SAMANLEARN:~# htop --help
htop 3.2.1
(C) 2004-2019 Hisham Muhammad. (C) 2020-2022 htop dev team.
Released under the GNU GPLv2+.

-C --no-color                   Use a monochrome color scheme
-d --delay=DELAY                Set the delay between updates, in tenths of seconds
-F --filter=FILTER              Show only the commands matching the given filter
-h --help                       Print this help screen
-H --highlight-changes[=DELAY]  Highlight new and old processes
-M --no-mouse                   Disable the mouse
-p --pid=PID[,PID,PID...]       Show only the given PIDs
   --readonly                   Disable all system and process changing features
-s --sort-key=COLUMN            Sort by COLUMN in list view (try --sort-key=help for a list)
-t --tree                       Show the tree view (can be combined with -s)
-u --user[=USERNAME]            Show only processes for a given user (or $USER)
-U --no-unicode                 Do not use unicode but plain ASCII
-V --version                    Print version info

Long options may be passed with a single dash.
```


man
```
NAME
       htop, pcp-htop - interactive process viewer

SYNOPSIS
       htop [-dCFhpustvH]
       pcp htop [-dCFhpustvH] [--host/-h host]

DESCRIPTION
       htop is a cross-platform ncurses-based process viewer.

       It  is  similar to top, but allows you to scroll vertically and horizontally, and interact
       using a pointing device (mouse).  You can observe all processes  running  on  the  system,
       along  with  their  command  line arguments, as well as view them in a tree format, select
       multiple processes and act on them all at once.

       Tasks related to processes (killing, renicing) can be done without entering their PIDs.

       pcp-htop is a version of htop built using the Performance Co-Pilot (PCP) Metrics API  (see
       PCPIntro(1),  PMAPI(3)), allowing to extend htop to display values from arbitrary metrics.
       See the section below titled CONFIG FILES for further details.

COMMAND-LINE OPTIONS
       Mandatory arguments to long options are mandatory for short options too.

       -d --delay=DELAY
              Delay between updates, in tenths of a second. If the delay value is less than 1, it
              is  increased to 1, i.e. 1/10 second. If the delay value is greater than 100, it is
              decreased to 100, i.e. 10 seconds.

       -C --no-color --no-colour
              Start htop in monochrome mode

       -F --filter=FILTER
              Filter processes by terms matching the commands. The terms are matched  case-insen‐
              sitive and as fixed strings (not regexs). You can separate multiple terms with "|".

       -h --help
              Display a help message and exit

       -p --pid=PID,PID...
              Show only the given PIDs

       -s --sort-key COLUMN
              Sort  by  this  column  (use --sort-key help for a column list).  This will force a
              list view unless you specify -t at the same time.

       -u --user=USERNAME|UID
              Show only the processes of a given user

       -U --no-unicode
              Do not use unicode but ASCII characters for graph meters

       -M --no-mouse
              Disable support of mouse control

       --readonly
              Disable all system and process changing features
       -V --version
              Output version information and exit

       -t --tree
              Show processes in tree view. This can be used to force a tree view when  requesting
              a sort order with -s.

       -H --highlight-changes=DELAY
              Highlight new and old processes

          --drop-capabilities[=off|basic|strict]
              Linux only; requires libcap support.
              Drop  unneeded  Linux capabilities.  In strict mode features like killing, changing
              process priorities, and reading process delay accounting information will not work,
              due to less capabilities held.

INTERACTIVE COMMANDS
       The following commands are supported while in htop:

       Tab, Shift-Tab
            Select  the  next  /  the previous screen tab to display.  You can enable showing the
            screen tab names in the Setup screen (F2).

       Up, Alt-k
            Select (highlight) the previous process in the process list. Scroll the list if  nec‐
            essary.

       Down, Alt-j
            Select  (highlight)  the  next process in the process list. Scroll the list if neces‐
            sary.

       Left, Alt-h
            Scroll the process list left.

       Right, Alt-l
            Scroll the process list right.

       PgUp, PgDn
            Scroll the process list up or down one window.

       Home Scroll to the top of the process list and select the first process.

       End  Scroll to the bottom of the process list and select the last process.

       Ctrl-A, ^
            Scroll left to the beginning of the process entry (i.e. beginning of line).

       Ctrl-E, $
            Scroll right to the end of the process entry (i.e. end of line).

       Space
            Tag or untag a process. Commands that can operate on multiple processes, like "kill",
            will  then  apply  over  the list of tagged processes, instead of the currently high‐
            lighted one.

       c    Tag the current process and its children. Commands that can operate on multiple  pro‐
            cesses,  like  "kill",  will then apply over the list of tagged processes, instead of
            the currently highlighted one.

       U    Untag all processes (remove all tags added with the Space or c keys).

       s    Trace process system calls: if strace(1) is installed, pressing this key will  attach
            it to the currently selected process, presenting a live update of system calls issued
            by the process.

       l    Display open files for a process: if lsof(1) is installed,  pressing  this  key  will
            display the list of file descriptors opened by the process.

       w    Display  the  command line of the selected process in a separate screen, wrapped onto
            multiple lines as needed.

       x    Display the active file locks of the selected process in a separate screen.

       F1, h, ?
            Go to the help screen

       F2, S
            Go to the setup screen, where you can configure the meters displayed at  the  top  of
            the screen, set various display options, choose among color schemes, and select which
            columns are displayed, in which order.

       F3, /
            Incrementally search the command lines of all the displayed processes. The  currently
            selected  (highlighted) command will update as you type. While in search mode, press‐
            ing F3 will cycle through matching occurrences.  Pressing Shift-F3 will  cycle  back‐
            wards.

            Alternatively  the search can be started by simply typing the command you are looking
            for, although for the first character normal key bindings take precedence.

       F4, \
            Incremental process filtering: type in part of a process command line and  only  pro‐
            cesses  whose names match will be shown. To cancel filtering, enter the Filter option
            again and press Esc.  The matching is done case-insensitive. Terms are fixed  strings
            (no regex).  You can separate multiple terms with "|".

       F5, t
            Tree view: organize processes by parenthood, and layout the relations between them as
            a tree. Toggling the key will switch between tree and your previously  selected  sort
            view. Selecting a sort view will exit tree view.

       F6, <, >
            Selects a field for sorting, also accessible through < and >.  The current sort field
            is indicated by a highlight in the header.

       F7, ]
            Increase the selected process's priority (subtract from 'nice' value).  This can only
            be done by the superuser.

       F8, [
            Decrease the selected process's priority (add to 'nice' value)

       Shift-F7, }
            Increase  the  selected  process's autogroup priority (subtract from autogroup 'nice'
       Shift-F8, {
            Decrease the selected process's autogroup priority (add to autogroup 'nice' value)

       F9, k
            "Kill" process: sends a signal which is selected in a menu, to one or a group of pro‐
            cesses.  If processes were tagged, sends the signal to all tagged processes.  If none
            is tagged, sends to the currently selected process.

       F10, q
            Quit

       I    Invert the sort order: if sort order is increasing, switch to decreasing,  and  vice-
            versa.


..........
..........
..........

```
