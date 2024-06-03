# /usr/lib/systemd/user/

The `/usr/lib/systemd/user/` directory is part of the `systemd` system and service manager on Linux systems. It contains system-wide unit files specifically for user services and targets. Unlike system-level services managed by `root`, user services are managed by regular users without requiring elevated privileges.

### Overview of `/usr/lib/systemd/user/`

The `/usr/lib/systemd/user/` directory contains unit files that define system-wide user services and targets. These services are typically started and managed by individual users rather than the system administrator (`root`). 

#### Common Directories and Files

- **Unit Files**: Files with `.service`, `.socket`, `.target`, etc., extensions that define the behavior of user services.
- **`default.target.wants/`**: Directory containing symlinks to unit files that are wanted by the default user target.
- **`graphical.target.wants/`**: Directory containing symlinks to unit files that are wanted by the graphical user interface (GUI) target.

### User Services and Targets

User services are similar to system services but are managed by individual users. They can perform various tasks or run applications specific to the user's needs. Targets in the user context define states that the user session can transition to, such as graphical or multi-user mode.

#### Example Unit File

A typical `.service` unit file for a user service may look like this:

```ini
[Unit]
Description=Example User Service

[Service]
ExecStart=/usr/bin/example-service
Restart=always
User=myuser

[Install]
WantedBy=default.target
```

### Enabling and Disabling User Units

Users can enable or disable their own units without requiring root privileges. 

#### Enable a User Unit

To enable a user unit (e.g., `example-service.service`):

```sh
systemctl --user enable example-service.service
```

#### Disable a User Unit

To disable a user unit:

```sh
systemctl --user disable example-service.service
```

### Starting and Stopping User Units

Users can start, stop, restart, or check the status of their own units.

#### Start a User Unit

To start a user unit:

```sh
systemctl --user start example-service.service
```

#### Stop a User Unit

To stop a user unit:

```sh
systemctl --user stop example-service.service
```

#### Restart a User Unit

To restart a user unit:

```sh
systemctl --user restart example-service.service
```

#### Check Status of a User Unit

To check the status of a user unit:

```sh
systemctl --user status example-service.service
```

### Viewing and Editing Unit Files

#### Viewing User Unit Files

To view the content of a user unit file:

```sh
systemctl --user cat example-service.service
```

#### Editing User Unit Files

Users can edit their own unit files using a text editor:

```sh
nano ~/.config/systemd/user/example-service.service
```

After editing, reload the user's `systemd` configuration:

```sh
systemctl --user daemon-reload
```

### Conclusion

The `/usr/lib/systemd/user/` directory provides a centralized location for system-wide unit files for user services and targets managed by `systemd`. Users can enable, disable, start, stop, and manage their own services without requiring root privileges. Understanding how to utilize this directory is crucial for users who want to customize their session environment and manage their own services.
