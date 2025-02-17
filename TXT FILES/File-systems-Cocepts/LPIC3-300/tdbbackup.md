# tdbbackup
The `tdbbackup` command is a tool provided by **Samba** for creating backups of **Trivial Database (TDB)** files. TDB is a simple database format used by Samba to store various types of configuration and state data, including user information, secrets, and other internal Samba management data.

### Overview of `tdbbackup`

The `tdbbackup` command allows administrators to create backups of TDB files that are critical to Samba's operation. This is an important step in ensuring the integrity and availability of data. By regularly backing up TDB files, you can prevent data loss in case of system failure or corruption of the TDB database.

### General Syntax:
```bash
tdbbackup [options] <db-file>
```

Where:
- `<db-file>`: The path to the TDB file you wish to back up (e.g., `secrets.tdb`, `passdb.tdb`).

### Key Options for `tdbbackup`

- `--help`: Displays help information about the command and available options.
- `--target`: Specifies the backup file name or location. If not provided, the backup will be created in the same directory with the `.bak` extension.
- `--no-rename`: Prevents renaming of existing backup files.
- `--force`: Forces the backup operation, even if the target file already exists.

### Example Usage

#### 1. Basic Backup of a TDB File

To create a backup of a specific TDB file (e.g., `secrets.tdb`), you can run the following command:

```bash
tdbbackup /var/lib/samba/private/secrets.tdb
```

This command will create a backup of the `secrets.tdb` file in the same directory with the extension `.bak` (i.e., `secrets.tdb.bak`).

#### 2. Specifying a Backup Location

If you want to store the backup in a specific directory, you can specify the target file or directory using the `--target` option:

```bash
tdbbackup --target=/path/to/backup/secrets.tdb.bak /var/lib/samba/private/secrets.tdb
```

This creates a backup of `secrets.tdb` and saves it as `/path/to/backup/secrets.tdb.bak`.

#### 3. Force Backup

If a backup file already exists at the target location and you want to overwrite it, use the `--force` option:

```bash
tdbbackup --force /var/lib/samba/private/secrets.tdb
```

#### 4. Displaying Help

To view the available options for `tdbbackup`:

```bash
tdbbackup --help
```

This will display detailed help information, including all available options and their usage.

### Use Cases for `tdbbackup`

1. **Backing up Samba Database Files**: Critical Samba configuration files like `secrets.tdb` (for machine secrets) and `passdb.tdb` (for user data) are stored in the TDB format. Regular backups of these files are essential to ensure that user data and Samba's internal configurations are not lost in the event of system failure.
   
2. **System Migration**: When migrating a Samba server to a new system or disk, backing up the TDB files ensures that you can preserve your domain's configuration, user data, and other settings.

3. **Disaster Recovery**: In case of data corruption or hardware failure, `tdbbackup` enables you to recover critical Samba database files from backups, ensuring minimal downtime and data loss.

### Example Scenario: Samba User Data Backup

You can use `tdbbackup` to back up the Samba user database (`passdb.tdb`) regularly. Here's an example command:

```bash
tdbbackup --target=/backups/passdb_backup.tdb /var/lib/samba/private/passdb.tdb
```

This command backs up the `passdb.tdb` file, which contains the Samba user information, to `/backups/passdb_backup.tdb`.

### Conclusion

The `tdbbackup` command is a valuable tool for Samba administrators who need to create backups of their critical TDB files, such as the user and secrets databases. Regular backups are essential for disaster recovery and system migration, and `tdbbackup` simplifies the process of safeguarding the important configuration and state data of Samba. By ensuring your Samba TDB files are backed up, you can mitigate the risks of data corruption, system failures, and unforeseen disasters.
