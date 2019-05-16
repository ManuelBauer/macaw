# Macaw
Simple colorful logging utility for Go

## Installation
```
go get -u github.com/ManuelBauer/macaw
```

## Usage

### Basic
```
logger := macaw.CreateLogger("MyApp")
logger.Debug("This is a debug message")
```

### Sub logger
```
logger := macaw.CreateLogger("MyApp")
logger.Debug("This is a debug message")

subLogger := logger.CreateSubLogger("Frontend")
subLogger.Info("Message from the sub logger")
```
