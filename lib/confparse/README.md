This is a small package the provides a ini style configuration parser. This is 
what is allowed:

- Comments start, either with the "#" or ":" anything after it, till newline is ignored
- Sections are written like the following [default] and contain a map of key values,
  anything between [] is a valid section.
- Key and values are like "ip=192.168.10.1" ,the separator is "=" otherwise will
  not be considered a key value.
- The Parser can handle bools, ints and floats (both 64bit), strings and string slices,
  as long as they are divided by columns.

Empty line are ignored, whitespaces are ignored as well. It hasn't really been
tested yet, it might have some bugs.

Suggestion: do not use this software in a production environment, it's not even
completed yet and hasn't really been tested.
