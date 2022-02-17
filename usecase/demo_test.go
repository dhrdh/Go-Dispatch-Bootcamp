package usecase

import (
	"Go-Dispatch-Bootcamp/mocks"
	"Go-Dispatch-Bootcamp/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	user = types.User{
		Id:         1,
		Username:   "test",
		Identifier: "test",
		FirstName:  "John",
		LastName:   "Doe",
	}
	feedUsers = [][]string{
		{"col1", "col2", "col3", "col4"},
		{"test", "test", "test", "test"},
	}
	feedUsersMapped = []types.User{
		{1, "test", "test", "test", "test"},
	}
)

func Test_demoUsecase_Fetch(t *testing.T) {
	tests := []struct {
		name    string
		dsData  []types.User
		dsError error
		want    *[]types.User
	}{
		{
			name:    "Fetch users. Success story.",
			dsData:  []types.User{user},
			dsError: nil,
			want:    &[]types.User{user},
		},
	}

	ds := &mocks.DemoService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.On("GetUsers", dataFileName).Return(&tt.dsData, tt.dsError)

			uc := demoUsecase{
				service: ds,
			}

			result, err := uc.Fetch()

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoUsecase_FetchById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		dsData  map[int]types.User
		dsError error
		want    *types.User
	}{
		{
			name:    "Fetch user by id. Success story.",
			id:      1,
			dsData:  map[int]types.User{1: user},
			dsError: nil,
			want:    &user,
		},
	}

	ds := &mocks.DemoService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.On("GetUsersMap", dataFileName).Return(tt.dsData, tt.dsError)

			uc := demoUsecase{
				service: ds,
			}

			result, err := uc.FetchById(tt.id)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoUsecase_Feed(t *testing.T) {
	tests := []struct {
		name    string
		dsData  [][]string
		dsError error
		want    [][]string
	}{
		{
			name:    "Feed. Success story.",
			dsData:  feedUsers,
			dsError: nil,
			want:    feedUsers,
		},
	}

	ds := &mocks.DemoService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.On("GetFeedUsers", feedFileName).Return(tt.dsData, tt.dsError)

			uc := demoUsecase{
				service: ds,
			}

			result, err := uc.Feed()

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoUsecase_UpdateUsersFromFeed(t *testing.T) {
	tests := []struct {
		name                      string
		dsFetchCsvFromRemoteData  [][]string
		dsFetchCsvFromRemoteError error
		dsUpdateUsersData         bool
		dsUpdateUsersError        error
		dsUpdateUsersParameter    *[]types.User
		want                      bool
	}{
		{
			name:                      "Update users from feed. Success story.",
			dsFetchCsvFromRemoteData:  feedUsers,
			dsFetchCsvFromRemoteError: nil,
			dsUpdateUsersData:         true,
			dsUpdateUsersError:        nil,
			dsUpdateUsersParameter:    &feedUsersMapped,
			want:                      true,
		},
	}

	ds := &mocks.DemoService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds.On("FetchCsvFromRemote", feedUrl).Return(tt.dsFetchCsvFromRemoteData, tt.dsFetchCsvFromRemoteError)
			ds.On("UpdateUsers", tt.dsUpdateUsersParameter, dataFileName).Return(tt.dsUpdateUsersData, tt.dsUpdateUsersError)

			uc := demoUsecase{
				service: ds,
			}

			result, err := uc.UpdateUsersFromFeed()

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ds demoService
	}
	tests := []struct {
		name string
		args args
		want *demoUsecase
	}{
		{
			name: "New controller test",
			args: args{
				ds: &mocks.DemoService{},
			},
			want: &demoUsecase{
				service: &mocks.DemoService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
