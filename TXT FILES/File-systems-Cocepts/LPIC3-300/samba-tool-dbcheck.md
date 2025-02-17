# samba-tool dbcheck
The `samba-tool dbcheck` command is a utility in Samba used to check the integrity and consistency of the Active Directory (AD) database. It is an essential tool for administrators to diagnose and fix issues related to the Samba domain controller’s database. This command can help detect inconsistencies or errors that may arise in the AD database due to corruption, improper replication, or other issues.

### Overview of `samba-tool dbcheck`

The `dbcheck` command verifies the structure and content of the Samba AD database (which is stored in **tdb** format) and checks for any database inconsistencies or corruption. If issues are found, the command can attempt to fix them automatically.

### Syntax

```bash
samba-tool dbcheck [options]
```

### Common Options for `samba-tool dbcheck`

#### 1. **`--fix`**
   - **Description**: This option attempts to automatically fix any issues or inconsistencies found in the database during the check.
   - **Usage**:
     ```bash
     samba-tool dbcheck --fix
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --fix
     ```
   - **Important**: Use this with caution, as it will attempt to make changes to the AD database. It's recommended to perform a backup of your domain controller before using the `--fix` option.

#### 2. **`--rebuild`**
   - **Description**: This option is used to rebuild the Samba database from scratch, useful if the database is corrupted and other repair attempts have failed. **This operation can take a significant amount of time** depending on the size of your AD.
   - **Usage**:
     ```bash
     samba-tool dbcheck --rebuild
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --rebuild
     ```

#### 3. **`--no-fix`**
   - **Description**: This option runs the database check but does not attempt to fix any issues. It's used when you want to identify problems but not make any changes yet.
   - **Usage**:
     ```bash
     samba-tool dbcheck --no-fix
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --no-fix
     ```
   - **Recommendation**: Use this option for a dry run to see what issues exist before making any changes.

#### 4. **`--deep`**
   - **Description**: Performs a more thorough check, including checking for errors in the database schema and records.
   - **Usage**:
     ```bash
     samba-tool dbcheck --deep
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --deep
     ```

#### 5. **`--repair`**
   - **Description**: This option is used to perform repairs and fix certain types of corruption in the AD database.
   - **Usage**:
     ```bash
     samba-tool dbcheck --repair
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --repair
     ```

#### 6. **`--extended`**
   - **Description**: Checks for a broader range of problems within the database, including extended attributes and other non-critical issues.
   - **Usage**:
     ```bash
     samba-tool dbcheck --extended
     ```
   - **Example**:
     ```bash
     samba-tool dbcheck --extended
     ```

### Example Use Cases for `samba-tool dbcheck`

#### 1. Checking the AD database for issues (without fixing)
If you suspect issues in your Samba AD database but do not want to automatically attempt repairs, you can run the following:

```bash
samba-tool dbcheck --no-fix
```

This will perform a check and show any problems but will not modify the database.

#### 2. Checking the AD database and automatically fixing any issues
If you've identified issues and want to attempt repairs, you can use the `--fix` option:

```bash
samba-tool dbcheck --fix
```

This command will attempt to fix any database issues it detects. You can specify additional options like `--deep` for more thorough checking.

#### 3. Rebuilding the database (if there is corruption)
In the case where the database is severely corrupted and repairs cannot fix it, the `--rebuild` option can help rebuild the database from scratch:

```bash
samba-tool dbcheck --rebuild
```

This operation can be time-consuming, depending on the size of your directory and the number of records in the AD database.

#### 4. Running an extended check for non-critical issues
You can also use the `--extended` option to check for extended issues that may not be immediately critical but could impact the domain controller's performance:

```bash
samba-tool dbcheck --extended
```

### Additional Notes

- **Backup Before Repairing**: Before running the `--fix` or `--rebuild` options, it’s recommended to back up your Samba AD database and related data. This ensures you can restore the system in case the fix process causes further issues.
  
- **Performance**: Running `samba-tool dbcheck` can be resource-intensive, especially if using the `--deep` or `--extended` options. It’s best to run these checks during off-peak hours.

- **Log Files**: The command will output details to the terminal. If you need to analyze results over time, consider redirecting the output to a log file for review:
  ```bash
  samba-tool dbcheck --fix > /var/log/samba/dbcheck.log 2>&1
  ```

### Conclusion

The `samba-tool dbcheck` command is a critical utility for diagnosing and fixing issues in the Samba AD database. It provides several options for checking the integrity of the database and repairing any inconsistencies that are found. Administrators should use this tool as part of regular maintenance to ensure the health of their Samba domain controllers, but always ensure backups are made before performing repairs or rebuilds.
