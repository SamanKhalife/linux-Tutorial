# openvassd

The `openvassd` command is the daemon process for OpenVAS (Open Vulnerability Assessment System). This daemon handles the actual scanning tasks, managing the execution of Network Vulnerability Tests (NVTs). It listens for scan requests, executes the scans, and returns the results.

#### Basic Usage

To start the `openvassd` daemon, you can use the following command:

```bash
sudo openvassd
```

This command starts the OpenVAS Scanner daemon, which will begin listening for scan requests.

#### Common Usage and Options

1. **Start Daemon**:
    ```bash
    sudo openvassd
    ```

2. **Configuration File**:
    You can specify a custom configuration file using the `-c` option:
    ```bash
    sudo openvassd -c /path/to/custom/openvassd.conf
    ```

3. **Log File**:
    To specify a log file for `openvassd`, use the `-l` option:
    ```bash
    sudo openvassd -l /var/log/openvas/openvassd.log
    ```

4. **Verbose Output**:
    For more detailed output, you can run `openvassd` in verbose mode with the `-v` option:
    ```bash
    sudo openvassd -v
    ```

5. **Help**:
    To see a list of available options and usage information, use the `-h` option:
    ```bash
    openvassd -h
    ```

#### Example

Here’s an example of starting `openvassd` with a custom configuration file and logging to a specific file:

```bash
sudo openvassd -c /etc/openvas/openvassd.conf -l /var/log/openvas/openvassd.log
```

#### Configuration File

The default configuration file for `openvassd` is typically located at `/etc/openvas/openvassd.conf`. This file includes various settings that control the behavior of the OpenVAS Scanner daemon.

Example contents of `openvassd.conf`:

```text
# Host and port for openvassd to listen on
host = 127.0.0.1
port = 9391

# Directory for log files
logfile = /var/log/openvas/openvassd.log

# Path to the NVTs directory
plugins_folder = /var/lib/openvas/plugins

# Other configuration settings...
```

#### Managing the Daemon

On systems with systemd, you can manage the `openvassd` service with the following commands:

- **Start the service**:
  ```bash
  sudo systemctl start openvas-scanner
  ```

- **Stop the service**:
  ```bash
  sudo systemctl stop openvas-scanner
  ```

- **Enable the service to start on boot**:
  ```bash
  sudo systemctl enable openvas-scanner
  ```

- **Check the status of the service**:
  ```bash
  sudo systemctl status openvas-scanner
  ```

#### Security Considerations

1. **Access Control**: Ensure that only authorized users can start and stop the `openvassd` daemon.
2. **Secure Configuration**: Regularly review the configuration file and ensure it adheres to your security policies.
3. **Log Monitoring**: Regularly monitor the log files for any suspicious activity or errors.

#### Conclusion

The `openvassd` daemon is a core component of the OpenVAS framework, responsible for executing vulnerability scans. Understanding how to configure, start, and manage `openvassd` is crucial for maintaining an effective vulnerability assessment system. By keeping the daemon properly configured and up to date, you ensure that your network scans are reliable and secure.
