# find
The `find` command in Unix and Linux is a powerful utility used to search for files and directories within a specified directory hierarchy. It allows you to locate files based on various criteria such as name, type, size, permissions, and timestamps.

### Basic Usage

The basic syntax for the `find` command is:

```sh
find [path...] [expression]
```

- **`path`**: Specifies the starting directory or directories to search. If not specified, `find` starts from the current directory.
- **`expression`**: Defines the search criteria and actions to be performed on matched files and directories.

### Examples

#### Finding Files by Name

To find files with a specific name:

```sh
find /path/to/search -name "filename.txt"
```

This command searches `/path/to/search` and its subdirectories for files named `filename.txt`.

#### Finding Directories

To find directories:

```sh
find /path/to/search -type d
```

This command lists all directories under `/path/to/search`.

#### Finding Files by Type

To find all regular files (excluding directories and symbolic links):

```sh
find /path/to/search -type f
```

This command searches for all regular files under `/path/to/search`.

#### Finding Files Based on Size

To find files larger than a specific size (e.g., 1MB):

```sh
find /path/to/search -type f -size +1M
```

This command searches for files larger than 1MB under `/path/to/search`.

#### Finding Files Modified Recently

To find files modified within the last 7 days:

```sh
find /path/to/search -type f -mtime -7
```

This command searches for files modified in the last 7 days under `/path/to/search`.

### Options and Expressions

#### `-name` Option: Match by Name

To match files by name pattern (supports wildcard characters):

```sh
find /path/to/search -name "*.log"
```

This command finds files ending in `.log` under `/path/to/search`.

#### `-type` Option: Match by Type

To specify the type of file or directory (`f` for regular files, `d` for directories):

```sh
find /path/to/search -type d
```

This command finds directories under `/path/to/search`.

#### `-size` Option: Match by Size

To specify file size in blocks (`c`), kilobytes (`k`), megabytes (`M`), gigabytes (`G`), or terabytes (`T`):

```sh
find /path/to/search -size +1M
```

This command finds files larger than 1MB under `/path/to/search`.

#### `-mtime` Option: Match by Modification Time

To specify files modified within a certain number of days (`+` for older, `-` for newer):

```sh
find /path/to/search -mtime -7
```

This command finds files modified in the last 7 days under `/path/to/search`.

### Practical Use Cases

#### Deleting Old Log Files

To delete log files older than 30 days:

```sh
find /var/log -name "*.log" -mtime +30 -exec rm {} \;
```

This command finds log files in `/var/log` older than 30 days and deletes them.

#### Finding and Copying Files

To find all `.txt` files and copy them to a new directory:

```sh
find /path/to/search -name "*.txt" -exec cp {} /path/to/destination \;
```

This command finds all `.txt` files under `/path/to/search` and copies them to `/path/to/destination`.

### Summary

The `find` command is a versatile tool for searching files and directories in Unix and Linux environments. Its ability to filter based on various criteria makes it invaluable for tasks such as file management, cleanup, and automation. Understanding its options and expressions allows you to perform precise searches and operations on files and directories efficiently.

# help

```
Usage: find [-H] [-L] [-P] [-Olevel] [-D debugopts] [path...] [expression]

Default path is the current directory; default expression is -print.
Expression may consist of: operators, options, tests, and actions.

Operators (decreasing precedence; -and is implicit where no others are given):
      ( EXPR )   ! EXPR   -not EXPR   EXPR1 -a EXPR2   EXPR1 -and EXPR2
      EXPR1 -o EXPR2   EXPR1 -or EXPR2   EXPR1 , EXPR2

Positional options (always true):
      -daystart -follow -nowarn -regextype -warn

Normal options (always true, specified before other expressions):
      -depth -files0-from FILE -maxdepth LEVELS -mindepth LEVELS
       -mount -noleaf -xdev -ignore_readdir_race -noignore_readdir_race

Tests (N can be +N or -N or N):
      -amin N -anewer FILE -atime N -cmin N -cnewer FILE -context CONTEXT
      -ctime N -empty -false -fstype TYPE -gid N -group NAME -ilname PATTERN
      -iname PATTERN -inum N -iwholename PATTERN -iregex PATTERN
      -links N -lname PATTERN -mmin N -mtime N -name PATTERN -newer FILE
      -nouser -nogroup -path PATTERN -perm [-/]MODE -regex PATTERN
      -readable -writable -executable
      -wholename PATTERN -size N[bcwkMG] -true -type [bcdpflsD] -uid N
      -used N -user NAME -xtype [bcdpfls]

Actions:
      -delete -print0 -printf FORMAT -fprintf FILE FORMAT -print 
      -fprint0 FILE -fprint FILE -ls -fls FILE -prune -quit
      -exec COMMAND ; -exec COMMAND {} + -ok COMMAND ;
      -execdir COMMAND ; -execdir COMMAND {} + -okdir COMMAND ;

Other common options:
      --help                   display this help and exit
      --version                output version information and exit

Valid arguments for -D:
exec, opt, rates, search, stat, time, tree, all, help
Use '-D help' for a description of the options, or see find(1)

```
man
```
NAME
       find - search for files in a directory hierarchy

SYNOPSIS
       find [-H] [-L] [-P] [-D debugopts] [-Olevel] [starting-point...] [expression]

DESCRIPTION
       This  manual page documents the GNU version of find.  GNU find searches the directory tree
       rooted at each given starting-point by evaluating the given expression from left to right,
       according  to  the rules of precedence (see section OPERATORS), until the outcome is known
       (the left hand side is false for and operations, true for or), at which point  find  moves
       on to the next file name.  If no starting-point is specified, `.' is assumed.

       If  you  are  using find in an environment where security is important (for example if you
       are using it to search directories that are writable by other users), you should read  the
       `Security  Considerations' chapter of the findutils documentation, which is called Finding
       Files and comes with findutils.  That document also includes a lot more detail and discus‐
       sion than this manual page, so you may find it a more useful source of information.

OPTIONS
       The -H, -L and -P options control the treatment of symbolic links.  Command-line arguments
       following these are taken to be names of files or directories to be examined,  up  to  the
       first  argument  that  begins with `-', or the argument `(' or `!'.  That argument and any
       following arguments are taken to be the expression describing what is to be searched  for.
       If  no paths are given, the current directory is used.  If no expression is given, the ex‐
       pression -print is used (but you should probably consider using -print0 instead, anyway).

       This manual page talks about `options' within the expression list.  These options  control
       the  behaviour  of  find but are specified immediately after the last path name.  The five
       `real' options -H, -L, -P, -D and -O must appear before the first path name, if at all.  A
       double  dash -- could theoretically be used to signal that any remaining arguments are not
       options, but this does not really work due to the way find determines the end of the  fol‐
       lowing  path  arguments: it does that by reading until an expression argument comes (which
       also starts with a `-').  Now, if a path argument would start with a `-', then find  would
       treat  it as expression argument instead.  Thus, to ensure that all start points are taken
       as such, and especially to prevent that wildcard patterns expanded by  the  calling  shell
       are  not mistakenly treated as expression arguments, it is generally safer to prefix wild‐
       cards or dubious path names with either `./' or to use absolute path names  starting  with
       '/'.   Alternatively,  it  is  generally  safe  though  non-portable to use the GNU option
       -files0-from to pass arbitrary starting points to find.

       -P     Never follow symbolic links.  This is the default behaviour.  When find examines or
              prints  information  about  files, and the file is a symbolic link, the information
              used shall be taken from the properties of the symbolic link itself.
       -L     Follow symbolic links.  When find examines or prints information about  files,  the
              information  used  shall be taken from the properties of the file to which the link
              points, not from the link itself (unless it is a broken symbolic link  or  find  is
              unable  to  examine the file to which the link points).  Use of this option implies
              -noleaf.  If you later use the -P option, -noleaf will still be in effect.   If  -L
              is  in  effect  and  find  discovers  a  symbolic link to a subdirectory during its
              search, the subdirectory pointed to by the symbolic link will be searched.

              When the -L option is in effect, the -type predicate will always match against  the
              type of the file that a symbolic link points to rather than the link itself (unless
              the symbolic link is broken).  Actions that can cause symbolic links to become bro‐
              ken while find is executing (for example -delete) can give rise to confusing behav‐
              iour.  Using -L causes the -lname and -ilname predicates always to return false.

       -H     Do not follow symbolic links, except while processing the command  line  arguments.
              When find examines or prints information about files, the information used shall be
              taken from the properties of the symbolic link itself.  The only exception to  this
              behaviour  is when a file specified on the command line is a symbolic link, and the
              link can be resolved.  For that situation, the information used is taken from what‐
              ever the link points to (that is, the link is followed).  The information about the
              link itself is used as a fallback if the file pointed to by the symbolic link  can‐
              not  be examined.  If -H is in effect and one of the paths specified on the command
              line is a symbolic link to a directory, the contents of that directory will be  ex‐
              amined (though of course -maxdepth 0 would prevent this).
       If  more  than  one of -H, -L and -P is specified, each overrides the others; the last one
       appearing on the command line takes effect.  Since it is the default, the -P option should
       be considered to be in effect unless either -H or -L is specified.

       GNU  find  frequently stats files during the processing of the command line itself, before
       any searching has begun.  These options also affect how  those  arguments  are  processed.
       Specifically,  there  are  a number of tests that compare files listed on the command line
       against a file we are currently considering.  In each case, the file specified on the com‐
       mand line will have been examined and some of its properties will have been saved.  If the
       named file is in fact a symbolic link, and the -P option is in effect (or  if  neither  -H
       nor  -L  were  specified),  the information used for the comparison will be taken from the
       properties of the symbolic link.  Otherwise, it will be taken from the properties  of  the
       file  the  link points to.  If find cannot follow the link (for example because it has in‐
       sufficient privileges or the link points to a nonexistent file) the properties of the link
       itself will be used.

       When  the  -H  or  -L  options are in effect, any symbolic links listed as the argument of
       -newer will be dereferenced, and the timestamp will be taken from the file  to  which  the
       symbolic link points.  The same consideration applies to -newerXY, -anewer and -cnewer.

       The  -follow  option has a similar effect to -L, though it takes effect at the point where
       it appears (that is, if -L is not used but -follow is, any symbolic links appearing  after
       -follow on the command line will be dereferenced, and those before it will not).

       -D debugopts
              Print  diagnostic  information;  this  can be helpful to diagnose problems with why
              find is not doing what you want.  The list of debug options should be  comma  sepa‐
              rated.   Compatibility  of  the debug options is not guaranteed between releases of
              findutils.  For a complete list of valid debug options,  see  the  output  of  find
              -D help.  Valid debug options include

              exec   Show diagnostic information relating to -exec, -execdir, -ok and -okdir

              opt    Prints diagnostic information relating to the optimisation of the expression
                     tree; see the -O option.

              rates  Prints a summary indicating how often each predicate succeeded or failed.

              search Navigate the directory tree verbosely.

              stat   Print messages as files are examined with the stat and lstat  system  calls.
                     The find program tries to minimise such calls.

              tree   Show the expression tree in its original and optimised form.

              all    Enable all of the other debug options (but help).

              help   Explain the debugging options.
      -Olevel
              Enables  query optimisation.  The find program reorders tests to speed up execution
              while preserving the overall effect; that is, predicates with side effects are  not
              reordered relative to each other.  The optimisations performed at each optimisation
              level are as follows.

              0      Equivalent to optimisation level 1.

              1      This is the default optimisation level and corresponds  to  the  traditional
                     behaviour.   Expressions are reordered so that tests based only on the names
                     of files (for example -name and -regex) are performed first.

              2      Any -type or -xtype tests are performed after any tests based  only  on  the
                     names  of  files, but before any tests that require information from the in‐
                     ode.  On many modern versions of Unix, file types are returned by  readdir()
                     and so these predicates are faster to evaluate than predicates which need to
                     stat the file first.  If you use the -fstype FOO  predicate  and  specify  a
                     filesystem  type FOO which is not known (that is, present in `/etc/mtab') at
                     the time find starts, that predicate is equivalent to -false.

              3      At this optimisation level, the full cost-based query optimiser is  enabled.
                     The order of tests is modified so that cheap (i.e. fast) tests are performed
                     first and more expensive ones are performed  later,  if  necessary.   Within
                     each  cost  band,  predicates  are  evaluated  earlier or later according to
                     whether they are likely to succeed or not.  For  -o,  predicates  which  are
                     likely  to  succeed  are evaluated earlier, and for -a, predicates which are
                     likely to fail are evaluated earlier.

              The cost-based optimiser has a fixed idea of how likely any given test is  to  suc‐
              ceed.   In  some  cases the probability takes account of the specific nature of the
              test (for example, -type f is assumed to be more likely to succeed  than  -type c).
              The cost-based optimiser is currently being evaluated.  If it does not actually im‐
              prove the performance of find, it will be removed again.  Conversely, optimisations
              that  prove  to be reliable, robust and effective may be enabled at lower optimisa‐
              tion levels over time.  However, the default behaviour (i.e. optimisation level  1)
              will not be changed in the 4.3.x release series.  The findutils test suite runs all
              the tests on find at each optimisation level and ensures that  the  result  is  the
              same.

EXPRESSION
       The part of the command line after the list of starting points is the expression.  This is
       a kind of query specification describing how we match files and what we do with the  files
       that were matched.  An expression is composed of a sequence of things:

       Tests  Tests return a true or false value, usually on the basis of some property of a file
              we are considering.  The -empty test for example is true only when the current file
              is empty.

       Actions
              Actions  have  side effects (such as printing something on the standard output) and
              return either true or false, usually based on whether or not they  are  successful.
              The  -print  action for example prints the name of the current file on the standard
              output.

       Global options
              Global options affect the operation of tests and actions specified on any  part  of
              the  command line.  Global options always return true.  The -depth option for exam‐
              ple makes find traverse the file system in a depth-first order.

       Positional options
              Positional options affect only tests or actions which follow them.  Positional  op‐
              tions  always return true.  The -regextype option for example is positional, speci‐
              fying the regular expression dialect for regular expressions occurring later on the
              command line.

       Operators
              Operators  join  together  the other items within the expression.  They include for
              example -o (meaning logical OR) and -a (meaning logical AND).  Where an operator is
              missing, -a is assumed.

       The -print action is performed on all files for which the whole expression is true, unless
       it contains an action other than -prune or  -quit.   Actions  which  inhibit  the  default
       -print are -delete, -exec, -execdir, -ok, -okdir, -fls, -fprint, -fprintf, -ls, -print and
       -printf.

       The -delete action also acts like an option (since it implies -depth).

   POSITIONAL OPTIONS
       Positional options always return true.  They affect only tests occurring later on the com‐
       mand line.

       -daystart
              Measure times (for -amin, -atime, -cmin, -ctime, -mmin, and -mtime) from the begin‐
              ning of today rather than from 24 hours ago.  This option only affects tests  which
              appear later on the command line.

   POSITIONAL OPTIONS
       Positional options always return true.  They affect only tests occurring later on the com‐
       mand line.

       -daystart
              Measure times (for -amin, -atime, -cmin, -ctime, -mmin, and -mtime) from the begin‐
              ning of today rather than from 24 hours ago.  This option only affects tests  which
              appear later on the command line.

       -follow
              Deprecated;  use  the  -L  option  instead.   Dereference  symbolic links.  Implies
              -noleaf.  The -follow option affects only those tests which appear after it on  the
              command  line.   Unless the -H or -L option has been specified, the position of the
              -follow option changes the behaviour of the -newer predicate; any files  listed  as
              the  argument  of -newer will be dereferenced if they are symbolic links.  The same
              consideration applies to -newerXY, -anewer and -cnewer.  Similarly, the -type pred‐
              icate will always match against the type of the file that a symbolic link points to
              rather than the link itself.  Using -follow causes the -lname  and  -ilname  predi‐
              cates always to return false.

       -regextype type
              Changes  the regular expression syntax understood by -regex and -iregex tests which
              occur later on the command line.  To see which regular expression types are  known,
              use -regextype help.  The Texinfo documentation (see SEE ALSO) explains the meaning
              of and differences between the various types of regular expression.

       -warn, -nowarn
              Turn warning messages on or off.  These warnings apply only to the command line us‐
              age,  not to any conditions that find might encounter when it searches directories.
              The default behaviour corresponds to -warn if standard  input  is  a  tty,  and  to
              -nowarn  otherwise.   If  a  warning message relating to command-line usage is pro‐
              duced, the exit status of find is not affected.  If the POSIXLY_CORRECT environment
              variable  is  set, and -warn is also used, it is not specified which, if any, warn‐
              ings will be active.

   GLOBAL OPTIONS
       Global options always return true.  Global options take effect even for tests which  occur
       earlier on the command line.  To prevent confusion, global options should specified on the
       command-line after the list of start points, just before the first test, positional option
       or  action.  If you specify a global option in some other place, find will issue a warning
       message explaining that this can be confusing.

       The global options occur after the list of start points, and so are not the same  kind  of
       option as -L, for example.
       -d     A synonym for -depth, for compatibility with FreeBSD, NetBSD, MacOS X and OpenBSD.

       -depth Process  each directory's contents before the directory itself.  The -delete action
              also implies -depth.

       -files0-from file
              Read the starting points from file instead of getting them on the command line.  In
              contrast  to  the known limitations of passing starting points via arguments on the
              command line, namely the limitation of the amount of file names, and  the  inherent
              ambiguity  of  file  names  clashing with option names, using this option allows to
              safely pass an arbitrary number of starting points to find.

              Using this option and passing starting points on the command line is  mutually  ex‐
              clusive, and is therefore not allowed at the same time.

              The  file  argument  is  mandatory.  One can use -files0-from - to read the list of
              starting points from the standard input stream, and e.g.  from  a  pipe.   In  this
              case,  the actions -ok and -okdir are not allowed, because they would obviously in‐
              terfere with reading from standard input in order to get a user confirmation.

              The starting points in file have to be separated by ASCII NUL characters.  Two con‐
              secutive NUL characters, i.e., a starting point with a Zero-length file name is not
              allowed and will lead to an error diagnostic  followed  by  a  non-Zero  exit  code
              later.

              In  the  case the given file is empty, find does not process any starting point and
              therefore will exit immediately after parsing the program arguments.  This  is  un‐
              like  the  standard invocation where find assumes the current directory as starting
              point if no path argument is passed.

              The processing of the starting points is otherwise as usual, e.g.   find  will  re‐
              curse into subdirectories unless otherwise prevented.  To process only the starting
              points, one can additionally pass -maxdepth 0.

              Further notes: if a file is listed more than once in the input file, it is unspeci‐
              fied whether it is visited more than once.  If the file is mutated during the oper‐
              ation of find, the result is unspecified  as  well.   Finally,  the  seek  position
              within the named file at the time find exits, be it with -quit or in any other way,
              is also unspecified.  By "unspecified" here is meant that it may or may not work or
              do  any specific thing, and that the behavior may change from platform to platform,
              or from findutils release to release.

       -help, --help
             Print a summary of the command-line usage of find and exit.

       -ignore_readdir_race
              Normally, find will emit an error message when it fails to stat  a  file.   If  you
              give  this option and a file is deleted between the time find reads the name of the
              file from the directory and the time it tries to stat the file,  no  error  message
              will be issued.  This also applies to files or directories whose names are given on
              the command line.  This option takes effect at the time the command line  is  read,
              which  means  that you cannot search one part of the filesystem with this option on
              and part of it with this option off (if you need to do that, you will need to issue
              two find commands instead, one with the option and one without it).

              Furthermore,  find  with  the -ignore_readdir_race option will ignore errors of the
              -delete action in the case the file has disappeared since the parent directory  was
              read:  it  will  not output an error diagnostic, and the return code of the -delete
              action will be true.

       -maxdepth levels
              Descend at most levels (a non-negative integer) levels  of  directories  below  the
              starting-points.   Using  -maxdepth 0 means only apply the tests and actions to the
              starting-points themselves.

       -mindepth levels
              Do not apply any tests or actions at levels less than levels (a non-negative  inte‐
              ger).  Using -mindepth 1 means process all files except the starting-points.

       -mount Don't  descend  directories on other filesystems.  An alternate name for -xdev, for
              compatibility with some other versions of find.

       -noignore_readdir_race
              Turns off the effect of -ignore_readdir_race.

       -noleaf
              Do not optimize by assuming that directories contain 2  fewer  subdirectories  than
              their  hard  link  count.  This option is needed when searching filesystems that do
              not follow the Unix directory-link convention, such as CD-ROM or MS-DOS filesystems
              or  AFS  volume  mount  points.   Each directory on a normal Unix filesystem has at
              least 2 hard links: its name and its `.' entry.  Additionally,  its  subdirectories
              (if any) each have a `..' entry linked to that directory.  When find is examining a
              directory, after it has statted 2 fewer subdirectories than  the  directory's  link
              count,  it  knows that the rest of the entries in the directory are non-directories
              (`leaf' files in the directory tree).  If only the files' names need  to  be  exam‐
              ined,  there  is  no need to stat them; this gives a significant increase in search
              speed.

       -version, --version
              Print the find version number and exit.

       -xdev  Don't descend directories on other filesystems.
  TESTS
       Some tests, for example -newerXY and -samefile, allow comparison  between  the  file  cur‐
       rently  being  examined and some reference file specified on the command line.  When these
       tests are used, the interpretation of the reference file is determined by the options  -H,
       -L  and  -P and any previous -follow, but the reference file is only examined once, at the
       time the command line is parsed.  If the reference file cannot be examined  (for  example,
       the  stat(2)  system call fails for it), an error message is issued, and find exits with a
       nonzero status.

       A numeric argument n can be specified to tests (like -amin, -mtime, -gid,  -inum,  -links,
       -size, -uid and -used) as

       +n     for greater than n,

       -n     for less than n,

       n      for exactly n.

       Supported tests:

       -amin n
              File was last accessed less than, more than or exactly n minutes ago.

       -anewer reference
              Time  of  the  last access of the current file is more recent than that of the last
              data modification of the reference file.  If reference is a symbolic link  and  the
              -H  option  or the -L option is in effect, then the time of the last data modifica‐
              tion of the file it points to is always used.

       -atime n
              File was last accessed less than, more than or exactly n*24 hours ago.   When  find
              figures out how many 24-hour periods ago the file was last accessed, any fractional
              part is ignored, so to match -atime +1, a file has to have been accessed  at  least
              two days ago.

       -cmin n
              File's status was last changed less than, more than or exactly n minutes ago.

       -cnewer reference
              Time  of the last status change of the current file is more recent than that of the
              last data modification of the reference file.  If reference is a symbolic link  and
              the -H option or the -L option is in effect, then the time of the last data modifi‐
              cation of the file it points to is always used.

       -ctime n
              File's status was last changed less than, more than or exactly n*24 hours ago.  See
              the  comments  for  -atime to understand how rounding affects the interpretation of
              file status change times.

       -empty File is empty and is either a regular file or a directory.

       -executable
              Matches files which are executable and directories which are searchable (in a  file
              name resolution sense) by the current user.  This takes into account access control
              lists and other permissions artefacts which the  -perm  test  ignores.   This  test
              makes  use  of the access(2) system call, and so can be fooled by NFS servers which
              do UID mapping (or root-squashing), since many systems implement access(2)  in  the
              client's  kernel  and so cannot make use of the UID mapping information held on the
              server.  Because this test is based only on the  result  of  the  access(2)  system
              call,  there  is no guarantee that a file for which this test succeeds can actually
              be executed.

       -false Always false.

       -fstype type
              File is on a filesystem of type type.  The valid filesystem types vary  among  dif‐
              ferent  versions  of Unix; an incomplete list of filesystem types that are accepted
              on some version of Unix or another is: ufs, 4.2, 4.3, nfs, tmp,  mfs,  S51K,  S52K.
              You can use -printf with the %F directive to see the types of your filesystems.
       -gid n File's numeric group ID is less than, more than or exactly n.

       -group gname
              File belongs to group gname (numeric group ID allowed).

       -ilname pattern
              Like  -lname,  but  the match is case insensitive.  If the -L option or the -follow
              option is in effect, this test returns false unless the symbolic link is broken.

       -iname pattern
              Like -name, but the match is case insensitive.  For example, the patterns `fo*' and
              `F??'  match  the  file names `Foo', `FOO', `foo', `fOo', etc.  The pattern `*foo*`
              will also match a file called '.foobar'.

       -inum n
              File has inode number smaller than, greater than or exactly n.  It is normally eas‐
              ier to use the -samefile test instead.

       -ipath pattern
              Like -path.  but the match is case insensitive.

       -iregex pattern
              Like -regex, but the match is case insensitive.

       -iwholename pattern
              See -ipath.  This alternative is less portable than -ipath.

       -links n
              File has less than, more than or exactly n hard links.

       -lname pattern
              File  is a symbolic link whose contents match shell pattern pattern.  The metachar‐
              acters do not treat `/' or `.' specially.  If the -L option or the  -follow  option
              is in effect, this test returns false unless the symbolic link is broken.

       -mmin n
              File's data was last modified less than, more than or exactly n minutes ago.

       -mtime n
              File's  data was last modified less than, more than or exactly n*24 hours ago.  See
              the comments for -atime to understand how rounding affects  the  interpretation  of
              file modification times.

       -name pattern
              Base  of  file  name  (the path with the leading directories removed) matches shell
              pattern pattern.  Because the leading directories are removed, the file names  con‐
              sidered  for  a  match  with  -name will never include a slash, so `-name a/b' will
              never match anything (you probably need to use -path instead).  A warning is issued
              if you try to do this, unless the environment variable POSIXLY_CORRECT is set.  The
              metacharacters (`*', `?', and `[]') match a `.' at the start of the base name (this
              is  a  change in findutils-4.2.2; see section STANDARDS CONFORMANCE below).  To ig‐
              nore a directory and the files under it, use -prune rather than checking every file
              in  the  tree;  see  an  example in the description of that action.  Braces are not
              recognised as being special, despite the fact that some shells including Bash imbue
              braces  with  a  special  meaning in shell patterns.  The filename matching is per‐
              formed with the use of the fnmatch(3) library function.  Don't  forget  to  enclose
              the pattern in quotes in order to protect it from expansion by the shell.

       -newer reference
              Time  of the last data modification of the current file is more recent than that of
              the last data modification of the reference file.  If reference is a symbolic  link
              and  the  -H  option  or the -L option is in effect, then the time of the last data
              modification of the file it points to is always used.

       -newerXY reference
              Succeeds if timestamp X of the file being considered is newer than timestamp  Y  of
              the file reference.  The letters X and Y can be any of the following letters:

              a   The access time of the file reference
              B   The birth time of the file reference
              c   The inode status change time of reference
              m   The modification time of the file reference
              t   reference is interpreted directly as a time

              Some combinations are invalid; for example, it is invalid for X to be t.  Some com‐
              binations are not implemented on all systems; for example B is not supported on all
              systems.   If an invalid or unsupported combination of XY is specified, a fatal er‐
              ror results.  Time specifications are interpreted as for the argument to the -d op‐
              tion  of  GNU  date.  If you try to use the birth time of a reference file, and the
              birth time cannot be determined, a fatal error message results.  If you  specify  a
              test  which  refers  to the birth time of files being examined, this test will fail
              for any files where the birth time is unknown.

       -nogroup
              No group corresponds to file's numeric group ID.

       -nouser
              No user corresponds to file's numeric user ID.

.......
.......
.......

```




 

