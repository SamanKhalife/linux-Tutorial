# while

The `while` loop in Unix-like operating systems is a control structure in shell scripting that repeatedly executes a set of commands as long as a specified condition is true. This can be particularly useful for tasks that need to continue until a particular state is achieved.

### Basic Syntax

```bash
while condition
do
  commands
done
```

- **condition:** A command or set of commands that return a true (0) or false (non-zero) exit status.
- **commands:** The commands to execute as long as the condition is true.

### Examples

#### Basic Example

```bash
count=1
while [ $count -le 5 ]
do
  echo "Count: $count"
  count=$((count + 1))
done
```

Output:
```bash
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
```

#### Reading User Input

```bash
input=""
while [ "$input" != "exit" ]
do
  echo "Type something (type 'exit' to quit):"
  read input
  echo "You typed: $input"
done
```

#### Checking a Command's Exit Status

```bash
while ping -c 1 example.com > /dev/null 2>&1
do
  echo "example.com is reachable"
  sleep 5
done

echo "example.com is not reachable"
```

### Practical Use Cases

#### Monitoring a Process

```bash
process_name="mysqld"
while pgrep $process_name > /dev/null
do
  echo "$process_name is running"
  sleep 10
done

echo "$process_name has stopped"
```

#### Waiting for a File to Exist

```bash
file="/path/to/file"
while [ ! -f $file ]
do
  echo "Waiting for $file to exist..."
  sleep 5
done

echo "$file exists"
```

#### Implementing a Simple Menu

```bash
choice=0
while [ $choice -ne 3 ]
do
  echo "Menu:"
  echo "1. Option 1"
  echo "2. Option 2"
  echo "3. Exit"
  read -p "Enter your choice: " choice

  case $choice in
    1)
      echo "You chose option 1"
      ;;
    2)
      echo "You chose option 2"
      ;;
    3)
      echo "Exiting..."
      ;;
    *)
      echo "Invalid choice"
      ;;
  esac
done
```

### Advanced Examples

#### Infinite Loop with Break Condition

```bash
while true
do
  echo "Infinite loop. Type 'break' to exit."
  read input
  if [ "$input" == "break" ]; then
    break
  fi
done
```

#### Looping with Multiple Conditions

```bash
count=0
while [ $count -lt 10 ] && [ $count -ne 5 ]
do
  echo "Count: $count"
  count=$((count + 1))
done
```

### Conclusion

The `while` loop is a powerful construct in shell scripting that allows you to execute commands repeatedly based on a condition. It is especially useful for tasks that need to be performed until a certain condition is met or as long as a certain condition remains true. By mastering the `while` loop, you can write more dynamic and responsive shell scripts.
