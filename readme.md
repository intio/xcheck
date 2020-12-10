# xcheck

Simple healthcheck client for X11 servers.

Try to open a connection to an X11 display. If connection succeeds,
exit with a 0 status. If any error occurs, log the error to stderr and
exit with a non-0 status. Display is taken from the `DISPLAY`
environment variable.

That's it.

This code is in public domain.
