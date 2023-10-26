# basename

The `basename` command in Linux is used to print the base name of a file or directory. The base name is the last component of the file or directory path.

The `basename` command is used in the following syntax:

```
basename [options] [file or directory]
```

The `file or directory` is the file or directory whose base name you want to print.

The `options` can be used to specify the following:

* `-s` : Strip any trailing suffix from the file name.
* `-z` : Use a NUL character (`\0`) as the line terminator instead of a newline character (`\n`).

For example, to print the base name of the file `/home/user/file.txt`, you would run the following command:

```
basename /home/user/file.txt
```

This command will print the base name of the file, which is `file.txt`.

To strip any trailing suffix from the file name, you would run the following command:

```
basename -s /home/user/file.txt
```

This command will print the base name of the file, without the trailing suffix `.txt`, which is `file`.

To use a NUL character (`\0`) as the line terminator instead of a newline character (`\n`), you would run the following command:

```
basename -z /home/user/file.txt
```

This command will print the base name of the file, with a NUL character (`\0`) as the line terminator, which is `file\0`.

The `basename` command is a simple and easy-to-use command that can be used to print the base name of a file or directory. It is a versatile command that can be used in a variety of contexts.

Here are some additional things to note about the `basename` command:

* The `basename` command can be used to print the base name of any file or directory.
* The `basename` command can be used to strip any trailing suffix from a file name.
* The `basename` command can be used to use a NUL character (`\0`) as the line terminator instead of a newline character (`\n`).
* The `basename` command is a powerful command that can be used in a variety of contexts.



# help 

```

```
