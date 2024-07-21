# smartd, smartctl

**`smartd`** and **`smartctl`** are utilities used for monitoring and managing the health of hard drives using the Self-Monitoring, Analysis, and Reporting Technology (SMART) system. SMART is built into most modern hard drives and SSDs to detect and report various indicators of drive reliability, allowing for early warnings of potential failures.

### `smartctl`

`smartctl` is a command-line utility used to control and monitor SMART-enabled hard drives and SSDs. It can query and control SMART attributes, run tests, and display drive health information.

#### Basic Syntax

```sh
smartctl [options] <device>
```

- **`[options]`**: Various command-line options to perform different operations.
- **`<device>`**: The device name (e.g., `/dev/sda`).

#### Common Options

- **`-a`**: Print all SMART information available for the drive.
- **`-i`**: Print drive information.
- **`-t`**: Run a SMART self-test (e.g., `short`, `long`, `offline`).
- **`-H`**: Print the overall health status of the drive.
- **`-l`**: Print log of SMART events.

#### Example Commands

1. **Check SMART Information**

   ```sh
   smartctl -a /dev/sda
   ```

   This command displays detailed SMART information for the drive `/dev/sda`, including attributes, error logs, and self-test results.

2. **Run a Short Self-Test**

   ```sh
   smartctl -t short /dev/sda
   ```

   Initiates a short self-test, which checks the drive for errors. The results will be available after the test completes.

3. **Print Health Status**

   ```sh
   smartctl -H /dev/sda
   ```

   Displays the overall health status of the drive, indicating if it is considered healthy or not.

### `smartd`

`smartd` is the daemon (background service) component of the SMART monitoring suite. It periodically monitors the SMART attributes of the drives, performs tests, and sends notifications if it detects potential issues.

#### Configuration

The configuration file for `smartd` is usually located at `/etc/smartd.conf`. This file defines how `smartd` monitors the drives and what actions to take.

#### Example Configuration

Here’s an example of a `smartd.conf` configuration:

```conf
# Monitor /dev/sda and email on error
/dev/sda -a -M exec /usr/local/sbin/smartd_notify

# Monitor /dev/sdb with specific thresholds
/dev/sdb -a -S on -o on -s (S/../.././02|L/../.././07) -l error -l selftest
```

- **`/dev/sda -a`**: Monitor `/dev/sda` and report all SMART attributes.
- **`-M exec /usr/local/sbin/smartd_notify`**: Execute a notification script when an issue is detected.
- **`/dev/sdb -S on`**: Enable SMART monitoring for `/dev/sdb`.
- **`-s (S/../.././02|L/../.././07)`**: Schedule short self-tests every 2 days and long self-tests every 7 days.
- **`-l error`**: Log errors.
- **`-l selftest`**: Log self-test results.

#### Starting and Stopping `smartd`

- **Start the Service**

  ```sh
  systemctl start smartd
  ```

- **Stop the Service**

  ```sh
  systemctl stop smartd
  ```

- **Enable at Boot**

  ```sh
  systemctl enable smartd
  ```

- **Disable at Boot**

  ```sh
  systemctl disable smartd
  ```

### Summary

- **`smartctl`**: Command-line tool to check and manage SMART data, perform self-tests, and get health status of hard drives and SSDs.
- **`smartd`**: Daemon that continuously monitors SMART data, performs scheduled self-tests, and can send notifications about drive health issues.

Together, `smartctl` and `smartd` provide robust tools for monitoring the health and performance of storage devices, helping to predict and prevent data loss due to drive failures.
