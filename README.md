# scoremaster

### Logging

The project uses a custom logging system defined in `logs.go`. It provides three log levels:
- `Info`: For informational messages.
- `Warning`: For warnings.
- `Error`: For errors.

#### FatalIf Method
The `Error` logger includes a `FatalIf` method for handling fatal errors. This method logs the error message and terminates the program if the error is not `nil`.

**Usage Example:**
```go
logs.Error.FatalIf(err, "Failed to start server")


### Todo
- Crear Makefile que borre fichero logs + sqlite.db