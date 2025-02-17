# REG_SZ and REG_MULTI_SZ
In Windows Registry, `REG_SZ` and `REG_MULTI_SZ` are two common data types used for storing different kinds of information.

### 1. **REG_SZ**:
- **Description**: 
  - `REG_SZ` (String Value) is a simple string data type. It is used to store a single, readable text string.
  - It is one of the most commonly used types in the registry for configuration values like file paths, environment variables, or other textual data.

- **Examples of Usage**:
  - Storing file paths: `C:\Program Files\ExampleApp`
  - Storing application version: `1.0.0`

- **Registry Example**:
  - A key might look like this:
    ```
    "ApplicationPath" = REG_SZ "C:\Program Files\MyApp"
    ```

### 2. **REG_MULTI_SZ**:
- **Description**: 
  - `REG_MULTI_SZ` (Multiple String Value) is used to store multiple strings in a single value. It can be thought of as an array of strings. Each string is separated by a null (`\0`) character, and the entire list is terminated with an additional null character.
  - It is useful for storing lists of items such as multi-line text, environment variables, or multiple file paths.

- **Examples of Usage**:
  - Storing a list of search paths.
  - Storing a list of DNS server IP addresses.

- **Registry Example**:
  - A key might look like this:
    ```
    "SearchPaths" = REG_MULTI_SZ "C:\Path1\0C:\Path2\0C:\Path3\0"
    ```

### Differences Between `REG_SZ` and `REG_MULTI_SZ`:
- **Data Type**: `REG_SZ` holds a single string, while `REG_MULTI_SZ` holds a list of strings.
- **Usage**: `REG_SZ` is used for simple single-value strings (like file paths or text data), while `REG_MULTI_SZ` is useful for when you need to store multiple values in a single registry entry (like lists of search paths or environment variables).
