# Psutils Client

## Build

### Linux
```bash
    GOOS=linux go build -o ./build/monitron-client -ldflags="-X main.AppVersion={version}" ./cmd/main.go
```

### Windows
```bash
    GOOS=windows go build -o ./build/monitron-client.exe -ldflags="-X main.AppVersion={version}" ./cmd/main.go
```

## Dependencies
- [Fiber](https://gofiber.io/)
- [NutsDB](https://github.com/nutsdb/nutsdb)
- [GoCron](https://github.com/go-co-op/gocron)
- [GoCache](https://github.com/eko/gocache)
- [GoPsUtils](https://github.com/shirou/gopsutil)

## Next Feature
- Secure connection using JWT and Client Authentication
- Default authentication
- Scheduled stats and save into NutsDB
- Web interface
- Caching JWT and client authentication