# Excalidraw Decrypt

## Purpose
This library allows to fetch Excalidraw digrams in Plain text, which can be used for storage and combined with other tools to render the diagrams.

## Background
The library relies on the parameters obtained on Excalidraw shareable links, which have the following format

https://excalidraw.com/#json=documentID,decryptionKey

## Usage

```
decrypter := excalidrawdecrypt.CreateShareableExcalidrawDecrypter()
decrypt, err := decrypter.Decrypt("pJK6JcJMr7LGOuy1NbCKP,YneEARvxllEU6vlDQfz81A")
```

## Decryption process
Using a public Excalidraw API to download the ciphered diagram:

https://json.excalidraw.com/api/v2/documentID

The downloaded diagram is decrypted using the decryptionKey from the shareable link.

The diagram in this shareable link paints the picture of the steps required to get the Excalidraw file in plaintext:
https://excalidraw.com/#json=pJK6JcJMr7LGOuy1NbCKP,YneEARvxllEU6vlDQfz81A

![Decryption Process](decryption_process.png "Decryption Process")
