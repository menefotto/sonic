This two packages implement two super simple functions:

- Package limiter does what is says, it wraps a function and runs it, limiting the
  number of goroutunes that can be run at the same time ( number can be toogle with
  a simple exposed variable ).

- Package retry does what is says as well if a function calls fails ( resulting in
  a error rather than a nil value ), runs is again.
  As arguments takes also the number of retries and a sleep value specified 
  in milliseconds.

The package limiter in basically a shameless copy from this golang talk, all the credits
go to the guy https://www.youtube.com/watch?v=yeetIgNeIkc. Tough I am almost sure
I am going to rewrite it in a near future. This is what a call nice to have middleware.
