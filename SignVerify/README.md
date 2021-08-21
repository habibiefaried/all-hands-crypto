# Description

How sign & verify works

# To do

Generate key pair

```
openssl genrsa -aes256 -out private.key 4096
openssl rsa -in private.key -pubout -out public.key
```

`private.key` is Private Key and `public.key` is Public Key.
For only demo purpose, I'm using passphrase `1234`. If you want to change, you need to change on source as well

`msg.txt` is the message output from `make sign` and `sig.txt` is the signature itself. Use `make verify` to verify those message.
