# tee

The `tee` command in Unix and Linux is used to read from standard input and write to standard output and files simultaneously. It allows you to redirect the output of a command or a series of commands to multiple destinations, such as files or other commands.

### Basic Usage

The basic syntax for the `tee` command is:

```sh
command | tee [options] [file(s)]
```

- **`command`**: The command whose output you want to capture and duplicate.
- **`options`**: Optional command-line options to modify the behavior of `tee`.
- **`file(s)`**: The name(s) of the file(s) where you want to redirect the output.

### Examples

#### Reading from Standard Input

To capture the output of a command and display it on the terminal:

```sh
ls -l | tee output.txt
```

- This command lists the files in the current directory (`ls -l`) and simultaneously writes the output to both the terminal and the file `output.txt`.

#### Appending to a File

To append output to an existing file:

```sh
echo "Additional line" | tee -a output.txt
```

- The `-a` option appends (`tee -a`) the output to `output.txt`, rather than overwriting it.

#### Outputting to Multiple Files

To output to multiple files simultaneously:

```sh
ls -l | tee file1.txt file2.txt
```

- This command lists the files in the current directory and writes the output to both `file1.txt` and `file2.txt`.

#### Using tee with sudo

When using `tee` with `sudo` to write to protected files:

```sh
echo "Some content" | sudo tee /path/to/protected/file.txt
```

- This command uses `sudo` to write the output of `echo` to `/path/to/protected/file.txt`, allowing you to write to files that require elevated privileges.

### Options

#### Append to Files

- **`-a`**: Append to the given files rather than overwriting them.

#### Ignore Interrupts

- **`-i`**: Ignore interrupt signals (`SIGINT`), which allows `tee` to continue writing even if interrupted.

#### Output to Standard Error

- **`-p`**: Output to standard error as well as standard output.

#### Suppress Standard Output

- **`-q`**: Suppress standard output.

### Practical Use Cases

#### Monitoring Command Output

When debugging or monitoring the output of a command, `tee` allows you to observe it on the terminal while also saving it for future reference.

#### Logging Output

In scripts or automated processes, `tee` can log command output to files for later analysis or debugging.

#### Pipeline Management

Using `tee` in complex pipelines helps manage and duplicate output streams effectively, allowing for versatile data processing.

### Summary

The `tee` command is a useful tool for redirecting standard output to both the terminal and files simultaneously in Unix and Linux systems. It offers flexibility with options for appending, handling interrupts, and managing output streams, making it valuable for various scripting and command-line tasks. Understanding its usage and options can enhance your ability to manage and manipulate command output efficiently.

# help 

```
Usage: tee [OPTION]... [FILE]...
Copy standard input to each FILE, and also to standard output.

  -a, --append              append to the given FILEs, do not overwrite
  -i, --ignore-interrupts   ignore interrupt signals
  -p                        diagnose errors writing to non pipes
      --output-error[=MODE]   set behavior on write error.  See MODE below
      --help     display this help and exit
      --version  output version information and exit

```



## breakdown

```
-a, --append: This option appends to the file instead of overwriting it.
-i, --ignore-eof: This option does not stop at end-of-file (EOF) on standard input.
-h, --help: This option shows this help message.
```
