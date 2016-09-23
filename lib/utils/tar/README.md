Package tarutils

This package provides a two functions one which extract from a tar archive the list 
of file and directories inside, and the other one checks if the given file is a valid
tar archive.

-TarExtractor(filename) returns given a valid tar file, a map[path/filename][]byte
 , so that each key is either only path or path/filename with a byte slice containing
 data, a len of 0 is a valid byte slice content indicating the key is only a path.

-IsTarFile(filename) returns true if the file is an actual tar file. If not returns
 false, ErrNotTarFile

Disclaimer:
-This is barely tested and you should not use on a production enviroment.
