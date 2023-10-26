# fc

The `fc` command in Linux is used to edit and execute previous commands. It is a powerful tool that can be used to repeat commands, to correct mistakes, and to debug scripts.

The `fc` command is used in the following syntax:

```
fc [options] [command]
```

The `command` is the command that you want to edit and execute.

The options can be used to specify the following:

* `-e` : Edit the command before executing it.
* `-l` : List the previous commands.
* `-r` : Repeat the previous command.
* `-s` : Substitute text in the command.

For example, the following code will edit the previous command and then execute it:

```
fc -e
```

This code will open the previous command in a text editor, so that you can edit it. Once you have edited the command, press `Ctrl`+`X` to exit the text editor and the command will be executed.

The `fc` command is a powerful tool that can be used to edit and execute previous commands. It is a valuable tool to know, especially if you frequently use the command line.

Here are some additional things to note about the `fc` command:

* The `fc` command can be used to edit any command that you have previously entered.
* The `fc` command can be used to correct mistakes in commands.
* The `fc` command can be used to debug scripts.
* The `fc` command should be used with caution, as it can change the meaning of commands if used incorrectly.
# help

```
fc: fc [-e ename] [-lnr] [first] [last] or fc -s [pat=rep] [command]
    Display or execute commands from the history list.
    
    fc is used to list or edit and re-execute commands from the history list.
    FIRST and LAST can be numbers specifying the range, or FIRST can be a
    string, which means the most recent command beginning with that
    string.
    
    Options:
      -e ENAME  select which editor to use.  Default is FCEDIT, then EDITOR,
                then vi
      -l        list lines instead of editing
      -n        omit line numbers when listing
      -r        reverse the order of the lines (newest listed first)
    
    With the `fc -s [pat=rep ...] [command]' format, COMMAND is
    re-executed after the substitution OLD=NEW is performed.
    
    A useful alias to use with this is r='fc -s', so that typing `r cc'
    runs the last command beginning with `cc' and typing `r' re-executes
    the last command.
    
    Exit Status:
    Returns success or status of executed command; non-zero if an error occurs.
```
