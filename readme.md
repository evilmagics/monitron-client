# Psutils Client

## Features
- Supported secure HTTPS protocols
- Supported websocket protocols usage stats
- Implement in-memory cache more efficiently for performance
- Lightweight and small resources but better performance

## Build

### Linux
```bash
    GOOS=linux go build -o ./build/monitron-client -ldflags="-s -X main.AppVersion={version}" ./cmd/main.go
```

### Windows
```bash
    GOOS=windows go build -o ./build/monitron-client.exe -ldflags="-s -X main.AppVersion={version}" ./cmd/main.go
```

## Dependencies
- [Fiber](https://gofiber.io/)
- [NutsDB](https://github.com/nutsdb/nutsdb)
- [GoCache](github.com/patrickmn/go-cache)
- [GoPsUtils](https://github.com/shirou/gopsutil)

## Next Feature
- Split disk partition usage into separate cache and api
- Split cpu usage and information into separate api and remove the information from cache
- Secure connection using JWT and Client Authentication
- Default authentication
- Scheduled stats and save into NutsDB
- Web interface
- Caching JWT and client authentication