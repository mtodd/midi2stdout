package main

import (
  "fmt"
  "log"

  "github.com/rakyll/portmidi"
)

type midimsg struct {
  a, b, c float64
}

func main() {
  portmidi.Initialize()
  defer portmidi.Terminate()

  deviceId := portmidi.GetDefaultInputDeviceId()

  in, err := portmidi.NewInputStream(deviceId, 1024)
  if err != nil {
    log.Fatal(err)
  }

  // or alternatively listen events
  events := in.Listen()

  for event := range events {
    fmt.Printf("%x\t%d\t%d\n", event.Status, event.Data1, event.Data2)
  }
}
