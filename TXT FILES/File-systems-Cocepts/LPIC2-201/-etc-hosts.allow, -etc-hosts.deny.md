# /etc/hosts.allow , /etc/hosts.deny

The `/etc/hosts.allow` and `/etc/hosts.deny` files are part of the TCP Wrappers security system on Unix-like operating systems. They are used to control access to various network services by specifying which hosts are allowed or denied access.

#### Format and Usage

Both files use the same format for specifying rules. The basic syntax is:

```
service_list : host_list [: shell_command]
```

- **service_list**: A list of services (daemons) to which the rule applies. Wildcards like `ALL` can be used to specify all services.
- **host_list**: A list of hostnames, IP addresses, or network addresses. Wildcards like `ALL` can be used to specify all hosts.
- **shell_command**: An optional command to be executed when a rule matches.

#### /etc/hosts.allow

The `/etc/hosts.allow` file specifies which hosts are allowed to access which services. Rules in this file are evaluated first.

Example:

```sh
# Allow all services from the local network
ALL : 192.168.1.0/24

# Allow SSH from a specific IP
sshd : 203.0.113.5
```

#### /etc/hosts.deny

The `/etc/hosts.deny` file specifies which hosts are denied access to services. Rules in this file are evaluated after `/etc/hosts.allow`.

Example:

```sh
# Deny all services from everywhere (default deny)
ALL : ALL

# Deny telnet from a specific IP
telnetd : 203.0.113.10
```

#### Example Configuration

A typical configuration might involve allowing specific hosts access to certain services while denying all others:

**/etc/hosts.allow**:

```sh
# Allow SSH from the local network
sshd : 192.168.1.0/24

# Allow HTTP from a specific IP
httpd : 203.0.113.5
```

**/etc/hosts.deny**:

```sh
# Deny all other access
ALL : ALL
```

In this example, SSH access is allowed from the local network, HTTP access is allowed from a specific IP, and all other access is denied.

#### Wildcards and Patterns

- **ALL**: Matches all services or hosts.
- **KNOWN**: Matches any host whose name and address are known.
- **UNKNOWN**: Matches any host whose name or address is unknown.
- **PARANOID**: Matches any host whose name does not match its address.

Example using wildcards:

```sh
# Allow all services from the local machine
ALL : localhost

# Deny access to all services from unknown hosts
ALL : UNKNOWN
```

#### Logging and Commands

You can include shell commands to be executed when a rule matches. This can be useful for logging or triggering alerts.

Example:

```sh
# Log and deny all access from a specific IP
ALL : 203.0.113.10 : echo "Access attempt from 203.0.113.10" >> /var/log/tcpwrappers.log
```

#### Conclusion

The `/etc/hosts.allow` and `/etc/hosts.deny` files provide a flexible way to control access to network services using TCP Wrappers. By understanding and utilizing these files, administrators can enhance the security of their systems by allowing or denying access to services based on specific criteria.
