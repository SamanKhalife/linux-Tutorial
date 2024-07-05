# df

The `df` command in Unix-like operating systems is used to display information about disk space usage on filesystems. It stands for "disk free". Here’s an overview of `df` and its common usage:

### Overview of `df`

**Purpose:** `df` is used to report the amount of disk space used and available on filesystems mounted in the system.

**Availability:** `df` is a standard command-line utility available on Unix-like systems, including Linux distributions, macOS, and BSD variants.

### Common `df` Commands and Usage

1. **Display Disk Space Usage for All Filesystems:**
   - To display information about all mounted filesystems:
     ```bash
     df
     ```
     By default, `df` displays information in 1 KB blocks.

2. **Display Disk Space Usage in Human-Readable Format:**
   - To display disk space usage in a more human-readable format (e.g., in kilobytes, megabytes, gigabytes):
     ```bash
     df -h
     ```
     The `-h` option (or `--human-readable`) converts sizes into a human-readable format.

3. **Display Disk Space Usage for Specific Filesystem:**
   - To display disk space usage for a specific filesystem (e.g., `/dev/sda1`):
     ```bash
     df -h /dev/sda1
     ```
     Replace `/dev/sda1` with the device or mount point you want to inspect.

4. **Display Filesystem Type:**
   - To display the type of filesystem along with disk space usage:
     ```bash
     df -T
     ```
     The `-T` option (or `--print-type`) adds a column showing the filesystem type.

5. **Display Inode Usage:**
   - To display inode usage (number of used and available inodes) instead of disk space:
     ```bash
     df -i
     ```
     The `-i` option (or `--inodes`) switches `df` to report inode usage instead of disk space.

6. **Display Total Disk Space Usage:**
   - To display the total disk space usage across all filesystems:
     ```bash
     df -h --total
     ```
     The `--total` option provides a summary at the end of the output, showing the total used and available disk space.

### Considerations

- **Reserved Space:** By default, `df` shows the amount of disk space available to ordinary users. Some filesystems reserve a percentage of disk space for root user (`root-reserved`). This reserved space is not shown as available to normal users.

- **Symbolic Links:** `df` does not follow symbolic links. It reports on the filesystem that contains the symlink's target.

- **Performance Impact:** Running `df` on large filesystems or with the `-i` option for inode information can be resource-intensive.

### Alternatives

- **pydf:** A Python-based tool that provides an enhanced and colorful output of disk usage, similar to `df`.

- **Disk Usage Analyzers:** Graphical tools like `Baobab` (for GNOME) and `Filelight` (for KDE) offer visual representations of disk usage.

### Conclusion

`df` is a fundamental command-line tool for monitoring disk space usage on Unix-like systems, providing essential information about available storage and filesystem types. It’s useful for managing disk resources, identifying available space, and planning storage allocation.

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


