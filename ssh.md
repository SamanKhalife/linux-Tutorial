# Config file parameters you can use 


best parameters to set 

```
# Common SSH configuration file
# Place this file under ~/.ssh/config for user-specific configurations
# or /etc/ssh/ssh_config for system-wide configurations.

# Global settings
Host *
    # Address family: any, inet, inet6
    AddressFamily any
    
    # Batch mode: yes or no (default: no)
    BatchMode no
    
    # Check host IP in known_hosts file: yes or no (default: yes)
    CheckHostIP yes
    
    # Clear all forwardings: yes or no (default: no)
    ClearAllForwardings no
    
    # Compression: yes or no (default: no)
    Compression no
    
    # Connection attempts before exiting (default: 1)
    ConnectionAttempts 3
    
    # Connect timeout in seconds
    ConnectTimeout 10
    
    # Enable SSH keysign for HostbasedAuthentication: yes or no (default: no)
    EnableSSHKeysign no
    
    # Forward agent connection: yes or no (default: no)
    ForwardAgent no
    
    # Forward X11 connections: yes or no (default: no)
    ForwardX11 no
    
    # Allow remote hosts to connect to local forwarded ports: yes or no (default: no)
    GatewayPorts no
    
    # Hash host names in known_hosts file: yes or no (default: no)
    HashKnownHosts yes
    
    # Try hostbased authentication: yes or no (default: no)
    HostbasedAuthentication no
    
    # Authentication identity file
    IdentityFile ~/.ssh/id_rsa
    
    # Keyboard-interactive authentication: yes or no (default: yes)
    KbdInteractiveAuthentication yes
    
    # Log level (default: INFO)
    LogLevel INFO
    
    # Password authentication: yes or no (default: yes)
    PasswordAuthentication yes
    
    # Port number to connect on the remote host (default: 22)
    Port 22
    
    # Preferred authentication methods (default shown below)
    PreferredAuthentications publickey,password,keyboard-interactive
    
    # Protocol versions to support (default: 2,1)
    Protocol 2
    
    # Public key authentication: yes or no (default: yes)
    PubkeyAuthentication yes
    
    # Maximum number of server alive messages sent without receiving a reply (default: 3)
    ServerAliveCountMax 3
    
    # Timeout interval in seconds to send a message through the encrypted channel (default: 0)
    ServerAliveInterval 60
    
    # Strict host key checking: yes or no (default: ask)
    StrictHostKeyChecking ask
    
    # TCP keepalive: yes or no (default: yes)
    TCPKeepAlive yes

# Host-specific settings (example for a specific host)
Host example.com
    # Use IPv4 only
    AddressFamily inet 
             # or inet6 
    
    # Enable agent forwarding
    ForwardAgent yes
    
    # Set specific port
    Port 2200
    
    # Use a specific identity file
    IdentityFile ~/.ssh/id_rsa_example
    
    # Set log level to DEBUG for this host
    LogLevel DEBUG
    
    # Set a local forward
    LocalForward 8080 localhost:80

# More host-specific settings can be added below following the same format

```




## AddressFamily
Specifies which address family to use when connecting. Valid arguments are ''any'', ''inet'' (use IPv4 only), or ''inet6'' (use IPv6 only).

## BatchMode
If set to ''yes'', passphrase/password querying will be disabled. This option is useful in scripts and other batch jobs where no user is present to supply the password. The argument must be ''yes'' or ''no''. The default is ''no''.

## BindAddress
Use the specified address on the local machine as the source address of the connection. Only useful on systems with more than one address. Note that this option does not work if UsePrivilegedPort is set to ''yes''.

## ChallengeResponseAuthentication
Specifies whether to use challenge-response authentication. The argument to this keyword must be ''yes'' or ''no''. The default is ''yes''.

## CheckHostIP
If this flag is set to ''yes'', ssh(1) will additionally check the host IP address in the known_hosts file. This allows ssh to detect if a host key changed due to DNS spoofing. If the option is set to ''no'', the check will not be executed. The default is ''yes''.

## Cipher
Specifies the cipher to use for encrypting the session in protocol version 1. Currently, ''blowfish'', ''3des'', and ''des'' are supported. des is only supported in the ssh(1) client for interoperability with legacy protocol 1 implementations that do not support the 3des cipher. Its use is strongly discouraged due to cryptographic weaknesses. The default is ''3des''.

## Ciphers
Specifies the ciphers allowed for protocol version 2 in order of preference. Multiple ciphers must be comma-separated. The supported ciphers are ''3des-cbc'', ''aes128-cbc'', ''aes192-cbc'', ''aes256-cbc'', ''aes128-ctr'', ''aes192-ctr'', ''aes256-ctr'', ''arcfour128'', ''arcfour256'', ''arcfour'', ''blowfish-cbc'', and ''cast128-cbc''. The default is:

## Compression
Specifies whether to use compression. The argument must be ''yes'' or ''no''. The default is ''no''.

## CompressionLevel
Specifies the compression level to use if compression is enabled. The argument must be an integer from 1 (fast) to 9 (slow, best). The default level is 6, which is good for most applications. The meaning of the values is the same as in gzip(1). Note that this option applies to protocol version 1 only.

## ConnectionAttempts
Specifies the number of tries (one per second) to make before exiting. The argument must be an integer. This may be useful in scripts if the connection sometimes fails. The default is 1.

## ConnectTimeout
Specifies the timeout (in seconds) used when connecting to the SSH server, instead of using the default system TCP timeout. This value is used only when the target is down or really unreachable, not when it refuses the connection.


## EnableSSHKeysign
Setting this option to ''yes'' in the global client configuration file /etc/ssh/ssh_config enables the use of the helper program ssh-keysign(8) during HostbasedAuthentication. The argument must be ''yes'' or ''no''. The default is ''no''. This option should be placed in the non-hostspecific section. See ssh-keysign(8) for more information.

## EscapeChar
Sets the escape character (default: '~'). The escape character can also be set on the command line. The argument should be a single character, '^' followed by a letter, or ''none'' to disable the escape character entirely (making the connection transparent for binary data).

## ExitOnForwardFailure
Specifies whether ssh(1) should terminate the connection if it cannot set up all requested dynamic, tunnel, local, and remote port forwardings. The argument must be ''yes'' or ''no''. The default is ''no''.

## ForwardAgent
Specifies whether the connection to the authentication agent (if any) will be forwarded to the remote machine. The argument must be ''yes'' or ''no''. The default is ''no''.

Agent forwarding should be enabled with caution. Users with the ability to bypass file permissions on the remote host (for the agent's Unix-domain socket) can access the local agent through the forwarded connection. An attacker cannot obtain key material from the agent, however they can perform operations on the keys that enable them to authenticate using the identities loaded into the agent.

## GatewayPorts
Specifies whether remote hosts are allowed to connect to local forwarded ports. By default, ssh(1) binds local port forwardings to the loopback address. This prevents other remote hosts from connecting to forwarded ports. GatewayPorts can be used to specify that ssh should bind local port forwardings to the wildcard address, thus allowing remote hosts to connect to forwarded ports. The argument must be ''yes'' or ''no''. The default is ''no''.

## GlobalKnownHostsFile
Specifies a file to use for the global host key database instead of /etc/ssh/ssh_known_hosts.

## HashKnownHosts
Indicates that ssh(1) should hash host names and addresses when they are added to ~/.ssh/known_hosts. These hashed names may be used normally by ssh(1) and sshd(8), but they do not reveal identifying information should the file's contents be disclosed. The default is ''no''. Note that existing names and addresses in known hosts files will not be converted automatically, but may be manually hashed using ssh-keygen(1).

## HostbasedAuthentication
Specifies whether to try rhosts based authentication with public key authentication. The argument must be ''yes'' or ''no''. The default is ''no''. This option applies to protocol version 2 only and is similar to RhostsRSAAuthentication.

## HostKeyAlgorithms
Specifies the protocol version 2 host key algorithms that the client wants to use in order of preference. The default for this option is: ''ssh-rsa,ssh-dss''.

## HostKeyAlias
Specifies an alias that should be used instead of the real host name when looking up or saving the host key in the host key database files. This option is useful for tunneling SSH connections or for multiple servers running on a single host.

## HostName
Specifies the real host name to log into. This can be used to specify nicknames or abbreviations for hosts. The default is the name given on the command line. Numeric IP addresses are also permitted (both on the command line and in HostName specifications).

## IdentitiesOnly
Specifies that ssh(1) should only use the authentication identity files configured in the ssh_config files, even if ssh-agent(1) offers more identities. The argument to this keyword must be ''yes'' or ''no''. This option is intended for situations where ssh-agent offers many different identities. The default is ''no''.

## IdentityFile
Specifies a file from which the user's RSA or DSA authentication identity is read. The default is ~/.ssh/identity for protocol version 1, and ~/.ssh/id_rsa and ~/.ssh/id_dsa for protocol version 2. Additionally, any identities represented by the authentication agent will be used for authentication.

The file name may use the tilde syntax to refer to a user's home directory or one of the following escape characters: '%d' (local user's home directory), '%u' (local user name), '%l' (local host name), '%h' (remote host name) or '%r' (remote user name).

It is possible to have multiple identity files specified in configuration files; all these identities will be tried in sequence.

## KbdInteractiveAuthentication
Specifies whether to use keyboard-interactive authentication. The argument to this keyword must be ''yes'' or ''no''. The default is ''yes''.

## KbdInteractiveDevices
Specifies the list of methods to use in keyboard-interactive authentication. Multiple method names must be comma-separated. The default is to use the server specified list. The methods available vary depending on what the server supports. For an OpenSSH server, it may be zero or more of: ''bsdauth'', ''pam'', and ''skey''.

## LocalCommand
Specifies a command to execute on the local machine after successfully connecting to the server. The command string extends to the end of the line, and is executed with the user's shell. The following escape character substitutions will be performed: '%d' (local user's home directory), '%h' (remote host name), '%l' (local host name), '%n' (host name as provided on the command line), '%p' (remote port), '%r' (remote user name) or '%u' (local user name). This directive is ignored unless PermitLocalCommand has been enabled.

## LocalForward
Specifies that a TCP port on the local machine be forwarded over the secure channel to the specified host and port from the remote machine. The first argument must be [
bind_address:]port and the second argument must be host:hostport. IPv6 addresses can be specified by enclosing addresses in square brackets or by using an alternative syntax: [
bind_address/]port and host/hostport. Multiple forwardings may be specified, and additional forwardings can be given on the command line. Only the superuser can forward privileged ports. By default, the local port is bound in accordance with the GatewayPorts setting. However, an explicit bind_address may be used to bind the connection to a specific address. The bind_address of ''localhost'' indicates that the listening port be bound for local use only, while an empty address or '*' indicates that the port should be available from all interfaces.

## LogLevel
Gives the verbosity level that is used when logging messages from ssh(1). The possible values are: QUIET, FATAL, ERROR, INFO, VERBOSE, DEBUG, DEBUG1, DEBUG2, and DEBUG3. The default is INFO. DEBUG and DEBUG1 are equivalent. DEBUG2 and DEBUG3 each specify higher levels of verbose output.

MACs' Specifies the MAC (message authentication code) algorithms in order of preference. The MAC algorithm is used in protocol version 2 for data integrity protection. Multiple algorithms must be comma-separated. 
