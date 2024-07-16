# /etc/strongswan.d/.

The `/etc/strongswan.d/` directory is used for modular configuration files for the strongSwan VPN suite. This directory allows for more granular and organized configuration management by splitting the settings into multiple files. Each configuration file within this directory typically corresponds to a specific plugin or feature within strongSwan.

#### Purpose

The purpose of the `/etc/strongswan.d/` directory is to house individual configuration files for various plugins and components of strongSwan. This modular approach facilitates easier management, customization, and troubleshooting of the VPN configuration.

#### Structure

The `/etc/strongswan.d/` directory contains subdirectories and files that configure different aspects of the strongSwan service. The files are usually in the INI format and are named after the plugin or component they configure.

#### Common Subdirectories and Files

- **charon/**: Configuration files related to the Charon daemon, which handles IKE (Internet Key Exchange) and IPsec.
  - **charon-logging.conf**: Configures logging for Charon.
  - **charon-plugins.conf**: Lists and configures plugins for Charon.
- **starter/**: Configuration files related to the starter daemon, which is responsible for starting strongSwan.
  - **starter-logging.conf**: Configures logging for the starter.
  - **starter-plugins.conf**: Lists and configures plugins for the starter.
- **plugins/**: Contains configuration files for individual plugins.
  - **load-tester.conf**: Configuration for the load tester plugin.
  - **attr-sql.conf**: Configuration for the attribute SQL plugin.

### Example Configuration Files

#### charon-logging.conf

This file configures the logging for the Charon daemon. 

```ini
charon {
    filelog {
        /var/log/charon.log {
            time_format = %b %e %T
            append = no
            default = 1
            flush_line = yes
        }
    }

    syslog {
        identifier = charon
        daemon {
            ike = 2
            cfg = 2
            knl = 2
            net = 2
        }
        auth {
            peer = 2
        }
    }
}
```

#### charon-plugins.conf

This file lists and configures the plugins for the Charon daemon.

```ini
charon {
    plugins {
        include strongswan.d/charon/*.conf
    }
}
```

#### load-tester.conf

This file configures the load tester plugin.

```ini
load-tester {
    initiate = 10
    delay = 100
    shutdown_when_complete = yes
    version = 1
    request_count = 1000
}
```

### Using `/etc/strongswan.d/`

When configuring strongSwan, you typically modify or create specific configuration files within `/etc/strongswan.d/` to enable or adjust the settings for the desired plugins or features. After making changes, you need to restart the strongSwan service to apply the new configurations:

```bash
sudo systemctl restart strongswan
```

### Conclusion

The `/etc/strongswan.d/` directory provides a structured and modular way to manage strongSwan configurations. By dividing the configuration into multiple files, it allows for better organization, easier customization, and more straightforward troubleshooting. Each file within this directory focuses on specific aspects of the strongSwan service, such as logging, plugins, or the Charon and starter daemons, thereby offering a flexible and powerful approach to configuring your VPN setup.
