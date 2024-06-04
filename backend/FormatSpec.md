# Serial Data Format Specification

## Serial Port Configuration

- BaudRate: 9600
- DataBits: 8
- Parity:   bug_serial.NoParity
- StopBits: bug_serial.OneStopBit

## Packet Data by Bytes

Data is sent in packets of 45 bytes.

The following table describes what data is sent by the ISC9000 device. 

| Encoded Data          | Bytes     | Type    | Notes 
| --------------------- | --------- | ------- | --------------------- 
| Prefix (0274)         | (0, 2]    | ASCII   | Prefixes Every Packet 
| Game Time             | (2, 6]    | uint8   | [Parsing Game Time](#parsing-game-time) 
| Preiod                | (6, 7]    | ASCII   | 
| Home Score            | (7, 9]    | ASCII   | 
| Away Score            | (9, 11]   | ASCII   | 
| Home Timeout          | (11, 12]  | Unknown | Unsure, not enough data
| Away Timeout          | (12, 13]  | Unknown | Unsure, not enough data
| Home Fouls            | (13, 14]  | Unknown | Unsure, not enough data
| Away Fouls            | (14, 15]  | Unknown | Unsure, not enough data
| "S" of Unkown Purpose | (15, 16]  | ASCII   | 
| Home Possession/Bonus | (16, 17]  | ASCII   | 
| Away Possession/Bonus | (17, 18]  | ASCII   | 
| Home Shots on Goal    | (18, 20]  | ASCII   | Unsure, not enough data
| Away Shots on Goal    | (20, 22]  | ASCII   | Unsure, not enough data
| Home Penalty Player 1 | (22, 24]  | ASCII   |  
| Home Penalty Time 1   | (24, 27]  | uint8   | [Parsing Game Time](#parsing-game-time) 
| Home Penalty Player 2 | (27, 29]  | ASCII   |  
| Home Penalty Time 2   | (29, 32]  | uint8   | [Parsing Game Time](#parsing-game-time) 
| Away Penalty Player 1 | (32, 34]  | ASCII   |  
| Away Penalty Time 1   | (34, 37]  | uint8   | [Parsing Game Time](#parsing-game-time) 
| Away Penalty Player 2 | (37, 39]  | ASCII   | 
| Away Penalty Time 2   | (39, 42]  | uint8   | [Parsing Game Time](#parsing-game-time) 
| Suffix (0d)           | (44, 45]  | ASCII   | Suffixes Every Packet 

The above table was referenced to create the code at backend/internal/serial/serail.go:parsePacket() and the functions it calls. If the information in the table is incorrect, please update the table and related code.

## Special Characters

- ASCII `:` (`3a` in hex) is used as a placeholder for data with no value.

## Parsing Game Time

While most fields are ASCII characters, Game Time and Penalty Time are the only fields that make use of specific bits of data.

### Game Time

Game Time is four bytes long. Call this information `data` with each byte accessed by `data[0]`, `data[1]`, or `data[2]`.

1. The first bit of the first byte of Game Time indicates whether the game is paused. If paused, the first bit will be enabled.

    > `data[0] & 0x80` is used to test if the game is paused. For example, if `data[0]` is `b1` (10110001), then the game is paused.

1. The first bit of the second byte of Game Time indicates whether the ":" delimiter is used or the "." delimiter is used. That is, it indicates if time is being measured in `minutes:seconds` or `seconds.deciseconds`.

    > `data[1] & 0x80` is used to test if the ":" delimiter is used.

1. It is unclear what the first bit of the second byte of Game Time indicates, but it is likely similar to the first bit of the second byte (":" delimiter)

1. The fourth byte does not encode data like the others; it is just an ASCII character value.

After reading the information contained in those bits, they can be ignored and treated as zeros. Then, ASCII game time information can be parsed.

> For example, `b1` (10110001) is read as the ASCII value "1" (`31` or 00110001). 
> This is done with the operation `data[0] = data[0] & 0x7F` for each byte of the first three bytes.

Now the game time can be read as ASCII. The first two bytes indicate the coarse time and the final two indicate the fine time.

> "Coarse Time" can be either minutes or seconds, depending on what level of precision and what delimiter is used.
> "Fine Time" can be either seconds or deciseconds, depending on what level of precision and what delimiter is used.
> For example, 05:30 means Five minutes and Thirty Seconds, while 05.30 means 5.30 seconds.

### Penalty Time

Penalty time is parsed very similarly to Game Time. Here are the differences:

- Penalty Time is three bytes long.
- There is only one byte to indicate coarse time. The final two bytes indicate the fine time.
- There is no "paused" bit.
- The first bit of the first byte indicates the delimiter, ":" if enabled and "." if disabled, much like the first bit of the second byte of Game Time.
- The entire field may be made of ":" characters, meaning there is no penalty.