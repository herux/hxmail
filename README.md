# hxmail
Email client library for Go

The ```hxmail``` package currently supports the following:
*  From, To, Bcc, Cc and ReplyTo fields
*  Email addresses supported "test1@example.com,test2@example.com,.."
*  Text and HTML Message Body
*  Attachments
*  More to come!

### installation 

```go
go get github.com/herux/hxmail
```

### Examples, 
tested using smtp.office365.com :

```go
email := hxmail.NewHxMail(host, 587, username, password)
email.To(to)
email.Subject("test hxmail")
email.SendMail()
```
