package main

import (
    "log"
)

func main() {
    app := wireApp()
    if err := app.Run(); err != nil {
        log.Fatalln(err)
    }
}
