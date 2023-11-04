# whatis

The `whatis` command is a Linux command that can be used to get a brief description of a command. It is a versatile command that can be used to quickly get a summary of a command's purpose and usage.

The `whatis` command is used as follows:

```
whatis command
```

* `command`: This is the name of the command that you want to get a description of.

For example, the following command will get a brief description of the `wc` command:

```
whatis wc
```

The output of the `whatis` command will be:

```
wc (1) - count lines, words, and characters
```

This output shows that the `wc` command is used to count the lines, words, and characters in a file.

The `whatis` command can be used to get a brief description of any command. It is a valuable tool for system administrators who need to learn about the commands that are available on their system. It can also be used by users to get a quick summary of a command's purpose and usage.

Here are some other examples of how the `whatis` command can be used:

* To get a description of the `ls` command:

```
whatis ls
```

* To get a description of the `grep` command:

```
whatis grep
```

* To get a description of the `man` command:

```
whatis man
```

The `whatis` command is a quick and easy way to get a brief description of a command. It is a valuable tool for anyone who needs to learn about the commands that are available on their system.

# help

```
Usage: whatis [OPTION...] KEYWORD...

  -d, --debug                emit debugging messages
  -v, --verbose              print verbose warning messages
  -r, --regex                interpret each keyword as a regex
  -w, --wildcard             the keyword(s) contain wildcards
  -l, --long                 do not trim output to terminal width
  -C, --config-file=FILE     use this user configuration file
  -L, --locale=LOCALE        define the locale for this search
  -m, --systems=SYSTEM       use manual pages from other systems
  -M, --manpath=PATH         set search path for manual pages to PATH
  -s, --sections=LIST, --section=LIST
                             search only these sections (colon-separated)
  -?, --help                 give this help list
      --usage                give a short usage message
  -V, --version              print program version

Mandatory or optional arguments to long options are also mandatory or optional
for any corresponding short options.
```

man
```
NAME
       whatis - display one-line manual page descriptions

SYNOPSIS
       whatis [-dlv?V] [-r|-w] [-s list] [-m system[,...]] [-M path] [-L locale] [-C file] name ...

DESCRIPTION
       Each  manual page has a short description available within it.  whatis searches the manual page names and displays the manual
       page descriptions of any name matched.

       name may contain wildcards (-w) or be a regular expression (-r).  Using these options, it may be necessary to quote the  name
       or escape (\) the special characters to stop the shell from interpreting them.

       index  databases  are used during the search, and are updated by the mandb program.  Depending on your installation, this may
       be run by a periodic cron job, or may need to be run manually after new manual pages have been installed.  To produce an  old
       style text whatis database from the relative index database, issue the command:

       whatis -M manpath -w '*' | sort > manpath/whatis

       where manpath is a manual page hierarchy such as /usr/man.

OPTIONS
       -d, --debug
              Print debugging information.

       -v, --verbose
              Print verbose warning messages.

       -r, --regex
              Interpret  each  name as a regular expression.  If a name matches any part of a page name, a match will be made.  This
              option causes whatis to be somewhat slower due to the nature of database searches.

       -w, --wildcard
              Interpret each name as a pattern containing shell style wildcards.  For a match to be  made,  an  expanded  name  must
              match the entire page name.  This option causes whatis to be somewhat slower due to the nature of database searches.

       -l, --long
              Do  not trim output to the terminal width.  Normally, output will be truncated to the terminal width to avoid ugly re‐
              sults from poorly-written NAME sections.

       -s list, --sections=list, --section=list
              Search only the given manual sections.  list is a colon- or comma-separated list of sections.  If an entry in list  is
              a  simple  section,  for  example  "3",  then  the  displayed list of descriptions will include pages in sections "3",
            "3perl", "3x", and so on; while if an entry in list has an extension, for example "3perl", then the list will only in‐
              clude pages in that exact part of the manual section.

       -m system[,...], --systems=system[,...]
              If  this  system has access to other operating systems' manual page names, they can be accessed using this option.  To
              search NewOS's manual page names, use the option -m NewOS.

              The system specified can be a combination of comma delimited operating system names.  To include a search of  the  na‐
              tive operating system's manual page names, include the system name man in the argument string.  This option will over‐
              ride the $SYSTEM environment variable.

       -M path, --manpath=path
              Specify an alternate set of colon-delimited manual page hierarchies to search.  By default, whatis uses  the  $MANPATH
              environment  variable,  unless  it  is empty or unset, in which case it will determine an appropriate manpath based on
              your $PATH environment variable.  This option overrides the contents of $MANPATH.

       -L locale, --locale=locale
              whatis will normally determine your current locale by a call to the C function setlocale(3) which interrogates various
              environment  variables,  possibly including $LC_MESSAGES and $LANG.  To temporarily override the determined value, use
              this option to supply a locale string directly to whatis.  Note that it will not take  effect  until  the  search  for
              pages actually begins.  Output such as the help message will always be displayed in the initially determined locale.
       -C file, --config-file=file
              Use this user configuration file rather than the default of ~/.manpath.

       -?, --help
              Print a help message and exit.

       --usage
              Print a short usage message and exit.

       -V, --version
              Display version information.

EXIT STATUS
       0      Successful program execution.

       1      Usage, syntax or configuration file error.

       2      Operational error.

       16     Nothing was found that matched the criteria specified.

ENVIRONMENT
       SYSTEM If $SYSTEM is set, it will have the same effect as if it had been specified as the argument to the -m option.

       MANPATH
              If $MANPATH is set, its value is interpreted as the colon-delimited manual page hierarchy search path to use.

              See  the  SEARCH  PATH section of manpath(5) for the default behaviour and details of how this environment variable is
              handled.

       MANWIDTH
              If $MANWIDTH is set, its value is used as the terminal width (see the --long option).  If it is not set, the  terminal
              width  will  be calculated using the value of $COLUMNS, and ioctl(2) if available, or falling back to 80 characters if
              all else fails.

..............................
```


## breakdown

```

```
