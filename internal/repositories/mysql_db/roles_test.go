package mysql_db_test

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/elSyarif/posnote-api.git/internal/core/domain"
	"github.com/elSyarif/posnote-api.git/internal/repositories/mysql_db"
	"github.com/jmoiron/sqlx"
)

func TestRoleRepoSave(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}

	type args struct {
		ctx   context.Context
		input domain.Roles
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       domain.Roles
		wantErr    bool
	}{
		{
			name: "fail save roles",
			args: args{
				ctx: context.TODO(),
				input: domain.Roles{
					Id:          "test-123",
					Name:        "test-admin",
					Description: "des",
				},
			},
			beforeTest: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO roles VALUES (?, ?, ?)`)).WithArgs("test-123", "test-admin", "des").WillReturnError(errors.New("wooops, error coy"))
			},
			wantErr: true,
		},
		{
			name: "success save roles",
			args: args{
				ctx: context.TODO(),
				input: domain.Roles{
					Id:          "test-123",
					Name:        "test-admin",
					Description: "des",
				},
			},
			beforeTest: func(mocksql sqlmock.Sqlmock) {
				mocksql.
					ExpectExec(regexp.QuoteMeta(`INSERT INTO roles VALUES (?, ?, ?)`)).
					WithArgs("test-123", "test-admin", "des").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: domain.Roles{
				Id:          "test-123",
				Name:        "test-admin",
				Description: "des",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &mysql_db.RoleRepositroy{
				DB: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSql)
			}

			got, err := r.Save(tt.args.ctx, &tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoleRepo.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want.Id) {
				t.Errorf("RoleRepo.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
