# ldbsearch

The `ldbsearch` command is used to search an LDB (LDAP-like Database) file in Samba, much like how an LDAP search would operate in a traditional LDAP server. It allows you to query the contents of an LDB database, such as a Samba Active Directory, and retrieve specific entries or attributes based on search criteria.

### Key Features of `ldbsearch`:
- **Purpose**: Searches for and displays entries stored in an LDB database. You can filter the results by specifying search filters and attributes.
- **Usage Context**: It's commonly used in Samba environments to query the Active Directory or domain controller database for user, group, and other directory objects.
- **Command Syntax**: The syntax for `ldbsearch` allows you to define the base DN (Distinguished Name) and filter to narrow the search results.

### General Syntax:
```bash
ldbsearch [OPTIONS] <LDB_FILE> <SEARCH_FILTER>
```
Where:
- **`<LDB_FILE>`**: Path to the LDB database file (e.g., `/var/lib/samba/private/sam.ldb`).
- **`<SEARCH_FILTER>`**: The filter or criteria used to search for entries (e.g., `(&(objectClass=user)(cn=john))`).
- **`<ATTRIBUTES>`**: Optional list of attributes to retrieve in the search results.

### Example Usage:

1. **Search for All Users**:
   To list all users in the LDB database, you can use a filter like `(objectClass=user)`. This will return entries that represent user objects.
   ```bash
   ldbsearch /var/lib/samba/private/sam.ldb "(objectClass=user)"
   ```
   This command will return all user entries from the LDB database.

2. **Search for a Specific User**:
   To find a user by their `cn` (common name), you can use a filter like `(cn=john.doe)`.
   ```bash
   ldbsearch /var/lib/samba/private/sam.ldb "(cn=john.doe)"
   ```
   This will return information for the user `john.doe` from the LDB database.

3. **Search for Specific Attributes**:
   You can specify which attributes you want to retrieve. For example, to search for all users and only display their `uid` and `mail` attributes, you can run:
   ```bash
   ldbsearch /var/lib/samba/private/sam.ldb "(objectClass=user)" uid mail
   ```
   This will return the `uid` and `mail` attributes for all user entries.

4. **Search for Groups**:
   You can also search for groups in the LDB database by using a filter like `(objectClass=group)`.
   ```bash
   ldbsearch /var/lib/samba/private/sam.ldb "(objectClass=group)"
   ```
   This command will return all group entries from the LDB database.

### Options:

- **`--help`**: Displays help information for the `ldbsearch` command.
   ```bash
   ldbsearch --help
   ```

   Example output:
   ```
   usage: ldbsearch [-h] [--help] LDB_FILE SEARCH_FILTER [ATTRIBUTES]
   Search entries in an LDB database.
   ```

- **`-v` or `--verbose`**: Provides more detailed output, useful for debugging or tracking the operation.
   ```bash
   ldbsearch --verbose /var/lib/samba/private/sam.ldb "(objectClass=user)"
   ```

- **`--no-filters`**: Disables any filtering and returns all entries from the database. Be careful, as this may result in large amounts of data.
   ```bash
   ldbsearch --no-filters /var/lib/samba/private/sam.ldb
   ```

- **`-p` or `--page-size`**: Specifies the number of results per page to return, useful for pagination when dealing with large databases.
   ```bash
   ldbsearch -p 100 /var/lib/samba/private/sam.ldb "(objectClass=user)"
   ```

- **`--debug-level=<level>`**: Sets the debug level for troubleshooting purposes. The higher the number, the more detailed the output.
   ```bash
   ldbsearch --debug-level=3 /var/lib/samba/private/sam.ldb "(objectClass=user)"
   ```

### Practical Use Cases:

1. **User Lookup**:
   `ldbsearch` is useful for querying user entries in a Samba-based Active Directory. For example, if an administrator needs to find details about a user, such as their email or group membership, this command can be used.

2. **Directory Auditing**:
   Administrators can use `ldbsearch` to audit directory contents, ensuring that the information in the directory is accurate and compliant with organizational policies.

3. **Troubleshooting**:
   If there are issues with user authentication or group membership, `ldbsearch` can be used to inspect the underlying directory data to ensure the correct information is present.

4. **Integration with Scripts**:
   `ldbsearch` can be combined with other commands or scripted operations to automate the management of Samba directory entries.

### Example: Search for Group Members
If you want to find all members of a specific group, for instance, a group named `admins`, you would use a filter based on the group's DN (Distinguished Name). For example:
```bash
ldbsearch /var/lib/samba/private/sam.ldb "(cn=admins)" member
```
This will return a list of all members of the `admins` group.

### Conclusion:
The `ldbsearch` command is an essential tool for querying entries in an LDB database in Samba environments. It allows administrators to filter and retrieve specific entries or attributes from a Samba-based Active Directory, making it valuable for user and group management, directory auditing, and troubleshooting. By understanding the search filters and options available, administrators can efficiently manage and interact with directory data.
