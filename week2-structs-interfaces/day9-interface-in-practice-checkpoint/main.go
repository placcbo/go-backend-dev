package main

import (
	"fmt"
	"time"
)

// ── INTERFACE ────────────────────────────────────────────────────
type Notifier interface {
	Send(to, subject, message string) error
}

// ── COMPILE-TIME CHECKS ───────────────────────────────────────────
var _ Notifier = EmailNotifier{}
var _ Notifier = SMSNotifier{}
var _ Notifier = LogNotifier{}

// ── EMAIL NOTIFIER ───────────────────────────────────────────────
type EmailNotifier struct {
	FromAddress string
}

func (e EmailNotifier) Send(to, subject, message string) error {
	fmt.Printf("📧 EMAIL from %s to %s\nSubject: %s\nBody: %s\n\n",
		e.FromAddress, to, subject, message)
	return nil
}

// ── SMS NOTIFIER ─────────────────────────────────────────────────
type SMSNotifier struct {
	Provider string
}

func (s SMSNotifier) Send(to, subject, message string) error {
	fmt.Printf("📱 SMS via %s to %s\nMessage: %s\n\n",
		s.Provider, to, message)
	return nil
}

// ── LOG NOTIFIER ─────────────────────────────────────────────────
type LogNotifier struct{}

func (l LogNotifier) Send(to, subject, message string) error {
	fmt.Printf("[%s] LOG NOTIFICATION → %s: %s\n\n",
		time.Now().Format("15:04:05"), to, message)
	return nil
}

// ── NOTIFICATION SERVICE ─────────────────────────────────────────
type NotificationService struct {
	notifiers []Notifier
}

func (ns *NotificationService) Register(n Notifier) {
	ns.notifiers = append(ns.notifiers, n)
}

func (ns *NotificationService) Broadcast(to, subject, message string) {
	fmt.Printf("── Broadcasting to %d notifiers ──\n\n", len(ns.notifiers))

	for _, n := range ns.notifiers {
		if err := n.Send(to, subject, message); err != nil {
			fmt.Println("Error sending:", err)
		}
	}
}

// ── TYPE SWITCH ───────────────────────────────────────────────────
func Describe(n Notifier) {
	switch v := n.(type) {

	case EmailNotifier:
		fmt.Println("Email notifier (from:", v.FromAddress, ")")

	case SMSNotifier:
		fmt.Println("SMS notifier (provider:", v.Provider, ")")

	case LogNotifier:
		fmt.Println("Log notifier (writes to stdout)")

	default:
		fmt.Printf("Unknown notifier type: %T\n", v)
	}
}

// ── MAIN ─────────────────────────────────────────────────────────
func main() {

	svc := &NotificationService{}

	// register different notifier types
	svc.Register(EmailNotifier{FromAddress: "alerts@system.com"})
	svc.Register(SMSNotifier{Provider: "Twilio"})
	svc.Register(LogNotifier{})

	// broadcast message
	svc.Broadcast(
		"admin@company.com",
		"Job Failed",
		"Job #1042 failed after 3 retries",
	)

	fmt.Println("── Describing registered notifiers ──")

	for _, n := range svc.notifiers {
		Describe(n)
	}
}
