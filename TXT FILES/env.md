# env

The `env` command in Unix and Linux is used to run a command in a modified environment or to display the environment variables. Environment variables are dynamic values that affect the way processes and applications run on a computer.

### Usage of `env`

#### Display All Environment Variables

To display all current environment variables, simply run:

```sh
env
```

This command will list all the environment variables and their values in the current shell session.

#### Running a Command with Modified Environment Variables

You can use `env` to run a command with one or more environment variables temporarily set. The syntax is:

```sh
env VAR1=value1 VAR2=value2 command
```

For example, to run a script with a modified `PATH` variable:

```sh
env PATH=/custom/path:/another/path myscript.sh
```

In this example, `myscript.sh` will be executed with the `PATH` variable set to `/custom/path:/another/path`, but the current shell's `PATH` will remain unchanged.

#### Clearing All Environment Variables

To clear all environment variables and then run a command:

```sh
env -i command
```

For example:

```sh
env -i /bin/bash
```

This will start a new Bash shell with an empty environment.

### Common Environment Variables

Here are some commonly used environment variables:

- **`PATH`**: Specifies a list of directories where the shell looks for executable files.
- **`HOME`**: The home directory of the current user.
- **`USER`**: The name of the current user.
- **`SHELL`**: The path to the current user's shell.
- **`LANG`**: Defines the language, localization, and character encoding settings.
- **`PWD`**: The current working directory.

### Example: Using `env`

#### Displaying Environment Variables

To list all environment variables:

```sh
env
```

#### Setting an Environment Variable for a Command

To temporarily set the `EDITOR` variable to `vim` for a session:

```sh
env EDITOR=vim crontab -e
```

This command opens the `crontab` editor using `vim`, but does not change the `EDITOR` variable for the current shell session.

#### Running a Command with a Clean Environment

To run a command with an empty environment:

```sh
env -i /usr/bin/env
```

This runs the `env` command itself in a completely clean environment, which will show no variables except for those set by `env` itself.

### Practical Use Cases

1. **Temporary Environment for a Script**:
   - Suppose you want to test a script with a different `PATH`:

     ```sh
     env PATH=/custom/path myscript.sh
     ```

2. **Debugging**:
   - If a program behaves differently when run in a clean environment, it can help identify issues related to environment variables:

     ```sh
     env -i /path/to/program
     ```

3. **Changing Locale Settings**:
   - Run a program with a different language setting:

     ```sh
     env LANG=es_ES.UTF-8 some_program
     ```

### Conclusion

The `env` command is a versatile tool for managing and interacting with environment variables in Unix and Linux systems. Whether you need to display, modify, or clear environment variables, `env` provides a simple and effective way to control the environment in which commands are executed. Understanding how to use `env` can help you better manage your system's behavior and debug issues related to environment variables.

# help 

```
env [options] [variable=value]

Print or set environment variables.

Options:

-i, --ignore-environment   Do not inherit environment variables from the parent shell.
-u, --unset   Unset environment variable.
-l, --list   List all environment variables.
-v, --verbose   Print each variable and its value.
-P, --path   Print the current $PATH.
-S, --secure   Print environment variables in a secure way.

Examples:

    env
    env HOME=/home/user
    env -i
    env -u HOME
    env -l
    env -v
    env -P
    env -S

```
