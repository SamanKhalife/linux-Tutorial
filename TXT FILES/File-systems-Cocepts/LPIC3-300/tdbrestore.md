# tdbrestore
The `tdbrestore` command is part of the **Samba** suite, specifically used for restoring **Trivial Database (TDB)** files. TDB is a simple database format used by Samba to store various configurations and state information. The most common use case for `tdbrestore` is to recover or restore the TDB database files, which can be crucial for Samba's operation, such as restoring the Samba **user database**, **secrets database**, and other related Samba state files.

### Overview of `tdbrestore`

`tdbrestore` is used primarily when a Samba TDB database is corrupted, missing, or needs to be restored from a backup. This tool allows administrators to restore TDB database files from a backup into their respective Samba configuration, enabling the system to recover from database failures.

### General Syntax:
```bash
tdbrestore [options] <db-file>
```

Where:
- `<db-file>` is the path to the TDB database file you want to restore.
- The tool restores data from the backup to the given TDB file.

### Common Use Cases
- **Restoring TDB Files After Corruption**: If a TDB file (e.g., `secrets.tdb` or `passdb.tdb`) becomes corrupted, `tdbrestore` allows you to restore from a previously created backup.
- **Recovering User Data**: If user data in a `passdb.tdb` file gets corrupted, it can be restored using the backup.
- **Restoring Samba Secrets**: Samba uses the `secrets.tdb` file to store various secrets, including the machine's password for authenticating with other services. If this file is corrupted, it can be restored.

### Key Options for `tdbrestore`

- `--help`: Displays help for the command and available options.
- `--backup`: Specifies the backup file to restore from.
- `--no-rename`: Optionally prevents renaming of files.
- `--recovery`: Used to restore a TDB file during a recovery process, such as after a crash.

### Example Usage

#### 1. Restoring from a Backup
Suppose you have a backup of your `passdb.tdb` file and want to restore it to its original location. Here's how you can do it:

```bash
tdbrestore --backup=/path/to/backup/passdb.tdb /var/lib/samba/private/passdb.tdb
```
This restores the `passdb.tdb` from the backup to its original location.

#### 2. Restoring a Specific TDB Database File
If you want to restore a specific TDB database file (like `secrets.tdb`), you would use:

```bash
tdbrestore /path/to/secrets.tdb
```

This restores the file to its default location based on the Samba configuration.

#### 3. Viewing Available Options
To see the available options for `tdbrestore` and get more details, run:

```bash
tdbrestore --help
```

This command will output detailed information about the command and the available options.

### Typical Use Cases for TDB Databases in Samba

- **`secrets.tdb`**: Stores secrets such as the machine password used for communication between the Samba server and the domain.
- **`passdb.tdb`**: Stores user account data, including passwords and other information related to Samba users.
- **`tdb` Files for Print Spoolers**: Samba uses TDB files for spool management when printing over the network.

### Conclusion

`tdbrestore` is a vital tool for administrators who need to recover Samba's TDB databases from backups. Whether it’s recovering user data, restoring Samba secrets, or dealing with corruption, `tdbrestore` helps ensure that Samba services can be quickly restored and remain operational. Regular backups of critical TDB files are highly recommended to avoid potential data loss.
