package main

import (
    "godoc/app/logger"
    "log"
)

func main() {
    app := wireApp(logger.NewAppLogger())
    if err := app.Run(); err != nil {
        log.Fatalln(err)
    }
}
