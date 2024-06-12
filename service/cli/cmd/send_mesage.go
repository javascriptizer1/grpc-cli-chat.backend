package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/javascriptizer1/grpc-cli-chat.backend/service/cli/internal/app"
	"github.com/spf13/cobra"
)

func newSendMessageCommand(ctx context.Context, sp *app.ServiceProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-message",
		Short: "Send a message to a chat",
		Run: func(cmd *cobra.Command, args []string) {
			chatID, _ := cmd.Flags().GetString("chat_id")
			text, _ := cmd.Flags().GetString("text")

			err := sp.ChatClient(ctx).SendMessage(context.Background(), text, chatID)

			if err != nil {
				log.Fatalf("Could not send message: %v", err)
			}

			fmt.Println("Message sent successfully")
		},
	}

	cmd.Flags().String("chat_id", "", "ID of the chat")
	cmd.Flags().String("text", "", "Text of the message")

	return cmd
}
