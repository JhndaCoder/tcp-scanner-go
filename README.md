# tcp-scanner-go

## The code includes the following

- **dial** : basic usage of net package and dial function
- **tcp-scanner-slow** : a basic tcp scanner
- **tcp-scanner-fast** : a tcp scanner but much faster as multiple ports are being scanned concurrently using goroutines
- **tcp-sync-scanner** : implementation of a worker pool that use a pool of goroutines to manage the concurrent work being performed.
- **tcp-scanner-final** : a highly efficient port scanner
- **echo-server** : a basic tcp server that reads and writes data from a socket
- **port-forwarder** : a basic port-forwarder to proxy a connection through an intermediary service or host
- **netcat-replica** : a simple implementation of netcat that allows stdin and stdout of any arbitrary program to be redirected over TCP
