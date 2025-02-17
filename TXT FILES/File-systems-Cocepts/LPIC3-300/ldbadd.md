
# ldbadd

The `ldbadd` command is used to add entries to an LDB (LDAP-like Database) file, which is commonly used by Samba for managing directory information, including user accounts, groups, and other related data. The LDB database is a key component in environments such as Samba-based Active Directory.

### Key Features of `ldbadd`:
- **Purpose**: Adds new entries to an LDB database.
- **Usage Context**: Primarily used in Samba environments for adding user, group, or other directory entries to the database.
- **Command Syntax**: Adds an entry, usually specified in LDIF (LDAP Data Interchange Format) or another format, to an LDB file.

### General Syntax:
```bash
ldbadd [OPTIONS] <LDB_FILE> <LDIF_FILE>
```
Where:
- **`<LDB_FILE>`**: The path to the LDB database file (e.g., `/var/lib/samba/private/sam.ldb`).
- **`<LDIF_FILE>`**: Path to the LDIF file containing the entries to be added.

### Example Usage:

1. **Adding an Entry for a User**:
   To add a new user entry to the Samba LDB database, you would typically write the user details in an LDIF file and then use `ldbadd` to add the entry to the database.

   **Example LDIF file (add-user.ldif)**:
   ```ldif
   dn: uid=john.doe,ou=users,dc=example,dc=com
   objectClass: inetOrgPerson
   uid: john.doe
   cn: John Doe
   sn: Doe
   givenName: John
   mail: john.doe@example.com
   userPassword: password
   ```

   **Command to add the user**:
   ```bash
   ldbadd /var/lib/samba/private/sam.ldb add-user.ldif
   ```

   This command will read the `add-user.ldif` file and add the user entry to the specified Samba LDB database (`sam.ldb`).

2. **Adding a Group**:
   Similarly, you can use `ldbadd` to add a group entry. Here is an example of how to add a group to the database.

   **Example LDIF file (add-group.ldif)**:
   ```ldif
   dn: cn=admins,ou=groups,dc=example,dc=com
   objectClass: posixGroup
   cn: admins
   gidNumber: 1000
   ```

   **Command to add the group**:
   ```bash
   ldbadd /var/lib/samba/private/sam.ldb add-group.ldif
   ```

### Options:

- **`--help`**: Displays help information for the `ldbadd` command.
   ```bash
   ldbadd --help
   ```

   Example output:
   ```
   usage: ldbadd [-h] [--help] LDB_FILE LDIF_FILE
   Add entries to an LDB database.
   ```

- **`--dry-run`**: This option allows you to simulate the command without actually making changes to the LDB database. It’s a good way to verify the LDIF data before committing to the changes.
   ```bash
   ldbadd --dry-run /var/lib/samba/private/sam.ldb add-user.ldif
   ```

- **`-v` or `--verbose`**: Provides more detailed output during execution, helpful for debugging or tracking the operation.
   ```bash
   ldbadd --verbose /var/lib/samba/private/sam.ldb add-user.ldif
   ```

### Practical Use Cases:
- **User and Group Management**: Adding new users and groups to a Samba-based Active Directory or domain controller.
- **Directory Imports**: Importing data from an LDIF file into the Samba LDB database. This is particularly useful during migrations or bulk user creation.
- **Directory Maintenance**: Adding additional entries such as contacts, organizational units (OUs), or other directory objects.

### Safety Considerations:
- **Verify LDIF Data**: Ensure that the LDIF file is properly formatted and contains valid data before running `ldbadd` to avoid errors or corrupt entries.
- **Backup First**: It's always a good idea to back up your LDB database (`sam.ldb`) before adding critical entries. This allows you to restore the database in case of any mistakes.
  
### Conclusion:
The `ldbadd` command is an essential tool for adding entries to the Samba LDB database. Whether you are managing users, groups, or other directory objects, `ldbadd` helps ensure that your Samba directory remains up-to-date and well-organized. By using LDIF files, you can efficiently add multiple entries at once, making it a powerful tool for bulk imports or new directory object creation. Always exercise caution when modifying your LDB database to avoid unnecessary changes or corruption.
