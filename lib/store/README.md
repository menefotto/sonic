Package store

This package implements a simple storage with two backends, an in memory backend 
and a persistent backends using boltdb.
To implement another backend using a database of your choice or data structure, 
you have to implement the DB interface.
The backend is passed to the graph during construction as a parameter.

Disclaimer:
- Do not use this software in a production environment, it hasn't been tested or
  is ready for production this is a very early release. It hasn't been finished
  yet even if it works as it should.

- It is quite limited for the moment only accepting strings as key, tough anything
  marshable by the json standard library can be passed as a edge. The persistent
  implementation is quite fast.

- Documentation is missing at the moment you can look at tests to see an example.

