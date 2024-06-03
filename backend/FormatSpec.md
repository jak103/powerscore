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
| Home Penalty Player   | (22, 24]  | ASCII   |  
| Home Penalty Time     | (24, 28]  | uint8   | 
| Away Penalty Player   | (32, 34]  | ASCII   | 
| Away Penalty Time     | (34, 38]  | uint8   | 
| Suffix (0d)           | (44, 45]  | ASCII   | Suffixes Every Packet 

The above table was referenced to create the code at backend/internal/serial/serail.go:parsePacket() and the functions it calls. If the information in the table is incorrect, please update the table and related code.

## Special Characters

- ASCII `:` (`3a` in hex) is used as a placeholder for data with no value.

## Parsing Game Time