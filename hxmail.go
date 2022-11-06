package hxmail

import (
	"bytes"
	"fmt"
	"io"
)

type HxAttachment struct {
}

type HxHeader struct {
}

type HxMail struct {
	to         string
	cc         string
	bcc        string
	subject    string
	from       string
	fromName   string
	replyTo    string
	headers    map[string]HxHeader
	attachment []HxAttachment
	sender     *HxMailSender

	buffer bytes.Buffer
}

func NewHxMail(server string, port int, username, password string) *HxMail {
	sender := newSender(server, port, username, password)
	return &HxMail{
		headers: make(map[string]HxHeader),
		sender:  sender,
	}
}

func (m *HxMail) To(to string) {
	m.to = to
}

func (m *HxMail) From(from, name string) {
	m.from = from
	m.fromName = name
}

func (m *HxMail) Cc(cc string) {
	m.cc = cc
}

func (m *HxMail) Bcc(bcc string) {
	m.bcc = bcc
}

func (m *HxMail) Subject(subject string) {
	m.subject = subject
}

func (m *HxMail) ReplyTo(replyto string) {
	m.replyTo = replyto
}

func (m *HxMail) toBytes() (msg []byte, err error) {
	return nil, nil
}

func (m *HxMail) AddAttachment(file HxAttachment) {
	m.attachment = append(m.attachment, file)
}

func (m *HxMail) SendMail() error {
	auth, err := m.sender.tryConnect()
	if err != nil {
		return err
	}

	err = m.sender.sendMail(auth, m)
	if err != nil {
		return err
	}

	return nil
}

func (m *HxMail) createHeaders(buf io.Writer) error {
	fromHeader := fmt.Sprintf("From: %s <%s>\r\n", m.fromName, m.from)
	if m.fromName == "" {
		fromHeader = fmt.Sprintf("From: %s\r\n", m.from)
	}
	_, err := buf.Write([]byte(fromHeader))
	if err != nil {
		return err
	}

	if _, err := buf.Write([]byte("Mime-Version: 1.0\r\n")); err != nil {
		return err
	}

	if m.replyTo != "" {
		fmt.Fprintf(buf, "Reply-To: %s\r\n", m.replyTo)
	}

	fmt.Fprintf(buf, "Subject: %s\r\n", m.Subject)
	for _, to := range m.to {
		fmt.Fprintf(buf, "To: %s\r\n", to)
	}

	for _, cc := range m.cc {
		fmt.Fprintf(buf, "CC: %s\r\n", cc)
	}

	for _, bcc := range m.bcc {
		fmt.Fprintf(buf, "BCC: %s\r\n", bcc)
	}

	return nil
}
