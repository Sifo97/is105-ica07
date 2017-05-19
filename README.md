# ICA07

## UDP 
Start `server.go` med kommadoen `go run server.go` i mappa UDP/Server/ 

Start `Client.go` med kommadoen `go run Client.go` i mappa UDP/Client

Dette sender en tekst-string til serveren ukryptert

## TCP
Samme som ovenfor, bare i mappa TCP.

Klienten kan sende egen definerte tekst-stringer til serveren, serveren responderer med teksten i UPPERCASE.

## UDP Encrypted
Gjør det samme som det står i `UDP` seksjonen.
Teksten blir encodet med Base64, kryptert mes AES, så sent over til serveren.

Serveren tar i mot informasjonen, dekrypterer det, å viser meldingen
