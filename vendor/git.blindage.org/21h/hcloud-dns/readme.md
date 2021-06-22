# Hetzner DNS golang library

I made this library to interact with Hetzner DNS API in most easy way. Hopefully in future it will be used for Hetzner external-dns provider. Check out [example](example) directory and [API_help.md](API_help.md).

Get your own token on Hetzner DNS and place it to `token` variable and run code

```go
token := "jcB2UywP9XtZGhvhSHpH5m"
zone := "vhSHpH5mjcB2UywP9XtZGh"

log.Println("Create new instance")
hdns := hclouddns.New(token)

log.Println("Get zone", zone)

allRecords, err := hdns.GetRecords(zone)
if err != nil {
	log.Fatalln(err)
}

log.Println(allRecords.Records)
log.Println(allRecords.Error)
```

At this moment library supports all API requests. If you found some bug mail me or register here and create issue.

---
Copyright by Vladimir Smagin (21h) 2020  
http://blindage.org email: 21h@blindage.org  
Project page: https://git.blindage.org/21h/hcloud-dns  