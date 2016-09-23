Package xzutils

This package provides 4 functions that give an easier to use interface to common
operations, the backend is go-liblzma.

-XzCompress, XzDecompress they either take []byte or string as input the give back
 a byte slice and a error( which is nil in case of success).

-FileXzCompress and FileXzDecompress take file names as input, Compress takes two
 files input and output file and writes the compressed content of input file to 
 output file. FileXzDecompress takes an input file name and gives back a decompressed
 array.

Disclaimer:
- Do not use this lib in production is barely tested.



