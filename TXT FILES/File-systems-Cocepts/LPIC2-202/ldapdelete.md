# ldapdelete

`ldapdelete` is a command-line utility used to delete entries from an LDAP directory. It communicates with an LDAP server to remove specified entries by their Distinguished Names (DNs).

### Usage of ldapdelete

#### Basic Usage

To delete entries from an LDAP directory, you typically provide the DNs of the entries to be deleted. The basic command is:

```sh
ldapdelete -x -D "cn=admin,dc=example,dc=com" -W "uid=john,dc=example,dc=com"
```

- `-x`: Use simple authentication.
- `-D "cn=admin,dc=example,dc=com"`: Specify the bind DN for authentication.
- `-W`: Prompt for the bind DN password.

#### Common Options and Parameters

- `-H ldap://ldap.example.com`: Specify the LDAP server URL.
- `-D "binddn"`: Use the specified bind DN to bind to the directory.
- `-w password`: Use the specified bind DN password (insecure, use only in scripts).
- `-c`: Continue on errors (skip entries that cause errors).
- `-v`: Run in verbose mode.
- `-f filename`: Read the DNs to delete from the specified file.
- `-r`: Delete recursively, removing an entry and all its subordinate entries.

#### Example Commands

1. **Delete a Single Entry**

   Delete a single entry specified by its DN:

   ```sh
   ldapdelete -x -D "cn=admin,dc=example,dc=com" -W "uid=john,dc=example,dc=com"
   ```

   This command will prompt for the bind DN password and then delete the specified entry.

2. **Delete Multiple Entries**

   Delete multiple entries specified by their DNs:

   ```sh
   ldapdelete -x -D "cn=admin,dc=example,dc=com" -W "uid=john,dc=example,dc=com" "uid=jane,dc=example,dc=com"
   ```

   This command will prompt for the bind DN password and then delete the specified entries.

3. **Delete Entries from a File**

   Delete entries whose DNs are listed in a file (`delete_entries.txt`):

   ```sh
   ldapdelete -x -D "cn=admin,dc=example,dc=com" -W -f delete_entries.txt
   ```

   The file `delete_entries.txt` should contain one DN per line:

   ```
   uid=john,dc=example,dc=com
   uid=jane,dc=example,dc=com
   ```

4. **Using a Secure Connection**

   Delete entries using a secure connection:

   ```sh
   ldapdelete -x -ZZ -H ldap://ldap.example.com -D "cn=admin,dc=example,dc=com" -W "uid=john,dc=example,dc=com"
   ```

5. **Delete Recursively**

   Delete an entry and all its subordinate entries:

   ```sh
   ldapdelete -x -r -D "cn=admin,dc=example,dc=com" -W "ou=department,dc=example,dc=com"
   ```

#### Security Considerations

1. **Avoid Hardcoding Passwords**: Do not hardcode passwords in scripts or command lines. Use prompts or secure methods to pass passwords.

2. **Use Secure Connections**: Always use secure connections (`-ZZ` for StartTLS) to protect data during transmission.

3. **Proper DN File Permissions**: Ensure that the file containing DNs has appropriate permissions to prevent unauthorized access.

4. **Limit Privileges**: Use the least privilege principle. Bind with an account that has only the necessary permissions to delete entries.

### Conclusion

The `ldapdelete` utility is essential for removing entries from an LDAP directory. By understanding its options and secure usage practices, administrators can effectively manage the deletion of entries while maintaining security and integrity. Properly configured, `ldapdelete` facilitates the efficient removal of user and resource entries in LDAP environments.
