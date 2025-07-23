package main

import (
	"fmt"
)

type NotifyType int

const (
	EMAILTYPE NotifyType = iota
	SMSTYPE
)

var notifyTypeMap = map[NotifyType]string{
	EMAILTYPE: "EMAIL",
	SMSTYPE:   "SMS",
}

func (n *NotifyType) String(nType NotifyType) string {
	return notifyTypeMap[nType]
}

type Customer struct {
	name            string
	parcelStatus    string
	isStatusChanged bool
	notifyType      NotifyType
}

type ParcelStatusNotifier interface {
	Notify(Customer)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (eNotifier EmailNotifier) Notify(customer Customer) {
	fmt.Printf("Email Notification is sent to Customer %s for %s status", customer.name, customer.parcelStatus)
	fmt.Println()
}

func (smsNotifier SMSNotifier) Notify(customer Customer) {
	fmt.Printf("SMS Notification is sent to Customer %s for %s status", customer.name, customer.parcelStatus)
	fmt.Println()
}

type NotificationService struct {
	notifier ParcelStatusNotifier
}

func NewNotificationService(notifier ParcelStatusNotifier) *NotificationService {
	return &NotificationService{notifier: notifier}
}

func (notifyService NotificationService) sendNotification(customer Customer) {
	// fmt.Printf("here")
	notifyService.notifier.Notify(customer)
}

func main() {
	customers := []Customer{
		Customer{
			name:            "Purnay Barge",
			parcelStatus:    "SHIPPED",
			isStatusChanged: true,
			notifyType:      0,
		},
		Customer{
			name:            "Vinay Yadav",
			parcelStatus:    "DELIVERED",
			isStatusChanged: false,
			notifyType:      0,
		},
		Customer{
			name:            "Amey Mule",
			parcelStatus:    "DELAYED",
			isStatusChanged: true,
			notifyType:      1,
		},
	}
	var notificationService *NotificationService
	for _, cust := range customers {
		if cust.isStatusChanged && cust.notifyType == EMAILTYPE {
			notificationService = NewNotificationService(EmailNotifier{})
			notificationService.sendNotification(cust)
		} else if cust.isStatusChanged && cust.notifyType == SMSTYPE {
			notificationService = NewNotificationService(SMSNotifier{})
			notificationService.sendNotification(cust)
		}
	}
}
