package service

import (
	"Go-Dispatch-Bootcamp/types"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *demoService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_readCsvFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		ts      *demoService
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.readCsvFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.readCsvFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoService.readCsvFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_FetchCsvFromRemote(t *testing.T) {
	tests := []struct {
		name    string
		ts      *demoService
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.FetchCsvFromRemote()
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.FetchCsvFromRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoService.FetchCsvFromRemote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_UpdateUsers(t *testing.T) {
	type args struct {
		users *[]types.User
	}
	tests := []struct {
		name    string
		ts      *demoService
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.UpdateUsers(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.UpdateUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("demoService.UpdateUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_GetUsers(t *testing.T) {
	tests := []struct {
		name    string
		ts      *demoService
		want    *[]types.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoService.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_GetUsersMap(t *testing.T) {
	tests := []struct {
		name    string
		ts      *demoService
		want    map[int]types.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.GetUsersMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.GetUsersMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoService.GetUsersMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoService_GetFeedUsers(t *testing.T) {
	tests := []struct {
		name    string
		ts      *demoService
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ts.GetFeedUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("demoService.GetFeedUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("demoService.GetFeedUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
