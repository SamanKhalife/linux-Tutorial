# ldbdel

The `ldbdel` command is used to delete entries from an LDB (LDAP-like Database) file, which is used in Samba for storing various types of data such as Samba user accounts, group memberships, and other directory-related information.

LDB (LDAP-like Database) is a lightweight, embedded database used by Samba for its internal use, often replacing traditional LDAP for handling directory information in environments like Active Directory or a Samba-based domain controller.

### Key Features of `ldbdel`:
- **Purpose**: Deletes an entry from an LDB database.
- **Usage Context**: Primarily used in Samba and its associated tools, including managing user and group entries in a Samba-based Active Directory.
- **Command Syntax**: Deletes a specific entry identified by its Distinguished Name (DN) or unique key.
  
### General Syntax:
```bash
ldbdel [OPTIONS] <LDB_FILE> <DN>
```
Where:
- **`<LDB_FILE>`**: Path to the LDB database file (e.g., `/var/lib/samba/private/sam.ldb`).
- **`<DN>`**: Distinguished Name (DN) of the entry to delete (e.g., `uid=john.doe,ou=users,dc=example,dc=com`).

### Example Usage:
1. **Delete a User Entry from the LDB Database**:
   If you want to delete a user entry from a Samba LDB database, you would specify the LDB file and the DN of the user you want to delete. For example:
   ```bash
   ldbdel /var/lib/samba/private/sam.ldb "uid=john.doe,ou=users,dc=example,dc=com"
   ```

   This will remove the user entry for `john.doe` from the Samba LDB database.

2. **Delete a Group Entry**:
   Similarly, you can delete a group entry:
   ```bash
   ldbdel /var/lib/samba/private/sam.ldb "cn=admins,ou=groups,dc=example,dc=com"
   ```

   This deletes the group `admins` from the LDB database.

### Options:
- **`--help`**: Displays the help message with the list of options.
  
   ```bash
   ldbdel --help
   ```

   Example Output:
   ```
   usage: ldbdel [-h] [--help] LDB_FILE DN
   Delete an entry from the LDB database.
   ```

### Practical Use Cases:
- **Managing Samba Directory**: Deleting user or group accounts from a Samba-based domain controller.
- **Clearing Out Old Entries**: Useful when cleaning up outdated or inactive entries in an Active Directory replica.
- **Database Maintenance**: Ensures that unnecessary or unwanted entries are properly removed from the LDB database, especially during migrations or after user deletions.

### Safety Considerations:
- **Irreversible Action**: Deleting an entry using `ldbdel` is a permanent operation. Once an entry is deleted, it cannot be recovered unless you have backups.
- **Backup First**: Always ensure you have a recent backup of your LDB database before deleting critical entries.

### Conclusion:
The `ldbdel` command is a powerful tool for removing entries from the Samba LDB database. It allows administrators to manage directory data in a Samba domain controller environment efficiently. Whether you're dealing with user accounts or groups, `ldbdel` helps to maintain a clean and organized LDB database. Always use caution when performing delete operations to avoid accidentally removing necessary entries.
