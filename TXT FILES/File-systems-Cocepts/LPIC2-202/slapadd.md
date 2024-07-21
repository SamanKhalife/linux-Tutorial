# slapadd

`slapadd` is a utility provided with OpenLDAP, used to populate an LDAP directory database from an LDIF (LDAP Data Interchange Format) file. Unlike `ldapadd`, which interacts with the LDAP server over the network, `slapadd` directly accesses the database files on the server, making it faster and more suitable for initial bulk loading of data.

### Usage of slapadd

#### Basic Usage

To add entries to an LDAP directory database using `slapadd`, you typically provide an LDIF file containing the entries to be added. The basic command is:

```sh
slapadd -v -l new_entries.ldif
```

- `-v`: Run in verbose mode.
- `-l new_entries.ldif`: Specify the LDIF file containing the entries to add.

#### Common Options and Parameters

- `-f slapd.conf`: Specify the configuration file to use.
- `-F slapd.d`: Specify the configuration directory to use.
- `-n database`: Specify the database number (default is 1).
- `-q`: Enable quick mode (faster but might bypass certain integrity checks).
- `-c`: Continue on errors (skip entries that cause errors).
- `-u`: Dry run (check syntax without making changes).
- `-s`: Use a specific suffix (DN) as the base for the imported data.

#### Example Commands

1. **Add Entries from LDIF File**

   Add entries defined in an LDIF file to the default database:

   ```sh
   slapadd -v -l new_entries.ldif
   ```

2. **Specify Configuration File**

   Use a specific configuration file:

   ```sh
   slapadd -v -f /etc/openldap/slapd.conf -l new_entries.ldif
   ```

3. **Specify Configuration Directory**

   Use a specific configuration directory:

   ```sh
   slapadd -v -F /etc/openldap/slapd.d -l new_entries.ldif
   ```

4. **Add Entries to a Specific Database**

   Add entries to a specific database number:

   ```sh
   slapadd -v -n 2 -l new_entries.ldif
   ```

5. **Quick Mode**

   Enable quick mode for faster processing:

   ```sh
   slapadd -v -q -l new_entries.ldif
   ```

6. **Continue on Errors**

   Continue processing even if there are errors in the LDIF file:

   ```sh
   slapadd -v -c -l new_entries.ldif
   ```

7. **Dry Run**

   Check the syntax of the LDIF file without making any changes to the database:

   ```sh
   slapadd -v -u -l new_entries.ldif
   ```

8. **Use a Specific Suffix**

   Import data with a specific base DN:

   ```sh
   slapadd -v -s "dc=example,dc=com" -l new_entries.ldif
   ```

#### Example LDIF File

An LDIF file (`new_entries.ldif`) might look like this:

```ldif
dn: uid=john,dc=example,dc=com
objectClass: inetOrgPerson
uid: john
sn: Doe
cn: John Doe
mail: john.doe@example.com
userPassword: secret

dn: uid=jane,dc=example,dc=com
objectClass: inetOrgPerson
uid: jane
sn: Smith
cn: Jane Smith
mail: jane.smith@example.com
userPassword: secret
```

This file defines two new entries to be added to the LDAP directory.

#### Security Considerations

1. **File Permissions**: Ensure that the LDIF file has appropriate permissions to prevent unauthorized access.
2. **Backup**: Always backup your LDAP database before running `slapadd` to avoid data loss in case of errors.
3. **Run as Root**: Ensure `slapadd` is run with appropriate permissions, typically as the root user, or with sudo.
4. **Database Consistency**: Be cautious with the `-q` option, as it may bypass some integrity checks.

### Conclusion

The `slapadd` utility is essential for efficiently populating an LDAP directory database with bulk data. By understanding its options and secure usage practices, administrators can effectively manage the addition of large amounts of data while maintaining security and integrity. Properly configured, `slapadd` facilitates the rapid and reliable population of LDAP directory databases in various environments.
