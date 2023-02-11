package client

import (
	"testing"
	"time"
)

func TestClient_BeforeCreate(t *testing.T) {
	type fields struct {
		ID            uint
		CreatedAt     time.Time
		UpdatedAt     time.Time
		Name          string
		GlobalService bool
		Homepage      string
		CallbackUri   string
		SecretKey     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Successful creation of secret key",
			fields: fields{
				ID:            1,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				Name:          "test client",
				GlobalService: false,
				Homepage:      "https://www.example.com",
				CallbackUri:   "https://www.example.com/callback",
				SecretKey:     "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ID:            tt.fields.ID,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
				Name:          tt.fields.Name,
				GlobalService: tt.fields.GlobalService,
				Homepage:      tt.fields.Homepage,
				CallbackUri:   tt.fields.CallbackUri,
				SecretKey:     tt.fields.SecretKey,
			}
			if err := c.BeforeCreate(nil); (err != nil) != tt.wantErr {
				t.Errorf("Client.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.SecretKey == "" {
				t.Errorf("Secret key was not created")
			}
		})
	}
}
