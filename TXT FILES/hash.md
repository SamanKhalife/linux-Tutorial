# hash

The hash command in Linux is used to maintain a hash table of recently executed commands. This can be useful for speeding up the execution of commands that you use frequently.


For example, the following command will add the command ls to the hash table:

`hash ls`

The next time you run the ls command, the shell will look for the command in the hash table. If the command is found in the hash table, the shell will execute the command directly. This can speed up the execution of the command.

The hash command also has a -r option that can be used to clear the hash table. This can be useful if you want to start over with a clean hash table.

# help 

```
hash: hash [-lr] [-p pathname] [-dt] [name ...]
    Remember or display program locations.
    
    Determine and remember the full pathname of each command NAME.  If
    no arguments are given, information about remembered commands is displayed.
    
    Options:
      -d        forget the remembered location of each NAME
      -l        display in a format that may be reused as input
      -p pathname       use PATHNAME as the full pathname of NAME
      -r        forget all remembered locations
      -t        print the remembered location of each NAME, preceding
                each location with the corresponding NAME if multiple
                NAMEs are given
```

