# sssd.conf
The `sssd.conf` file is the main configuration file for the System Security Services Daemon (SSSD), which provides access to different identity and authentication providers, such as LDAP, Active Directory, and FreeIPA. SSSD can cache credentials and provide offline authentication, making it a valuable tool for environments with intermittent network connectivity.

### Structure of sssd.conf

#### Location
- **Location**: The `sssd.conf` file is typically located in `/etc/sssd/sssd.conf`.

#### Basic Structure

The `sssd.conf` file is organized into sections, each defining a specific aspect of SSSD's operation. The most common sections are:

- `[sssd]`: Global settings for the SSSD daemon.
- `[domain/<domain_name>]`: Settings for a specific domain or provider.
- `[nss]`: Configuration for the Name Service Switch (NSS) module.
- `[pam]`: Configuration for the Pluggable Authentication Modules (PAM) module.

#### Example Configuration

Here is an example of a basic `sssd.conf` file:

```ini
[sssd]
services = nss, pam
config_file_version = 2
domains = example.com

[domain/example.com]
id_provider = ldap
auth_provider = ldap
chpass_provider = ldap
ldap_uri = ldap://ldap.example.com
ldap_search_base = dc=example,dc=com
ldap_user_search_base = ou=People,dc=example,dc=com
ldap_group_search_base = ou=Groups,dc=example,dc=com
ldap_tls_reqcert = demand
cache_credentials = true
enumerate = true

[nss]
filter_groups = root
filter_users = root
reconnection_retries = 3

[pam]
offline_credentials_expiration = 2
offline_failed_login_attempts = 3
offline_failed_login_delay = 5
```

### Explanation of Sections and Options

#### [sssd]
This section contains global settings for SSSD.

- `services`: Lists the services SSSD will provide (e.g., `nss`, `pam`).
- `config_file_version`: Specifies the version of the configuration file format.
- `domains`: Lists the domains or identity providers that SSSD will manage.

#### [domain/<domain_name>]
This section configures a specific domain or identity provider.

- `id_provider`: Specifies the identity provider backend (e.g., `ldap`, `ad`).
- `auth_provider`: Specifies the authentication provider backend (e.g., `ldap`, `krb5`).
- `chpass_provider`: Specifies the change password provider backend (e.g., `ldap`).
- `ldap_uri`: The URI of the LDAP server.
- `ldap_search_base`: The base DN for LDAP searches.
- `ldap_user_search_base`: The base DN for user searches within LDAP.
- `ldap_group_search_base`: The base DN for group searches within LDAP.
- `ldap_tls_reqcert`: Specifies the requirement level for server certificates (e.g., `never`, `allow`, `try`, `demand`).
- `cache_credentials`: Enables or disables credential caching.
- `enumerate`: Specifies whether to enumerate users and groups.

#### [nss]
This section configures the Name Service Switch module.

- `filter_groups`: A list of groups to filter out from NSS lookups.
- `filter_users`: A list of users to filter out from NSS lookups.
- `reconnection_retries`: Number of times SSSD should retry connecting to the provider if it fails.

#### [pam]
This section configures the Pluggable Authentication Modules module.

- `offline_credentials_expiration`: Specifies the number of days offline credentials remain valid.
- `offline_failed_login_attempts`: Number of failed login attempts allowed while offline.
- `offline_failed_login_delay`: Delay in seconds between allowed failed login attempts while offline.

### Best Practices

- **Security**: Ensure the `sssd.conf` file is readable only by the root user. This can be achieved with the command `chmod 600 /etc/sssd/sssd.conf`.
- **Backups**: Always backup the `sssd.conf` file before making changes.
- **Testing**: Test changes in a non-production environment if possible.

### Advanced Configuration

#### Integrating with Active Directory
To integrate SSSD with an Active Directory (AD) domain, the configuration might look like this:

```ini
[sssd]
services = nss, pam
config_file_version = 2
domains = example.com

[domain/example.com]
id_provider = ad
auth_provider = ad
chpass_provider = ad
ad_server = ad.example.com
ad_domain = example.com
ldap_id_mapping = true
cache_credentials = true

[nss]
filter_groups = root
filter_users = root

[pam]
offline_credentials_expiration = 2
offline_failed_login_attempts = 3
offline_failed_login_delay = 5
```

#### Using Multiple Domains
If you need to configure SSSD for multiple domains, you can list them in the `[sssd]` section and create corresponding `[domain/<domain_name>]` sections:

```ini
[sssd]
services = nss, pam
config_file_version = 2
domains = example.com, anotherdomain.com

[domain/example.com]
id_provider = ldap
auth_provider = ldap
ldap_uri = ldap://ldap.example.com
ldap_search_base = dc=example,dc=com

[domain/anotherdomain.com]
id_provider = ldap
auth_provider = ldap
ldap_uri = ldap://ldap.anotherdomain.com
ldap_search_base = dc=anotherdomain,dc=com

[nss]
filter_groups = root
filter_users = root

[pam]
offline_credentials_expiration = 2
offline_failed_login_attempts = 3
offline_failed_login_delay = 5
```

### Conclusion

The `sssd.conf` file is crucial for configuring SSSD, enabling it to interact with various identity and authentication providers. Understanding the structure and options available in this file allows you to leverage SSSD's full capabilities, providing a robust and flexible authentication system for your Unix-like environment.
