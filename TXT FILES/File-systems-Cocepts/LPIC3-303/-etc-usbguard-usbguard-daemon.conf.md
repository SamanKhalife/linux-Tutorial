# /etc/usbguard/usbguard-daemon.conf
The `/etc/usbguard/usbguard-daemon.conf` file is used to configure the USBGuard daemon. USBGuard is a software framework for implementing USB device authorization policies, which helps to protect your system against unauthorized USB devices. By defining rules and configuring the daemon, you can control which USB devices are allowed or denied access to your system.

### Overview of `/etc/usbguard/usbguard-daemon.conf`

#### Purpose

The `usbguard-daemon.conf` file provides configuration options for the USBGuard daemon, including settings for logging, rule files, device enumeration, and more. This configuration file allows administrators to customize the behavior of the USBGuard daemon to suit their security needs.

### Basic Format

The general format of the `/etc/usbguard/usbguard-daemon.conf` file is:

```plaintext
# Comment
option = value
```

Each line in the file specifies a configuration option and its value. Comments can be added using the `#` character.

### Example Entries and Their Descriptions

Here are some common configuration options and their descriptions:

1. **Logging Options**

   - **LogLevel**: Sets the logging level for the daemon. Possible values are `TRACE`, `DEBUG`, `INFO`, `WARNING`, `ERROR`, and `FATAL`.
     ```plaintext
     LogLevel = INFO
     ```

   - **LogFile**: Specifies the file where logs will be written.
     ```plaintext
     LogFile = /var/log/usbguard/usbguard-daemon.log
     ```

2. **Rule Options**

   - **RuleFile**: Specifies the file containing USB authorization rules.
     ```plaintext
     RuleFile = /etc/usbguard/rules.conf
     ```

3. **Device Enumeration Options**

   - **DeviceEnumeration**: Controls whether devices are enumerated on daemon startup. Possible values are `smart`, `yes`, and `no`.
     ```plaintext
     DeviceEnumeration = smart
     ```

4. **Policy Options**

   - **ImplicitPolicyTarget**: Defines the default policy target for devices that do not match any rules. Possible values are `allow` and `block`.
     ```plaintext
     ImplicitPolicyTarget = block
     ```

5. **IPC Options**

   - **IPCAllowedGroups**: Specifies the groups that are allowed to communicate with the daemon via the IPC interface.
     ```plaintext
     IPCAllowedGroups = usbguard
     ```

   - **IPCAccessControlFiles**: Specifies access control files for IPC.
     ```plaintext
     IPCAccessControlFiles = /etc/usbguard/IPCAccessControl.conf
     ```

### Example Configuration

Here is an example of a complete `/etc/usbguard/usbguard-daemon.conf` file:

```plaintext
# USBGuard daemon configuration file

# Set the logging level
LogLevel = INFO

# Specify the log file
LogFile = /var/log/usbguard/usbguard-daemon.log

# Specify the rule file
RuleFile = /etc/usbguard/rules.conf

# Control device enumeration on daemon startup
DeviceEnumeration = smart

# Define the default policy target for unmatched devices
ImplicitPolicyTarget = block

# Specify the groups allowed to communicate with the daemon via IPC
IPCAllowedGroups = usbguard

# Specify access control files for IPC
IPCAccessControlFiles = /etc/usbguard/IPCAccessControl.conf
```

### Applying Changes

After modifying the `/etc/usbguard/usbguard-daemon.conf` file, restart the USBGuard daemon to apply the changes:

```bash
sudo systemctl restart usbguard
```

### Conclusion

The `/etc/usbguard/usbguard-daemon.conf` file is a critical component for configuring the USBGuard daemon, allowing administrators to define how USB devices are handled and authorized. Proper configuration of this file enhances the security of your system by controlling USB device access based on specified policies.
