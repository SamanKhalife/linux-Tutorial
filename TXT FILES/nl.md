# nl

The `nl` command in Unix and Linux is used to number lines in a file or standard input. It's particularly useful for adding line numbers to text files or when processing data that requires line-by-line numbering.

### Basic Usage

The basic syntax for the `nl` command is:

```sh
nl [options] [file]
```

- **`options`**: Command-line options to control the behavior of `nl`.
- **`file`**: The file to be processed. If no file is specified, `nl` reads from standard input.

### Examples

#### Numbering Lines in a File

To number lines in a file:

```sh
nl file.txt
```

This command outputs each line of `file.txt` preceded by its line number.

Example output:
```
     1  line 1
     2  line 2
     3  line 3
```

#### Displaying Line Numbers Only

To display only line numbers without the content of each line:

```sh
nl -n ln file.txt
```

Example output:
```
     1
     2
     3
```

#### Customizing Line Number Formatting

To customize the format of line numbers, use the `-n` option followed by a format specifier:

- `ln` - Left justified, no leading zeros.
- `rn` - Right justified, no leading zeros.
- `rz` - Right justified, leading zeros.

For example, to right-justify line numbers with leading zeros:

```sh
nl -n rz file.txt
```

Example output:
```
000001  line 1
000002  line 2
000003  line 3
```

#### Adding Headers

To add a header to each page of output:

```sh
nl -s '. ' -b a file.txt
```

This command adds line numbers only to non-empty lines and separates them from the content with a period and space.

Example output:
```
     1. line 1
     2. line 2
     3. line 3
```

### Options

#### `-b` Option: Body Selection

- `a` - Number all lines.
- `t` - Number only non-empty lines.
- `n` - Do not number lines.

#### `-s` Option: Separator

Specifies the string to use as a separator between the line number and the line contents.

#### `-w` Option: Line Width

Specifies the number of columns to use for the line number. Useful for adjusting spacing.

#### `-v` Option: Starting Line Number

Specifies the starting line number. By default, it starts at 1.

### Practical Use Cases

#### Numbering Code Files

When reviewing code or scripts, numbering lines helps reference specific lines for discussion or debugging.

#### Adding Line Numbers to Reports

When generating reports or documents, `nl` can be used to add line numbers for easier navigation and referencing.

#### Formatting Data Files

When working with data files that lack structure, `nl` can add structure by numbering each line, making it easier to process and analyze.

### Summary

The `nl` command is a straightforward yet powerful utility for numbering lines in text files or data streams. It provides several options for customizing line numbering formats, handling empty lines, and adding headers, making it versatile for various text processing tasks. Understanding these options can significantly enhance your ability to manipulate and analyze textual data efficiently.

# help

```
Usage: nl [OPTION]... [FILE]...
Write each FILE to standard output, with line numbers added.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -b, --body-numbering=STYLE      use STYLE for numbering body lines
  -d, --section-delimiter=CC      use CC for logical page delimiters
  -f, --footer-numbering=STYLE    use STYLE for numbering footer lines
  -h, --header-numbering=STYLE    use STYLE for numbering header lines
  -i, --line-increment=NUMBER     line number increment at each line
  -l, --join-blank-lines=NUMBER   group of NUMBER empty lines counted as one
  -n, --number-format=FORMAT      insert line numbers according to FORMAT
  -p, --no-renumber               do not reset line numbers for each section
  -s, --number-separator=STRING   add STRING after (possible) line number
  -v, --starting-line-number=NUMBER  first line number for each section
  -w, --number-width=NUMBER       use NUMBER columns for line numbers
      --help     display this help and exit
      --version  output version information and exit

Default options are: -bt -d'\:' -fn -hn -i1 -l1 -n'rn' -s<TAB> -v1 -w6

CC are two delimiter characters used to construct logical page delimiters;
a missing second character implies ':'.

STYLE is one of:

  a      number all lines
  t      number only nonempty lines
  n      number no lines
  pBRE   number only lines that contain a match for the basic regular
         expression, BRE

FORMAT is one of:

  ln     left justified, no leading zeros
  rn     right justified, no leading zeros
  rz     right justified, leading zeros


GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Report any translation bugs to <https://translationproject.org/team/>
Full documentation <https://www.gnu.org/software/coreutils/nl>
or available locally via: info '(coreutils) nl invocation'
```
