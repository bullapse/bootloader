# bootloader #
- Golang package that runs a bootloader from a text file using a log or output to the console

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

## LICENSE
MIT License

Copyright (c) 2016 Spencer Bull

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
