package main
import (
    "fmt"
    "github.com/vee2xx/camtron"
)
func main() {
    fmt.Println("Starting...")
    camtron.StartStreamToFileConsumer()
    fmt.Println("Working...")
    camtron.StartCam()
    fmt.Println("Ending...")
}