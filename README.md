# vslogger
Very Simple Logger for GoLang, add support of levels for the standard "log" library

## Goals
 * Being small
 * Being simple
 * Using default log library

##  Usage example

```
package main

import (
        logger "github.com/shengis/vslogger"
        "os"
)

func main() {
        log := logger.New(logger.LogWarning, os.Stderr, "Test: ", logger.Ldate | logger.Ltime)
        
        log.Debug("Foo")
        log.Warning("Bar")
        log.Info("Biz")
}
```
Output :
```
Test: 2018/06/06 15:14:19 [W] Bar
```
