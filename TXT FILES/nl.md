# nl

The `nl` command in Linux is used to number lines in a file. It is a simple command, but it can be very useful for debugging programs and for making files easier to read.

The syntax of the `nl` command is as follows:

```
nl [options] file
```

The `file` argument specifies the file that you want to number the lines for.

The `options` argument controls the output of the `nl` command. The most common options are as follows:

* `-b`: Numbers the lines starting from 1.
* `-n`: Specifies the number of columns to use for the line numbers.
* `-s`: Specifies the separator to use between the line numbers and the text.

For example, the following command will number the lines in the file `file.txt`, starting from 1, and using a space as the separator:

```
nl -b -n 5 -s ' ' file.txt
```

This command will number the lines in the file `file.txt`, starting from 1, and using a space as the separator. The numbers will be displayed in columns of 5.

The `nl` command is a simple command, but it can be very useful for debugging programs and for making files easier to read.

Here are some additional things to keep in mind about the `nl` command:

* The `nl` command can be used to number the lines in any type of file, including text files, binary files, and compressed files.
* The `nl` command can be used to number the lines in a file that is piped to it.
* The `nl` command can be used to number the lines in a file that is being read from standard input.

It is important to be aware of these limitations when using the `nl` command, so that you do not get confused by the output.


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
