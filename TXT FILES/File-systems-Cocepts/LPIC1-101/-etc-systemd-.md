# /etc/systemd

The `/etc/systemd` directory is a critical part of the `systemd` init system on Linux. It contains configuration files and subdirectories that control the behavior of system and service units. Understanding the structure and purpose of this directory is essential for effective system management.

### Overview of `/etc/systemd`

The `/etc/systemd` directory typically contains the following subdirectories:

- **`/etc/systemd/system/`**: Local unit files and custom unit configurations.
- **`/etc/systemd/user/`**: User-specific unit files and custom unit configurations.

### `/etc/systemd/system/`

The `/etc/systemd/system/` directory is where administrators place unit files to override the default ones provided by the system or create custom units. This directory takes precedence over the system's default unit files located in `/lib/systemd/system/`.

#### Common Directories and Files

- **Unit Files**: Files that define how services, sockets, targets, and other units are managed.
- **`multi-user.target.wants/`**: Directory containing symlinks to unit files that are wanted by the `multi-user.target`.
- **`graphical.target.wants/`**: Directory containing symlinks to unit files that are wanted by the `graphical.target`.

### Creating and Managing Unit Files

Unit files in `systemd` describe the configuration and behavior of services and other system resources. They typically have the extension `.service`, `.socket`, `.target`, etc.

#### Structure of a Unit File

A typical `.service` unit file contains sections like `[Unit]`, `[Service]`, and `[Install]`.

```ini
[Unit]
Description=My Custom Service
After=network.target

[Service]
ExecStart=/usr/bin/my-service
Restart=always
User=myuser

[Install]
WantedBy=multi-user.target
```

### Enabling and Disabling Units

Enabling a unit means creating symlinks in the appropriate `.wants` directory so that the unit starts automatically at boot.

#### Enable a Unit

To enable a unit (e.g., `my-service.service`):

```sh
sudo systemctl enable my-service.service
```

#### Disable a Unit

To disable a unit:

```sh
sudo systemctl disable my-service.service
```

### Reloading and Restarting Units

When changes are made to unit files, you need to reload the `systemd` configuration and potentially restart the affected service.

#### Reload `systemd` Configuration

To reload the `systemd` manager configuration:

```sh
sudo systemctl daemon-reload
```

#### Restart a Service

To restart a service:

```sh
sudo systemctl restart my-service.service
```

### Viewing and Editing Unit Files

#### Viewing Unit Files

To view the content of a unit file:

```sh
systemctl cat my-service.service
```

#### Editing Unit Files

To edit a unit file (using a text editor like `nano`):

```sh
sudo nano /etc/systemd/system/my-service.service
```

After editing, reload the `systemd` configuration:

```sh
sudo systemctl daemon-reload
```

### Overriding Default Unit Files

To override or extend the default unit files without modifying them directly, you can create drop-in files.

#### Creating Drop-In Files

To create a drop-in file for `my-service.service`:

```sh
sudo mkdir -p /etc/systemd/system/my-service.service.d
sudo nano /etc/systemd/system/my-service.service.d/override.conf
```

In the `override.conf` file, you can add or modify specific settings:

```ini
[Service]
Environment="MY_ENV_VAR=some_value"
```

### Managing User-Specific Units

The `/etc/systemd/user/` directory is used for user-specific unit files. These are typically managed by non-root users and allow user-level control over services.

#### Enabling User Units

To enable a user unit:

```sh
systemctl --user enable my-user-service.service
```

#### Starting User Units

To start a user unit:

```sh
systemctl --user start my-user-service.service
```

### Conclusion

The `/etc/systemd` directory is crucial for managing `systemd` configurations on a Linux system. It allows administrators to customize, enable, disable, and override system services and other units. Understanding how to navigate and utilize this directory is essential for effective system administration with `systemd`.
