# ad-exchange
 
## Installation
Download dependencies:
```console
go mod download
```

Install fresh:
```console
go install github.com/zzwx/fresh@latest
```

Create .env file:
```dotenv
JWT_SIGNING_KEY=YOUR_KEY
```

Run app using fresh:
```console
fresh
```

... or with `go run` command:
```console
go run .
```

## TODOs
- Validate forms (validator doesn't work for some reason)
- Profile view and settings
- Fix dropdown on PC
