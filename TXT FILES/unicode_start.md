# unicode_start

The `unicode_start` function is a Python function that returns the Unicode code point for the first character in a string. The function is defined as follows:

```python
def unicode_start(string):
    """Returns the Unicode code point for the first character in a string."""

    if not isinstance(string, str):
        raise TypeError("Argument must be a string.")

    return ord(string[0])
```

The `unicode_start` function takes a single argument, which is the string that you want to get the Unicode code point for. The function returns the Unicode code point for the first character in the string.

For example, the following code will print the Unicode code point for the first character in the string "hello":

```python
>>> unicode_start("hello")
104
```

The output of the code is 104, which is the Unicode code point for the letter "h".

The `unicode_start` function is a useful function for getting the Unicode code point for the first character in a string. It can be used for a variety of tasks, such as finding the encoding of a string or determining whether a string contains a certain character.
