
The if command is a conditional statement that can be used to execute commands based on the value of an expression.

The expression argument is a Boolean expression that evaluates to either true or false. If the expression evaluates to true, the then block of commands is executed. If the expression evaluates to false, the then block of commands is skipped.

The commands argument is a list of commands that are executed if the expression evaluates to true.

For example, the following command will check if the variable name is set. If it is, the command echo "Hello, $name" is executed. If it is not, the command is skipped.


# help 

```
-n: This option tests if a variable is not empty.
-z: This option tests if a variable is empty.
-e: This option tests if a file or directory exists.
-f: This option tests if a file is a regular file.
-d: This option tests if a file is a directory.
```

