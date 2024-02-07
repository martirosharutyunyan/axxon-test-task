make tidy for uploading all packages

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/pressly/goose/v3/cmd/goose@latest

`
export GOPRIVATE=github.com/go-boilerplate
git config \
--global \
url."https://${user}:${personal_access_token}@github.com".insteadOf \
"https://github.com"
`
for downloading private packages

pre commit:
    pip install pre-commit
    pre-commit autoupdate
    pre-commit install
