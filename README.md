Go + React Boilerplate, using [`fiber`](https://github.com/gofiber/fiber)

---

# Table of Contents

 - [Commands](#project-commands)
 - [Structure](#project-structure)

---

# Project Commands

* **NOTE:** all commands should be invoked from the root of the project!

| Command | Result | Notes |
| --- | --- | --- |
| `make start` | Builds and runs project | Same as running `make build` and then `make run-build` |
| `make build` | Builds project using `go build` | |
| `make run-build` | Runs compiled binary | Run `make build` prior to running this command |
| `make run` | Uses the `go run` command to build (in memory) and run project | |

---

# Project Structure

| Path | Purpose |
| --- | --- |
| `/cmd` | Entry point, where `main.go` lives |
| `/public` | Static file path for web assets |
| `/bin` | Output location of compiled files |