This package is a think wrapper around the standard library package syslog. Since
logging is a blocking operation can become quite expensive, specially in logging
heavy applications. This package is born in order to reduce to minum or totally
avoid, blocking in fact offers two main functions Log and MustLog.
The first of the two offers no blocking but does not grant that the message will
be logged, MustLog should instead grant the message will be logged.


-Package syslogger provides a NewSysLogger() func that given a prefix string
 and a syslog priority construct without failing a SysLogger defined as a struct
 embeding a chan of messages, a done channel and a logger.

-The other provided functions are Log() and MustLog() wich sends the message, 
 and Close() which close the logger is it important to close the logger otherwise 
 we would leak the gourutine who does the job in the background.



Disclaimer and caveats:
-Do not use this package in production it has not been tested, it's in early 
 developement stages. More test on slower machines are needed.

-If logging with Log and closing immeditelly after, the message might not be,
 logged in such a case use MustLog
