# capsh

The `capsh` (capabilities shell) command in Linux is a tool used to examine and modify the capabilities of the current process. It can be used to set or drop capabilities, enter a capabilities-aware shell, and more. This tool is particularly useful for debugging and testing capabilities.

### Overview of `capsh`

#### Purpose

`capsh` provides a way to:
- Launch a shell or execute commands with specified capabilities.
- Drop or add capabilities for the current shell session.
- View the current capabilities of a process.

### Basic Usage

The general syntax for `capsh` is:

```bash
capsh [options]
```

### Example Usages

1. **Starting a Shell with Specific Capabilities**

   To start a shell with specific capabilities, you can use the `--caps` option followed by the list of capabilities:

   ```bash
   capsh --caps="cap_net_raw,cap_net_admin+ep" -- -c /bin/bash
   ```

   This command starts a new shell with `cap_net_raw` and `cap_net_admin` capabilities.

2. **Dropping All Capabilities**

   To drop all capabilities from the current shell, you can use the `--drop` option:

   ```bash
   capsh --drop=all -- -c /bin/bash
   ```

   This command starts a new shell with all capabilities dropped.

3. **Viewing Current Capabilities**

   To view the current capabilities of the shell, you can use the `--print` option:

   ```bash
   capsh --print
   ```

   This command prints the current capabilities of the process.

4. **Entering a Shell with Inherited Capabilities**

   To enter a shell that inherits the current capabilities:

   ```bash
   capsh --inh="cap_net_raw+eip" -- -c /bin/bash
   ```

   This command starts a new shell with the `cap_net_raw` capability set as inheritable, effective, and permitted.

### Capabilities Format

Capabilities can be specified in the following format:

- `capability_name`: The name of the capability.
- `+`: Indicates that the capability should be added.
- `-`: Indicates that the capability should be removed.
- `=`: Indicates that the capability should be set.
- `e`: Effective flag, the capability is in the set of effective capabilities.
- `p`: Permitted flag, the capability is in the set of permitted capabilities.
- `i`: Inheritable flag, the capability is in the set of inheritable capabilities.

### Common Capabilities

Here are some common capabilities and their purposes:

- `cap_net_raw`: Perform raw network operations.
- `cap_net_admin`: Perform various network-related administrative tasks.
- `cap_sys_admin`: Perform various system administrative tasks.
- `cap_sys_ptrace`: Trace arbitrary processes using `ptrace`.

### Example Capabilities and Their Meanings

- `cap_net_raw+ep`: Grants the process the ability to perform raw network operations with effective (`e`) and permitted (`p`) flags.
- `cap_net_admin+eip`: Grants the process network administration capabilities with effective (`e`), inheritable (`i`), and permitted (`p`) flags.

### Conclusion

`capsh` is a versatile tool for managing and debugging process capabilities in Linux. It allows administrators and developers to fine-tune the privileges of processes, thereby enhancing system security by limiting the scope of capabilities granted to executables and shell sessions.
