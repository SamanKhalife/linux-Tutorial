# if


The `if` statement in Unix-like operating systems is a fundamental control structure used in shell scripting to perform conditional execution of commands. It allows you to execute a set of commands only if a specified condition is true, and optionally execute another set of commands if the condition is false.

### Basic Syntax

```bash
if condition
then
  commands
fi
```

- **condition:** A command or test expression that returns a true (0) or false (non-zero) exit status.
- **commands:** The commands to execute if the condition is true.

### Extended Syntax

You can extend the `if` statement with `else` and `elif` (else if) clauses to handle multiple conditions:

```bash
if condition1
then
  commands1
elif condition2
then
  commands2
else
  commands3
fi
```

- **condition1:** The primary condition to test.
- **commands1:** Commands to execute if condition1 is true.
- **condition2:** The secondary condition to test if condition1 is false.
- **commands2:** Commands to execute if condition2 is true.
- **commands3:** Commands to execute if neither condition1 nor condition2 is true.

### Examples

#### Basic `if` Statement

```bash
if [ $USER == "root" ]
then
  echo "You are the root user."
fi
```

#### `if-else` Statement

```bash
if [ $USER == "root" ]
then
  echo "You are the root user."
else
  echo "You are not the root user."
fi
```

#### `if-elif-else` Statement

```bash
if [ $USER == "root" ]
then
  echo "You are the root user."
elif [ $USER == "admin" ]
then
  echo "You are the admin user."
else
  echo "You are a regular user."
fi
```

### Practical Use Cases

#### Checking File Existence

```bash
file="/path/to/file"
if [ -f $file ]
then
  echo "$file exists."
else
  echo "$file does not exist."
fi
```

#### Checking Directory Existence

```bash
directory="/path/to/directory"
if [ -d $directory ]
then
  echo "$directory exists."
else
  echo "$directory does not exist."
fi
```

#### Comparing Numbers

```bash
num1=10
num2=20
if [ $num1 -lt $num2 ]
then
  echo "$num1 is less than $num2."
else
  echo "$num1 is not less than $num2."
fi
```

#### Checking Command Exit Status

```bash
ping -c 1 example.com > /dev/null 2>&1
if [ $? -eq 0 ]
then
  echo "example.com is reachable."
else
  echo "example.com is not reachable."
fi
```

### Advanced Examples

#### Nested `if` Statements

```bash
if [ $USER == "root" ]
then
  if [ -f "/etc/passwd" ]
  then
    echo "Root user and /etc/passwd exists."
  else
    echo "Root user but /etc/passwd does not exist."
  fi
else
  echo "You are not the root user."
fi
```

#### Combining Conditions with `&&` and `||`

```bash
if [ $USER == "root" ] && [ -f "/etc/passwd" ]
then
  echo "Root user and /etc/passwd exists."
else
  echo "Either not root user or /etc/passwd does not exist."
fi
```

### Using `test` Command

The `test` command (or its alias `[ ... ]`) is used to evaluate expressions:

```bash
if test $USER = "root"
then
  echo "You are the root user."
fi
```

### Conclusion

The `if` statement is a crucial tool in shell scripting, enabling conditional execution of commands based on the evaluation of expressions. By using `if`, `else`, and `elif` clauses, you can create scripts that respond dynamically to different conditions and states.

# help 

```
-n: This option tests if a variable is not empty.
-z: This option tests if a variable is empty.
-e: This option tests if a file or directory exists.
-f: This option tests if a file is a regular file.
-d: This option tests if a file is a directory.
```

