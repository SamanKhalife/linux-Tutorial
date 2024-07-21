# /etc/auto.master
The `/etc/auto.master` file is a configuration file for the `autofs` service on Unix-like systems. `autofs` is a service used to automatically mount filesystems when they are accessed and to unmount them when they are no longer needed. This can be useful for managing network shares and other filesystems without manually mounting them.

### Structure and Purpose of `/etc/auto.master`

The `/etc/auto.master` file defines the mount points and their associated maps (or configuration files) for `autofs`. It specifies where to look for additional map files that define the details of the filesystems to be mounted.

#### File Format

The `/etc/auto.master` file consists of a list of lines where each line represents a mount point and the map file associated with it. The basic syntax is:

```
<mount_point> <map_file> [options]
```

- **`<mount_point>`**: The directory where the filesystem or share will be mounted.
- **`<map_file>`**: The path to the map file or configuration file that contains the details of what should be mounted at `<mount_point>`.
- **`[options]`**: Optional parameters to control the behavior of `autofs` for that mount point.

### Example `/etc/auto.master` Entries

#### Basic Example

```plaintext
/mnt/auto /etc/auto.misc
```

In this example:
- `/mnt/auto` is the mount point where filesystems will be mounted.
- `/etc/auto.misc` is the map file that contains the details for the filesystems to be mounted under `/mnt/auto`.

#### Example with Options

```plaintext
/home /etc/auto.home --timeout=60 --ghost
```

In this example:
- `/home` is the mount point.
- `/etc/auto.home` is the map file.
- `--timeout=60` specifies that filesystems should be unmounted after 60 seconds of inactivity.
- `--ghost` ensures that empty directories are created for the mount points even if the filesystem is not currently mounted.

### Related Map Files

The map files specified in `/etc/auto.master` define what should be mounted at each mount point. For example, the `/etc/auto.misc` file might contain entries like:

```plaintext
# /etc/auto.misc
share -fstype=nfs,rw,soft server:/path/to/share
```

In this map file:
- `share` is the directory under `/mnt/auto` where the NFS share will be mounted.
- `-fstype=nfs,rw,soft` specifies the filesystem type (NFS), read-write mode, and soft mount options.
- `server:/path/to/share` is the network path to the NFS share.

### Key Points

- **Dynamic Mounting**: `autofs` dynamically mounts filesystems when accessed and unmounts them after a period of inactivity.
- **Flexibility**: Allows for easy configuration of multiple mount points and filesystems.
- **Map Files**: Define specific details for each mount point, such as filesystem type and mount options.

### Summary

The `/etc/auto.master` file is a critical component in configuring the `autofs` service. It provides a centralized location for defining mount points and their associated map files. Proper configuration ensures efficient and dynamic management of filesystems and network shares, reducing the need for manual mounting and unmounting.
