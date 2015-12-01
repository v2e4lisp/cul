cul
=======

cul - count uniq lines.

A faster version of `sort | uniq | wc -l`


Usage
----------

```bash

# single file
cul /tmp/file1

# multiple files
cul /tmp/file1 /tmp/file2

# pipe
cat /tmp/file1 | cul

# multiple files and pipe
cat /tmp/file1 | cul /tmp/file2 /tmp/file3

```





