# go-fiber-template

This "template" was originally made during development of my and my friend's commercial website, but I've decided to put it on public repo in its current state.

There's a lot of work to do to make it easy for a random people to understand, such as documentation etc.

In my opinion, this project is a bit overkill for a template, but having somehow working auth right out of the box might be useful.

Enjoy :)

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
