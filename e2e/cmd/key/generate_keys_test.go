package key

import (
	"crypto/rsa"
	"reflect"
	"testing"

	uuid "github.com/samborkent/uuidv8"
)

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
		want uuid.UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUUID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name  string
		want  *rsa.PrivateKey
		want1 *rsa.PublicKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenerateKey()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GenerateKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
