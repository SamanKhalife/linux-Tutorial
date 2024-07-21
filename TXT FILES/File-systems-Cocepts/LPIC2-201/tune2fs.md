# tune2fs

`tune2fs` is a command-line utility used to adjust and manage parameters of ext2, ext3, and ext4 filesystems on Unix-like operating systems. It allows users to modify filesystem attributes such as mount options, reserved space, and filesystem labels without needing to reformat the filesystem.

### Key Features

- **Filesystem Parameter Tuning**: Adjust various parameters of ext2/ext3/ext4 filesystems.
- **Reserved Space Management**: Set aside a portion of the filesystem for the root user to prevent non-root users from filling the disk.
- **Filesystem Labeling**: Set or change filesystem labels for easier identification.
- **Filesystem Check Settings**: Configure automatic filesystem checks at boot time.

### Basic Usage

The general syntax for `tune2fs` is:

```sh
tune2fs [options] <device>
```

- **`[options]`**: Command-line options for modifying filesystem parameters.
- **`<device>`**: The disk partition or device to modify (e.g., `/dev/sdX1`).

### Common Options

- **`-l`**: List filesystem parameters.
- **`-m <percent>`**: Set the percentage of filesystem space reserved for the root user.
- **`-L <label>`**: Set the filesystem label.
- **`-e <errors>`**: Set the behavior when filesystem errors are detected (e.g., `continue`, `panic`, `remount-ro`).
- **`-c <count>`**: Set the number of mounts between filesystem checks.
- **`-i <bytes>`**: Set the number of bytes between inode table checks.
- **`-o <option>`**: Set or modify filesystem options (e.g., `discard` for SSDs).

### Examples

#### List Filesystem Parameters

To list the current parameters of a filesystem on `/dev/sda1`:

```sh
tune2fs -l /dev/sda1
```

#### Set Filesystem Label

To set the label of the filesystem on `/dev/sda1` to `mydata`:

```sh
tune2fs -L mydata /dev/sda1
```

#### Adjust Reserved Space

To set the reserved space percentage to 2% for the root user:

```sh
tune2fs -m 2 /dev/sda1
```

#### Configure Filesystem Check Interval

To set the filesystem to be checked every 50 mounts:

```sh
tune2fs -c 50 /dev/sda1
```

#### Configure Filesystem Check Interval by Time

To set the filesystem to be checked every 30 days:

```sh
tune2fs -i 30d /dev/sda1
```

### Important Considerations

- **Unmount Filesystem**: While many options can be applied to a mounted filesystem, it is generally safer to unmount the filesystem before making changes with `tune2fs`, especially for certain operations.
- **Reserved Space**: The reserved space is intended to prevent the filesystem from becoming completely full, which helps maintain system performance and prevents potential issues. Adjust the reserved space with care.
- **Filesystem Label**: Changing the filesystem label can be useful for identifying filesystems, especially in systems with multiple disks or partitions.

### Summary

`tune2fs` is a versatile tool for managing and tuning ext2, ext3, and ext4 filesystems. It allows users to adjust filesystem parameters, manage reserved space, and configure filesystem checks. By using `tune2fs`, system administrators can optimize filesystem performance, maintain system reliability, and make filesystem management more convenient.
