# tdbtool

The `tdbtool` command is used for interacting with **Trivial Database (TDB)** files in Samba. TDB is a simple database format used by Samba to store various data, such as user and group information, share definitions, and other configurations that need to be stored persistently. 

`tdbtool` is a utility that allows administrators to inspect, query, and modify the contents of TDB files, making it useful for troubleshooting and managing Samba databases.

### Key Features of `tdbtool`:
- **Database Management**: `tdbtool` allows administrators to open, read, and modify TDB files.
- **Usage Context**: It's primarily used for debugging or modifying Samba's internal databases, such as `smbpasswd`, `secrets.tdb`, `user.tdb`, and `group.tdb`.
- **Command Syntax**: The syntax is straightforward, offering various subcommands for inspecting and modifying TDB files.

### General Syntax:
```bash
tdbtool <TDB_FILE> [SUBCOMMAND]
```
Where:
- **`<TDB_FILE>`**: The path to the TDB database file (e.g., `/var/lib/samba/private/secrets.tdb`).
- **`[SUBCOMMAND]`**: The subcommand or operation you want to perform on the TDB file (e.g., `list`, `get`, `put`, etc.).

### Common `tdbtool` Subcommands:
1. **`list`**: Lists all the entries in the TDB file.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb list
   ```
   This will show all the keys and their corresponding values in the specified TDB file.

2. **`get`**: Retrieves the value associated with a specified key in the TDB database.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb get <key>
   ```
   Replace `<key>` with the actual key you're trying to retrieve. This command will show the value stored for the given key.

3. **`put`**: Puts or modifies a key-value pair in the TDB file.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb put <key> <value>
   ```
   Replace `<key>` and `<value>` with the appropriate key and value to be stored in the TDB database.

4. **`del`**: Deletes a specific key-value pair from the TDB file.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb del <key>
   ```
   This removes the specified key and its associated value from the TDB file.

5. **`dump`**: Dumps the contents of the TDB file in a human-readable format.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb dump
   ```
   This command will provide a more detailed, human-readable format of the database entries.

6. **`close`**: Closes the TDB file after the operation is complete.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb close
   ```
   This command ensures that the TDB file is properly closed after any operations are performed.

7. **`verify`**: Checks the integrity of the TDB file to ensure there are no corruption issues.
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb verify
   ```
   This command will verify that the TDB file is structurally sound and not corrupted.

### Example Usage:

1. **List All Entries in `secrets.tdb`**:
   To view all the key-value pairs in the `secrets.tdb` file, which contains secrets and authentication-related information:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb list
   ```

2. **Get a Specific Value**:
   If you want to get the value associated with a key (e.g., `secrets.ldb`), you can use:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb get secrets.ldb
   ```

3. **Modify or Add a Key-Value Pair**:
   If you need to add or modify a key-value pair in the `secrets.tdb` file, use:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb put secret_key new_value
   ```

4. **Delete an Entry**:
   If you want to delete a specific key-value pair, use:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb del secret_key
   ```

5. **Verify Database Integrity**:
   To check the integrity of the `secrets.tdb` file:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb verify
   ```

6. **Dump the Contents**:
   For a more detailed human-readable output of the database:
   ```bash
   tdbtool /var/lib/samba/private/secrets.tdb dump
   ```

### Options:
- **`-h` or `--help`**: Displays help information for `tdbtool`.
   ```bash
   tdbtool --help
   ```
   Output will provide a list of available subcommands and options.

### Practical Use Cases:

1. **Debugging Samba Authentication Issues**:
   TDB files store various secrets and user information related to Samba’s authentication system (e.g., `secrets.tdb` or `smbpasswd.tdb`). If there are issues with Samba authentication, `tdbtool` can be used to inspect or modify the contents of these files to troubleshoot.

2. **Managing Samba Database Entries**:
   Administrators can use `tdbtool` to directly modify Samba's internal TDB files, such as adding new user accounts or modifying existing ones, or clearing certain keys that might be causing issues.

3. **Repairing Corrupted Databases**:
   If a TDB file becomes corrupted, `tdbtool` can help verify and possibly repair the file. The `verify` subcommand checks the file for any inconsistencies.

4. **Data Extraction**:
   In cases where specific data needs to be extracted from a TDB file (for example, user credentials or secrets), `tdbtool` can be used to dump or extract the information.

5. **Backup and Restoration**:
   `tdbtool` can be used as part of a backup strategy, allowing administrators to extract and store key-value pairs or dump the entire TDB database, making it easier to restore if needed.

### Conclusion:
`tdbtool` is a powerful utility for managing and troubleshooting the Trivial Database files used by Samba. It allows system administrators to perform operations like querying, modifying, and verifying Samba-related TDB files, which are critical for storing user authentication information, share definitions, and various other configuration data. Whether for debugging or making direct modifications, `tdbtool` provides an essential interface for working with Samba’s internal database.
