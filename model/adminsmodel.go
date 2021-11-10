package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	adminsFieldNames          = builderx.RawFieldNames(&Admins{})
	adminsRows                = strings.Join(adminsFieldNames, ",")
	adminsRowsExpectAutoSet   = strings.Join(stringx.Remove(adminsFieldNames, "`id`", "`created_at`", "`updated_at`", "`deleted_at`"), ",")
	adminsRowsWithPlaceHolder = strings.Join(stringx.Remove(adminsFieldNames, "`id`", "`created_at`", "`updated_at`", "`deleted_at`"), "=?,") + "=?"
)

type (
	AdminsModel interface {
		Insert(data Admins) (sql.Result, error)
		FindOne(id int64) (*Admins, error)
		FindOneByPhoneDeletedAt(phone string, deletedAt sql.NullTime) (*Admins, error)
		Update(data Admins) error
		Delete(id int64) error
		FindPhoneOne(phone string) (*Admins, error)
	}

	defaultAdminsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Admins struct {
		Id        int64        `db:"id"`         // 主键ID
		Name      string       `db:"name"`       // 员工名称
		Phone     string       `db:"phone"`      // 手机号
		Password  string       `db:"password"`   // 密码
		Status    int64        `db:"status"`     // 状态 1、启用 2、关闭
		CreatedAt time.Time    `db:"created_at"` // 注册时间
		UpdatedAt time.Time    `db:"updated_at"` // 修改时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewAdminsModel(conn sqlx.SqlConn) AdminsModel {
	return &defaultAdminsModel{
		conn:  conn,
		table: "`admins`",
	}
}

func (m *defaultAdminsModel) Insert(data Admins) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, adminsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Phone, data.Password, data.Status)
	return ret, err
}

func (m *defaultAdminsModel) FindOne(id int64) (*Admins, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminsRows, m.table)
	var resp Admins
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminsModel) FindOneByPhoneDeletedAt(phone string, deletedAt sql.NullTime) (*Admins, error) {
	var resp Admins
	query := fmt.Sprintf("select %s from %s where `phone` = ? and `deleted_at` = ? limit 1", adminsRows, m.table)
	err := m.conn.QueryRow(&resp, query, phone, deletedAt)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminsModel) Update(data Admins) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Phone, data.Password, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultAdminsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultAdminsModel) FindPhoneOne(phone string) (*Admins, error) {
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", adminsRows, m.table)
	var resp Admins
	err := m.conn.QueryRow(&resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
