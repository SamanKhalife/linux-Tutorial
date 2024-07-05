# mount

The `mount` command in Unix and Linux is used to attach a filesystem to a specified directory in the directory tree. Conversely, the `umount` command detaches a filesystem. Understanding how to use these commands is crucial for managing filesystems, storage devices, and partitions.

### Basic Usage

The basic syntax for `mount` is:

```sh
mount [options] device dir
```

- **`device`**: The device or partition to mount (e.g., `/dev/sda1`).
- **`dir`**: The directory where the device is to be mounted.

### Mounting a Filesystem

To mount a filesystem, you need to specify the device and the directory where you want it to be mounted. You must have the necessary permissions (typically root) to perform the mount operation.

#### Example

Mounting `/dev/sda1` to `/mnt`:

```sh
sudo mount /dev/sda1 /mnt
```

### Viewing Mounted Filesystems

To view all currently mounted filesystems:

```sh
mount
```

Or using a more concise format:

```sh
findmnt
```

### Options

The `mount` command comes with various options that can be used to modify its behavior.

#### Common Options

- **`-t`** *type*: Specifies the filesystem type (e.g., `ext4`, `vfat`, `ntfs`).
- **`-o`** *options*: Specifies mount options, such as `rw` (read/write), `ro` (read-only), `noexec` (do not execute binaries), `nosuid` (ignore set-user-identifier or set-group-identifier bits), and many more.
- **`-a`**: Mount all filesystems mentioned in `/etc/fstab`.
- **`-r`**: Mount the filesystem as read-only.

#### Example

Mounting with specific options (read-only):

```sh
sudo mount -o ro /dev/sda1 /mnt
```

Mounting with a specified filesystem type:

```sh
sudo mount -t ext4 /dev/sda1 /mnt
```

### Unmounting a Filesystem

To unmount a filesystem, use the `umount` command followed by the mount point or the device name.

#### Example

Unmounting the filesystem mounted at `/mnt`:

```sh
sudo umount /mnt
```

Or using the device name:

```sh
sudo umount /dev/sda1
```

### Mounting Filesystems Automatically

Filesystems can be automatically mounted at boot time by adding an entry to the `/etc/fstab` file. Each line in `/etc/fstab` describes a filesystem to be mounted, using the following format:

```sh
device    mountpoint    fstype    options    dump    pass
```

#### Example `/etc/fstab` Entry

Mounting `/dev/sda1` to `/mnt` with read-only access:

```sh
/dev/sda1    /mnt    ext4    defaults,ro    0    2
```

### Examples with Explanations

#### Mounting a USB Drive

Assume you have a USB drive at `/dev/sdb1` and you want to mount it to `/media/usb`.

1. **Create the Mount Point**:

    ```sh
    sudo mkdir -p /media/usb
    ```

2. **Mount the USB Drive**:

    ```sh
    sudo mount /dev/sdb1 /media/usb
    ```

3. **Verify the Mount**:

    ```sh
    mount | grep /media/usb
    ```

#### Mounting an ISO File

To mount an ISO file (`example.iso`) to `/mnt/iso`:

1. **Create the Mount Point**:

    ```sh
    sudo mkdir -p /mnt/iso
    ```

2. **Mount the ISO File**:

    ```sh
    sudo mount -o loop example.iso /mnt/iso
    ```

3. **Verify the Mount**:

    ```sh
    mount | grep /mnt/iso
    ```

#### Unmounting a Network Filesystem

Assume you have a network filesystem mounted at `/mnt/nfs`:

```sh
sudo umount /mnt/nfs
```

### Practical Use Cases

- **Mounting External Drives**: Attach external storage devices like USB drives, external hard drives, or SD cards.
- **Network Storage**: Mount network file systems such as NFS (Network File System) or CIFS (Common Internet File System).
- **ISO Files**: Access the contents of ISO files without burning them to a disk.
- **Read-Only Mounts**: Mount filesystems in read-only mode for safety.

### Summary

The `mount` and `umount` commands are fundamental tools for managing filesystems in Unix and Linux systems. They allow you to attach and detach filesystems, providing flexibility in how storage devices and partitions are used. Understanding these commands helps in managing storage effectively, whether it's for daily use, troubleshooting, or system administration tasks.

# help

```
Usage:
 mount [-lhV]
 mount -a [options]
 mount [options] [--source] <source> | [--target] <directory>
 mount [options] <source> <directory>
 mount <operation> <mountpoint> [<target>]

Mount a filesystem.

Options:
 -a, --all               mount all filesystems mentioned in fstab
 -c, --no-canonicalize   don't canonicalize paths
 -f, --fake              dry run; skip the mount(2) syscall
 -F, --fork              fork off for each device (use with -a)
 -T, --fstab <path>      alternative file to /etc/fstab
 -i, --internal-only     don't call the mount.<type> helpers
 -l, --show-labels       show also filesystem labels
 -m, --mkdir[=<mode>]    alias to '-o X-mount.mkdir[=<mode>]'
 -n, --no-mtab           don't write to /etc/mtab
     --options-mode <mode>
                         what to do with options loaded from fstab
     --options-source <source>
                         mount options source
     --options-source-force
                         force use of options from fstab/mtab
 -o, --options <list>    comma-separated list of mount options
 -O, --test-opts <list>  limit the set of filesystems (use with -a)
 -r, --read-only         mount the filesystem read-only (same as -o ro)
 -t, --types <list>      limit the set of filesystem types
     --source <src>      explicitly specifies source (path, label, uuid)
     --target <target>   explicitly specifies mountpoint
     --target-prefix <path>
                         specifies path used for all mountpoints
 -v, --verbose           say what is being done
 -w, --rw, --read-write  mount the filesystem read-write (default)
 -N, --namespace <ns>    perform mount in another namespace

 -h, --help              display this help
 -V, --version           display version

Source:
 -L, --label <label>     synonym for LABEL=<label>
 -U, --uuid <uuid>       synonym for UUID=<uuid>
 LABEL=<label>           specifies device by filesystem label
 UUID=<uuid>             specifies device by filesystem UUID
 PARTLABEL=<label>       specifies device by partition label
 PARTUUID=<uuid>         specifies device by partition UUID
 ID=<id>                 specifies device by udev hardware ID
 <device>                specifies device by path
 <directory>             mountpoint for bind mounts (see --bind/rbind)
 <file>                  regular file for loopdev setup

Operations:
 -B, --bind              mount a subtree somewhere else (same as -o bind)
 -M, --move              move a subtree to some other place
 -R, --rbind             mount a subtree and all submounts somewhere else
 --make-shared           mark a subtree as shared
 --make-slave            mark a subtree as slave
 --make-private          mark a subtree as private
 --make-unbindable       mark a subtree as unbindable
 --make-rshared          recursively mark a whole subtree as shared
 --make-rslave           recursively mark a whole subtree as slave
 --make-rprivate         recursively mark a whole subtree as private
 --make-runbindable      recursively mark a whole subtree as unbindable

For more details see mount(8).
```
