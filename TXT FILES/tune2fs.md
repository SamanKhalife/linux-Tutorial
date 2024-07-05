# tune2fs
`**tune2fs**` is a command-line utility in Linux used to adjust various tunable filesystem parameters on ext2, ext3, and ext4 filesystems. This tool is part of the e2fsprogs package and is commonly used by system administrators to modify filesystem attributes after the filesystem has been created.

### Overview of `tune2fs`

#### Basic Syntax
```sh
tune2fs [options] <device>
```

- `<device>`: The device or partition containing the filesystem to be modified (e.g., `/dev/sda1`).

### Commonly Used Options

#### Adjusting Filesystem Parameters
- `-l` : Lists the current parameters of the filesystem.
- `-c max-mount-counts` : Sets the number of times the filesystem can be mounted before a consistency check is forced.
- `-i interval-between-checks[d|m|w]` : Sets the interval between forced filesystem checks.
- `-r reserved-blocks-count` : Sets the number of reserved filesystem blocks.
- `-m reserved-blocks-percentage` : Sets the percentage of the filesystem blocks reserved for the super-user.
- `-e errors-behavior` : Sets the behavior when an error is detected (`continue`, `remount-ro`, or `panic`).
- `-L volume-label` : Sets the volume label for the filesystem.
- `-U UUID` : Sets the universally unique identifier (UUID) for the filesystem.
- `-o mount-options` : Sets default mount options for the filesystem.
- `-E extended-options` : Sets extended options like `stripe`, `resize`, etc.

#### Example Usage

##### Listing Filesystem Parameters
To list all current parameters of the filesystem on `/dev/sda1`:
```sh
sudo tune2fs -l /dev/sda1
```

##### Setting Maximum Mount Count
To set the maximum mount count to 50 for `/dev/sda1`:
```sh
sudo tune2fs -c 50 /dev/sda1
```

##### Setting Interval Between Checks
To set the interval between forced filesystem checks to 30 days:
```sh
sudo tune2fs -i 30d /dev/sda1
```

##### Setting Reserved Blocks Count
To set the number of reserved blocks to 5000:
```sh
sudo tune2fs -r 5000 /dev/sda1
```

##### Setting Reserved Blocks Percentage
To set the reserved blocks percentage to 1%:
```sh
sudo tune2fs -m 1 /dev/sda1
```

##### Setting Error Behavior
To set the filesystem to remount as read-only on error:
```sh
sudo tune2fs -e remount-ro /dev/sda1
```

##### Setting Volume Label
To set the volume label to "MY_LABEL":
```sh
sudo tune2fs -L MY_LABEL /dev/sda1
```

##### Setting UUID
To set a new UUID for the filesystem:
```sh
sudo tune2fs -U random /dev/sda1
```

##### Setting Default Mount Options
To set default mount options to `noatime` and `nodiratime`:
```sh
sudo tune2fs -o noatime,nodiratime /dev/sda1
```

##### Setting Extended Options
To set an extended option such as `resize` (resize inode size):
```sh
sudo tune2fs -E resize=256 /dev/sda1
```

### Advanced Examples

#### Reducing Reserved Space
By default, ext2/3/4 filesystems reserve 5% of the space for the root user. This is useful on root filesystems to avoid running out of space completely, but on large storage partitions, you might want to reduce this:

To set the reserved blocks percentage to 1% on a large storage partition:
```sh
sudo tune2fs -m 1 /dev/sdb1
```

#### Setting Filesystem Check Intervals
You might want to ensure that your filesystem is checked for consistency every 3 months:
```sh
sudo tune2fs -i 3m /dev/sda1
```

#### Modifying Error Behavior
Setting the filesystem to panic on error can be useful for critical systems where data integrity is paramount:
```sh
sudo tune2fs -e panic /dev/sda1
```

### Summary

`**tune2fs**` is a versatile tool that allows administrators to optimize and control various aspects of ext2, ext3, and ext4 filesystems. Understanding and using `tune2fs` can help in maintaining filesystem integrity, optimizing performance, and ensuring efficient use of disk space.

