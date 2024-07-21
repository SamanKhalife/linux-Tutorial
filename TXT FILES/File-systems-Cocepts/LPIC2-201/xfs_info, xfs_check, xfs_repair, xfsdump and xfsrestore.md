# Xfs_info, Xfs_check, Xfs_repair, Xfsdump and Xfsrestore

XFS is a high-performance filesystem designed for handling large files and large amounts of data. It offers advanced features like journaling, high capacity, and scalability. Here’s an overview of several key XFS utilities:

### `xfs_info`

The `xfs_info` command provides detailed information about an XFS filesystem. It can be used to check filesystem parameters, configuration, and status.

#### Basic Syntax

```sh
xfs_info <mount-point>
```

- **`<mount-point>`**: The mount point of the XFS filesystem (e.g., `/mnt/xfs`).

#### Example

```sh
xfs_info /mnt/xfs
```

#### Output

The output will show details such as:
- **Filesystem Geometry**: Block size, number of blocks, etc.
- **File System State**: Status of the filesystem, including whether it is mounted.

### `xfs_check`

The `xfs_check` command was used to verify the integrity of an XFS filesystem. However, this utility has been deprecated and replaced by `xfs_repair` for filesystem checking and repair tasks.

#### Basic Syntax

```sh
xfs_check <device>
```

- **`<device>`**: The device or partition of the XFS filesystem (e.g., `/dev/sdX1`).

#### Example

```sh
xfs_check /dev/sdX1
```

**Note:** Use `xfs_repair` for modern filesystem checks.

### `xfs_repair`

The `xfs_repair` command is used to repair and check the integrity of XFS filesystems. It is a crucial tool for fixing filesystem inconsistencies and errors.

#### Basic Syntax

```sh
xfs_repair [options] <device>
```

- **`[options]`**: Command-line options for specific repair operations.
- **`<device>`**: The device or partition to repair (e.g., `/dev/sdX1`).

#### Common Options

- **`-n`**: No changes mode; performs a read-only check.
- **`-L`**: Clear the log and repair the filesystem (use with caution as it can result in data loss).
- **`-d`**: Enable debugging messages.

#### Example

To perform a repair on `/dev/sdX1`:

```sh
xfs_repair /dev/sdX1
```

### `xfsdump`

The `xfsdump` command is used to create backups of XFS filesystems. It supports incremental backups, allowing for efficient storage and restoration.

#### Basic Syntax

```sh
xfsdump [options] <dump-device> <filesystem>
```

- **`[options]`**: Command-line options for backup operation.
- **`<dump-device>`**: The destination for the backup (e.g., `/dev/st0` for tape).
- **`<filesystem>`**: The XFS filesystem to backup (e.g., `/mnt/xfs`).

#### Common Options

- **`-0`**: Full backup.
- **`-1`**: Incremental backup.
- **`-u`**: Update the backup (when used with incremental backups).

#### Example

To create a full backup of the XFS filesystem mounted at `/mnt/xfs` to a tape drive `/dev/st0`:

```sh
xfsdump -0 -f /dev/st0 /mnt/xfs
```

### `xfsrestore`

The `xfsrestore` command is used to restore files from backups created by `xfsdump`. It supports both full and incremental restores.

#### Basic Syntax

```sh
xfsrestore [options] <restore-device> <filesystem>
```

- **`[options]`**: Command-line options for restoration.
- **`<restore-device>`**: The source for the backup (e.g., `/dev/st0` for tape).
- **`<filesystem>`**: The directory to restore files to (e.g., `/mnt/xfs`).

#### Common Options

- **`-f`**: Specify the file or device from which to restore.
- **`-r`**: Restore the file system to a specified directory.
- **`-v`**: Verbose mode, displays detailed information during restoration.

#### Example

To restore a backup from tape `/dev/st0` to `/mnt/xfs`:

```sh
xfsrestore -f /dev/st0 /mnt/xfs
```

### Summary

- **`xfs_info`**: Displays detailed information about an XFS filesystem.
- **`xfs_check`**: Deprecated; previously used for checking filesystem integrity.
- **`xfs_repair`**: Repairs and checks the integrity of XFS filesystems.
- **`xfsdump`**: Creates backups of XFS filesystems, supports full and incremental backups.
- **`xfsrestore`**: Restores files from backups created by `xfsdump`.

These tools are essential for managing, backing up, and repairing XFS filesystems, providing comprehensive support for both maintenance and disaster recovery operations.
