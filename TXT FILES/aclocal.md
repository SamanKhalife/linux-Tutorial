help
```
Usage: aclocal [OPTION]...

Generate 'aclocal.m4' by scanning 'configure.ac' or 'configure.in'

Options:
      --automake-acdir=DIR  directory holding automake-provided m4 files
      --system-acdir=DIR    directory holding third-party system-wide files
      --diff[=COMMAND]      run COMMAND [diff -u] on M4 files that would be
                            changed (implies --install and --dry-run)
      --dry-run             pretend to, but do not actually update any file
      --force               always update output file
      --help                print this help, then exit
  -I DIR                    add directory to search list for .m4 files
      --install             copy third-party files to the first -I directory
      --output=FILE         put output in FILE (default aclocal.m4)
      --print-ac-dir        print name of directory holding system-wide
                              third-party m4 files, then exit
      --verbose             don't be silent
      --version             print version number, then exit
  -W, --warnings=CATEGORY   report the warnings falling in CATEGORY
Warning categories include:
  cross                  cross compilation issues
  gnu                    GNU coding standards (default in gnu and gnits modes)
  obsolete               obsolete features or constructions (default)
  override               user redefinitions of Automake rules or variables
  portability            portability issues (default in gnu and gnits modes)
  portability-recursive  nested Make variables (default with -Wportability)
  extra-portability      extra portability issues related to obscure tools
  syntax                 dubious syntactic constructs (default)
  unsupported            unsupported or incomplete features (default)
  all                    all the warnings
  no-CATEGORY            turn off warnings in CATEGORY
  none                   turn off all the warnings
  error                  treat warnings as errors
Report bugs to <bug-automake@gnu.org>.
GNU Automake home page: <https://www.gnu.org/software/automake/>.
General help using GNU software: <https://www.gnu.org/gethelp/>.
```