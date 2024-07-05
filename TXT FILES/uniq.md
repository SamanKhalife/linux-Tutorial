# uniq

The `uniq` command in Unix and Linux is used to report or filter out repeated lines in a file. By default, it only removes adjacent repeated lines. Therefore, it's often used in conjunction with the `sort` command to ensure that duplicate lines are adjacent.

### Basic Usage

The basic syntax for the `uniq` command is:

```sh
uniq [options] [input_file [output_file]]
```

- **`options`**: Command-line options to control the behavior of `uniq`.
- **`input_file`**: The file to be read. If no file is specified, `uniq` reads from standard input.
- **`output_file`**: The file to write the result to. If no file is specified, `uniq` writes to standard output.

### Examples

#### Removing Adjacent Duplicate Lines

To remove adjacent duplicate lines from a file:

```sh
uniq input.txt
```

This command removes adjacent duplicate lines from `input.txt` and outputs the result.

#### Counting Duplicate Lines

To display a count of repeated lines:

```sh
uniq -c input.txt
```

This command prefixes each line with the number of times it occurred.

Example output:
```
  1 First line
  2 Second line
  3 Third line
```

#### Ignoring Case

To compare lines case-insensitively:

```sh
uniq -i input.txt
```

This command ignores case differences when comparing lines.

#### Only Printing Duplicates

To print only duplicate lines (one for each group):

```sh
uniq -d input.txt
```

This command outputs only the lines that are repeated in the input.

#### Only Printing Unique Lines

To print only unique lines (lines that are not repeated):

```sh
uniq -u input.txt
```

This command outputs only the lines that are not repeated.

### Practical Use Cases

#### Sorting and Removing Duplicates

When working with unsorted files, use `sort` before `uniq` to remove all duplicates:

```sh
sort input.txt | uniq
```

This ensures that all duplicate lines are adjacent and can be removed by `uniq`.

#### Counting Unique Lines

To count the number of unique lines in a file:

```sh
sort input.txt | uniq | wc -l
```

This command sorts the file, removes duplicates, and counts the unique lines.

#### Filtering Log Files

When analyzing log files, `uniq` can help identify repeated entries:

```sh
sort logfile.log | uniq -c | sort -nr
```

This command sorts the log file, counts the occurrences of each line, and sorts the result in numerical order, showing the most frequent entries first.

### Options Summary

- `-c`: Prefix lines by the number of occurrences.
- `-d`: Only print duplicate lines.
- `-u`: Only print unique lines.
- `-i`: Ignore case when comparing lines.
- `-f N`: Skip the first N fields when comparing lines.
- `-s N`: Skip the first N characters when comparing lines.
- `-w N`: Compare no more than N characters in lines.

### Summary

The `uniq` command is a powerful tool for removing or identifying repeated lines in files. Its combination with other commands like `sort` makes it versatile for various text processing tasks. Understanding its options and practical use cases can greatly enhance your ability to manipulate and analyze text data in Unix and Linux environments.

# help 

```

```
