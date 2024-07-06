# read

The `read` command in Unix-like operating systems is used in shell scripting to take input from the user or from a file. It is a simple but powerful tool for making scripts interactive or for processing input data line by line.

### Basic Syntax

```bash
read [options] variable
```

- **variable:** The name of the variable that will store the input.
- **options:** Various flags that modify the behavior of the `read` command.

### Examples

#### Reading User Input

```bash
echo "Enter your name:"
read name
echo "Hello, $name!"
```

#### Reading Multiple Variables

```bash
echo "Enter your first name and last name:"
read first_name last_name
echo "Hello, $first_name $last_name!"
```

#### Reading Input with a Prompt

You can use the `-p` option to display a prompt without using `echo`:

```bash
read -p "Enter your name: " name
echo "Hello, $name!"
```

#### Silent Input (e.g., Passwords)

The `-s` option hides the input, useful for passwords:

```bash
read -sp "Enter your password: " password
echo
echo "Password entered."
```

### Reading Input from a File

You can use the `read` command to process a file line by line:

```bash
while read line
do
  echo "Processing: $line"
done < input_file.txt
```

### Advanced Examples

#### Reading Input with a Timeout

The `-t` option sets a timeout for input:

```bash
if read -t 5 -p "Enter your name within 5 seconds: " name
then
  echo "Hello, $name!"
else
  echo "Timed out."
fi
```

#### Reading a Single Character

The `-n` option reads a specified number of characters (usually one):

```bash
read -n 1 -p "Press any key to continue..."
echo
echo "Key pressed."
```

#### Reading Input with a Delimiter

The `-d` option allows specifying a delimiter other than newline:

```bash
read -d ':' -p "Enter data separated by colons: " data
echo "You entered: $data"
```

#### Reading an Array

```bash
echo "Enter names separated by spaces:"
read -a names
echo "Names entered: ${names[@]}"
```

#### Using Default Values

You can provide a default value if the user input is empty:

```bash
read -p "Enter your name: " name
name=${name:-"Default Name"}
echo "Hello, $name!"
```

### Practical Use Cases

#### Menu Selection

```bash
echo "Select an option:"
echo "1. Option One"
echo "2. Option Two"
echo "3. Option Three"
read -p "Enter your choice [1-3]: " choice

case $choice in
  1) echo "You chose Option One." ;;
  2) echo "You chose Option Two." ;;
  3) echo "You chose Option Three." ;;
  *) echo "Invalid choice." ;;
esac
```

#### Confirmation Prompt

```bash
read -p "Are you sure you want to continue? (y/n): " confirmation
if [[ $confirmation == "y" || $confirmation == "Y" ]]
then
  echo "Continuing..."
else
  echo "Aborting..."
fi
```

### Conclusion

The `read` command is a versatile and essential tool for shell scripting, enabling scripts to interact with users or process input data effectively. By understanding its various options and use cases, you can create more dynamic and responsive scripts.
