The "mkfs" command in Linux is used to create a file system on a storage device or partition. It is typically used to format a disk or partition with a specific file system type, preparing it for use with the operating system.

# help 

```
Usage:
 mkfs [options] [-t <type>] [fs-options] <device> [<size>]

Make a Linux filesystem.

Options:
 -t, --type=<type>  filesystem type; when unspecified, ext2 is used
     fs-options     parameters for the real filesystem builder
     <device>       path to the device to be used
     <size>         number of blocks to be used on the device
 -V, --verbose      explain what is being done;
                      specifying -V more than once will cause a dry-run
 -h, --help         display this help
 -V, --version      display version

For more details see mkfs(8).
```

