# samba-tool dns

The `samba-tool dns` command group is a set of Samba utilities for managing DNS records in an Active Directory environment. It is used to query, add, delete, update, and check DNS records on a DNS server (which may be Samba’s internal DNS server or an external one like BIND9 with DLZ). Maintaining accurate DNS records is critical for proper domain controller discovery, Kerberos authentication, and overall Active Directory functionality.

## Overview

In an AD environment, DNS records (including A, PTR, and SRV records) enable clients to locate domain controllers and services. The `samba-tool dns` commands allow administrators to:

- **Query DNS records** to verify current settings.
- **Add new records** when services are deployed or IP addresses change.
- **Delete obsolete records** that are no longer needed.
- **Update existing records** to reflect current network information.
- **Check zone integrity** to detect potential issues.

## Common Subcommands

### 1. `query`
- **Purpose**: Retrieve DNS records from a specific zone.
- **Usage**:
  ```bash
  samba-tool dns query <server> <zone> <node> [<record type>] [options]
  ```
- **Example**:
  ```bash
  samba-tool dns query dc.example.com example.com @ A -U Administrator
  ```
  This queries the A records for the root of the `example.com` zone on the DNS server `dc.example.com`.

### 2. `add`
- **Purpose**: Add a new DNS record.
- **Usage**:
  ```bash
  samba-tool dns add <server> <zone> <node> <record type> <record data> [options]
  ```
- **Example**:
  ```bash
  samba-tool dns add dc.example.com example.com www A 192.168.1.100 -U Administrator
  ```
  This adds an A record for `www.example.com` with the IP address `192.168.1.100`.

### 3. `delete`
- **Purpose**: Remove an existing DNS record.
- **Usage**:
  ```bash
  samba-tool dns delete <server> <zone> <node> <record type> <record data> [options]
  ```
- **Example**:
  ```bash
  samba-tool dns delete dc.example.com example.com www A 192.168.1.100 -U Administrator
  ```
  This deletes the specified A record from the zone.

### 4. `update`
- **Purpose**: Update an existing DNS record.
- **Usage**:
  ```bash
  samba-tool dns update <server> <zone> <node> <record type> <old record data> <new record data> [options]
  ```
- **Example**:
  ```bash
  samba-tool dns update dc.example.com example.com www A 192.168.1.100 192.168.1.101 -U Administrator
  ```
  This command updates the A record for `www.example.com` from the old IP to a new IP address.

### 5. `check`
- **Purpose**: Verify the integrity or health of DNS records within a zone.
- **Usage**:
  ```bash
  samba-tool dns check <server> <zone> [options]
  ```
- **Example**:
  ```bash
  samba-tool dns check dc.example.com example.com -U Administrator
  ```
  This checks the `example.com` zone for potential DNS issues.

### 6. `list`
- **Purpose**: List all DNS records for a given node or zone.
- **Usage**:
  ```bash
  samba-tool dns list <server> <zone> <node> [options]
  ```
- **Example**:
  ```bash
  samba-tool dns list dc.example.com example.com @ -U Administrator
  ```
  This lists all DNS records at the zone root of `example.com`.

## Additional Options

- **`-U <username>`**:  
  Specifies the administrator account for authentication.

- **`--verbose`**:  
  Provides detailed output for troubleshooting.

- **Other Options**:  
  Depending on your Samba version, additional options may be available to control output formatting or to target specific record attributes.

## Use Cases

- **Domain Controller Discovery**:  
  Ensuring that SRV records for domain controllers are accurate so that clients can locate them.

- **Service Registration**:  
  Adding or updating DNS records for new services or after network changes.

- **Zone Maintenance**:  
  Periodically querying and checking the DNS zone to confirm that records are up-to-date and correct.

## Conclusion

The `samba-tool dns` command group is essential for managing DNS records in a Samba Active Directory environment. Whether you need to query current records, add new ones, delete outdated entries, update existing records, or perform health checks, these subcommands provide the necessary functionality. Proper DNS management via `samba-tool dns` is key to ensuring reliable domain controller discovery, seamless Kerberos authentication, and overall network stability.
