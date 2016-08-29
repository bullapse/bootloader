# bootloader #

## Get and Import
- `go get github.com/bullapse/bootloader`
- `import github.com/bullapse/bootloader`

## Usage
- `Run` takes in a txt file path that you want to run as a bootloader for your app to the console
- `bootloader.Run("/path/to/{bootloader.txt}")`
- File in the same directory Example
    - `bootloader.Run("./{bootloader.txt})`
## Usage with log and lumberjack
- `RunWithLog` takes in a txt file path that you want to run as a bootloader and a logger

Example:
```
var logger *log.Logger

logger = log.New(&lumberjack.Logger{
   		Filename:   "/path/to/your/log/{logName}.log",
   		MaxSize:    500, // megabytes
   		MaxBackups: 3,
   		MaxAge:     28, //days
   	}, "", log.Ldate|log.Ltime|log.Lshortfile)
   
bootloader.RunWithLog("./bootloader.txt", logger)
```
## Links
- A good place to create ASCII Art can be foudn [here](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20)