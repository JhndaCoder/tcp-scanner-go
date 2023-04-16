# tcp-scanner-go

## The code includes the following

- **dial** : basic usage of net package and dial function
- **tcp-scanner-slow** : a basic tcp scanner
- **tcp-scanner-fast** : a tcp scanner but much faster as multiple ports are being scanned concurrently using goroutines
- **tcp-sync-scanner** : implementation of a worker pool that use a pool of goroutines to manage the concurrent work being performed.
- **tcp-scanner-final** : a highly efficient port scanner
