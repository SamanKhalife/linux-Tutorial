# info

The info command is a Linux command that can be used to display information about a variety of system topics. It is a powerful tool that can be used to troubleshoot problems, or to learn more about how the system works.

The info command will display a hierarchical tree of information about the topic. You can use the arrow keys to navigate through the tree, and the q key to exit

# help

```
Usage: info [OPTION]... [MENU-ITEM...]

Read documentation in Info format.

Frequently-used options:
  -a, --all                    use all matching manuals
  -k, --apropos=STRING         look up STRING in all indices of all manuals
  -d, --directory=DIR          add DIR to INFOPATH
  -f, --file=MANUAL            specify Info manual to visit
  -h, --help                   display this help and exit
      --index-search=STRING    go to node pointed by index entry STRING
  -n, --node=NODENAME          specify nodes in first visited Info file
  -o, --output=FILE            output selected nodes to FILE
  -O, --show-options, --usage  go to command-line options node
      --subnodes               recursively output menu items
  -v, --variable VAR=VALUE     assign VALUE to Info variable VAR
      --version                display version information and exit
  -w, --where, --location      print physical location of Info file

The first non-option argument, if present, is the menu entry to start from;
it is searched for in all 'dir' files along INFOPATH.
If it is not present, info merges all 'dir' files and shows the result.
Any remaining arguments are treated as the names of menu
items relative to the initial node visited.

For a summary of key bindings, type H within Info.

Examples:
  info                         show top-level dir menu
  info info-stnd               show the manual for this Info program
  info emacs                   start at emacs node from top-level dir
  info emacs buffers           select buffers menu entry in emacs manual
  info emacs -n Files          start at Files node within emacs manual
  info '(emacs)Files'          alternative way to start at Files node
  info --show-options emacs    start at node with emacs' command line options
  info --subnodes -o out.txt emacs
                               dump entire emacs manual to out.txt
  info -f ./foo.info           show file ./foo.info, not searching dir
```


man

```
NAME
       info - read Info documents

SYNOPSIS
       info [OPTION]... [MENU-ITEM...]

DESCRIPTION
       Read documentation in Info format.

   Frequently-used options:
       -a, --all
              use all matching manuals

       -k, --apropos=STRING
              look up STRING in all indices of all manuals

       -d, --directory=DIR
              add DIR to INFOPATH

       -f, --file=MANUAL
              specify Info manual to visit

       -h, --help
              display this help and exit

       --index-search=STRING
              go to node pointed by index entry STRING

       -n, --node=NODENAME
              specify nodes in first visited Info file

       -o, --output=FILE
              output selected nodes to FILE

       -O, --show-options, --usage
              go to command-line options node

       --subnodes
              recursively output menu items

       -v, --variable VAR=VALUE
              assign VALUE to Info variable VAR

       --version
              display version information and exit

       -w, --where, --location
              print physical location of Info file

       The first non-option argument, if present, is the menu entry to start from; it is searched
       for in all 'dir' files along INFOPATH.  If it is not present, info merges all 'dir'  files
       and shows the result.  Any remaining arguments are treated as the names of menu items rel‐
       ative to the initial node visited.

       For a summary of key bindings, type H within Info.
EXAMPLES
       info   show top-level dir menu

       info info-stnd
              show the manual for this Info program

       info emacs
              start at emacs node from top-level dir

       info emacs buffers
              select buffers menu entry in emacs manual

       info emacs -n Files
              start at Files node within emacs manual

       info '(emacs)Files'
              alternative way to start at Files node

       info --show-options emacs
              start at node with emacs' command line options

       info --subnodes -o out.txt emacs
              dump entire emacs manual to out.txt

       info -f ./foo.info
              show file ./foo.info, not searching dir
```
