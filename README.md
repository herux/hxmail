# hxmail
Email client library for Go

example, tested using smtp.office365.com :

```
email := hxmail.NewHxMail(host, 587, username, password)
email.To(to)
email.Subject("test hxmail")
email.SendMail()
```
