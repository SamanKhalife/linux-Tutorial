# /etc/dovecot

**Dovecot** is a popular open-source IMAP and POP3 server for Unix-like operating systems. It is known for its security and performance features. Configuration for Dovecot is done via various files located in the `/etc/dovecot/` directory. Here's a detailed overview of the key configuration files and their usage:

### Key Dovecot Configuration Files

1. **`/etc/dovecot/dovecot.conf`**
   - **Purpose:** This is the main configuration file for Dovecot. It includes basic settings and other configuration files.
   - **Key Settings:**
     - **`listen`:** Specifies the IP addresses to listen on.
     - **`protocols`:** Specifies which protocols (IMAP, POP3, etc.) Dovecot should support.
     - **`mail_location`:** Defines where the mailboxes are located.
   - **Example:**
     ```sh
     listen = *
     protocols = imap pop3
     mail_location = maildir:~/Maildir
     ```

2. **`/etc/dovecot/conf.d/10-auth.conf`**
   - **Purpose:** Configures authentication settings for Dovecot.
   - **Key Settings:**
     - **`auth_mechanisms`:** Specifies the authentication mechanisms to use (e.g., `plain`, `login`, `digest-md5`, `cram-md5`).
     - **`!include auth-sql.conf.ext`:** Includes additional configuration files for SQL-based authentication.
   - **Example:**
     ```sh
     auth_mechanisms = plain login
     ```

3. **`/etc/dovecot/conf.d/10-ssl.conf`**
   - **Purpose:** Configures SSL/TLS settings for secure communication.
   - **Key Settings:**
     - **`ssl`:** Whether SSL is enabled (`yes` or `no`).
     - **`ssl_cert`:** Path to the SSL certificate.
     - **`ssl_key`:** Path to the SSL key.
   - **Example:**
     ```sh
     ssl = yes
     ssl_cert = </etc/ssl/certs/dovecot.pem
     ssl_key = </etc/ssl/private/dovecot.key
     ```

4. **`/etc/dovecot/conf.d/10-master.conf`**
   - **Purpose:** Configures master process settings and services.
   - **Key Settings:**
     - **`service imap-login`:** Settings for the IMAP login service.
     - **`service pop3-login`:** Settings for the POP3 login service.
     - **`service auth`:** Settings for the authentication service.
   - **Example:**
     ```sh
     service imap-login {
       inet_listener imap {
         port = 0
       }
       inet_listener imaps {
         port = 993
         ssl = yes
       }
     }
     ```

5. **`/etc/dovecot/conf.d/10-mail.conf`**
   - **Purpose:** Configures mail storage settings.
   - **Key Settings:**
     - **`mail_location`:** Specifies the format and location of mail storage (e.g., `maildir`, `mbox`).
     - **`mail_privileged_group`:** Group with access to mail directories.
   - **Example:**
     ```sh
     mail_location = maildir:~/Maildir
     mail_privileged_group = mail
     ```

6. **`/etc/dovecot/conf.d/10-logging.conf`**
   - **Purpose:** Configures logging settings.
   - **Key Settings:**
     - **`log_path`:** Path to the log file.
     - **`info_log_path`:** Path to the info log file.
   - **Example:**
     ```sh
     log_path = /var/log/dovecot.log
     info_log_path = /var/log/dovecot-info.log
     ```

7. **`/etc/dovecot/conf.d/20-lmtp.conf`**
   - **Purpose:** Configures the LMTP (Local Mail Transfer Protocol) settings.
   - **Key Settings:**
     - **`service lmtp`:** Settings for the LMTP service.
   - **Example:**
     ```sh
     service lmtp {
       unix_listener /var/spool/postfix/private/dovecot-lmtp {
         mode = 0666
         user = postfix
         group = postfix
       }
     }
     ```

8. **`/etc/dovecot/conf.d/20-auth-sql.conf.ext`**
   - **Purpose:** Configures SQL-based authentication.
   - **Key Settings:**
     - **`driver`:** SQL driver (e.g., `mysql`, `pgsql`).
     - **`connect`:** Connection string to the database.
   - **Example:**
     ```sh
     driver = mysql
     connect = host=localhost dbname=mail user=mailuser password=secret
     ```

### Example Workflow

1. **Edit Main Configuration File:**
   ```sh
   sudo nano /etc/dovecot/dovecot.conf
   ```

2. **Configure Authentication:**
   ```sh
   sudo nano /etc/dovecot/conf.d/10-auth.conf
   ```

3. **Set Up SSL/TLS:**
   ```sh
   sudo nano /etc/dovecot/conf.d/10-ssl.conf
   ```

4. **Define Mail Storage:**
   ```sh
   sudo nano /etc/dovecot/conf.d/10-mail.conf
   ```

5. **Restart Dovecot Service:**
   ```sh
   sudo systemctl restart dovecot
   ```

### Conclusion

Dovecot’s configuration is highly modular, with different aspects of the service managed in separate files within the `/etc/dovecot/` directory. Understanding and configuring these files allows for tailored mail server setups with appropriate security, authentication, and logging configurations. Regularly reviewing and updating configurations ensures a secure and efficient mail server environment.
