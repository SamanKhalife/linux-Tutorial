# bzgrep

The bzgrep command is a command-line utility that can be used to search for text in bzip2-compressed files. It is similar to the grep command, but it can be used to search for text in compressed files.

The bzgrep command has a number of options that can be used to control the behavior of the command. Some of the most commonly used options are:

-c: This option prints the number of matches that were found.
-i: This option ignores case differences when searching for text.
-n: This option prints the line number of each match.
-v: This option prints all lines that do not contain the specified pattern.
For example, the following command will search for the text "hello" in the bzip2-compressed file file.bz2:

`bzgrep "hello" file.bz2`

The output of the bzgrep command will be a list of all the lines in the file that contain the text "hello".



# help 

```
bzgrep [options] PATTERN FILE

Search for PATTERN in FILE, which may be a compressed file.

Options:

-c, --count   Print the number of matches.
-i, --ignore-case   Ignore case differences when searching for text.
-n, --line-number   Print the line number of each match.
-v, --invert-match   Print all lines that do not contain the specified pattern.
-h, --help    Display this help message.

Examples:

    bzgrep "hello" file.bz2
    bzgrep -i "hello" file.bz2
    bzgrep -c "hello" file.bz2
    bzgrep -n "hello" file.bz2
    bzgrep -v "hello" file.bz2
```
