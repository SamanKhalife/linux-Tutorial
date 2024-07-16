# pulledpork.pl

**pulledpork.pl** is a popular Perl script used to manage Snort rules. It automates the process of downloading, processing, and managing Snort rule sets from various sources. PulledPork helps ensure that your Snort installation has the most up-to-date rules, which is crucial for maintaining an effective intrusion detection and prevention system.

### Key Features of PulledPork

1. **Rule Management**: PulledPork downloads and processes Snort rule sets, including rules from the Snort community, Emerging Threats, and other sources.

2. **Rule Categorization**: It categorizes and sorts rules, enabling better organization and management of the rule sets.

3. **Automatic Updates**: PulledPork can be scheduled to run periodically, ensuring that Snort rules are always current.

4. **Rule Merging and Disabling**: It merges new rules with existing ones and can disable certain rules based on user-defined criteria, reducing false positives and tuning the rule set to your specific environment.

5. **Compatibility**: Supports multiple versions of Snort and various rule sources, making it a versatile tool for rule management.

6. **Logging and Reporting**: Provides detailed logs and reports of its operations, helping administrators understand what changes were made and if any issues occurred during the update process.

### Installation

#### Prerequisites

Ensure you have Perl installed on your system. You may also need additional Perl modules, such as LWP::UserAgent, Crypt::SSLeay, and others.

#### Steps to Install PulledPork

1. **Download PulledPork**: Download the latest version from the official GitHub repository or other trusted sources.

   ```bash
   git clone https://github.com/shirkdog/pulledpork.git
   ```

2. **Navigate to the PulledPork Directory**:

   ```bash
   cd pulledpork
   ```

3. **Install Required Perl Modules**:

   ```bash
   cpan install LWP::UserAgent Crypt::SSLeay
   ```

4. **Configure PulledPork**: Copy the sample configuration file and edit it to suit your environment.

   ```bash
   cp etc/pulledpork.conf /etc/snort/pulledpork.conf
   nano /etc/snort/pulledpork.conf
   ```

   Update the configuration file with your Oinkcode (if using Snort rules) and other settings specific to your environment.

### Basic Usage

#### Downloading and Updating Rules

To run PulledPork and update Snort rules, use the following command:

```bash
perl pulledpork.pl -c /etc/snort/pulledpork.conf -l
```

This command downloads the latest rules, processes them, and updates your Snort configuration.

#### Command-line Options

- **`-c <config_file>`**: Specify the path to the configuration file.
- **`-l`**: Enable verbose logging.
- **`-v`**: Show the version of PulledPork.
- **`-V <snort_version>`**: Specify the Snort version (useful if you have multiple Snort versions).

### Configuration File (`pulledpork.conf`)

Here are some key configuration options in the `pulledpork.conf` file:

- **`rule_url`**: URLs for downloading Snort rules.
- **`ignore`**: List of rule IDs to ignore.
- **`local_rules`**: Path to local rules file.
- **`sid_msg_map`**: Path to the `sid-msg.map` file.
- **`distro`**: Distribution type (e.g., `ubuntu`, `centos`).
- **`snort_path`**: Path to the Snort binary.
- **`oinkcode`**: Your Oinkcode from Snort.org.

Example configuration snippet:

```ini
rule_url=https://www.snort.org/rules/snortrules-snapshot-29120.tar.gz|oinkcode
rule_url=https://rules.emergingthreats.net/open/suricata/emerging.rules.tar.gz|open
ignore=1:9999999,1:1111111
local_rules=/etc/snort/rules/local.rules
sid_msg_map=/etc/snort/sid-msg.map
distro=ubuntu
snort_path=/usr/local/bin/snort
oinkcode=your_oinkcode_here
```

### Scheduling PulledPork

To keep your Snort rules up-to-date, you can schedule PulledPork to run periodically using cron jobs.

Example cron job to run PulledPork daily at midnight:

```bash
0 0 * * * /usr/bin/perl /path/to/pulledpork.pl -c /etc/snort/pulledpork.conf -l
```

### Security Considerations

- **Secure Configuration**: Ensure the configuration file is securely stored and accessible only to authorized users.
- **Regular Updates**: Schedule regular updates to maintain the effectiveness of your intrusion detection/prevention system.

### Conclusion

PulledPork is a vital tool for managing Snort rules, automating the process of downloading, updating, and organizing rule sets. By using PulledPork, administrators can ensure that their Snort deployment remains effective against the latest threats with minimal manual intervention. For specific deployment scenarios, advanced configuration options, and best practices, consulting the official PulledPork documentation and community resources is recommended.
