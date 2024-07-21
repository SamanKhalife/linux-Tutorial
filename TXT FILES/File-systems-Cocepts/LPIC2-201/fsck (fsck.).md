# fsck

`fsck` (file system check) is a command-line utility in Unix-like operating systems used to check and repair filesystems. It verifies the integrity of filesystems, detects and fixes errors, and ensures that the filesystem structure is consistent.

### Key Features

- **Filesystem Integrity Checking**: Scans filesystems for inconsistencies and errors.
- **Repair Capabilities**: Attempts to automatically repair detected issues.
- **Support for Various Filesystems**: Works with different filesystem types using specific front-end commands.

### Basic Usage

The general syntax for `fsck` is:

```sh
fsck [options] <device>
```

- **`[options]`**: Command-line options for configuring the behavior of `fsck`.
- **`<device>`**: The disk partition or device to check (e.g., `/dev/sdX1`).

### Common Filesystem Types and Their Commands

#### 1. **ext2/ext3/ext4**

For ext2, ext3, and ext4 filesystems, `fsck` can be used with the specific filesystem type option:

```sh
fsck.ext4 /dev/sdX1
```

Options:
- **`-n`**: No changes; just show what would be done.
- **`-y`**: Automatically answer "yes" to all questions.
- **`-f`**: Force a check, even if the filesystem seems clean.

#### 2. **xfs**

For `xfs`, use the `xfs_repair` utility instead of `fsck`:

```sh
xfs_repair /dev/sdX1
```

Options:
- **`-n`**: No modifications; just show what would be done.
- **`-L`**: Log file damage repair.

#### 3. **btrfs**

For `btrfs`, use the `btrfs check` command:

```sh
btrfs check /dev/sdX1
```

Options:
- **`--repair`**: Attempt to repair the filesystem.
- **`--readonly`**: Perform a read-only check.

#### 4. **vfat**

For `vfat` (FAT32) filesystems, use `fsck.vfat`:

```sh
fsck.vfat /dev/sdX1
```

Options:
- **`-a`**: Automatically fix errors.
- **`-n`**: No changes; just show what would be done.

### Examples

#### Check an ext4 Filesystem

To check the filesystem on `/dev/sda1`:

```sh
fsck.ext4 /dev/sda1
```

#### Force Check and Repair

To force a check and automatically repair errors:

```sh
fsck -f -y /dev/sda1
```

#### Check an XFS Filesystem

For XFS, use `xfs_repair`:

```sh
xfs_repair /dev/sda1
```

### Important Considerations

- **Unmount Filesystem**: It is generally recommended to unmount the filesystem before running `fsck` to avoid data corruption. If the filesystem cannot be unmounted (e.g., it is the root filesystem), you may need to boot from a live CD or enter recovery mode.
- **Data Backup**: Always back up important data before performing filesystem checks and repairs, as repairs can potentially lead to data loss.
- **Filesystem Type**: Ensure you are using the correct `fsck` variant for your filesystem type. Different filesystems have different tools and options.

### Summary

`fsck` is a powerful utility for maintaining filesystem integrity by checking and repairing file systems. It supports various filesystem types through specific front-end commands and offers options for customizing checks and repairs. Proper use of `fsck` helps ensure that filesystems remain consistent and reliable, preventing data loss and system issues.
