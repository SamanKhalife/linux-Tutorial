# include 
The **`include`** directive in the Samba configuration (`smb.conf`) allows you to reference and include additional configuration files. This is useful for modularizing your Samba configuration, managing complex setups, or dynamically including different configurations based on conditions like user, host, or operating system.

### Key Points about `include`:

1. **Purpose**:
   - To split your configuration into multiple files, making it easier to manage complex setups.
   - To include different configuration settings based on dynamic variables such as the machine's IP address, the user accessing the share, or specific operating systems.

2. **Usage in `smb.conf`**:
   - The `include` directive is followed by a path to another configuration file.
   - You can also use variable substitutions (like `%U` for the username, `%H` for the home directory, etc.) to dynamically include files based on the environment.

   **Basic Example**:
   ```ini
   include = /etc/samba/smb.custom.conf
   ```

3. **Dynamic Inclusion**:
   - You can dynamically include different configuration files based on the client, user, or other variables. For example, you can include configuration files based on the accessing user or the machine's IP address.

   **Example of Dynamic Inclusion**:
   ```ini
   include = /etc/samba/smb.conf.%U  # Includes a user-specific config file
   ```

4. **Common Use Cases**:
   - **Separate Configurations for Different Hosts**: Different machines or subnets might have separate configurations.
   - **User-Specific Configurations**: Customize settings based on who is accessing the share.
   - **Operating System-Specific Configurations**: Different platforms may require different Samba settings.
   
   **Example**:
   ```ini
   include = /etc/samba/smb.conf.%m  # Includes config based on the machine's NetBIOS name
   ```

   **Example with Host IP**:
   ```ini
   include = /etc/samba/smb.conf.%I  # Includes config based on the client's IP address
   ```

### Variables in `include`:

Samba allows the use of variables to provide dynamic behavior in configuration files. These variables can be substituted based on the environment or connection details. Here are some commonly used variables:

- **`%U`**: The username of the connected user.
- **`%m`**: The NetBIOS name of the client machine.
- **`%I`**: The IP address of the client.
- **`%L`**: The server's NetBIOS name.
- **`%H`**: The home directory of the user.
- **`%D`**: The domain or workgroup of the user.
- **`%T`**: The current date and time.

### Example Scenario:

Suppose you want to have user-specific configurations so that different users accessing Samba shares can have their own custom configurations. You can use the `include` directive to point to user-specific config files.

```ini
[global]
   include = /etc/samba/smb.conf.%U

[share]
   path = /srv/samba/shared
   browseable = yes
   read only = no
```

This setup will load `/etc/samba/smb.conf.user1` when `user1` connects, and `/etc/samba/smb.conf.user2` when `user2` connects.

### Advantages of `include`:

- **Modularity**: Large or complex configurations can be split into smaller, more manageable files.
- **Dynamic Configuration**: Different users, machines, or subnets can have different configurations without needing to maintain a single massive `smb.conf` file.
- **Scalability**: Easier to scale and maintain Samba setups in environments with multiple users, machines, or operating systems.

### Conclusion:
The `include` directive in Samba's `smb.conf` is a powerful tool that allows you to manage complex or dynamic Samba configurations. It enables the modularization of settings, and with variable substitution, you can dynamically include different configuration files based on various parameters such as users or machine IP addresses.
