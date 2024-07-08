# update-rc.d
The `update-rc.d` command is used in Debian-based Linux distributions to manage runlevel configuration for services, particularly those using the SysV init system. Here's a detailed explanation of its purpose and usage:

### Purpose of `update-rc.d`

1. **Service Configuration**:
   - `update-rc.d` is used to install or remove System-V style init script links to the `/etc/init.d/` directory.
   - It allows administrators to control which services should be started or stopped automatically at different runlevels during system startup or shutdown.

2. **Runlevels**:
   - Runlevels represent different operating states of the system, such as single-user mode, multi-user mode with networking, and shutdown.
   - Services can be configured to start (`enable`) or not start (`disable`) at specific runlevels.

### Usage Examples

- **Viewing Current Configuration**:
  - To see the current runlevel configuration of a service:

    ```bash
    update-rc.d <service_name> defaults
    ```

    This command shows the current status of `<service_name>` in each runlevel.

- **Enabling a Service**:
  - To configure a service to start automatically at boot:

    ```bash
    update-rc.d <service_name> defaults
    ```

    This command creates the necessary symbolic links in the appropriate `/etc/rc*.d/` directories to start `<service_name>` in the default runlevels.

- **Disabling a Service**:
  - To prevent a service from starting automatically at boot:

    ```bash
    update-rc.d -f <service_name> remove
    ```

    This command removes the symbolic links from the `/etc/rc*.d/` directories that start `<service_name>`.

### `update-rc.d` with `systemd`

- **Compatibility**:
  - `update-rc.d` is primarily used with SysV init scripts. However, modern Debian-based distributions often use `systemd` as the default init system.
  - `systemd` provides backward compatibility with SysV init scripts and can manage services using `systemctl` commands (`systemctl enable`, `systemctl start`, etc.).

### Conclusion

`update-rc.d` is a useful tool for managing services in Debian-based Linux distributions that use the SysV init system. It provides a straightforward way to configure services to start or stop automatically at different runlevels. As Linux distributions continue to evolve, administrators may encounter both SysV init and `systemd` systems, each requiring different management commands. For `systemd`-based systems, it's recommended to use `systemctl` for managing services for better integration and compatibility with modern Linux features.
