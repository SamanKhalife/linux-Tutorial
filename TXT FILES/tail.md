# tail
The `tail` command in Unix and Linux is used to display the last part of files. By default, `tail` outputs the last 10 lines of a file, but you can customize the output with various options. This command is particularly useful for monitoring log files and other files that are continuously updated.

### Basic Usage

The basic syntax for the `tail` command is:

```sh
tail [options] [file...]
```

- **`file`**: The file to be read. If no file is specified, `tail` reads from standard input.

### Examples

#### Displaying the Last 10 Lines

To display the last 10 lines of a file:

```sh
tail file.txt
```

#### Displaying a Specific Number of Lines

To display a specific number of lines from the end of a file, use the `-n` option followed by the number of lines:

```sh
tail -n 20 file.txt
```

This command displays the last 20 lines of `file.txt`.

Alternatively, you can use the shorthand form:

```sh
tail -20 file.txt
```

#### Displaying Lines from a Specific Point

To display lines starting from a specific position in the file, use the `-c` option followed by the number of bytes:

```sh
tail -c 100 file.txt
```

This command displays the last 100 bytes of `file.txt`.

### Monitoring File Changes

The `-f` option is used to follow a file as it grows. This is useful for monitoring log files in real-time:

```sh
tail -f /var/log/syslog
```

With this command, `tail` will display the last 10 lines of `/var/log/syslog` and continue to output new lines as they are added to the file.

To stop following the file, press `Ctrl+C`.

### Combining Options

You can combine the `-f` option with the `-n` option to start displaying from a specific number of lines and continue monitoring the file:

```sh
tail -n 20 -f /var/log/syslog
```

This command displays the last 20 lines of `/var/log/syslog` and continues to output new lines as they are added.

### Practical Use Cases

#### Monitoring Log Files

System administrators often use `tail` to monitor log files for troubleshooting purposes. For example, to monitor the Apache web server log:

```sh
tail -f /var/log/apache2/access.log
```

#### Debugging

Developers can use `tail` to monitor application logs while debugging to see real-time output and error messages:

```sh
tail -f application.log
```

#### Displaying the End of a Long File

When working with large files, you might want to quickly view the most recent entries:

```sh
tail large_file.txt
```

### Summary

The `tail` command is a versatile tool for viewing the end of files and monitoring changes in real-time. It is particularly useful for monitoring log files and other continuously updated files. By understanding the various options and usage scenarios, you can effectively use `tail` to streamline your workflow and enhance your productivity.
# help 

```
Usage: tail [OPTION]... [FILE]...
Print the last 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[+]NUM       output the last NUM bytes; or use -c +NUM to
                             output starting with byte NUM of each file
  -f, --follow[={name|descriptor}]
                           output appended data as the file grows;
                             an absent option argument means 'descriptor'
  -F                       same as --follow=name --retry
  -n, --lines=[+]NUM       output the last NUM lines, instead of the last 10;
                             or use -n +NUM to output starting with line NUM
      --max-unchanged-stats=N
                           with --follow=name, reopen a FILE which has not
                             changed size after N (default 5) iterations
                             to see if it has been unlinked or renamed
                             (this is the usual case of rotated log files);
                             with inotify, this option is rarely useful
      --pid=PID            with -f, terminate after process ID, PID dies
  -q, --quiet, --silent    never output headers giving file names
      --retry              keep trying to open a file if it is inaccessible
  -s, --sleep-interval=N   with -f, sleep for approximately N seconds
                             (default 1.0) between iterations;
                             with inotify and --pid=P, check process P at
                             least once every N seconds
  -v, --verbose            always output headers giving file names
  -z, --zero-terminated    line delimiter is NUL, not newline
      --help     display this help and exit
      --version  output version information and exit

NUM may have a multiplier suffix:
b 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,
GB 1000*1000*1000, G 1024*1024*1024, and so on for T, P, E, Z, Y.
Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

With --follow (-f), tail defaults to following the file descriptor, which
means that even if a tail'ed file is renamed, tail will continue to track
its end.  This default behavior is not desirable when you really want to
track the actual name of the file, not the file descriptor (e.g., log
rotation).  Use --follow=name in that case.  That causes tail to track the
named file in a way that accommodates renaming, removal and creation.

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Report any translation bugs to <https://translationproject.org/team/>
Full documentation <https://www.gnu.org/software/coreutils/tail>
or available locally via: info '(coreutils) tail invocation'
root@Saman:~# tail --help
Usage: tail [OPTION]... [FILE]...
Print the last 10 lines of each FILE to standard output.
With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -c, --bytes=[+]NUM       output the last NUM bytes; or use -c +NUM to
                             output starting with byte NUM of each file
  -f, --follow[={name|descriptor}]
                           output appended data as the file grows;
                             an absent option argument means 'descriptor'
  -F                       same as --follow=name --retry
  -n, --lines=[+]NUM       output the last NUM lines, instead of the last 10;
                             or use -n +NUM to output starting with line NUM
      --max-unchanged-stats=N
                           with --follow=name, reopen a FILE which has not
                             changed size after N (default 5) iterations
                             to see if it has been unlinked or renamed
                             (this is the usual case of rotated log files);
                             with inotify, this option is rarely useful
      --pid=PID            with -f, terminate after process ID, PID dies
  -q, --quiet, --silent    never output headers giving file names
      --retry              keep trying to open a file if it is inaccessible
  -s, --sleep-interval=N   with -f, sleep for approximately N seconds
                             (default 1.0) between iterations;
                             with inotify and --pid=P, check process P at
                             least once every N seconds
  -v, --verbose            always output headers giving file names
  -z, --zero-terminated    line delimiter is NUL, not newline
      --help     display this help and exit
      --version  output version information and exit

```

