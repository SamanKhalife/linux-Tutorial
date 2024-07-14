# sysctl

The `sysctl` command in Linux is used to view and modify kernel parameters at runtime. These parameters are used to configure the kernel and system behavior, covering aspects such as network settings, virtual memory management, and various kernel subsystems.

### Overview of `sysctl`

#### Purpose

`sysctl` is primarily used for:
- Viewing current kernel parameters.
- Modifying kernel parameters temporarily at runtime.
- Making permanent changes to kernel parameters by editing configuration files.

### Basic Usage

The general syntax for `sysctl` is:

```bash
sysctl [options] [variable[=value]]
```

- `[options]`: Various options to control the behavior of `sysctl`.
- `[variable]`: The kernel parameter to view or modify.
- `[value]`: The value to set the kernel parameter to.

### Common `sysctl` Commands

1. **Viewing Kernel Parameters**

   To view the value of a specific kernel parameter, use:

   ```bash
   sysctl variable
   ```

   Example:
   ```bash
   sysctl net.ipv4.ip_forward
   ```

   This command displays whether IP forwarding is enabled (`1`) or disabled (`0`).

   To view all kernel parameters and their values, use:

   ```bash
   sysctl -a
   ```

2. **Setting Kernel Parameters**

   To set a kernel parameter temporarily, use:

   ```bash
   sysctl variable=value
   ```

   Example:
   ```bash
   sudo sysctl net.ipv4.ip_forward=1
   ```

   This command enables IP forwarding.

3. **Persisting Kernel Parameters**

   To make changes permanent, add the parameter and value to the `/etc/sysctl.conf` file or a file in the `/etc/sysctl.d/` directory.

   Example:
   ```bash
   echo "net.ipv4.ip_forward = 1" | sudo tee -a /etc/sysctl.conf
   sudo sysctl -p
   ```

   The `-p` option reloads the `/etc/sysctl.conf` file to apply changes.

4. **Reloading Configuration Files**

   To reload settings from `/etc/sysctl.conf` or a specific configuration file, use:

   ```bash
   sudo sysctl -p [file]
   ```

   Example:
   ```bash
   sudo sysctl -p /etc/sysctl.d/99-custom.conf
   ```

5. **Writing Directly to `/proc/sys/`**

   Kernel parameters can also be set by writing directly to the corresponding files in the `/proc/sys/` directory.

   Example:
   ```bash
   echo 1 | sudo tee /proc/sys/net/ipv4/ip_forward
   ```

### Commonly Used Kernel Parameters

1. **Networking Parameters**

   - `net.ipv4.ip_forward`: Enable or disable IP forwarding.
   - `net.ipv4.conf.all.rp_filter`: Enable or disable reverse path filtering.
   - `net.core.somaxconn`: Set the maximum number of pending connections.

2. **Virtual Memory Parameters**

   - `vm.swappiness`: Set the kernel's swappiness value.
   - `vm.overcommit_memory`: Control the kernel's memory overcommit behavior.
   - `vm.dirty_ratio`: Set the maximum amount of system memory that can be filled with dirty pages.

3. **File System Parameters**

   - `fs.file-max`: Set the maximum number of open file descriptors.
   - `fs.inotify.max_user_watches`: Set the maximum number of inotify watches per user.

### Example Configurations

1. **Enable IP Forwarding**

   Temporary:
   ```bash
   sudo sysctl net.ipv4.ip_forward=1
   ```

   Permanent:
   ```bash
   echo "net.ipv4.ip_forward = 1" | sudo tee -a /etc/sysctl.conf
   sudo sysctl -p
   ```

2. **Increase Maximum Number of Open File Descriptors**

   Temporary:
   ```bash
   sudo sysctl fs.file-max=100000
   ```

   Permanent:
   ```bash
   echo "fs.file-max = 100000" | sudo tee -a /etc/sysctl.conf
   sudo sysctl -p
   ```

### Conclusion

`sysctl` is a powerful tool for managing kernel parameters in Linux, providing both temporary and permanent configuration options. By understanding and using `sysctl`, administrators can fine-tune the system's behavior to meet specific needs and improve performance or security.
