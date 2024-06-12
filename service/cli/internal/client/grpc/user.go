package client

import (
	"context"

	userv1 "github.com/javascriptizer1/grpc-cli-chat.backend/pkg/grpc/user_v1"
	"github.com/javascriptizer1/grpc-cli-chat.backend/service/cli/internal/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserClient struct {
	client userv1.UserServiceClient
}

func NewUserClient(client userv1.UserServiceClient) *UserClient {
	return &UserClient{client: client}
}

func (c *UserClient) GetUserInfo(ctx context.Context) (*domain.UserInfo, error) {
	res, err := c.client.GetUserInfo(ctx, &emptypb.Empty{})

	if err != nil {
		return nil, err
	}

	ui := domain.NewUserInfo(
		res.GetId(),
		res.GetName(),
		res.GetEmail(),
		uint16(res.GetRole()),
	)

	return ui, nil

}
