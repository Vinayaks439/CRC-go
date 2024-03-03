# CRC-go [![Go](https://github.com/Vinayaks439/CRC-go/actions/workflows/go.yml/badge.svg)](https://github.com/Vinayaks439/CRC-go/actions/workflows/go.yml)
CRC encode and decode practice in golang

# Usage

This is a CLI tool for illustrating CRC encode and decode in data link layer

## Steps

1. Clone the Repo using `git clone https://github.com/Vinayaks439/CRC-go.git`
2. build the binary using `go build -o crc`
3. To Encode the data use `./crc encode --data 1001 --key 1011`
4. To Decode the data use `./crc decode --frame 1001110 --key 1011`
