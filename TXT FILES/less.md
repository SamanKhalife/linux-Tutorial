# less

The `less` command in Unix and Linux is a pager program used to view the contents of text files interactively. It allows you to navigate through files, search for specific content, and perform various operations without loading the entire file into memory, making it efficient for viewing large files.

### Basic Usage

The basic syntax for the `less` command is:

```sh
less [options] [file]
```

- **`options`**: Command-line options to control the behavior of `less`.
- **`file`**: The file to be viewed. If no file is specified, `less` starts with standard input.

### Examples

#### Viewing a File

To view a file using `less`:

```sh
less file.txt
```

This command opens `file.txt` in the `less` pager, allowing you to scroll through its contents.

#### Navigating Within `less`

Once inside `less`, you can navigate using:

- **Arrow keys**: Scroll up and down.
- **Spacebar**: Scroll one page down.
- **Backspace**: Scroll one page up.
- **G**: Move to the end of the file.
- **1G** or **g**: Move to the beginning of the file.
- **/pattern**: Search for `pattern` forward.
- **?pattern**: Search for `pattern` backward.
- **n**: Move to the next occurrence of the search pattern.
- **N**: Move to the previous occurrence of the search pattern.
- **q**: Quit `less`.

#### Viewing Multiple Files

You can view multiple files in sequence by specifying them as arguments:

```sh
less file1.txt file2.txt
```

Press **:n** to move to the next file and **:p** to move to the previous file when viewing multiple files.

#### Viewing Standard Input

You can also pipe output from other commands into `less`:

```sh
ls -l | less
```

This command displays the output of `ls -l` in `less`, allowing you to scroll through the directory listing.

### Options

Some useful options for `less` include:

- **`-N`**: Display line numbers.
- **`-i`**: Ignore case in searches.
- **`-S`**: Chop long lines instead of wrapping.
- **`-R`**: Display ANSI color escape sequences in color.

### Practical Use Cases

#### Reading Log Files

When analyzing log files, `less` allows you to quickly navigate and search for specific entries without loading the entire file into memory.

#### Browsing Documentation

`less` is commonly used to view man pages (`man` command), providing a convenient way to read and search through system documentation.

#### Reviewing Code or Configuration Files

When reviewing code or configuration files, `less` enables easy navigation and searching through the content, facilitating efficient code review or troubleshooting.

### Summary

The `less` command is a versatile and efficient pager for viewing and navigating through text files interactively in Unix and Linux environments. Its ability to handle large files gracefully and provide powerful navigation and search capabilities makes it an essential tool for system administrators, developers, and anyone working with textual data.

# help

```

                   SUMMARY OF LESS COMMANDS

      Commands marked with * may be preceded by a number, N.
      Notes in parentheses indicate the behavior if N is given.
      A key preceded by a caret indicates the Ctrl key; thus ^K is ctrl-K.

  h  H                 Display this help.
  q  :q  Q  :Q  ZZ     Exit.
 ---------------------------------------------------------------------------

                           MOVING

  e  ^E  j  ^N  CR  *  Forward  one line   (or N lines).
  y  ^Y  k  ^K  ^P  *  Backward one line   (or N lines).
  f  ^F  ^V  SPACE  *  Forward  one window (or N lines).
  b  ^B  ESC-v      *  Backward one window (or N lines).
  z                 *  Forward  one window (and set window to N).
  w                 *  Backward one window (and set window to N).
  ESC-SPACE         *  Forward  one window, but don't stop at end-of-file.
  d  ^D             *  Forward  one half-window (and set half-window to N).
  u  ^U             *  Backward one half-window (and set half-window to N).
  ESC-)  RightArrow *  Right one half screen width (or N positions).
  ESC-(  LeftArrow  *  Left  one half screen width (or N positions).
  ESC-}  ^RightArrow   Right to last column displayed.
  ESC-{  ^LeftArrow    Left  to first column.
  F                    Forward forever; like "tail -f".
  ESC-F                Like F but stop when search pattern is found.
  r  ^R  ^L            Repaint screen.
  R                    Repaint screen, discarding buffered input.
        ---------------------------------------------------
        Default "window" is the screen height.
        Default "half-window" is half of the screen height.
 ---------------------------------------------------------------------------

                          SEARCHING

  /pattern          *  Search forward for (N-th) matching line.
  ?pattern          *  Search backward for (N-th) matching line.
  n                 *  Repeat previous search (for N-th occurrence).
  N                 *  Repeat previous search in reverse direction.
  ESC-n             *  Repeat previous search, spanning files.
  ESC-N             *  Repeat previous search, reverse dir. & spanning files.
  ESC-u                Undo (toggle) search highlighting.
  ESC-U                Clear search highlighting.
  &pattern          *  Display only matching lines.
        ---------------------------------------------------
        A search pattern may begin with one or more of:
        ^N or !  Search for NON-matching lines.
        ^E or *  Search multiple files (pass thru END OF FILE).
        ^F or @  Start search at FIRST file (for /) or last file (for ?).
        ^K       Highlight matches, but don't move (KEEP position).
        ^R       Don't use REGULAR EXPRESSIONS.
        ^W       WRAP search if no match found.
 ---------------------------------------------------------------------------

                           JUMPING

  g  <  ESC-<       *  Go to first line in file (or line N).
  G  >  ESC->       *  Go to last line in file (or line N).
  p  %              *  Go to beginning of file (or N percent into file).
  t                 *  Go to the (N-th) next tag.
  T                 *  Go to the (N-th) previous tag.
  {  (  [           *  Find close bracket } ) ].
  }  )  ]           *  Find open bracket { ( [.
  ESC-^F <c1> <c2>  *  Find close bracket <c2>.
  ESC-^B <c1> <c2>  *  Find open bracket <c1>.
        ---------------------------------------------------
        Each "find close bracket" command goes forward to the close bracket 
          matching the (N-th) open bracket in the top line.
        Each "find open bracket" command goes backward to the open bracket 
          matching the (N-th) close bracket in the bottom line.

  m<letter>            Mark the current top line with <letter>.
  M<letter>            Mark the current bottom line with <letter>.
  '<letter>            Go to a previously marked position.
  ''                   Go to the previous position.
  ^X^X                 Same as '.
  ESC-M<letter>        Clear a mark.
        ---------------------------------------------------
        A mark is any upper-case or lower-case letter.
        Certain marks are predefined:
             ^  means  beginning of the file
             $  means  end of the file
 ---------------------------------------------------------------------------
 
                        CHANGING FILES

  :e [file]            Examine a new file.
  ^X^V                 Same as :e.
  :n                *  Examine the (N-th) next file from the command line.
  :p                *  Examine the (N-th) previous file from the command line.
  :x                *  Examine the first (or N-th) file from the command line.
  :d                   Delete the current file from the command line list.
  =  ^G  :f            Print current file name.
 ---------------------------------------------------------------------------

                    MISCELLANEOUS COMMANDS

  -<flag>              Toggle a command line option [see OPTIONS below].
  --<name>             Toggle a command line option, by name.
  _<flag>              Display the setting of a command line option.
  __<name>             Display the setting of an option, by name.
  +cmd                 Execute the less cmd each time a new file is examined.

  !command             Execute the shell command with $SHELL.
  |Xcommand            Pipe file between current pos & mark X to shell command.
  s file               Save input to a file.
  v                    Edit the current file with $VISUAL or $EDITOR.
  V                    Print version number of "less".
 ---------------------------------------------------------------------------

                           OPTIONS

        Most options may be changed either on the command line,
        or from within less by using the - or -- command.
        Options may be given in one of two forms: either a single
        character preceded by a -, or a name preceded by --.

  -?  ........  --help
                  Display help (from command line).
  -a  ........  --search-skip-screen
                  Search skips current screen.
  -A  ........  --SEARCH-SKIP-SCREEN
                  Search starts just after target line.
  -b [N]  ....  --buffers=[N]
                  Number of buffers.
  -B  ........  --auto-buffers
                  Don't automatically allocate buffers for pipes.
  -c  ........  --clear-screen
                  Repaint by clearing rather than scrolling.
  -d  ........  --dumb
                  Dumb terminal.
  -D xcolor  .  --color=xcolor
                  Set screen colors.
  -e  -E  ....  --quit-at-eof  --QUIT-AT-EOF
                  Quit at end of file.
  -f  ........  --force
                  Force open non-regular files.
  -F  ........  --quit-if-one-screen
                  Quit if entire file fits on first screen.
  -g  ........  --hilite-search
                  Highlight only last match for searches.
  -G  ........  --HILITE-SEARCH
                  Don't highlight any matches for searches.
  -h [N]  ....  --max-back-scroll=[N]
                  Backward scroll limit.
  -i  ........  --ignore-case
                  Ignore case in searches that do not contain uppercase.
  -I  ........  --IGNORE-CASE
                  Ignore case in all searches.
  -j [N]  ....  --jump-target=[N]
                  Screen position of target lines.
  -J  ........  --status-column
                  Display a status column at left edge of screen.
  -k [file]  .  --lesskey-file=[file]
                  Use a lesskey file.
  -K  ........  --quit-on-intr
                  Exit less in response to ctrl-C.
  -L  ........  --no-lessopen
                  Ignore the LESSOPEN environment variable.
  -m  -M  ....  --long-prompt  --LONG-PROMPT
                  Set prompt style.
  -n  -N  ....  --line-numbers  --LINE-NUMBERS
                  Don't use line numbers.
  -o [file]  .  --log-file=[file]
                  Copy to log file (standard input only).
  -O [file]  .  --LOG-FILE=[file]
                  Copy to log file (unconditionally overwrite).
  -p [pattern]  --pattern=[pattern]
                  Start at pattern (from command line).
  -P [prompt]   --prompt=[prompt]
                  Define new prompt.
  -q  -Q  ....  --quiet  --QUIET  --silent --SILENT
                  Quiet the terminal bell.
  -r  -R  ....  --raw-control-chars  --RAW-CONTROL-CHARS
                  Output "raw" control characters.
  -s  ........  --squeeze-blank-lines
                  Squeeze multiple blank lines.
  -S  ........  --chop-long-lines
                  Chop (truncate) long lines rather than wrapping.
  -t [tag]  ..  --tag=[tag]
                  Find a tag.
  -T [tagsfile] --tag-file=[tagsfile]
                  Use an alternate tags file.
  -u  -U  ....  --underline-special  --UNDERLINE-SPECIAL
                  Change handling of backspaces.
  -V  ........  --version
                  Display the version number of "less".
  -w  ........  --hilite-unread
                  Highlight first new line after forward-screen.
  -W  ........  --HILITE-UNREAD
                  Highlight first new line after any forward movement.
  -x [N[,...]]  --tabs=[N[,...]]
                  Set tab stops.
  -X  ........  --no-init
                  Don't use termcap init/deinit strings.
  -y [N]  ....  --max-forw-scroll=[N]
                  Forward scroll limit.
  -z [N]  ....  --window=[N]
                  Set size of window.

```


# man
```
NAME
       less - opposite of more

SYNOPSIS
       less -?
       less --help
       less -V
       less --version
       less [-[+]aABcCdeEfFgGiIJKLmMnNqQrRsSuUVwWX~]
            [-b space] [-h lines] [-j line] [-k keyfile]
            [-{oO} logfile] [-p pattern] [-P prompt] [-t tag]
            [-T tagsfile] [-x tab,...] [-y lines] [-[z] lines]
            [-# shift] [+[+]cmd] [--] [filename]...
       (See the OPTIONS section for alternate option syntax with long option names.)

DESCRIPTION
       Less  is  a program similar to more(1), but it has many more features.  Less does not have
       to read the entire input file before starting, so with large  input  files  it  starts  up
       faster  than text editors like vi(1).  Less uses termcap (or terminfo on some systems), so
       it can run on a variety of terminals.  There is even limited support for  hardcopy  termi‐
       nals.  (On a hardcopy terminal, lines which should be printed at the top of the screen are
       prefixed with a caret.)

       Commands are based on both more and vi.  Commands may be preceded  by  a  decimal  number,
       called N in the descriptions below.  The number is used by some commands, as indicated.

COMMANDS
       In the following descriptions, ^X means control-X.  ESC stands for the ESCAPE key; for ex‐
       ample ESC-v means the two character sequence "ESCAPE", then "v".

       h or H Help: display a summary of these commands.  If you forget all the  other  commands,
              remember this one.

       SPACE or ^V or f or ^F
              Scroll  forward  N  lines,  default one window (see option -z below).  If N is more
              than the screen size, only the final screenful is displayed.  Warning: some systems
              use ^V as a special literalization character.

       z      Like SPACE, but if N is specified, it becomes the new window size.

       ESC-SPACE
              Like  SPACE,  but  scrolls  a full screenful, even if it reaches end-of-file in the
              process.

       ENTER or RETURN or ^N or e or ^E or j or ^J
              Scroll forward N lines, default 1.  The entire N lines are displayed, even if N  is
              more than the screen size.

       d or ^D
              Scroll forward N lines, default one half of the screen size.  If N is specified, it
              becomes the new default for subsequent d and u commands.

       b or ^B or ESC-v
              Scroll backward N lines, default one window (see option -z below).  If  N  is  more
              than the screen size, only the final screenful is displayed.

       w      Like ESC-v, but if N is specified, it becomes the new window size.

       y or ^Y or ^P or k or ^K
              Scroll backward N lines, default 1.  The entire N lines are displayed, even if N is
              more than the screen size.  Warning: some systems use ^Y as a special  job  control
              character.

       u or ^U
              Scroll  backward  N lines, default one half of the screen size.  If N is specified,
              it becomes the new default for subsequent d and u commands.

       J      Like j, but continues to scroll beyond the end of the file.

       K or Y Like k, but continues to scroll beyond the beginning of the file.

       ESC-) or RIGHTARROW
              Scroll horizontally right N characters, default half the screen width (see  the  -#
              option).   If a number N is specified, it becomes the default for future RIGHTARROW
              and LEFTARROW commands.  While the text is scrolled, it acts as though the  -S  op‐
              tion (chop lines) were in effect.

       ESC-( or LEFTARROW
              Scroll  horizontally  left  N characters, default half the screen width (see the -#
              option).  If a number N is specified, it becomes the default for future  RIGHTARROW
              and LEFTARROW commands.

       ESC-} or ^RIGHTARROW
              Scroll horizontally right to show the end of the longest displayed line.

       ESC-{ or ^LEFTARROW
              Scroll horizontally left back to the first column.

       r or ^R or ^L
              Repaint the screen.

       R      Repaint  the  screen,  discarding  any buffered input.  That is, reload the current
              file.  Useful if the file is changing while it is being viewed.

       F      Scroll forward, and keep trying to read when the end of file is reached.   Normally
              this  command  would  be  used when already at the end of the file.  It is a way to
              monitor the tail of a file which is growing while it is being viewed.  (The  behav‐
              ior is similar to the "tail -f" command.)  To stop waiting for more data, enter the
              interrupt character (usually ^C).  On some systems you can also use ^X.
       ESC-F  Like F, but as soon as a line is found which matches the last search  pattern,  the
              terminal bell is rung and forward scrolling stops.

       g or < or ESC-<
              Go  to  line  N  in the file, default 1 (beginning of file).  (Warning: this may be
              slow if N is large.)

       G or > or ESC->
              Go to line N in the file, default the end of the file.  (Warning: this may be  slow
              if  N is large, or if N is not specified and standard input, rather than a file, is
              being read.)

       ESC-G  Same as G, except if no number N is specified and the input is standard input, goes
              to the last line which is currently buffered.

       p or % Go  to  a position N percent into the file.  N should be between 0 and 100, and may
              contain a decimal point.

       P      Go to the line containing byte offset N in the file.

       {      If a left curly bracket appears in the top line displayed on the screen, the { com‐
              mand will go to the matching right curly bracket.  The matching right curly bracket
              is positioned on the bottom line of the screen.  If there is  more  than  one  left
              curly  bracket  on the top line, a number N may be used to specify the N-th bracket
              on the line.

       }      If a right curly bracket appears in the bottom line displayed on the screen, the  }
              command  will  go  to  the  matching  left  curly bracket.  The matching left curly
              bracket is positioned on the top line of the screen.  If there  is  more  than  one
              right  curly  bracket  on  the top line, a number N may be used to specify the N-th
              bracket on the line.

       (      Like {, but applies to parentheses rather than curly brackets.

       )      Like }, but applies to parentheses rather than curly brackets.
```



