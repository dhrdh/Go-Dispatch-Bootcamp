package service

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"Go-Dispatch-Bootcamp/types"

	"github.com/stretchr/testify/assert"
)

var (
	fileData = [][]string{
		{"1", "booker12", "9012", "Rachel", "Booker"},
		{"2", "grey07", "2070", "Laura", "Grey"},
		{"3", "johnson81", "4081", "Craig", "Johnson"},
		{"4", "jenkins46", "9346", "Mary", "Jenkins"},
		{"5", "smith79", "5079", "Jamie", "Smith"},
	}

	feedFileData = [][]string{
		{"username", "id", "first_name", "last_name"},
		{"booker12", "9012", "Rachel", "Booker"},
		{"grey07", "2070", "Laura", "Grey"},
		{"johnson81", "4081", "Craig", "Johnson"},
		{"jenkins46", "9346", "Mary", "Jenkins"},
		{"smith79", "5079", "Jamie", "Smith"},
	}

	users = []types.User{
		{1, "booker12", "9012", "Rachel", "Booker"},
		{2, "grey07", "2070", "Laura", "Grey"},
		{3, "johnson81", "4081", "Craig", "Johnson"},
		{4, "jenkins46", "9346", "Mary", "Jenkins"},
		{5, "smith79", "5079", "Jamie", "Smith"},
	}

	usersMap = map[int]types.User{
		1: {1, "booker12", "9012", "Rachel", "Booker"},
		2: {2, "grey07", "2070", "Laura", "Grey"},
		3: {3, "johnson81", "4081", "Craig", "Johnson"},
		4: {4, "jenkins46", "9346", "Mary", "Jenkins"},
		5: {5, "smith79", "5079", "Jamie", "Smith"},
	}

	feedUrl      = "http://localhost:8080/api/v1/feed"
	dataFileName = "../data/data.csv"
	feedFileName = "../data/feed.csv"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *demoService
	}{
		{
			name: "New controller test",
			want: &demoService{},
		},
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
		name string
		args args
		ts   *demoService
		want [][]string
	}{
		{
			name: "Read csv from file. Success story.",
			args: args{path: dataFileName},
			ts:   &demoService{},
			want: fileData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.readCsvFromFile(tt.args.path)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoService_UpdateUsers(t *testing.T) {
	type args struct {
		users *[]types.User
	}
	tests := []struct {
		name string
		ts   *demoService
		args args
		want bool
	}{
		{
			name: "Update users. Success story.",
			args: args{users: &users},
			ts:   &demoService{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.UpdateUsers(tt.args.users, dataFileName)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoService_GetUsers(t *testing.T) {
	tests := []struct {
		name string
		ts   *demoService
		want *[]types.User
	}{
		{
			name: "Get users. Success story.",
			ts:   &demoService{},
			want: &users,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.GetUsers(dataFileName)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoService_GetUsersMap(t *testing.T) {
	tests := []struct {
		name string
		ts   *demoService
		want map[int]types.User
	}{
		{
			name: "Get users map. Success story.",
			ts:   &demoService{},
			want: usersMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.GetUsersMap(dataFileName)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

func Test_demoService_GetFeedUsers(t *testing.T) {
	tests := []struct {
		name string
		ts   *demoService
		want [][]string
	}{
		{
			name: "Get users feed. Success story.",
			ts:   &demoService{},
			want: feedFileData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.GetFeedUsers(feedFileName)

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}

// To make this test pass you need to run the server
func Test_demoService_FetchCsvFromRemote(t *testing.T) {
	tests := []struct {
		name string
		ts   *demoService
		want [][]string
	}{
		{
			name: "Get users. Success story.",
			ts:   &demoService{},
			want: feedFileData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.ts.FetchCsvFromRemote(feedUrl)

			if strings.Contains(err.Error(), "connection refused") {
				log.Printf("error: %v", err)
				log.Fatalf("Make sure the server is up on port 8080")
			}

			assert.Equal(t, tt.want, result)
			assert.NoError(t, err)
		})
	}
}
