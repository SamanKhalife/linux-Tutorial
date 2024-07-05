# cat

The `cat` command in Unix and Linux is used to concatenate and display the content of files. It's one of the most frequently used commands for viewing file contents, combining files, and redirecting output.

### Basic Usage

The basic syntax for the `cat` command is:

```sh
cat [options] [file...]
```

- **`options`**: Command-line options to control the behavior of `cat`.
- **`file`**: The file(s) to be read. If no file is specified, `cat` reads from standard input.

### Examples

#### Displaying File Contents

To display the contents of a file:

```sh
cat file.txt
```

This command outputs the entire content of `file.txt` to the standard output (usually the terminal).

#### Concatenating Files

To concatenate multiple files and display their contents:

```sh
cat file1.txt file2.txt
```

This command displays the contents of `file1.txt` followed by `file2.txt`.

#### Redirecting Output to a File

To concatenate files and redirect the output to a new file:

```sh
cat file1.txt file2.txt > combined.txt
```

This command combines `file1.txt` and `file2.txt` into a new file named `combined.txt`.

#### Appending to a File

To append the contents of one file to another:

```sh
cat file1.txt >> file2.txt
```

This command appends the contents of `file1.txt` to the end of `file2.txt`.

### Options

#### `-n` Option: Numbering Lines

To number all output lines:

```sh
cat -n file.txt
```

This command displays the content of `file.txt` with line numbers.

Example output:
```
     1  First line
     2  Second line
     3  Third line
```

#### `-b` Option: Numbering Non-Empty Lines

To number only non-empty lines:

```sh
cat -b file.txt
```

This command displays the content of `file.txt` with line numbers for non-empty lines only.

Example output:
```
     1  First line

     2  Second line
     3  Third line
```

#### `-s` Option: Squeeze Blank Lines

To suppress repeated empty lines:

```sh
cat -s file.txt
```

This command reduces multiple consecutive blank lines to a single blank line.

#### `-E` Option: Displaying End-of-Line Characters

To display a `$` at the end of each line:

```sh
cat -E file.txt
```

This command helps visualize the end of each line.

Example output:
```
First line$
Second line$
Third line$
```

### Practical Use Cases

#### Viewing File Contents

The most common use of `cat` is to quickly view the contents of a file.

```sh
cat /etc/passwd
```

#### Creating Files

You can create a new file by using `cat` and redirecting input from the terminal:

```sh
cat > newfile.txt
```

Type the content you want to add to `newfile.txt` and press `Ctrl+D` to save and exit.

#### Combining Files

`cat` is useful for combining multiple files into one:

```sh
cat part1.txt part2.txt part3.txt > complete.txt
```

This combines `part1.txt`, `part2.txt`, and `part3.txt` into `complete.txt`.

#### Displaying Line Numbers

To help debug scripts or code, you can display line numbers:

```sh
cat -n script.sh
```

### Summary

The `cat` command is a simple yet powerful tool for viewing, combining, and manipulating file content in Unix and Linux environments. Its versatility makes it an essential command for file management and text processing tasks.
# help

```
Usage: cat [OPTION]... [FILE]...
Concatenate FILE(s) to standard output.

With no FILE, or when FILE is -, read standard input.

  -A, --show-all           equivalent to -vET
  -b, --number-nonblank    number nonempty output lines, overrides -n
  -e                       equivalent to -vE
  -E, --show-ends          display $ at end of each line
  -n, --number             number all output lines
  -s, --squeeze-blank      suppress repeated empty output lines
  -t                       equivalent to -vT
  -T, --show-tabs          display TAB characters as ^I
  -u                       (ignored)
  -v, --show-nonprinting   use ^ and M- notation, except for LFD and TAB
      --help     display this help and exit
      --version  output version information and exit

Examples:
  cat f - g  Output f's contents, then standard input, then g's contents.
  cat        Copy standard input to standard output.
```

man

```
NAME
       cat - concatenate files and print on the standard output

SYNOPSIS
       cat [OPTION]... [FILE]...

DESCRIPTION
       Concatenate FILE(s) to standard output.

       With no FILE, or when FILE is -, read standard input.

       -A, --show-all
              equivalent to -vET

       -b, --number-nonblank
              number nonempty output lines, overrides -n

       -e     equivalent to -vE

       -E, --show-ends
              display $ at end of each line

       -n, --number
              number all output lines

       -s, --squeeze-blank
              suppress repeated empty output lines

       -t     equivalent to -vT

       -T, --show-tabs
              display TAB characters as ^I

       -u     (ignored)

       -v, --show-nonprinting
              use ^ and M- notation, except for LFD and TAB

       --help display this help and exit

       --version
              output version information and exit

EXAMPLES
       cat f - g
              Output f's contents, then standard input, then g's contents.

       cat    Copy standard input to standard output.
```
