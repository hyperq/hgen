# hgen

a tool for wood to generate go code with mysql data table

## Install

```bash
go get github.com/hyperq/hgen
```

## How to use

```bash
hgen -d dbname -t user
```

```bash
a tool for wood to generate go code with mysql data table
a tool for wood to generate go code with mysql data table

Usage:
  hgen [flags]

Flags:
  -e, --cache              is cache
  -c, --comment            swagger comment
  -d, --dbname string      dbname
  -h, --help               help for hgen
  -i, --ip string          ip (default "127.0.0.1")
  -p, --password string    mysql password (default "123456ab")
      --port int           mysql port (default 3306)
  -t, --tablename string   table names,you yan use , join mult
  -g, --tags string        tags
  -u, --username string    mysql username (default "root")
```
