# systemd-delta
The command `systemd-delta` is used in Linux systems to compare and display differences between the distribution-provided systemd unit files and any local modifications made by the system administrator. Here’s a detailed overview of `systemd-delta` and its usage:

### Purpose of `systemd-delta`

1. **Configuration Management**:
   - `systemd-delta` helps administrators manage configuration changes by showing differences between the original unit files provided by the distribution and any local modifications made.

2. **Troubleshooting**:
   - It aids in troubleshooting systemd-related issues by highlighting changes that could affect service behavior, startup sequences, dependencies, or other critical configurations.

### Usage

- **Listing Differences**:
  - Running `systemd-delta` without arguments lists all systemd unit files where local modifications exist compared to the distribution-provided versions.

  ```bash
  systemd-delta
  ```

- **Viewing Specific Unit Changes**:
  - To view differences for a specific unit file (e.g., `sshd.service`), specify the unit name after `systemd-delta`.

  ```bash
  systemd-delta sshd.service
  ```

### Example Output

When you run `systemd-delta`, it typically displays output similar to the following:

```
[EXTENDED] 0 overridden configuration files found.
[EXTENDED] No override file found for unit 'systemd-resolved.service'.
[EXTENDED] No drop-in files found for 'systemd-resolved.service'.
[EXTENDED] No override file found for unit 'sshd.service'.
1 overridden configuration files found.
Use 'systemctl cat <unit>' to see them.
```

### Interpretation

- **No Overrides**: Indicates units where no local modifications (`override`) or customizations (`drop-in files`) have been made compared to the distribution-provided files.

- **Overrides Found**: Indicates units where local modifications (`override`) or customizations (`drop-in files`) exist, potentially altering the default behavior defined by the distribution.

### Benefits

- **Configuration Transparency**: Provides clarity on which systemd unit files have been customized locally, aiding in system administration and maintenance.

- **Consistency**: Helps maintain consistency across systems by highlighting deviations from distribution-provided defaults.

### Conclusion

`systemd-delta` is a valuable tool for system administrators managing systemd-based Linux systems. By showing differences in unit file configurations, it promotes clarity, transparency, and effective management of system services and configurations. Using `systemd-delta` allows administrators to identify, review, and manage local modifications efficiently, ensuring system stability and adherence to operational requirements.
