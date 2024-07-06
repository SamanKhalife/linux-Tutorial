# seq

The `seq` command in Unix-like operating systems generates sequences of numbers, either in increasing or decreasing order. It's useful in shell scripting for iterating over a range of numbers or generating lists.

### Basic Syntax

```bash
seq [options] [start] [increment] end
```

- **options:** Optional flags that modify the behavior of the `seq` command.
- **start:** The starting number of the sequence (default is 1).
- **increment:** The increment between numbers (default is 1).
- **end:** The ending number of the sequence.

### Examples

#### Generating a Sequence of Numbers

```bash
# Generate numbers from 1 to 5
seq 1 5
```

Output:
```
1
2
3
4
5
```

#### Specifying Start and End

```bash
# Generate numbers from 5 to 10
seq 5 10
```

Output:
```
5
6
7
8
9
10
```

#### Generating Numbers with Increment

```bash
# Generate numbers from 1 to 10 with increment of 2
seq 1 2 10
```

Output:
```
1
3
5
7
9
```

#### Generating a Sequence in Reverse Order

```bash
# Generate numbers from 10 to 1
seq 10 -1 1
```

Output:
```
10
9
8
7
6
5
4
3
2
1
```

### Practical Use Cases

#### Looping with `for` and `seq`

```bash
# Loop through numbers from 1 to 5
for i in $(seq 1 5)
do
  echo "Number: $i"
done
```

#### Creating Directories

```bash
# Create directories dir1 to dir5
for i in $(seq 1 5)
do
  mkdir "dir$i"
done
```

#### File Operations

```bash
# Create empty files file1.txt to file10.txt
for i in $(seq 1 10)
do
  touch "file$i.txt"
done
```

### Advanced Examples

#### Using `seq` in a `for` Loop

```bash
# Print squares of numbers from 1 to 5
for i in $(seq 1 5)
do
  echo "Square of $i is $(($i * $i))"
done
```

#### Generating a Fixed-Length Sequence

```bash
# Generate a sequence of 10 numbers starting from 100
seq 100 109
```

Output:
```
100
101
102
103
104
105
106
107
108
109
```

### Conclusion

The `seq` command is a handy tool in shell scripting for generating sequences of numbers efficiently. It simplifies tasks such as looping through ranges of numbers, creating directories or files in batches, and performing mathematical operations in scripts. By leveraging `seq`, you can make your shell scripts more dynamic and versatile.
