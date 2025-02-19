# klist

`klist` is a command-line utility for displaying Kerberos tickets held in a user's credential cache. It is used to view details about active Kerberos tickets, including their principal names, validity periods, and flags. This tool is essential for troubleshooting Kerberos authentication issues and ensuring that your Kerberos environment (used in Active Directory, Samba, etc.) is functioning correctly.

## Key Features

- **Display Credential Cache**:  
  Shows the current Kerberos tickets in the cache, including the Ticket Granting Ticket (TGT) and service tickets.

- **Ticket Details**:  
  Displays detailed information such as the default principal, ticket start and expiration times, and various ticket flags.

- **Multiple Cache Support**:  
  Allows you to specify a particular credential cache if more than one exists.

- **Script-Friendly**:  
  Options like `-s` (silent mode) enable integration in scripts to check for valid tickets.

## Basic Syntax

```bash
klist [options]
```

## Common Options

- **No Options**:  
  Running `klist` without any options displays the current default credential cache:
  ```bash
  klist
  ```
  Example output:
  ```
  Ticket cache: FILE:/tmp/krb5cc_1000
  Default principal: alice@EXAMPLE.COM

  Valid starting       Expires              Service principal
  05/01/2023 08:45:12  05/01/2023 18:45:12  krbtgt/EXAMPLE.COM@EXAMPLE.COM
  05/01/2023 08:50:30  05/01/2023 09:50:30  host/server.example.com@EXAMPLE.COM
  ```

- **`-c <cache_name>`**:  
  Specify a particular credential cache:
  ```bash
  klist -c /tmp/krb5cc_1000
  ```

- **`-f`**:  
  Display additional flags associated with each ticket (e.g., forwardable, renewable):
  ```bash
  klist -f
  ```

- **`-s`**:  
  Silent mode; useful in scripts. It returns a zero exit status if valid tickets are present, and a non-zero status otherwise:
  ```bash
  klist -s
  ```

- **`-V`**:  
  Verbose mode; provides extra details about each ticket:
  ```bash
  klist -V
  ```

## Use Cases

- **Troubleshooting Kerberos Authentication**:  
  If you suspect issues with Kerberos, such as expired tickets or failure to obtain a TGT, use `klist` to inspect your current credentials.

- **Verifying Ticket Renewal**:  
  Check the expiration times of your tickets to ensure that they will renew in time for long-running operations.

- **Scripting Checks**:  
  Use `klist -s` within scripts to verify that a valid ticket is present before proceeding with operations that require Kerberos authentication.

## Conclusion

`klist` is a vital tool in Kerberos environments, allowing administrators and users to inspect the status of their authentication tickets. Its straightforward output and various options make it indispensable for diagnosing and resolving Kerberos-related issues.
