



# help 

```
pushd: pushd [-n] [+N | -N | dir]
    Add directories to stack.
    
    Adds a directory to the top of the directory stack, or rotates
    the stack, making the new top of the stack the current working
    directory.  With no arguments, exchanges the top two directories.
    
    Options:
      -n        Suppresses the normal change of directory when adding
                directories to the stack, so only the stack is manipulated.
    
    Arguments:
      +N        Rotates the stack so that the Nth directory (counting
                from the left of the list shown by `dirs', starting with
                zero) is at the top.
    
      -N        Rotates the stack so that the Nth directory (counting
                from the right of the list shown by `dirs', starting with
                zero) is at the top.
    
      dir       Adds DIR to the directory stack at the top, making it the
                new current working directory.
    
    The `dirs' builtin displays the directory stack.
    
    Exit Status:
    Returns success unless an invalid argument is supplied or the directory
    change fails.
```



## breakdown

```

```