# /proc/mounts

The `/proc/mounts` file in Linux provides information about all the mounted filesystems. It is a part of the `/proc` filesystem, which is a virtual filesystem that presents information about the system's processes and other kernel-related information.

### Structure and Contents

The contents of `/proc/mounts` are presented in a tabular format, with each line representing a mounted filesystem. The typical columns found in `/proc/mounts` are as follows:

1. **Device**: The device or special file associated with the filesystem.
2. **Mount Point**: The directory where the filesystem is mounted.
3. **Filesystem Type**: The type of filesystem (e.g., `ext4`, `nfs`, `tmpfs`).
4. **Mount Options**: Options used to mount the filesystem.
5. **Dump**: Used by the `dump` command to determine which filesystems need to be dumped.
6. **Pass**: Used by the `fsck` command to determine the order in which filesystems should be checked at boot time.

### Example Output

Here is an example of what the `/proc/mounts` file might look like:

```plaintext
/dev/sda1 / ext4 rw,relatime,errors=remount-ro 0 1
/dev/sda2 /home ext4 rw,relatime 0 2
tmpfs /run tmpfs rw,nosuid,nodev,noexec,relatime,size=8156160k,mode=755 0 0
tmpfs /dev/shm tmpfs rw,nosuid,nodev 0 0
```

### Columns Explained

- **Device**: This could be a physical device (like `/dev/sda1`), a virtual device (like `tmpfs`), or a network file system (like `nfs`).
- **Mount Point**: The directory where the filesystem is accessible. The root filesystem is typically mounted at `/`, while other filesystems may be mounted at directories like `/home` or `/run`.
- **Filesystem Type**: Indicates the type of filesystem, such as `ext4` for the ext4 filesystem, `tmpfs` for temporary filesystems, and `nfs` for Network File System.
- **Mount Options**: Options used during the mounting process. Common options include `rw` (read-write), `ro` (read-only), `relatime` (update inode access times relative to modify/change times), `nosuid` (disallow set-user-identifier or set-group-identifier bits), and others.
- **Dump**: A numeric value indicating whether the filesystem should be dumped (1 for yes, 0 for no).
- **Pass**: A numeric value indicating the order in which `fsck` should check filesystems at boot time. The root filesystem is usually checked first (pass 1), and other filesystems are checked afterwards (pass 2).

### Practical Uses

#### Checking Mounted Filesystems

To quickly see which filesystems are currently mounted, you can simply view the contents of `/proc/mounts`:

```bash
cat /proc/mounts
```

#### Using `/proc/mounts` in Scripts

Since `/proc/mounts` provides a reliable and up-to-date list of mounted filesystems, it is often used in scripts to check for specific mounts or to gather information about the filesystem state.

For example, to check if `/home` is mounted:

```bash
grep " /home " /proc/mounts
```

#### Monitoring Filesystem Mounts

System administrators can monitor `/proc/mounts` to detect changes in the mounted filesystems. This can be useful for debugging issues related to mounts or to ensure that critical filesystems are always mounted.

### Difference Between `/proc/mounts` and `/etc/mtab`

- **`/proc/mounts`**: This file is maintained by the kernel and always contains an up-to-date list of all mounted filesystems.
- **`/etc/mtab`**: This file is maintained by the `mount` command and other utilities. It may not always be up-to-date or accurate if mounts are performed by methods other than `mount`.

### Conclusion

The `/proc/mounts` file is a valuable resource for system administrators and scripts that need to obtain real-time information about the mounted filesystems. Understanding its structure and content allows for effective monitoring and management of filesystem mounts on a Linux system.
