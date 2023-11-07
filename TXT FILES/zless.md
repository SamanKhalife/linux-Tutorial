# zless

The zless command in Linux is a filter that can be used to display compressed files one screenful at a time. The zless command is a part of the gzip suite of tools.

Here are some examples of how to use the zless command:

```
# To display the file `file.gz` one screenful at a time:
zless file.gz

# To overwrite the existing file `file.txt.gz`:
zless -f file.gz

# To show help for the `zless` command:
zless --help
```

# help

```
Usage: /usr/bin/zless [OPTION]... [FILE]...
Like 'less', but operate on the uncompressed contents of any compressed FILEs.

Options are the same as for 'less'
```
## breakdown

```
-f, --force: This option overwrites the existing file.
-h, --help: This option shows this help message.
-V, --version: This option prints version information.

```
