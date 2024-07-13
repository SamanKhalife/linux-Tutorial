# delv

`delv` is a command-line utility used for DNS troubleshooting and debugging. It is part of the BIND (Berkeley Internet Name Domain) DNS software suite and provides a way to perform DNS lookups and validate DNS responses.

### Overview of `delv`

#### Purpose

The main purpose of `delv` is to perform DNS queries and validate DNSSEC-signed zones. It allows administrators and users to manually inspect DNS records, troubleshoot DNS resolution issues, and verify the integrity of DNS responses using DNSSEC.

### Basic Usage

The general syntax for `delv` is:

```bash
delv [options] [@server] <name> [<type> [<class>]]
```

- `[options]`: Various options to control the behavior of `delv`.
- `[@server]`: Optional parameter specifying the DNS server to query.
- `<name>`: The domain name to query.
- `<type>`: Optional DNS record type to query (default is `A` for IPv4 address).
- `<class>`: Optional DNS class (default is `IN` for Internet).

#### Example Usages

1. **Basic Query**:

   To query an `A` record for a domain (e.g., `example.com`) using a specific DNS server (`8.8.8.8`):

   ```bash
   delv @8.8.8.8 example.com A
   ```

2. **Query with DNSSEC Validation**:

   To perform a DNSSEC-enabled query for an `A` record for a domain (e.g., `example.com`):

   ```bash
   delv -D example.com A
   ```

   This command validates DNSSEC signatures in the responses received.

3. **Query for Other Record Types**:

   To query for other types of DNS records, such as `MX` (Mail Exchange) records:

   ```bash
   delv example.com MX
   ```

### Features and Options

- **DNSSEC Validation (`-D`)**: Enables DNSSEC validation for the queried records.
- **Verbose Output (`-v`)**: Provides more detailed information about the DNS query and response.
- **Specifying DNS Server (`@server`)**: Allows querying a specific DNS server instead of using the system's default resolver.
- **Querying Specific Record Types**: Supports querying for various DNS record types like `A`, `AAAA`, `MX`, `TXT`, `NS`, etc.
- **Debugging and Troubleshooting**: Useful for debugging DNS issues, checking DNS configuration, and verifying DNS responses.

### Integration with DNSSEC

`delv` is particularly useful for validating DNSSEC-signed zones. By enabling DNSSEC validation (`-D` option), `delv` checks the authenticity and integrity of DNS responses using digital signatures.

### Conclusion

`delv` is a versatile tool for DNS troubleshooting and debugging, providing administrators and users with the ability to manually query DNS servers, inspect DNS records, and validate DNSSEC-signed zones. Understanding its usage and options is valuable for diagnosing and resolving DNS-related issues effectively.
