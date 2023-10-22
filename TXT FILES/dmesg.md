# dmesg

The `dmesg` command in Linux is used to display the kernel ring buffer. The kernel ring buffer is a circular buffer that contains messages from the kernel. These messages can be used to troubleshoot problems with the kernel or with hardware devices.

The `dmesg` command is used in the following syntax:

```
dmesg [options]
```

The options can be used to specify the following:

* `-c` : Clear the kernel ring buffer.
* `-n` : Display the last `n` lines of the kernel ring buffer.
* `-s` : Display the kernel ring buffer in short format.

For example, the following code will clear the kernel ring buffer and display the last 10 lines:

```
dmesg -c | dmesg -n 10
```

This code will clear the kernel ring buffer and then display the last 10 lines.

The `dmesg` command is a powerful and versatile tool that can be used to troubleshoot problems with the kernel or with hardware devices. It is a valuable command to know, especially if you are a Linux system administrator.

Here are some additional things to note about the `dmesg` command:

* The `dmesg` command can be used to troubleshoot problems with the kernel or with hardware devices.
* The `dmesg` command can be used to clear the kernel ring buffer.
* The `dmesg` command can be used to display the kernel ring buffer in short format.

I hope this helps! Let me know if you have any other questions.

# help

```
Options:
 -C, --clear                 clear the kernel ring buffer
 -c, --read-clear            read and clear all messages
 -D, --console-off           disable printing messages to console
 -E, --console-on            enable printing messages to console
 -F, --file <file>           use the file instead of the kernel log buffer
 -f, --facility <list>       restrict output to defined facilities
 -H, --human                 human readable output
 -J, --json                  use JSON output format
 -k, --kernel                display kernel messages
 -L, --color[=<when>]        colorize messages (auto, always or never)
                               colors are enabled by default
 -l, --level <list>          restrict output to defined levels
 -n, --console-level <level> set level of messages printed to console
 -P, --nopager               do not pipe output into a pager
 -p, --force-prefix          force timestamp output on each line of multi-line messages
 -r, --raw                   print the raw message buffer
     --noescape              don't escape unprintable character
 -S, --syslog                force to use syslog(2) rather than /dev/kmsg
 -s, --buffer-size <size>    buffer size to query the kernel ring buffer
 -u, --userspace             display userspace messages
 -w, --follow                wait for new messages
 -W, --follow-new            wait and print only new messages
 -x, --decode                decode facility and level to readable string
 -d, --show-delta            show time delta between printed messages
 -e, --reltime               show local time and time delta in readable format
 -T, --ctime                 show human-readable timestamp (may be inaccurate!)
 -t, --notime                don't show any timestamp with messages
     --time-format <format>  show timestamp using the given format:
                               [delta|reltime|ctime|notime|iso]
Suspending/resume will make ctime and iso timestamps inaccurate.
     --since <time>          display the lines since the specified time
     --until <time>          display the lines until the specified time

 -h, --help                  display this help
 -V, --version               display version

Supported log facilities:
    kern - kernel messages
    user - random user-level messages
    mail - mail system
  daemon - system daemons
    auth - security/authorization messages
  syslog - messages generated internally by syslogd
     lpr - line printer subsystem
    news - network news subsystem

Supported log levels (priorities):
   emerg - system is unusable
   alert - action must be taken immediately
    crit - critical conditions
     err - error conditions
    warn - warning conditions
  notice - normal but significant condition
    info - informational
   debug - debug-level messages


NAME
       dmesg - print or control the kernel ring buffer

SYNOPSIS
       dmesg [options]

       dmesg --clear

       dmesg --read-clear [options]

       dmesg --console-level level

       dmesg --console-on

       dmesg --console-off

DESCRIPTION
       dmesg is used to examine or control the kernel ring buffer.

       The default action is to display all messages from the kernel ring buffer.

OPTIONS
       The --clear, --read-clear, --console-on, --console-off, and --console-level options are mutually exclusive.

       -C, --clear
           Clear the ring buffer.

       -c, --read-clear
           Clear the ring buffer after first printing its contents.

       -D, --console-off
           Disable the printing of messages to the console.

       -d, --show-delta
           Display the timestamp and the time delta spent between messages. If used together with --notime then only the time delta
           without the timestamp is printed.

       -E, --console-on
           Enable printing messages to the console.

       -e, --reltime
           Display the local time and the delta in human-readable format. Be aware that conversion to the local time could be
           inaccurate (see -T for more details).

       -F, --file file
           Read the syslog messages from the given file. Note that -F does not support messages in kmsg format. The old syslog format
           is supported only.

       -f, --facility list
           Restrict output to the given (comma-separated) list of facilities. For example:

           dmesg --facility=daemon

           will print messages from system daemons only. For all supported facilities see the --help output.

       -H, --human
           Enable human-readable output. See also --color, --reltime and --nopager.

       -J, --json
           Use JSON output format. The time output format is in "sec.usec" format only, log priority level is not decoded by default
           (use --decode to split into facility and priority), the other options to control the output format or time format are
           silently ignored.

       -k, --kernel
           Print kernel messages.

       -L, --color[=when]
           Colorize the output. The optional argument when can be auto, never or always. If the when argument is omitted, it defaults
           to auto. The colors can be disabled; for the current built-in default see the --help output. See also the COLORS section
           below.

       -l, --level list
           Restrict output to the given (comma-separated) list of levels. For example:

           dmesg --level=err,warn

           will print error and warning messages only. For all supported levels see the --help output.

       -n, --console-level level
           Set the level at which printing of messages is done to the console. The level is a level number or abbreviation of the level
           name. For all supported levels see the --help output.

           For example, -n 1 or -n emerg prevents all messages, except emergency (panic) messages, from appearing on the console. All
           levels of messages are still written to /proc/kmsg, so syslogd(8) can still be used to control exactly where kernel messages
           appear. When the -n option is used, dmesg will not print or clear the kernel ring buffer.

       --noescape
           The unprintable and potentially unsafe characters (e.g., broken multi-byte sequences, terminal controlling chars, etc.) are
           escaped in format \x<hex> for security reason by default. This option disables this feature at all. It’s usable for example
           for debugging purpose together with --raw. Be careful and don’t use it by default.

       -P, --nopager
           Do not pipe output into a pager. A pager is enabled by default for --human output.

       -p, --force-prefix
           Add facility, level or timestamp information to each line of a multi-line message.

       -r, --raw
           Print the raw message buffer, i.e., do not strip the log-level prefixes, but all unprintable characters are still escaped
           (see also --noescape).

           Note that the real raw format depends on the method how dmesg reads kernel messages. The /dev/kmsg device uses a different
           format than syslog(2). For backward compatibility, dmesg returns data always in the syslog(2) format. It is possible to read
           the real raw data from /dev/kmsg by, for example, the command 'dd if=/dev/kmsg iflag=nonblock'.

       -S, --syslog
           Force dmesg to use the syslog(2) kernel interface to read kernel messages. The default is to use /dev/kmsg rather than
           syslog(2) since kernel 3.5.0.

       -s, --buffer-size size
           Use a buffer of size to query the kernel ring buffer. This is 16392 by default. (The default kernel syslog buffer size was
           4096 at first, 8192 since 1.3.54, 16384 since 2.1.113.) If you have set the kernel buffer to be larger than the default,
           then this option can be used to view the entire buffer.
       -T, --ctime
           Print human-readable timestamps.

           Be aware that the timestamp could be inaccurate! The time source used for the logs is not updated after system
           SUSPEND/RESUME. Timestamps are adjusted according to current delta between boottime and monotonic clocks, this works only
           for messages printed after last resume.

       --since time
           Display record since the specified time. The time is possible to specify in absolute way as well as by relative notation
           (e.g. '1 hour ago'). Be aware that the timestamp could be inaccurate and see --ctime for more details.

       --until time
           Display record until the specified time. The time is possible to specify in absolute way as well as by relative notation
           (e.g. '1 hour ago'). Be aware that the timestamp could be inaccurate and see --ctime for more details.

       -t, --notime
           Do not print kernel’s timestamps.

       --time-format format
           Print timestamps using the given format, which can be ctime, reltime, delta or iso. The first three formats are aliases of
           the time-format-specific options. The iso format is a dmesg implementation of the ISO-8601 timestamp format. The purpose of
           this format is to make the comparing of timestamps between two systems, and any other parsing, easy. The definition of the
           iso timestamp is: YYYY-MM-DD<T>HH:MM:SS,<microseconds>←+><timezone offset from UTC>.

           The iso format has the same issue as ctime: the time may be inaccurate when a system is suspended and resumed.

       -u, --userspace
           Print userspace messages.

       -w, --follow
           Wait for new messages. This feature is supported only on systems with a readable /dev/kmsg (since kernel 3.5.0).

       -W, --follow-new
           Wait and print only new messages.

       -x, --decode
           Decode facility and level (priority) numbers to human-readable prefixes.

       -h, --help
           Display help text and exit.

       -V, --version
           Print version and exit.
```
