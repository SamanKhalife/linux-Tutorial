# blkid
The `blkid` command in Unix and Linux is used to locate or print block device attributes. It is particularly useful for displaying the universally unique identifier (UUID) and file system type of block devices, which can be critical for system administration tasks such as mounting file systems and managing storage.

### Basic Usage

The basic syntax for `blkid` is:

```sh
blkid [options] [device...]
```

- **`device`**: The specific device to query. If omitted, `blkid` will list all available block devices.

### Examples

#### List All Block Devices with Attributes

To list all block devices with their attributes:

```sh
blkid
```

This will output details such as UUID, file system type, and label.

#### Query a Specific Device

To query a specific device, such as `/dev/sda1`:

```sh
blkid /dev/sda1
```

This will display the attributes of `/dev/sda1`.

### Options

- **`-c`** *cachefile*: Use the specified cache file.
- **`-g`**: Display devices without file systems.
- **`-k`**: Kernel-based detection.
- **`-L`** *label*: Find a device by its label.
- **`-U`** *uuid*: Find a device by its UUID.
- **`-o`** *output format*: Specify the output format (full, value, export, device).
- **`-p`**: Low-level probing mode.
- **`-s`** *tag*: Show only the value of the specified tag (e.g., UUID, TYPE).
- **`-t`** *token*: Display devices matching the specified token.
- **`-u`**: Update the cache file.
- **`-v`**: Enable verbose mode.

### Practical Use Cases

#### Identifying File System Types

To identify the file system type of a device:

```sh
blkid /dev/sdb1
```

This will show the type of file system on `/dev/sdb1`, such as ext4, xfs, etc.

#### Finding a Device by UUID

To find a device by its UUID:

```sh
blkid -U a1b2c3d4-e5f6-7890-1234-56789abcdef0
```

This will display the device that matches the specified UUID.

#### Finding a Device by Label

To find a device by its label:

```sh
blkid -L mylabel
```

This will display the device that has the label "mylabel".

### Examples with Explanations

#### Displaying All Attributes in Export Format

To display all attributes of block devices in a format suitable for importing into a shell script:

```sh
blkid -o export
```

- **`-o export`**: Outputs the attributes in an exportable format.

#### Showing Only the UUID of a Device

To show only the UUID of `/dev/sda1`:

```sh
blkid -s UUID -o value /dev/sda1
```

- **`-s UUID`**: Selects only the UUID attribute.
- **`-o value`**: Outputs only the value of the specified attribute.

#### Updating the Cache File

To update the cache file with the current information about all block devices:

```sh
blkid -u
```

- **`-u`**: Updates the cache file.

### Summary

The `blkid` command is a powerful utility for displaying and managing block device attributes in Unix and Linux systems. It provides essential information such as UUIDs and file system types, which are critical for system administration tasks. Mastering `blkid` helps in effectively managing storage devices, mounting file systems, and troubleshooting related issues.

Here is an example table of `blkid` output for a better understanding:

| Device    | UUID                                    | TYPE   | LABEL  |
|-----------|-----------------------------------------|--------|--------|
| /dev/sda1 | a1b2c3d4-e5f6-7890-1234-56789abcdef0    | ext4   | rootfs |
| /dev/sdb1 | b1c2d3e4-f5g6-7890-2345-6789abcdef01    | xfs    | data   |
