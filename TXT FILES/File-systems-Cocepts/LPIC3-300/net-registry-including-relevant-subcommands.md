# net registry
The `net` command in Linux is part of the **Samba** suite of tools and is used to interact with network shares, domains, and more. One important component of the `net` command is **net registry**, which allows administrators to query and manage the **Windows registry** on remote systems or Samba servers. While the `net registry` command has many useful subcommands for handling registry tasks, it is primarily used in environments where Samba is deployed to integrate Linux/UNIX systems with Windows environments.

### `net registry` Overview:
The `net registry` command enables remote manipulation and querying of the Windows registry via the Samba protocol. You can use it to query registry values, add or modify entries, and delete registry keys remotely, especially useful when managing Windows environments or Samba servers configured with registry backends.

### Key Subcommands of `net registry`:
Here are some of the most relevant subcommands for managing Windows registry entries remotely:

1. **`net registry enumerate`**:
   - Lists the contents of a registry key, including all subkeys and values.
   - Example:
     ```bash
     net registry enumerate \\remote_machine HKLM\Software
     ```
     This command will list the contents of the **HKEY_LOCAL_MACHINE\Software** key on the remote machine `\\remote_machine`.

2. **`net registry get`**:
   - Retrieves the value of a specific registry entry.
   - Example:
     ```bash
     net registry get \\remote_machine HKLM\Software\Example\Value
     ```
     This command retrieves the value of the `Example\Value` registry key on the remote machine.

3. **`net registry set`**:
   - Modifies or sets a value in the registry.
   - Example:
     ```bash
     net registry set \\remote_machine HKLM\Software\Example\NewValue DWORD 1
     ```
     This command sets the `NewValue` to `1` (a DWORD value) in the **HKEY_LOCAL_MACHINE\Software\Example** key on the remote machine.

4. **`net registry delete`**:
   - Deletes a specific registry key or value.
   - Example:
     ```bash
     net registry delete \\remote_machine HKLM\Software\Example\OldValue
     ```
     This command deletes the `OldValue` registry entry from **HKEY_LOCAL_MACHINE\Software\Example**.

5. **`net registry export`**:
   - Exports a registry key or hive to a file for backup or replication purposes.
   - Example:
     ```bash
     net registry export \\remote_machine HKLM\Software /path/to/exportfile.reg
     ```
     This exports the **HKEY_LOCAL_MACHINE\Software** registry key to a file named `exportfile.reg`.

6. **`net registry import`**:
   - Imports a previously exported registry file back into the registry.
   - Example:
     ```bash
     net registry import \\remote_machine /path/to/importfile.reg
     ```
     This imports the registry values from `importfile.reg` back into the remote system's registry.

7. **`net registry query`**:
   - Allows for advanced querying of the registry based on specific conditions.
   - Example:
     ```bash
     net registry query \\remote_machine HKLM\Software\Example WHERE "ValueName = 'Sample'"
     ```
     This queries the registry key for specific entries that match the condition.

8. **`net registry add`**:
   - Adds a new registry key or value.
   - Example:
     ```bash
     net registry add \\remote_machine HKLM\Software\Example NewValue DWORD 0
     ```
     This adds a new DWORD value called `NewValue` with a value of `0` to the **HKEY_LOCAL_MACHINE\Software\Example** key.

### Practical Use Cases:
- **Remote Management**: Administrators can use the `net registry` commands to remotely administer the Windows registry on networked machines, avoiding the need for direct access to the systems.
- **Automating Tasks**: The commands can be included in scripts to automate routine registry tasks such as setting configuration values, deploying new keys, or backing up important settings.
- **Troubleshooting**: Use `net registry` to query and modify registry settings on remote systems when troubleshooting issues related to software or network configurations.
- **Samba Configuration**: In environments where Samba uses the registry backend for storing configuration, `net registry` can be used to manage Samba-related settings directly from the command line.

### Conclusion:
The `net registry` command offers powerful functionality for managing and interacting with the Windows registry remotely from a Linux or Samba environment. Whether you're performing basic registry queries or making complex modifications, it provides administrators with a versatile toolset to manage systems in mixed-OS environments, offering flexibility and efficiency for large-scale IT infrastructure management.
