# now  
A simple cli tool to convert unix timestamps or human readable dates.

# Install
```bash
go install github.com/gloomyzerg/now
```

# Usage
```bash
now #output now unix timestamp
now 1637644574 #output 2021-11-23 13:16:14
now 2021-11-23 #output 1637596800
now 2021-11-23 10:00 #output 1637632800
```

```bash
now -h

Usage:
  now [timestamp]
  now [YYYY-MM-DD]
  now [YYYY-MM-DD H:i]
  now [YYYY-MM-DD H:i:s]

Flags:
  -h, --help   help for now
```