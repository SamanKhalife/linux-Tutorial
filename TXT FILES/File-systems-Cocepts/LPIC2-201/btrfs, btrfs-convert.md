# btrfs, btrfs-convert

`btrfs` and `btrfs-convert` are commands used to manage Btrfs filesystems on Linux. Btrfs (B-Tree Filesystem) is a modern filesystem that provides advanced features such as snapshots, dynamic inode allocation, and integrated volume management.

### `btrfs` Command

The `btrfs` command is a versatile utility for managing Btrfs filesystems. It provides various subcommands for creating, managing, and querying Btrfs filesystems.

#### Basic Syntax

```sh
btrfs [subcommand] [options] [arguments]
```

- **`[subcommand]`**: Specific action to perform (e.g., `subvolume`, `device`, `filesystem`).
- **`[options]`**: Command-line options.
- **`[arguments]`**: Targets or additional parameters.

#### Common Subcommands

- **`btrfs check`**: Check and repair a Btrfs filesystem.

  ```sh
  btrfs check /dev/sdX1
  ```

  Options:
  - **`--repair`**: Attempt to repair the filesystem.
  - **`--readonly`**: Perform a read-only check.

- **`btrfs device`**: Manage devices in a Btrfs filesystem.

  ```sh
  btrfs device add /dev/sdX2 /mnt/btrfs
  ```

  Options:
  - **`add`**: Add a device to the filesystem.
  - **`delete`**: Remove a device from the filesystem.

- **`btrfs filesystem`**: Manage Btrfs filesystems.

  ```sh
  btrfs filesystem df /mnt/btrfs
  ```

  Options:
  - **`df`**: Show disk space usage.
  - **`show`**: Display filesystem status.

- **`btrfs subvolume`**: Manage subvolumes in a Btrfs filesystem.

  ```sh
  btrfs subvolume create /mnt/btrfs/subvolume
  ```

  Options:
  - **`create`**: Create a new subvolume.
  - **`delete`**: Delete an existing subvolume.

- **`btrfs send`**: Send a subvolume to another location or backup.

  ```sh
  btrfs send /mnt/btrfs/subvolume | ssh user@remote 'btrfs receive /mnt/remote'
  ```

- **`btrfs balance`**: Rebalance the data in a Btrfs filesystem.

  ```sh
  btrfs balance start /mnt/btrfs
  ```

  Options:
  - **`start`**: Start a rebalance operation.
  - **`status`**: Check the status of a rebalance operation.

### `btrfs-convert` Command

The `btrfs-convert` command is used to convert an existing ext3/ext4 filesystem to Btrfs. This tool allows for a transition from older filesystems to Btrfs without losing data.

#### Basic Syntax

```sh
btrfs-convert [options] <device>
```

- **`[options]`**: Command-line options.
- **`<device>`**: The device or partition to convert.

#### Common Options

- **`--rootfs`**: Convert the root filesystem.
- **`--no-compat`**: Disable compatibility features.
- **`--revert`**: Revert a Btrfs filesystem back to ext3/ext4 (requires a clean conversion and has some limitations).

#### Example

To convert an ext4 filesystem on `/dev/sdX1` to Btrfs:

```sh
btrfs-convert /dev/sdX1
```

### Important Considerations

- **Backup Data**: Always back up important data before performing conversions or significant changes with `btrfs` commands.
- **Filesystem Compatibility**: Ensure that the filesystem you are converting is not mounted or in use. You may need to boot from a live environment or unmount the filesystem.
- **Revertibility**: While `btrfs-convert` allows conversion from ext3/ext4 to Btrfs, reverting back is complex and may require manual intervention. Ensure you understand the implications of conversion.

### Summary

- **`btrfs` Command**: A comprehensive tool for managing Btrfs filesystems, offering capabilities such as checking, repairing, managing devices and subvolumes, and rebalancing data.
- **`btrfs-convert` Command**: Facilitates conversion from ext3/ext4 filesystems to Btrfs, enabling advanced features of Btrfs without data loss.

Both commands are essential for effectively managing and utilizing the Btrfs filesystem, offering powerful tools for modern filesystem management and migration.
