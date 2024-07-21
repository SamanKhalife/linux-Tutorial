# slapd-config

`slapd-config` is the dynamic runtime configuration engine for OpenLDAP's slapd (Standalone LDAP Daemon). It allows administrators to manage LDAP server configurations without having to stop and restart the server. The configuration is stored in a special LDAP backend, often referred to as `cn=config`.

### Overview

The `slapd-config` system uses a hierarchical structure to store configuration settings. These settings can be modified using standard LDAP operations. This approach provides several advantages over the traditional `slapd.conf` file, including:

- **Dynamic Configuration:** Changes can be made without restarting the server.
- **Fine-Grained Control:** Individual settings can be adjusted as needed.
- **LDAP Interface:** Configuration can be managed using LDAP tools.

### Structure of slapd-config

The `slapd-config` database is typically organized under the `cn=config` entry, with various subentries representing different aspects of the configuration.

#### Key Entries and Attributes

1. **Global Configuration (cn=config):**
   - **olcLogLevel:** Sets the logging level.
   - **olcThreads:** Specifies the number of threads.
   - **olcAccessLog:** Configuration for the access log.

2. **Database Configuration (olcDatabase):**
   - **olcDatabase:** Type of database (e.g., bdb, hdb, mdb).
   - **olcSuffix:** Specifies the database suffix (base DN).
   - **olcRootDN:** The root DN for administrative access.
   - **olcRootPW:** Password for the root DN.
   - **olcIndex:** Indexing options for the database.

3. **Backend Configuration (olcBackend):**
   - **olcBackend:** Specifies the backend type (e.g., hdb, bdb, mdb).

4. **Schema Configuration (olcSchemaConfig):**
   - **olcAttributeTypes:** Defines attribute types.
   - **olcObjectClasses:** Defines object classes.

### Managing slapd-config

#### Viewing Configuration

To view the current configuration, you can use LDAP search operations. For example, to view the global configuration:

```sh
ldapsearch -x -H ldap://localhost -D "cn=admin,cn=config" -W -b "cn=config"
```

#### Modifying Configuration

To modify the configuration, you use LDAP modify operations. Below are some common examples:

1. **Changing the Log Level:**

Create a file `modify.ldif` with the following content:

```ldif
dn: cn=config
changetype: modify
replace: olcLogLevel
olcLogLevel: stats
```

Then apply the changes using `ldapmodify`:

```sh
ldapmodify -x -H ldap://localhost -D "cn=admin,cn=config" -W -f modify.ldif
```

2. **Adding an Index:**

Create a file `addindex.ldif` with the following content:

```ldif
dn: olcDatabase={1}mdb,cn=config
changetype: modify
add: olcIndex
olcIndex: cn,sn eq
```

Then apply the changes using `ldapmodify`:

```sh
ldapmodify -x -H ldap://localhost -D "cn=admin,cn=config" -W -f addindex.ldif
```

3. **Adding a New Database:**

Create a file `adddb.ldif` with the following content:

```ldif
dn: olcDatabase={2}mdb,cn=config
objectClass: olcDatabaseConfig
olcDatabase: {2}mdb
olcSuffix: dc=example,dc=com
olcRootDN: cn=admin,dc=example,dc=com
olcRootPW: secret
olcDbDirectory: /var/lib/ldap/example
olcDbIndex: objectClass eq
```

Then apply the changes using `ldapadd`:

```sh
ldapadd -x -H ldap://localhost -D "cn=admin,cn=config" -W -f adddb.ldif
```

#### Removing Configuration Entries

To remove configuration entries, you use the LDAP delete operation. For example, to remove an index:

Create a file `removeindex.ldif` with the following content:

```ldif
dn: olcDatabase={1}mdb,cn=config
changetype: modify
delete: olcIndex
olcIndex: cn,sn eq
```

Then apply the changes using `ldapmodify`:

```sh
ldapmodify -x -H ldap://localhost -D "cn=admin,cn=config" -W -f removeindex.ldif
```

### Security Considerations

1. **Access Control:** Ensure that access controls are properly configured to restrict who can modify the `cn=config` entries.
2. **Backups:** Regularly back up your configuration data. The `cn=config` directory is typically located in `/etc/ldap/slapd.d`.
3. **Audit Changes:** Keep an audit trail of configuration changes for troubleshooting and compliance purposes.

### Conclusion

The `slapd-config` system provides a powerful and flexible way to manage OpenLDAP server configurations. By leveraging LDAP operations to modify settings dynamically, administrators can achieve fine-grained control and ensure that their LDAP services are responsive and adaptable to changing requirements. Properly managing `slapd-config` involves understanding its structure, using appropriate tools and commands, and adhering to security best practices.
