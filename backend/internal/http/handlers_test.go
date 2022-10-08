package http

import (
	"lillybox-backend/internal/database"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/livepeer/go-api-client"
)

func TestHandlers_GetOnDemands(t *testing.T) {
	type fields struct {
		LivepeerClient *api.Client
		Database       database.Database
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handlers{
				LivepeerClient: tt.fields.LivepeerClient,
				Database:       &tt.fields.Database,
			}
			if err := h.GetOnDemands(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Handlers.GetOnDemands() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
