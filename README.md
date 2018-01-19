### Challenge Https in Go

launch using `go run https.go`

Or to deploy to a server:
```
GOOS=linux go build https.go
#Then on the server
sudo ./https
```

### What's the difference between 'A' and 'CNAME' DNS records ?

Like explained [here](https://support.dnsimple.com/articles/differences-between-a-cname-alias-url/)
- The A record maps a name to one or more IP addresses, when the IP are known and stable.
- The CNAME record maps a name to another name. It should only be used when there are no other records on that name.

### Links

https://crypto.stackexchange.com/questions/43697/what-is-the-difference-between-pem-csr-key-and-crt
https://quay.io/repository/letsencrypt/letsencrypt?tab=info
