# df

The df command is a command-line utility that can be used to display the disk space usage of a Linux system. It can be used to see how much disk space is available, how much disk space is used, and how much disk space is used by each filesystem.

The df command is a powerful tool that can be used to troubleshoot disk space problems and to manage disk space.

For example, the following output from the df command shows that the filesystem / has a total size of 100 GB, 80 GB of used space, and 20 GB of available space:

```
Filesystem     1K-blocks   Used  Available Use% Mounted on
/dev/sda1      10048864   8027520   2021344   80% /
```

You can also use the df command to display the disk space usage for a specific filesystem. To do this, you can use the -T option. For example, the following command will display the disk space usage for the filesystem /:

`df -T /`


# help 

```
df [options] [filesystems]

Display filesystem disk space usage.

Options:

-a, --all    show all filesystems
-h, --human-readable   print sizes in human-readable format
-k, --kilobytes    use 1024-byte blocks
-m, --megabytes   use 1048576-byte blocks
-t, --type=TYPE   show only filesystems of the specified type
-T, --print-type   print the filesystem type
-H, --si    use powers of 1000 instead of 1024
-i, --inodes   print inode usage
-x, --exclude=FSTYPE   exclude filesystems of the specified type
-P, --portability   print in a format compatible with other systems
-S, --sync    force filesystems to be synced before displaying information
-v, --verbose   print more information
-?, --help     display this help message

Examples:

    df
    df -h
    df -T
    df -i
    df -x ext4
```


