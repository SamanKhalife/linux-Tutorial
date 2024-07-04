# history

The `history` command in Unix and Linux is used to display the list of previously executed commands in the terminal session. This is particularly useful for recalling commands without having to retype them, and for searching through your command history to find specific commands you ran previously.

### Basic Usage

The syntax for the `history` command is:

```sh
history [options] [n]
```

- **`n`**: Display the last `n` commands.

When used without any options or arguments, `history` will display all commands in the history list.

### Examples

#### Displaying Command History

To display the complete command history:

```sh
history
```

This will output a numbered list of commands you have executed in the current terminal session:

```
1  ls
2  cd /var
3  cat /etc/passwd
4  history
```

#### Displaying the Last N Commands

To display the last 5 commands, for example:

```sh
history 5
```

This will output the last 5 commands:

```
10  ls -l
11  cd /tmp
12  mkdir test
13  cd test
14  history 5
```

### Using History for Command Recall

#### Re-executing Commands

You can re-execute a command from the history list using the exclamation mark (`!`) followed by the command number. For example, to re-execute command number 12:

```sh
!12
```

This will re-run `mkdir test`.

#### Repeating the Last Command

To repeat the last command you executed:

```sh
!!
```

This will re-run the most recent command.

#### Using History Search

You can search through your command history using `Ctrl` + `r`. This will start a reverse incremental search. Start typing the command you are looking for, and it will appear:

```sh
(reverse-i-search)`cat': cat /etc/passwd
```

Press `Enter` to execute the command.

### History Substitution

#### Running a Command with Modifications

You can re-run a previous command with modifications. For example, if you previously ran:

```sh
ls -l /var
```

You can re-run it with a different directory:

```sh
^/var^/tmp
```

This will execute:

```sh
ls -l /tmp
```

### Options

#### Clearing the History

To clear the history list:

```sh
history -c
```

This will clear the history of the current session.

#### Appending to the History File

By default, the history command saves the history of each session in a file specified by the `HISTFILE` environment variable (usually `~/.bash_history`). To append the current session’s history to this file:

```sh
history -a
```

### Configuration

You can customize the behavior of the command history using environment variables:

- **`HISTSIZE`**: The number of commands to remember in the current session.
- **`HISTFILESIZE`**: The maximum number of lines contained in the history file.
- **`HISTIGNORE`**: A colon-separated list of patterns used to decide which command lines should be saved in the history list.

For example, to set the history size to 500 commands and to ignore commands starting with `ls`:

```sh
export HISTSIZE=500
export HISTIGNORE="ls*"
```

### Practical Use Cases

#### Auditing Commands

The history command is useful for auditing the commands that were run on a system. By checking the history, you can see what actions were performed in previous sessions.

#### Scripting and Automation

When creating scripts, it is helpful to recall and reuse complex commands you’ve run before. The history command makes it easy to find and repeat those commands.

### Conclusion

The `history` command is a powerful tool for managing and recalling the commands you’ve executed in Unix and Linux systems. With its various options and the ability to re-execute, modify, and search through past commands, it greatly enhances the efficiency of working in the terminal. By understanding and utilizing the history command, you can save time and avoid repetitive typing. If you have any further questions or need assistance with anything else, feel free to ask!a

# help
```
history: history [-c] [-d offset] [n] or history -anrw [filename] or history -ps arg [arg...]
    Display or manipulate the history list.
    
    Display the history list with line numbers, prefixing each modified
    entry with a `*'.  An argument of N lists only the last N entries.
    
    Options:
      -c        clear the history list by deleting all of the entries
      -d offset delete the history entry at position OFFSET. Negative
                offsets count back from the end of the history list
    
      -a        append history lines from this session to the history file
      -n        read all history lines not already read from the history file
                and append them to the history list
      -r        read the history file and append the contents to the history
                list
      -w        write the current history to the history file
    
      -p        perform history expansion on each ARG and display the result
                without storing it in the history list
      -s        append the ARGs to the history list as a single entry
    
    If FILENAME is given, it is used as the history file.  Otherwise,
    if HISTFILE has a value, that is used, else ~/.bash_history.
    
    If the HISTTIMEFORMAT variable is set and not null, its value is used
    as a format string for strftime(3) to print the time stamp associated
    with each displayed history entry.  No time stamps are printed otherwise.
    
    Exit Status:
    Returns success unless an invalid option is given or an error occurs.
```
