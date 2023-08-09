In Linux, an alias is a command that is defined to be another command. This can be useful for shortening long commands or for creating shortcuts to commands that you use frequently.


You can also create aliases for commands that take arguments. For example, the following command creates an alias for the grep command:

`alias grep = "grep -n"`

This alias will now run the grep command with the -n option, which will print the line numbers of the matches.

Aliases are stored in the .bashrc file. This file is located in your home directory. You can edit the .bashrc file to add or remove aliases.

When you open a new terminal, the aliases in your .bashrc file will be loaded. This means that you can use your aliases immediately after opening a new terminal.

# help 

```
alias: alias [-p] [name[=value] ... ]
    Define or display aliases.
    
    Without arguments, `alias' prints the list of aliases in the reusable
    form `alias NAME=VALUE' on standard output.
    
    Otherwise, an alias is defined for each NAME whose VALUE is given.
    A trailing space in VALUE causes the next word to be checked for
    alias substitution when the alias is expanded.
    
    Options:
      -p        print all defined aliases in a reusable format
    
    Exit Status:
    alias returns true unless a NAME is supplied for which no alias has been
    defined.
```

