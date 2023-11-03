# return

In Linux, the `return` command is used to exit a shell function or script. It is a built-in command that is supported by all Linux distributions.

The `return` command is used in the following syntax:

```
return [exit status]
```

The `exit status` is an integer that is used to indicate the success or failure of the shell function or script. The exit status can be any number from 0 to 255.

For example, the following shell function will print the message "Hello, world!" and then return the exit status 0:

```
function hello_world() {
  echo "Hello, world!"
  return 0
}
```

The following script will print the message "This is a script" and then return the exit status 1:

```
#!/bin/bash

echo "This is a script"
return 1
```

The `return` command is a useful tool for controlling the flow of execution in a shell function or script. It can be used to exit a function or script early, or to return an exit status to the calling function or script.

Here are some additional things to note about the `return` command:

* The `return` command can only be used inside a shell function or script.
* The `return` command cannot be used in the interactive shell.
* The `return` command will always exit the shell function or script, even if it is nested inside another shell function or script.
* The `return` command can be used to return a non-zero exit status to the calling function or script. This can be useful for debugging or for signaling an error condition.

The `return` command is a versatile tool that can be used to control the flow of execution in a shell function or script. It is a built-in command that is supported by all Linux distributions.



# help 

```
return: return [n]
    Return from a shell function.
    
    Causes a function or sourced script to exit with the return value
    specified by N.  If N is omitted, the return status is that of the
    last command executed within the function or script.
    
    Exit Status:
    Returns N, or failure if the shell is not executing a function or script.
```
