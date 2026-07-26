package user

import (
	"reflect"
	"testing"

	"github.com/opd-ai/toxcore"
)

func TestCreateNewProfile(t *testing.T) {
	tests := []struct {
		name    string
		want    *toxcore.Tox
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNewProfile()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNewProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadOrCreateProfile(t *testing.T) {
	tests := []struct {
		name    string
		want    *toxcore.Tox
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadOrCreateProfile()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadOrCreateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadOrCreateProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadExistingProfile(t *testing.T) {
	type args struct {
		savedata []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *toxcore.Tox
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadExistingProfile(tt.args.savedata)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadExistingProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadExistingProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveProfile(t *testing.T) {
	type args struct {
		tox *toxcore.Tox
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveProfile(tt.args.tox); (err != nil) != tt.wantErr {
				t.Errorf("SaveProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
