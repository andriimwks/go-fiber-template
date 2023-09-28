# go-fiber-template

## Installation
Download dependencies:
```console
go mod download
```

Create .env file:
```dotenv
JWT_SIGNING_KEY=YOUR_KEY
```

Run app using fresh:
```console
go install github.com/zzwx/fresh@latest
fresh
```

... or with `go run` command:
```console
go run ./cmd
```

## TODOs
- Documentation
- Validate forms
- Profile view and settings
