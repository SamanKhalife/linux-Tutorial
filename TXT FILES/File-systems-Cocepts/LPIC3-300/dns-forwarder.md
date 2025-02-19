# dns-forwarder

The `dns forwarder` parameter in Samba specifies the IP address of an external DNS server to which unresolved DNS queries will be forwarded. This setting is especially important in environments where Samba is acting as an Active Directory Domain Controller with its internal DNS service, yet needs to resolve queries beyond its local zones.

## Purpose

- **Extended DNS Resolution**:  
  Enables the Samba DNS server to forward queries it cannot resolve locally to another DNS server (such as a public DNS server or an enterprise DNS server).

- **Integration with External DNS**:  
  Provides a bridge between the internal domain DNS records and the wider internet or corporate DNS infrastructure.

- **Improved Network Functionality**:  
  Ensures that clients can resolve both local and external names, which is critical for seamless Active Directory operation and user connectivity.

## Configuration

The `dns forwarder` parameter is defined in the `[global]` section of your Samba configuration file (`smb.conf`). Here’s an example of how to configure it:

```ini
[global]
   workgroup = EXAMPLE
   realm = EXAMPLE.COM
   security = ADS
   encrypt passwords = yes

   # Internal DNS settings
   dns forwarder = 8.8.8.8
```

- **`dns forwarder = 8.8.8.8`**:  
  In this example, any DNS query that the Samba DNS server cannot resolve locally will be forwarded to Google's public DNS server at `8.8.8.8`. Replace this with the IP address of your preferred DNS server.

## How It Works

- **Query Resolution Process**:  
  When a client requests DNS resolution, Samba’s internal DNS server first checks its local zones. If it cannot resolve the query, it forwards the request to the DNS server specified by `dns forwarder`.

- **Fallback Mechanism**:  
  This ensures that even if the Samba server lacks complete information for external domains, clients still receive proper DNS responses.

## Use Cases

- **Active Directory Environments**:  
  Samba AD Domain Controllers often manage local domain records while relying on an external DNS server to resolve internet addresses or other non-local names.

- **Small Networks**:  
  In workgroup or small office environments using Samba’s internal DNS, setting a DNS forwarder can provide users with full DNS resolution capabilities without the need for a separate DNS infrastructure.

- **Hybrid Networks**:  
  Useful in mixed environments where local name resolution is managed internally by Samba, but queries for external resources need to be handled by an upstream DNS server.

## Considerations

- **Network Reliability**:  
  Ensure the DNS forwarder you configure is reliable and has high availability, as it becomes a critical component for resolving external names.

- **Security**:  
  Using an external DNS server means that your forwarded queries will be handled by that server. Verify that it meets your organization’s security and privacy requirements.

- **Multiple Forwarders**:  
  Some Samba versions support specifying multiple forwarders (using a space-separated list). This can provide redundancy in case one forwarder becomes unavailable.

## Conclusion

The `dns forwarder` parameter is essential for extending the DNS resolution capabilities of a Samba Active Directory Domain Controller that uses its internal DNS service. By forwarding unresolved queries to an external DNS server, it ensures comprehensive name resolution for both internal and external resources, thereby supporting seamless network operations and client connectivity.
