# htop

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
