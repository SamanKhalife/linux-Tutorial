# head
The `head` command in Unix and Linux is used to display the beginning of a file or piped data. By default, it shows the first 10 lines of each file it processes, but you can customize the number of lines or bytes displayed.

### Basic Usage

The basic syntax for the `head` command is:

```sh
head [options] [file...]
```

- **`options`**: Command-line options to control the behavior of `head`.
- **`file`**: The file(s) to be processed. If no file is specified, `head` reads from standard input.

### Examples

#### Displaying the First 10 Lines of a File

To display the first 10 lines of a file:

```sh
head file.txt
```

#### Displaying a Custom Number of Lines

To display the first `n` lines of a file, use the `-n` option followed by the number of lines:

```sh
head -n 5 file.txt
```

This command displays the first 5 lines of `file.txt`.

#### Displaying a Custom Number of Bytes

To display the first `n` bytes of a file, use the `-c` option followed by the number of bytes:

```sh
head -c 20 file.txt
```

This command displays the first 20 bytes of `file.txt`.

#### Displaying Multiple Files

To display the first 10 lines of multiple files:

```sh
head file1.txt file2.txt
```

Output:
```
==> file1.txt <==
first 10 lines of file1.txt

==> file2.txt <==
first 10 lines of file2.txt
```

#### Piping Output to `head`

You can pipe the output of other commands into `head` to view the first few lines of the output:

```sh
ls -l | head
```

This command displays the first 10 lines of the `ls -l` output.

### Options

#### `-n` Option: Number of Lines

To specify the number of lines to display:

```sh
head -n 15 file.txt
```

This command displays the first 15 lines of `file.txt`.

#### `-c` Option: Number of Bytes

To specify the number of bytes to display:

```sh
head -c 50 file.txt
```

This command displays the first 50 bytes of `file.txt`.

#### `-v` Option: Always Print Headers

To always print headers when multiple files are provided, use the `-v` option:

```sh
head -v file1.txt file2.txt
```

This command ensures that headers are printed even if only one file is processed.

#### `-q` Option: Never Print Headers

To suppress headers when multiple files are provided, use the `-q` option:

```sh
head -q file1.txt file2.txt
```

This command prevents headers from being printed for each file.

### Practical Use Cases

#### Previewing Files

When you want to quickly preview the beginning of a file, `head` provides a convenient way to do so without opening the entire file in an editor.

#### Monitoring Log Files

When analyzing log files, you can use `head` to view the most recent entries (especially in combination with `tail`).

#### Debugging Scripts

When debugging scripts that produce a large amount of output, you can pipe the output to `head` to see only the initial part of the output for quick analysis.

### Summary

The `head` command is a simple yet powerful utility for viewing the beginning of files or piped data in Unix and Linux environments. Its flexibility in displaying a custom number of lines or bytes, combined with its ability to handle multiple files, makes it an essential tool for file and data inspection.
## help
```
Usage: head \[OPTION]... \[FILE]... Print the first 10 lines of each FILE to standard output. With more than one FILE, precede each with a header giving the file name.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too. -c, --bytes=\[-]NUM print the first NUM bytes of each file; with the leading '-', print all but the last NUM bytes of each file -n, --lines=\[-]NUM print the first NUM lines instead of the first 10; with the leading '-', print all but the last NUM lines of each file -q, --quiet, --silent never print headers giving file names -v, --verbose always print headers giving file names -z, --zero-terminated line delimiter is NUL, not newline --help display this help and exit --version output version information and exit

NUM may have a multiplier suffix: b 512, kB 1000, K 1024, MB 1000_1000, M 1024_1024, GB 1000_1000_1000, G 1024_1024_1024, and so on for T, P, E, Z, Y. Binary prefixes can be used, too: KiB=K, MiB=M, and so on.
```

## man

NAME head - output the first part of files

SYNOPSIS head \[OPTION]... \[FILE]...

DESCRIPTION Print the first 10 lines of each FILE to standard output. With more than one FILE, pre‐ cede each with a header giving the file name.


```
   With no FILE, or when FILE is -, read standard input.

   Mandatory arguments to long options are mandatory for short options too.

   -c, --bytes=[-]NUM
          print the first NUM bytes of each file; with the leading '-',  print  all  but  the
          last NUM bytes of each file

   -n, --lines=[-]NUM
          print  the first NUM lines instead of the first 10; with the leading '-', print all
          but the last NUM lines of each file

   -q, --quiet, --silent
          never print headers giving file names

   -v, --verbose
          always print headers giving file names

   -z, --zero-terminated
          line delimiter is NUL, not newline

   --help display this help and exit

   --version
          output version information and exit

   NUM may have a multiplier suffix: b 512, kB 1000, K 1024, MB 1000*1000,  M  1024*1024,  GB
   1000*1000*1000,  G  1024*1024*1024,  and  so on for T, P, E, Z, Y.  Binary prefixes can be
   used, too: KiB=K, MiB=M, and so on.
```
