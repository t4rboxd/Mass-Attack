# Flood Attack Script

This script is designed to use a DDoS attack script on the VPS/VDS to attack a list of hosts in a text file

## Prerequisites

To use this script, you will need the following:

- A VPS/Dedicated server running a Linux operating system.
- A script to attack with.

## Usage

1. Clone this repository to your server using the command `git clone https://github.com/t4rboxd/Mass-Attack`
2. Navigate to the `Mas-attack` directory using the command `cd Mass-Attack`
3. Create a file named `ips.txt` in the same directory as the script. This file should contain a list of IP addresses:port (one per line) that you want to attack.
4. Start flooding the hosts, using `python3 main.py`, or `go run main.go`

## Notes

- This script is for educational purposes only and should not be used to harm any systems.
- Using this script for malicious purposes may result in legal consequences.
