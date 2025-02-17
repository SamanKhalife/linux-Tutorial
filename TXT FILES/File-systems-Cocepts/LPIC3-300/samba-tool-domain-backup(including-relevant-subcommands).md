# samba-tool domain backup
The `samba-tool domain backup` command is part of the **Samba** suite of utilities, specifically designed for managing and backing up a **Samba Active Directory (AD)** domain controller. This command provides tools for creating and restoring backups of the domain controller, ensuring that you can recover your domain controller’s state if something goes wrong.

### Overview of `samba-tool domain backup`

This command helps administrators back up various critical components of the Samba AD Domain Controller. Backups are essential for ensuring that the Active Directory (AD) configuration, user data, and other vital information can be restored in case of failure or disaster.

The `samba-tool domain backup` command has several subcommands that allow for different types of backups, such as full backups or just specific components like the database or SYSVOL (a folder for storing system data and scripts).

### Syntax

```bash
samba-tool domain backup <subcommand> [options]
```

Where `<subcommand>` could be one of the relevant backup-related commands.

### Common Subcommands and Options

#### 1. **`samba-tool domain backup offline`**
   - **Description**: This subcommand allows you to create an offline backup of the domain controller. The backup is performed while the domain controller is offline, ensuring that no changes occur during the backup process.
   - **Usage**:
     ```bash
     samba-tool domain backup offline /path/to/backup
     ```
   - **Options**:
     - `--targetdir=<path>`: The path where the backup will be stored. (Required)
     - `--level=<level>`: The level of backup, e.g., full or incremental.
     - `--include=<component>`: Optionally specify which components to include in the backup (e.g., database, sysvol, etc.).

   - **Example**:
     ```bash
     samba-tool domain backup offline /var/lib/samba/backup --level=full
     ```

#### 2. **`samba-tool domain backup online`**
   - **Description**: Performs an online backup of the domain controller. This backup is done while the domain controller is still running, and no downtime is required.
   - **Usage**:
     ```bash
     samba-tool domain backup online /path/to/backup
     ```
   - **Options**:
     - `--targetdir=<path>`: The path where the backup will be stored. (Required)
     - `--include=<component>`: Optionally specify components to include in the backup (e.g., database, sysvol, etc.).
     - `--level=<level>`: Specifies whether to perform a full backup or an incremental backup.

   - **Example**:
     ```bash
     samba-tool domain backup online /var/lib/samba/backup --level=full
     ```

#### 3. **`samba-tool domain backup restore`**
   - **Description**: This subcommand is used to restore a backup of the domain controller. You can restore from either an online or offline backup. It is essential for disaster recovery operations.
   - **Usage**:
     ```bash
     samba-tool domain backup restore /path/to/backup
     ```
   - **Options**:
     - `--targetdir=<path>`: The directory where the backup is located. (Required)
     - `--no-preserve-timestamps`: Optionally disables preserving timestamps during restore.
   
   - **Example**:
     ```bash
     samba-tool domain backup restore /var/lib/samba/backup --no-preserve-timestamps
     ```

#### 4. **`samba-tool domain backup status`**
   - **Description**: Displays the status of the most recent backup, including information about whether the backup was successful, and its components.
   - **Usage**:
     ```bash
     samba-tool domain backup status
     ```

#### 5. **`samba-tool domain backup delete`**
   - **Description**: Removes a specific backup from the system. This is useful when cleaning up old backups that are no longer needed.
   - **Usage**:
     ```bash
     samba-tool domain backup delete /path/to/backup
     ```
   - **Example**:
     ```bash
     samba-tool domain backup delete /var/lib/samba/backup
     ```

#### 6. **`samba-tool domain backup check`**
   - **Description**: Verifies the integrity of a backup and checks if the files and database are consistent.
   - **Usage**:
     ```bash
     samba-tool domain backup check /path/to/backup
     ```

### Backup Components
When performing backups, you have the ability to specify which components you want to back up. Common components include:
- **`--include=database`**: Backs up the Active Directory database.
- **`--include=sysvol`**: Backs up the SYSVOL share (important for group policies and logon scripts).
- **`--include=secrets`**: Backs up the secrets database, which includes domain controller credentials and keys.

### Example Backup Procedures

#### 1. Full Offline Backup
A full offline backup ensures that the backup is taken when the domain controller is not making any changes. This prevents data corruption during the backup.

```bash
samba-tool domain backup offline /path/to/backup --level=full
```

This command would create a complete offline backup in the `/path/to/backup` directory.

#### 2. Online Backup (without downtime)
An online backup is performed while the domain controller is still running. This is useful for systems that require high availability.

```bash
samba-tool domain backup online /path/to/backup --level=full
```

This would create a full backup while the system remains online.

#### 3. Restore Backup
To restore a backup, simply use the restore subcommand with the backup location.

```bash
samba-tool domain backup restore /path/to/backup
```

This command would restore the backup located at `/path/to/backup`.

### Common Use Cases for `samba-tool domain backup`
- **Disaster Recovery**: Regularly backing up the Samba domain controller ensures that you can restore the domain in case of a failure or corruption.
- **Migration**: Backing up and restoring can be a crucial part of migrating from one domain controller to another, ensuring that no data is lost.
- **Replication Testing**: If testing domain controller replication or upgrading Samba, performing a backup before and after the test ensures you can easily roll back changes.

### Important Notes
- **Backup Frequency**: It's essential to regularly back up your Samba domain controller, especially before major changes or upgrades.
- **Backup Integrity**: Always check that your backups are valid and can be restored when needed. You can use the `samba-tool domain backup check` command for this purpose.
- **Backup Location**: Store backups in a secure, off-site location to ensure they are safe from local disasters like hardware failure or ransomware attacks.

### Conclusion
The `samba-tool domain backup` command is a powerful tool for managing backups of your Samba domain controllers. It allows for both offline and online backups and provides various options for flexibility in backing up specific components like the database or SYSVOL. Regular backups and testing restores are critical to ensure that you can recover your domain in case of failure.
