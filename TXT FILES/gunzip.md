# gunzip


The `gunzip` command in Unix and Linux is used to decompress files that have been compressed with gzip. It is a straightforward tool for handling gzip-compressed files and is commonly used for file decompression tasks.

### Basic Usage

The basic syntax for the `gunzip` command is:

```sh
gunzip [options] [file(s)]
```

- **`options`**: Optional command-line options to control the behavior of `gunzip`.
- **`file(s)`**: The name(s) of the file(s) to decompress.

### Examples

#### Decompressing a File

To decompress a file using `gunzip`:

```sh
gunzip filename.txt.gz
```

This command decompresses `filename.txt.gz` and restores it to `filename.txt`.

#### Decompressing with `-d` Option

The `-d` option can also be used to achieve the same result:

```sh
gunzip -d filename.txt.gz
```

#### Decompressing Multiple Files

To decompress multiple files at once:

```sh
gunzip file1.txt.gz file2.txt.gz
```

This command decompresses `file1.txt.gz` and `file2.txt.gz`, resulting in `file1.txt` and `file2.txt`.

### Options

#### Keep Original File

- **`-k`**: Keep the original compressed file after decompression.

#### Verbose Output

- **`-v`**: Verbose mode, display decompression statistics.

### Practical Use Cases

#### Automated Decompression

To decompress all `.gz` files in a directory and its subdirectories:

```sh
gunzip -r /path/to/files/*.gz
```

This command recursively decompresses all `.gz` files under `/path/to/files`.

#### Keeping Original Files

To keep the original `.gz` files after decompression:

```sh
gunzip -k file1.txt.gz file2.txt.gz
```

This command decompresses `file1.txt.gz` and `file2.txt.gz` while keeping the original compressed files.

### Summary

The `gunzip` command is a simple yet effective tool for decompressing gzip-compressed files in Unix and Linux environments. It offers options for maintaining original files, providing verbose output, and handling multiple files simultaneously. Understanding its usage and options can help you efficiently manage file decompression tasks on your system.

# help

```
Usage: /usr/bin/gunzip [OPTION]... [FILE]...
Uncompress FILEs (by default, in-place).

Mandatory arguments to long options are mandatory for short options too.

  -c, --stdout      write on standard output, keep original files unchanged
  -f, --force       force overwrite of output file and compress links
  -k, --keep        keep (don't delete) input files
  -l, --list        list compressed file contents
  -n, --no-name     do not save or restore the original name and timestamp
  -N, --name        save or restore the original name and timestamp
  -q, --quiet       suppress all warnings
  -r, --recursive   operate recursively on directories
  -S, --suffix=SUF  use suffix SUF on compressed files
      --synchronous synchronous output (safer if system crashes, but slower)
  -t, --test        test compressed file integrity
  -v, --verbose     verbose mode
      --help        display this help and exit
      --version     display version information and exit

With no FILE, or when FILE is -, read standard input.

Report bugs to <bug-gzip@gnu.org>.
```
