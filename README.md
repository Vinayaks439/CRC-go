# CRC-go [![Go](https://github.com/Vinayaks439/CRC-go/actions/workflows/go.yml/badge.svg)](https://github.com/Vinayaks439/CRC-go/actions/workflows/go.yml)
CRC encode and decode practice in golang

# Usage

This is a CLI tool for illustrating CRC encode and decode in data link layer

## Steps

1. Clone the Repo using `git clone https://github.com/Vinayaks439/CRC-go.git`
2. build the binary using `go build -o crc`
3. To Encode the data use `./crc encode --data 1001 --key 1011`
4. To Decode the data use `./crc decode --frame 1001110 --key 1011`

# Example 1

```bash
$ ./crc encode --data 1001 --key 1011
#The Encoded DataFrame after applying CRC is 1001110
$ ./crc decode --frame 1001110 --key 1011
#The Checksum after decoding the frame with the key is 000
#The Frame is Correct

$ ./crc encode -d 101101 -k 111011
#The Encoded DataFrame after applying CRC is 10110101001
$ ./crc decode -f 10110101001 -k 111011
#The Checksum after decoding the frame with the key is 00000
#The Frame is Correct
```

# Example 2
Let us now take a 10 bit data and a key P = 110101, and encode the data and decode the data

data = 1011011011
key = 110101
```bash
$ ./crc encode -d 1011011011 -k 110101
#The Encoded DataFrame after applying CRC is 101101101101101

#This Encoded DataFrame is now 15 bit long and we can now transmit this data
#Now during the transmission, the data is corrupted and we need to check if the data is correct or not

#corrupted_data = 101001001101001

#Now the receiver will decode the data using the key and check if the data is correct or not

$ ./crc decode -f 101001001101001 -k 110101
#The Checksum after decoding the frame with the key is 11100
#The Frame is Incorrect since the output is not 00000

#So the data is corrupted and we need to retransmit the data or discard the data
```