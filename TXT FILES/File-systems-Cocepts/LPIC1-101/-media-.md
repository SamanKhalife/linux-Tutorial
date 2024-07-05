# /media
The `/media` directory in Unix and Linux systems is used as a common mount point for removable media, such as USB drives, CDs, DVDs, and other external storage devices. When you insert a removable device, many Linux distributions automatically mount it under the `/media` directory, creating a subdirectory named after the device label or identifier.

### Overview of `/media`

- **Purpose**: To provide a standardized location for mounting removable media.
- **Usage**: Typically used by the operating system's automounter to mount devices when they are inserted.
- **Directory Structure**: Subdirectories within `/media` are often named after the label of the device or a unique identifier to distinguish between multiple devices.

### Example Structure

When you plug in a USB drive labeled `USB_DRIVE`, the automounter might create a directory like `/media/username/USB_DRIVE` and mount the drive there.

```sh
/media/
└── username/
    ├── USB_DRIVE/
    └── ANOTHER_DRIVE/
```

Here, `username` is your current user's name.

### Automatic Mounting

Many desktop environments (such as GNOME, KDE, and others) automatically mount removable devices to `/media` using tools like `udisks` or `gvfs`. This automatic mounting is convenient for everyday users who need to access removable media without manually running mount commands.

### Manual Mounting

While automatic mounting is convenient, you may sometimes need to manually mount a device to the `/media` directory. Here’s how you can do it:

#### Step-by-Step Example

1. **Identify the Device**:
   Use the `lsblk` or `blkid` command to find the device name.

    ```sh
    lsblk
    ```

    Output example:

    ```sh
    NAME   MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
    sda      8:0    0 465.8G  0 disk 
    ├─sda1   8:1    0  93.1G  0 part /
    ├─sda2   8:2    0     1K  0 part 
    └─sda5   8:5    0 372.7G  0 part /home
    sdb      8:16   1  14.9G  0 disk 
    └─sdb1   8:17   1  14.9G  0 part 
    ```

    Here, `/dev/sdb1` is the USB drive.

2. **Create a Mount Point**:
   Create a directory under `/media` where you want to mount the device.

    ```sh
    sudo mkdir -p /media/username/USB_DRIVE
    ```

3. **Mount the Device**:
   Use the `mount` command to mount the device to the created directory.

    ```sh
    sudo mount /dev/sdb1 /media/username/USB_DRIVE
    ```

4. **Verify the Mount**:
   Check that the device is mounted correctly.

    ```sh
    ls /media/username/USB_DRIVE
    ```

### Unmounting

To unmount a device, use the `umount` command followed by the mount point or the device name.

#### Example

Unmount the previously mounted USB drive:

```sh
sudo umount /media/username/USB_DRIVE
```

### Automatic Mounting Configuration

To configure automatic mounting behavior, you can use `/etc/fstab` for persistent mounts, or rely on desktop environment settings and automount tools.

#### Example `/etc/fstab` Entry

To automatically mount a USB drive with UUID `a1b2c3d4-e5f6-7890-1234-56789abcdef0` at boot:

```sh
UUID=a1b2c3d4-e5f6-7890-1234-56789abcdef0 /media/username/USB_DRIVE ext4 defaults 0 0
```

### Summary

The `/media` directory plays a crucial role in managing removable media in Unix and Linux systems. It provides a standardized location for mounting external devices, which can be handled automatically by the operating system or manually by the user. Understanding how to use and configure `/media` is essential for effectively managing storage devices and ensuring seamless access to external media.
