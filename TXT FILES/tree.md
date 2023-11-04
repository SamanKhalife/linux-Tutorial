# tree

The `tree` command in Linux is used to display the directory structure of a file system. The `tree` command is a recursive command, which means that it will display the contents of all subdirectories as well.

The `tree` command is used as follows:

```
tree [options] [directory]
```

* `options`: These are optional flags that can be used to control the behavior of the `tree` command.
* `directory`: This is the directory that will be displayed. The default directory is the current directory.

The `tree` command has a number of options that can be used to control the output of the command. Some of the most commonly used `tree` options are:

* `-d`: This option specifies that only directories should be displayed.
* `-f`: This option specifies that the output should be formatted in a tree-like format.
* `-L`: This option specifies the maximum depth of the tree.
* `-s`: This option specifies that the size of each file or directory should be displayed.

For example, the following command will display the contents of the current directory in a tree-like format:

```
tree
```

The following command will display the contents of the `/etc` directory in a tree-like format, with a maximum depth of 2:

```
tree -L 2 /etc
```

The `tree` command is a valuable tool for system administrators and users who need to view the directory structure of a file system. It can be used to troubleshoot problems, to find files, and to understand the file system layout.


# help 

```

usage: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L level [-R]] [-H  baseHREF]
        [-T title] [-o filename] [-P pattern] [-I pattern] [--gitignore]
        [--matchdirs] [--metafirst] [--ignore-case] [--nolinks] [--inodes]
        [--device] [--sort[=]<name>] [--dirsfirst] [--filesfirst]
        [--filelimit #] [--si] [--du] [--prune] [--charset X]
        [--timefmt[=]format] [--fromfile] [--noreport] [--version] [--help]
        [--] [directory ...]
  ------- Listing options -------
  -a            All files are listed.
  -d            List directories only.
  -l            Follow symbolic links like directories.
  -f            Print the full path prefix for each file.
  -x            Stay on current filesystem only.
  -L level      Descend only level directories deep.
  -R            Rerun tree when max dir level reached.
  -P pattern    List only those files that match the pattern given.
  -I pattern    Do not list files that match the given pattern.
  --gitignore   Filter by using .gitignore files.
  --ignore-case Ignore case when pattern matching.
  --matchdirs   Include directory names in -P pattern matching.
  --metafirst   Print meta-data at the beginning of each line.
  --prune       Prune empty directories from the output.
  --info        Print information about files found in .info files.
  --noreport    Turn off file/directory count at end of tree listing.
  --charset X   Use charset X for terminal/HTML and indentation line output.
  --filelimit # Do not descend dirs with more than # files in them.
  -o filename   Output to file instead of stdout.
  ------- File options -------
  -q            Print non-printable characters as '?'.
  -N            Print non-printable characters as is.
  -Q            Quote filenames with double quotes.
  -p            Print the protections for each file.
  -u            Displays file owner or UID number.
  -g            Displays file group owner or GID number.
  -s            Print the size in bytes of each file.
  -h            Print the size in a more human readable way.
  --si          Like -h, but use in SI units (powers of 1000).
  --du          Compute size of directories by their contents.
  -D            Print the date of last modification or (-c) status change.
  --timefmt <f> Print and format time according to the format <f>.
  -F            Appends '/', '=', '*', '@', '|' or '>' as per ls -F.
  --inodes      Print inode number of each file.
  --device      Print device ID number to which each file belongs.
  ------- Sorting options -------
  -v            Sort files alphanumerically by version.
  -t            Sort files by last modification time.
  -c            Sort files by last status change time.
  -U            Leave files unsorted.
  -r            Reverse the order of the sort.
  --dirsfirst   List directories before files (-U disables).
  --filesfirst  List files before directories (-U disables).
  --sort X      Select sort: name,version,size,mtime,ctime.
  ------- Graphics options -------
  -i            Don't print indentation lines.
  -A            Print ANSI lines graphic indentation lines.
  -S            Print with CP437 (console) graphics indentation lines.
  -n            Turn colorization off always (-C overrides).
  -C            Turn colorization on always.
  ------- XML/HTML/JSON options -------
  -X            Prints out an XML representation of the tree.
  -J            Prints out an JSON representation of the tree.
  -H baseHREF   Prints out HTML format with baseHREF as top directory.
  -T string     Replace the default HTML title and H1 header with string.
  --nolinks     Turn off hyperlinks in HTML output.
  ------- Input options -------
  --fromfile    Reads paths from files (.=stdin)
  ------- Miscellaneous options -------
  --version     Print version and exit.
  --help        Print usage and this help message and exit.
  --            Options processing terminator.
```
man 

```
NAME
       tree - list contents of directories in a tree-like format.

SYNOPSIS
       tree  [-acdfghilnpqrstuvxACDFJQNSUX]  [-L level [-R]] [-H baseHREF] [-T title] [-o filename] [-P pattern] [-I pattern] [--gitig‐
       nore] [--matchdirs] [--metafirst] [--ignore-case] [--nolinks] [--inodes] [--device] [--sort[=]name] [--dirsfirst] [--filesfirst]
       [--filelimit  #]  [--si] [--du] [--prune] [--timefmt[=]format] [--fromfile] [--info] [--noreport] [--version] [--help] [--] [di‐
       rectory ...]

DESCRIPTION
       Tree is a recursive directory listing program that produces a depth indented listing of files, which is colorized ala  dircolors
       if  the  LS_COLORS environment variable is set and output is to tty.  With no arguments, tree lists the files in the current di‐
       rectory.  When directory arguments are given, tree lists all the files and/or directories found in the given directories each in
       turn.  Upon completion of listing all files/directories found, tree returns the total number of files and/or directories listed.

       By default, when a symbolic link is encountered, the path that the symbolic link refers to is printed after the name of the link
       in the format:

           name -> real-path

       If the `-l' option is given and the symbolic link refers to an actual directory, then tree will follow the path of the  symbolic
       link as if it were a real directory.

OPTIONS
       Tree understands the following command line switches:

LISTING OPTIONS
       -a     All  files  are printed.  By default tree does not print hidden files (those beginning with a dot `.').  In no event does
              tree print the file system constructs `.' (current directory) and `..' (previous directory).

       -d     List directories only.

       -l     Follows symbolic links if they point to directories, as if they were directories. Symbolic links that will result in  re‐
              cursion are avoided when detected.

       -f     Prints the full path prefix for each file.

       -x     Stay on the current file-system only.  Ala find -xdev.

       -L level
              Max display depth of the directory tree.

       -R     Recursively cross down the tree each level directories (see -L option), and at each of them execute tree again adding `-o
              00Tree.html' as a new option.

       -P pattern
              List only those files that match the wild-card pattern.  You may have multiple -P options. Note: you must use the -a  op‐
              tion  to  also consider those files beginning with a dot `.' for matching.  Valid wildcard operators are `*' (any zero or
              more characters), `**` (any zero or more characters as well as null /'s, i.e. /**/ may match a single /), `?' (any single
              character), `[...]' (any single character listed between brackets (optional - (dash) for character range may be used: ex:
              [A-Z]), and `[^...]' (any single character not listed in brackets) and `|' separates alternate patterns. A '/' at the end
              of the pattern matches directories, but not files.

       -I pattern
              Do  not  list those files that match the wild-card pattern.  You may have multiple -I options.  See -P above for informa‐
              tion on wildcard patterns.
       --gitignore
              Uses git .gitignore files for filtering files and directories.  Also uses $GIT_DIR/info/exclude if present.

       --ignore-case
              If a match pattern is specified by the -P or -I option, this will cause the pattern to match without regards to the  case
              of each letter.

       --matchdirs
              If  a match pattern is specified by the -P option, this will cause the pattern to be applied to directory names (in addi‐
              tion to filenames).  In the event of a match on the directory name, matching is disabled for the directory's contents. If
              the --prune option is used, empty folders that match the pattern will not be pruned.

       --metafirst
              Print the meta-data information at the beginning of the line rather than after the indentation lines.

       --prune
              Makes  tree  prune  empty directories from the output, useful when used in conjunction with -P or -I.  See BUGS AND NOTES
              below for more information on this option.

       --info Prints file comments found in .info files.  See .INFO FILES below for more information on the format of .info files.

       --noreport
              Omits printing of the file and directory report at the end of the tree listing.

       --charset charset
              Set the character set to use when outputting HTML and for line drawing.

       --filelimit #
              Do not descend directories that contain more than # entries.

       --timefmt format
              Prints (implies -D) and formats the date according to the format string which uses the strftime(3) syntax.

       -o filename
              Send output to filename.

FILE OPTIONS
       -q     Print non-printable characters in filenames as question marks instead of the default.

       -N     Print non-printable characters as is instead of as escaped octal numbers.

       -Q     Quote the names of files in double quotes.

       -p     Print the file type and permissions for each file (as per ls -l).

       -u     Print the username, or UID # if no username is available, of the file.

       -g     Print the group name, or GID # if no group name is available, of the file.

       -u     Print the username, or UID # if no username is available, of the file.

       -g     Print the group name, or GID # if no group name is available, of the file.

       -s     Print the size of each file in bytes along with the name.

       -h     Print  the  size of each file but in a more human readable way, e.g. appending a size letter for kilobytes (K), megabytes (M), gigabytes (G), terabytes (T),
              petabytes (P) and exabytes (E).

       --si   Like -h but use SI units (powers of 1000) instead.

       --du   For each directory report its size as the accumulation of sizes of all its files and sub-directories (and their files, and so on).  The total amount of used
              space is also given in the final report (like the 'du -c' command.) This option requires tree to read the entire directory tree before emitting it, see BUGS
              AND NOTES below.  Implies -s.

       -D     Print the date of the last modification time or if -c is used, the last status change time for the file listed.

       -F     Append a `/' for directories, a `=' for socket files, a `*' for executable files, a `>' for doors (Solaris) and a `|' for FIFO's, as per ls -F

       --inodes
              Prints the inode number of the file or directory

       --device
              Prints the device number to which the file or directory belongs

SORTING OPTIONS
       -v     Sort the output by version.

       -t     Sort the output by last modification time instead of alphabetically.

       -c     Sort the output by last status change instead of alphabetically.  Modifies the -D option (if used) to print the last status change instead  of  modification
              time.

       -U     Do not sort.  Lists files in directory order. Disables --dirsfirst.

       -r     Sort the output in reverse order.  This is a meta-sort that alter the above sorts.  This option is disabled when -U is used.

       --dirsfirst
              List directories before files. This is a meta-sort that alters the above sorts.  This option is disabled when -U is used.

       --filesfirst
              List files before directories. This is a meta-sort that alters the above sorts.  This option is disabled when -U is used.

       --sort[=]type
              Sort the output by type instead of name. Possible values are: ctime (-c), mtime (-t), size, or version (-v).

GRAPHICS OPTIONS
       -i     Makes  tree not print the indentation lines, useful when used in conjunction with the -f option.  Also removes as much whites
pace as possible when used with
              the -J or -X options.

       -A     Turn on ANSI line graphics hack when printing the indentation lines.

       -S     Turn on CP437 line graphics (useful when using Linux console mode fonts). This option is now equivalent to `--charset=IBM437'
 and may eventually be depreci‐
              ated.

       -n     Turn colorization off always, over-ridden by the -C option, however overrides CLICOLOR_FORCE if present.

       -C     Turn  colorization on always, using built-in color defaults if the LS_COLORS or TREE_COLORS environment variables are not set
.  Useful to colorize output to
              a pipe.

XML/JSON/HTML OPTIONS
       -X     Turn on XML output. Outputs the directory tree as an XML formatted file.

       -J     Turn on JSON output. Outputs the directory tree as a JSON formatted array.

       -H baseHREF
              Turn on HTML output, including HTTP references. Useful for ftp sites.  baseHREF gives the base ftp location when using HTML o
utput. That is, the  local  di‐
              rectory  may be `/local/ftp/pub', but it must be referenced as `ftp://hostname.organization.domain/pub' (baseHREF should be `
ftp://hostname.organization.do‐
              main'). Hint: don't use ANSI lines with this option, and don't give more than one directory in the directory list. If you wis
h to use colors via CSS  style-
              sheet, use the -C option in addition to this option to force color output.

       -T title
              Sets the title and H1 header string in HTML output mode.

       --nolinks
              Turns off hyperlinks in HTML output.

INPUT OPTIONS
       --fromfile Reads a directory listing from a file rather than the file-system.  Paths provided on the command line are files to read from rather than directories to
       search.  The dot (.) directory indicates that tree should read paths from standard input. NOTE: this is only suitable for reading the output of a program  such  as
       find, not 'tree -fi' as symlinks cannot (at least as yet) be distinguished from files that simply contain ' -> ' as part of the filename.

MISC OPTIONS
       --help Outputs a verbose usage listing.

       --version
              Outputs the version of tree.

       --     Option processing terminator.  No further options will be processed after this.

.INFO FILES
       .info files are similar to .gitignore files, if a .info file is found while scanning a directory it is read and added to a stack of .info information. Each file is
       composed of comments (lines starting with hash marks (#),) or wild-card patterns which may match a file relative to the directory the .info file is found in.  If a
       file  should  match  a pattern, the tab indented comment that follows the pattern is used as the file comment.  A comment is terminated by a non-tab indented line.
       Multiple patterns, each to a line, may share the same comment.

FILES
       /etc/DIR_COLORS          System color database.
       ~/.dircolors             Users color database.
       .gitignore               Git exclusion file
       $GIT_DIR/info/exclude    Global git file exclusion list
       .info                    File comment file
       /usr/share/finfo/global_info  Global file comment file

ENVIRONMENT
       LS_COLORS      Color information created by dircolors
       TREE_COLORS    Uses this for color information over LS_COLORS if it is set.
       TREE_CHARSET   Character set for tree to use in HTML mode.
       CLICOLOR       Enables colorization even if TREE_COLORS or LS_COLORS is not set.
       CLICOLOR_FORCE Always enables colorization (effectively -C)
       LC_CTYPE       Locale for filename output.
       LC_TIME        Locale for timefmt output, see strftime(3).
       TZ             Timezone for timefmt output, see strftime(3).
       STDDATA_FD     Enable the stddata feature, optionally set descriptor to use.

```
