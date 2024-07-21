# ldapadd

`ldapadd` is a command-line utility used to add entries to an LDAP directory. It reads the data to be added from a file or from standard input and communicates with an LDAP server to perform the additions.

### Usage of ldapadd

#### Basic Usage

To add entries to an LDAP directory, you typically provide an LDIF (LDAP Data Interchange Format) file containing the entries to be added. The basic command is:

```sh
ldapadd -x -D "cn=admin,dc=example,dc=com" -W -f new_entries.ldif
```

- `-x`: Use simple authentication.
- `-D "cn=admin,dc=example,dc=com"`: Specify the bind DN for authentication.
- `-W`: Prompt for the bind DN password.
- `-f new_entries.ldif`: Specify the LDIF file containing the entries to add.

#### Common Options and Parameters

- `-H ldap://ldap.example.com`: Specify the LDAP server URL.
- `-D "binddn"`: Use the specified bind DN to bind to the directory.
- `-w password`: Use the specified bind DN password (insecure, use only in scripts).
- `-f filename`: Read the entries to be added from the specified file.
- `-ZZ`: Require a secure TLS connection.
- `-c`: Continue on errors (skip entries that cause errors).

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

#### Example Commands

1. **Add Entries from LDIF File**

   Add entries defined in an LDIF file:

   ```sh
   ldapadd -x -D "cn=admin,dc=example,dc=com" -W -f new_entries.ldif
   ```

   This command will prompt for the bind DN password and then add the entries from the file.

2. **Add Entries from Standard Input**

   You can also provide the LDIF data directly via standard input:

   ```sh
   ldapadd -x -D "cn=admin,dc=example,dc=com" -W <<EOF
   dn: uid=jane,dc=example,dc=com
   objectClass: inetOrgPerson
   uid: jane
   sn: Smith
   cn: Jane Smith
   mail: jane.smith@example.com
   userPassword: secret
   EOF
   ```

3. **Using a Secure Connection**

   Add entries using a secure connection:

   ```sh
   ldapadd -x -ZZ -H ldap://ldap.example.com -D "cn=admin,dc=example,dc=com" -W -f new_entries.ldif
   ```

4. **Continue on Errors**

   Add entries and continue on errors:

   ```sh
   ldapadd -x -c -D "cn=admin,dc=example,dc=com" -W -f new_entries.ldif
   ```

   This command will skip any entries that cause errors and continue with the rest.

#### Security Considerations

1. **Avoid Hardcoding Passwords**: Do not hardcode passwords in scripts or command lines. Use prompts or secure methods to pass passwords.

2. **Use Secure Connections**: Always use secure connections (`-ZZ` for StartTLS) to protect data during transmission.

3. **Proper LDIF File Permissions**: Ensure that the LDIF file has appropriate permissions to prevent unauthorized access.

4. **Limit Privileges**: Use the least privilege principle. Bind with an account that has only the necessary permissions to add entries.

### Conclusion

The `ldapadd` utility is essential for adding entries to an LDAP directory. By understanding its options and secure usage practices, administrators can effectively manage and expand LDAP directories while maintaining security and integrity. Properly configured, `ldapadd` facilitates the efficient addition of user and resource entries in LDAP environments.
