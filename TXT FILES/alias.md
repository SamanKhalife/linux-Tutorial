help 

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

