# 

The tee command in Linux is used to copy standard input to one or more files. It can also be used to pipe the output of a command to one or more files.

For example, the following command will copy standard input to the file output.txt and to the standard output:

`tee output.txt`

The following command will pipe the output of the ls command to the file output.txt and to the standard output:

`ls | tee output.txt`

The tee command is a powerful tool that can be used to save the output of a command to a file. However, it is important to use it carefully, as it can be easy to overwrite existing files.

# help 

```
Usage: tee [OPTION]... [FILE]...
Copy standard input to each FILE, and also to standard output.

  -a, --append              append to the given FILEs, do not overwrite
  -i, --ignore-interrupts   ignore interrupt signals
  -p                        diagnose errors writing to non pipes
      --output-error[=MODE]   set behavior on write error.  See MODE below
      --help     display this help and exit
      --version  output version information and exit

```



## breakdown

```
-a, --append: This option appends to the file instead of overwriting it.
-i, --ignore-eof: This option does not stop at end-of-file (EOF) on standard input.
-h, --help: This option shows this help message.
```