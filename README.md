## How to use

```zsh
$ cp config.sample.ini config.ini
```

- Rewite config.ini as necessary
- Open Discord

```zsh
$ go run main.go
```

## Directory

<pre>
.
├── README.md
├── apis
│   ├── discord
│   │   └── discord.go
│   ├── geocoding
│   │   └── geocoding.go
│   └── openweather
│       └── openweather.go
├── config
│   └── config.go
├── config.ini
├── config.sample.ini
├── go.mod
├── go.sum
├── main.go
└── server
    └── server.go

6 directories, 11 files
</pre>
