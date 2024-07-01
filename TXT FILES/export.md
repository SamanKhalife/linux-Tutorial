# export

The `export` command in Unix and Linux is used to set environment variables and make them available to child processes. Environment variables are dynamic values that can affect the way running processes behave on a computer. By exporting a variable, you ensure that it is passed to any child processes created by the shell.

### Basic Usage

The syntax for the `export` command is:

```sh
export VARIABLE_NAME=VALUE
```

### Examples

#### Setting and Exporting a Variable

To create a variable and export it so that it is available to child processes:

```sh
export MY_VAR="Hello, World!"
```

Now `MY_VAR` is available to any child process started from this shell. For example:

```sh
echo $MY_VAR
```

This will output:

```
Hello, World!
```

#### Exporting an Existing Variable

If you have already defined a variable and you want to export it:

```sh
MY_VAR="Hello, World!"
export MY_VAR
```

#### Unsetting an Exported Variable

To remove an environment variable, you can use the `unset` command:

```sh
unset MY_VAR
```

### Practical Use Cases

#### Exporting Path Variables

A common use of `export` is to add directories to the `PATH` variable, which determines where the shell looks for executable files:

```sh
export PATH=$PATH:/new/directory
```

Now, the shell will include `/new/directory` when searching for executable files.

#### Setting Variables for Scripts

When running a script, you might want to pass environment variables to it:

```sh
export CONFIG_FILE=/etc/myconfig.conf
./myscript.sh
```

Inside `myscript.sh`, you can access `CONFIG_FILE`:

```sh
#!/bin/bash
echo "Using config file: $CONFIG_FILE"
```

#### Making Variables Available to Sub-Shells

If you start a sub-shell or another script from your current shell, exported variables will be available to those sub-shells:

```sh
export GREETING="Hello"
bash -c 'echo $GREETING'
```

This will output:

```
Hello
```

### Combining with Other Commands

#### Setting and Exporting in One Line

You can set and export a variable in a single command:

```sh
export MY_VAR="New Value"
```

#### Using Export with Command Substitution

You can use command substitution to set a variable to the output of a command and export it:

```sh
export CURRENT_DATE=$(date)
echo $CURRENT_DATE
```

This will output the current date and time.

### Best Practices

1. **Use Meaningful Names**: Give your environment variables meaningful names that clearly indicate their purpose.
2. **Uppercase Naming**: By convention, environment variable names are usually in uppercase.
3. **Avoid Overwriting System Variables**: Be cautious not to overwrite important system variables like `PATH`, `HOME`, `USER`, etc., unless you are sure of the consequences.

### Conclusion

The `export` command is a powerful tool for managing environment variables in Unix and Linux systems. By understanding how to use `export`, you can control the environment in which your commands and scripts run, making it possible to pass important information to child processes and ensure consistent behavior across different parts of your system. Whether you're setting up paths, configuring scripts, or managing complex environments, `export` is an essential command to master.



# help 

```

```
