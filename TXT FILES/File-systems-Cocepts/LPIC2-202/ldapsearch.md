# ldapsearch

`ldapsearch` is a command-line utility for querying and searching LDAP (Lightweight Directory Access Protocol) directories. It is commonly used by system administrators to retrieve and display information stored in LDAP directories, such as user and group information.

### Usage of ldapsearch

#### Basic Usage

To perform a basic search with `ldapsearch`, open a terminal and type:

```sh
ldapsearch -x -H ldap://ldap.example.com -b "dc=example,dc=com"
```

- `-x`: Use simple authentication (default).
- `-H ldap://ldap.example.com`: Specify the LDAP server URL.
- `-b "dc=example,dc=com"`: Specify the search base DN (Distinguished Name).

#### Common Options and Parameters

- `-D "cn=admin,dc=example,dc=com"`: Specify the bind DN for authentication.
- `-W`: Prompt for the password.
- `-LLL`: Print the search results in a simplified format without comments and version information.
- `-s base|one|sub`: Specify the scope of the search (`base` for the base object, `one` for one level, `sub` for the subtree).
- `-f filterfile`: Read the search filter from a file.
- `-A`: Retrieve only the attribute names.
- `-a never|always|search|find`: Specify the alias dereferencing policy.

#### Example Commands

1. **Simple Search**

   Search for all entries under the base DN:

   ```sh
   ldapsearch -x -H ldap://ldap.example.com -b "dc=example,dc=com"
   ```

2. **Search with a Filter**

   Search for entries with the attribute `uid=john`:

   ```sh
   ldapsearch -x -H ldap://ldap.example.com -b "dc=example,dc=com" "(uid=john)"
   ```

3. **Search with Authentication**

   Perform a search with authentication:

   ```sh
   ldapsearch -x -D "cn=admin,dc=example,dc=com" -W -H ldap://ldap.example.com -b "dc=example,dc=com" "(uid=john)"
   ```

   You will be prompted to enter the password for the bind DN.

4. **Retrieve Specific Attributes**

   Search and retrieve only the `cn` and `mail` attributes for entries with `uid=john`:

   ```sh
   ldapsearch -x -H ldap://ldap.example.com -b "dc=example,dc=com" "(uid=john)" cn mail
   ```

5. **Search with Subtree Scope**

   Perform a subtree search:

   ```sh
   ldapsearch -x -H ldap://ldap.example.com -b "dc=example,dc=com" -s sub "(uid=john)"
   ```

6. **Output in Simplified Format**

   Retrieve the results in a simplified format:

   ```sh
   ldapsearch -x -LLL -H ldap://ldap.example.com -b "dc=example,dc=com" "(uid=john)"
   ```

### Filters

LDAP filters are used to specify criteria for the search. Filters are enclosed in parentheses and can combine multiple conditions using logical operators.

**Basic Filter Syntax:**

- `(attribute=value)`: Simple equality filter.
- `(|(attr1=value1)(attr2=value2))`: Logical OR.
- `(&(attr1=value1)(attr2=value2))`: Logical AND.
- `(!(attribute=value))`: Logical NOT.
- `(attribute<=value)`: Less than or equal filter.
- `(attribute>=value)`: Greater than or equal filter.
- `(attribute=value*)`: Presence filter (matches if attribute exists).

**Examples:**

- Search for entries with `uid` equal to `john`:

  ```sh
  (uid=john)
  ```

- Search for entries with `uid` equal to `john` or `cn` equal to `John Doe`:

  ```sh
  (|(uid=john)(cn=John Doe))
  ```

- Search for entries with `uid` equal to `john` and `departmentNumber` equal to `1001`:

  ```sh
  (&(uid=john)(departmentNumber=1001))
  ```

### Security Considerations

When using `ldapsearch` with authentication (`-D` and `-W` options), it is crucial to protect the credentials used for binding to the LDAP server. Avoid hardcoding passwords in scripts or command lines, and prefer using environment variables or prompting for passwords.

### Conclusion

The `ldapsearch` utility is a powerful tool for querying and retrieving information from LDAP directories. With its flexible options and powerful filtering capabilities, it allows administrators to efficiently manage and troubleshoot LDAP-based authentication systems. Proper understanding and use of `ldapsearch` can significantly enhance the management of directory services.
