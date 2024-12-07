package services

import (
	"fmt"
	"multi-user-approval/models"
)

// NotificationService will simulate sending email notifications
type NotificationService struct{}

// NewNotificationService returns a new instance of NotificationService
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// SendEmail simulates sending an email notification
func (ns *NotificationService) SendEmail(to string, subject string, message string) error {
	// In a real application, you would integrate with an email service here (e.g., SendGrid, SMTP, etc.)
	fmt.Printf("Sending email to: %s\n", to)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Message: %s\n", message)
	return nil
}

// NotifyApprovers sends an email notification to the task approvers
func (ns *NotificationService) NotifyApprovers(taskID string, approvers []models.User) {
	for _, approver := range approvers {
		subject := fmt.Sprintf("Task %s requires your approval", taskID)
		message := fmt.Sprintf("Hello %s, you have a task that requires your approval. Please review and approve it.", approver.Name)
		err := ns.SendEmail(approver.Email, subject, message)
		if err != nil {
			fmt.Printf("Failed to send email to %s: %s\n", approver.Email, err)
		}
	}
}

// NotifyCreator sends an email notification to the task creator
func (ns *NotificationService) NotifyCreator(taskID string, creator models.User) {
	subject := fmt.Sprintf("Task %s has been approved by all approvers", taskID)
	message := fmt.Sprintf("Hello %s, your task has been approved by all approvers.", creator.Name)
	err := ns.SendEmail(creator.Email, subject, message)
	if err != nil {
		fmt.Printf("Failed to send email to %s: %s\n", creator.Email, err)
	}
}
