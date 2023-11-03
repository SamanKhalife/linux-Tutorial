# pr

The "pr" command in Linux is used to prepare and format files for printing or for display purposes. It allows you to format text files, adding headers, footers, page breaks, and other formatting elements

# help

```
Usage: pr [OPTION]... [FILE]...
Paginate or columnate FILE(s) for printing.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  +FIRST_PAGE[:LAST_PAGE], --pages=FIRST_PAGE[:LAST_PAGE]
                    begin [stop] printing with page FIRST_[LAST_]PAGE
  -COLUMN, --columns=COLUMN
                    output COLUMN columns and print columns down,
                    unless -a is used. Balance number of lines in the
                    columns on each page
  -a, --across      print columns across rather than down, used together
                    with -COLUMN
  -c, --show-control-chars
                    use hat notation (^G) and octal backslash notation
  -d, --double-space
                    double space the output
  -D, --date-format=FORMAT
                    use FORMAT for the header date
  -e[CHAR[WIDTH]], --expand-tabs[=CHAR[WIDTH]]
                    expand input CHARs (TABs) to tab WIDTH (8)
  -F, -f, --form-feed
                    use form feeds instead of newlines to separate pages
                    (by a 3-line page header with -F or a 5-line header
                    and trailer without -F)
  -h, --header=HEADER
                    use a centered HEADER instead of filename in page header,
                    -h "" prints a blank line, don't use -h""
  -i[CHAR[WIDTH]], --output-tabs[=CHAR[WIDTH]]
                    replace spaces with CHARs (TABs) to tab WIDTH (8)
  -J, --join-lines  merge full lines, turns off -W line truncation, no column
                    alignment, --sep-string[=STRING] sets separators
  -l, --length=PAGE_LENGTH
                    set the page length to PAGE_LENGTH (66) lines
                    (default number of lines of text 56, and with -F 63).
                    implies -t if PAGE_LENGTH <= 10
  -m, --merge       print all files in parallel, one in each column,
                    truncate lines, but join lines of full length with -J
  -n[SEP[DIGITS]], --number-lines[=SEP[DIGITS]]
                    number lines, use DIGITS (5) digits, then SEP (TAB),
                    default counting starts with 1st line of input file
  -N, --first-line-number=NUMBER
                    start counting with NUMBER at 1st line of first
                    page printed (see +FIRST_PAGE)
  -o, --indent=MARGIN
                    offset each line with MARGIN (zero) spaces, do not
                    affect -w or -W, MARGIN will be added to PAGE_WIDTH
  -r, --no-file-warnings
                    omit warning when a file cannot be opened
  -s[CHAR], --separator[=CHAR]
                    separate columns by a single character, default for CHAR
                    is the <TAB> character without -w and 'no char' with -w.
                    -s[CHAR] turns off line truncation of all 3 column
                    options (-COLUMN|-a -COLUMN|-m) except -w is set
  -S[STRING], --sep-string[=STRING]
                    separate columns by STRING,
                    without -S: Default separator <TAB> with -J and <space>
                    otherwise (same as -S" "), no effect on column options
  -t, --omit-header  omit page headers and trailers;
                     implied if PAGE_LENGTH <= 10
  -T, --omit-pagination
                    omit page headers and trailers, eliminate any pagination
                    by form feeds set in input files
  -v, --show-nonprinting
                    use octal backslash notation
  -w, --width=PAGE_WIDTH
                    set page width to PAGE_WIDTH (72) characters for
                    multiple text-column output only, -s[char] turns off (72)
  -W, --page-width=PAGE_WIDTH
                    set page width to PAGE_WIDTH (72) characters always,
                    truncate lines, except -J option is set, no interference
                    with -S or -s
      --help     display this help and exit
      --version  output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
```
