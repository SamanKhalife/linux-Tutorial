# hostnamectl

The hostnamectl command is a Linux command that can be used to manage the hostname of a system. The hostname is the name that is used to identify a system on a network.

Here are some examples of how to use the hostnamectl command:


```
# To show the current hostname:
hostnamectl

# To set the hostname to a static value:
hostnamectl set-hostname myhostname

# To set the hostname to a transient value:
hostnamectl set-hostname --transient myhostname

# To show the full hostname:
hostnamectl --full

# To show the domain name:
hostnamectl --domain
```

# help 

```
hostnamectl [OPTIONS...] COMMAND ...

Query or change system hostname.

Commands:
  status                 Show current hostname settings
  hostname [NAME]        Get/set system hostname
  icon-name [NAME]       Get/set icon name for host
  chassis [NAME]         Get/set chassis type for host
  deployment [NAME]      Get/set deployment environment for host
  location [NAME]        Get/set location for host

Options:
  -h --help              Show this help
     --version           Show package version
     --no-ask-password   Do not prompt for password
  -H --host=[USER@]HOST  Operate on remote host
  -M --machine=CONTAINER Operate on local container
     --transient         Only set transient hostname
     --static            Only set static hostname
     --pretty            Only set pretty hostname
     --json=pretty|short|off
                         Generate JSON output
```

## breakdown

```
-h, --help: This option shows this help message.
-s, --static: This option tells hostnamectl to set the hostname to a static value.
-i, --transient: This option tells hostnamectl to set the hostname to a transient value.
-f, --full: This option tells hostnamectl to print the full hostname.
-d, --domain: This option tells hostnamectl to print the domain name.
-q, --quiet: This option tells hostnamectl to do not output any message.
-V, --version: This option prints program version.
```

