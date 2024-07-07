# /etc/mtab

1. **Purpose**:
   - `/etc/mtab` (mounted table) is a file used to record currently mounted filesystems and their options on a Linux system.
   - It dynamically reflects the current state of mounted filesystems and is updated automatically whenever a filesystem is mounted or unmounted.

2. **Format**:
   - Each line in `/etc/mtab` contains information about a mounted filesystem:
     ```
     /dev/sda1 / ext4 rw,relatime 0 0
     ```
     - **Column 1**: Device or filesystem path.
     - **Column 2**: Mount point in the directory tree.
     - **Column 3**: Filesystem type (e.g., ext4, nfs).
     - **Column 4**: Mount options (e.g., rw, relatime).
     - **Column 5**: Dump frequency (used by `dump` backup utility).
     - **Column 6**: Filesystem check order (used by `fsck` utility).

3. **Usage**:
   - Provides a real-time snapshot of currently mounted filesystems.
   - Used by system utilities and commands like `mount`, `df`, and `umount` to determine the current state of mounted filesystems.

4. **Differences with `/etc/fstab`**:
   - **`/etc/fstab`**: Static file used to define filesystems and their mount options for automatic mounting at system startup.
   - **`/etc/mtab`**: Dynamic file updated in real-time to reflect currently mounted filesystems.

5. **Maintenance**:
   - Normally managed automatically by the system.
   - Editing `/etc/mtab` directly is not recommended; changes may not persist across reboots.

6. **Access**:
   - Read access is typically required by system utilities to determine filesystem usage and status.

### Conclusion

Understanding `/etc/mtab` provides insights into the current status of mounted filesystems on a Linux system. It serves as a vital reference for system administrators and utilities to manage and monitor filesystem usage dynamically. Always ensure to consult system documentation and use appropriate commands (`mount`, `umount`, `df`) to interact with mounted filesystems rather than editing `/etc/mtab` directly.
