# function

In Unix-like operating systems, a shell function is a set of commands grouped together under a single name, allowing you to execute the group of commands by calling the function name. Shell functions are useful for simplifying complex scripts, reusing code, and improving script readability.

### Defining and Using Shell Functions

### Basic Syntax

```bash
function_name() {
  commands
}
```

or

```bash
function function_name {
  commands
}
```

### Example of a Simple Function

Here’s a simple example to demonstrate how to define and use a shell function:

```bash
greet() {
  echo "Hello, $1!"
}

# Call the function
greet "World"
```

Output:
```bash
Hello, World!
```

### Detailed Example with Explanation

Let's create a more detailed example with a function that adds two numbers.

```bash
add_numbers() {
  local num1=$1
  local num2=$2
  local sum=$((num1 + num2))
  echo "The sum of $num1 and $num2 is: $sum"
}

# Call the function with arguments
add_numbers 5 7
```

Output:
```bash
The sum of 5 and 7 is: 12
```

#### Explanation:
- **Function Definition:** `add_numbers()` defines the function.
- **Local Variables:** `local num1=$1` assigns the first argument to `num1`, and `local num2=$2` assigns the second argument to `num2`.
- **Calculation:** `local sum=$((num1 + num2))` calculates the sum of `num1` and `num2`.
- **Output:** `echo` prints the result.

### Using Functions in Scripts

You can use functions within scripts to organize code better and reuse functionality.

```bash
#!/bin/bash

# Function to check if a directory exists
check_directory() {
  if [ -d "$1" ]; then
    echo "Directory $1 exists."
  else
    echo "Directory $1 does not exist."
  fi
}

# Function to create a new directory
create_directory() {
  mkdir -p "$1"
  echo "Directory $1 created."
}

# Main script logic
DIR="/path/to/directory"

check_directory "$DIR"
create_directory "$DIR"
check_directory "$DIR"
```

### Advanced Features

#### Returning Values

In shell functions, you cannot return values like in other programming languages. Instead, you can use `echo` to output values and capture the output using command substitution.

```bash
get_sum() {
  local num1=$1
  local num2=$2
  echo $((num1 + num2))
}

# Capture the output of the function
sum=$(get_sum 5 7)
echo "The sum is: $sum"
```

#### Using Exit Status

You can use the `return` statement to set the exit status of a function, which can be checked using `$?`.

```bash
is_even() {
  local num=$1
  if (( num % 2 == 0 )); then
    return 0
  else
    return 1
  fi
}

# Call the function and check the exit status
is_even 4
if [ $? -eq 0 ]; then
  echo "The number is even."
else
  echo "The number is odd."
fi
```

### Recursion

Shell functions can be recursive, but this should be used with caution due to potential stack overflow issues.

```bash
factorial() {
  local num=$1
  if [ $num -le 1 ]; then
    echo 1
  else
    local prev=$(factorial $((num - 1)))
    echo $((num * prev))
  fi
}

# Call the recursive function
factorial 5
```

Output:
```bash
120
```

### Conclusion

Shell functions are powerful tools for structuring and reusing code in shell scripts. They help make scripts more modular, readable, and maintainable. By understanding how to define and use shell functions, you can significantly improve your scripting capabilities.

# help 

```
iotop [options]

Display a dynamic real-time view of the I/O activity of processes.

Options:

-d, --delay=NUMBER   Set the delay between updates in seconds.
-n, --number=NUMBER  Set the number of processes to display.
-p, --pid=PID        Only show processes with the specified PID.
-u, --user=USER      Only show processes owned by the specified user.
-b, --background     Show all processes, even those in the background.
-h, --help           Show this help message.

Examples:

    iotop
    iotop -d 2
    iotop -p 1234
    iotop -u root
```


## breakdown

```
-d, --delay=NUMBER: This option sets the delay between updates in seconds. The default is 2 seconds.
-n, --number=NUMBER: This option sets the number of processes to display. The default is 15.
-p, --pid=PID: This option only shows processes with the specified PID.
-u, --user=USER: This option only shows processes owned by the specified user.
-b, --background: This option shows all processes, even those in the background.
-h, --help: This option shows this help message.
```

