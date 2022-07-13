package pkg

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
)

func TestStorage_AllTasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.AllTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("AllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TasksByAuthor(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.TasksByAuthor(tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TasksByAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}