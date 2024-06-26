// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	likeFieldNames          = builder.RawFieldNames(&Like{})
	likeRows                = strings.Join(likeFieldNames, ",")
	likeRowsExpectAutoSet   = strings.Join(stringx.Remove(likeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	likeRowsWithPlaceHolder = strings.Join(stringx.Remove(likeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheBeyondLikeLikeIdPrefix = "cache:beyondLike:like:id:"
	cacheBeyondLikeLikeBizIdTargetIdUserIdPrefix = "cache:beyondLike:likeRecord:bizId:targetId:userId:"
)

type (
	likeModel interface {
		Insert(ctx context.Context, data *Like) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Like, error)
		FindOneByBizIdObjIdUserId(ctx context.Context, bizId string, objId int64, userId int64) (*Like, error)
		Update(ctx context.Context, data *Like) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikeModel struct {
		sqlc.CachedConn
		table string
	}

	Like struct {
		Id         int64     `db:"id"`          // 主键ID
		BizId      string    `db:"biz_id"`      // 业务ID
		TargetId   int64     `db:"target_id"`   // 点赞目标id
		UserId     int64     `db:"user_id"`     // 用户ID
		Type       int64     `db:"type"`        // 类型 0:点赞 1:点踩
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 最后修改时间
	}
)

func newLikeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLikeModel {
	return &defaultLikeModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`like`",
	}
}

func (m *defaultLikeModel) withSession(session sqlx.Session) *defaultLikeModel {
	return &defaultLikeModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`like`",
	}
}

func (m *defaultLikeModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	beyondLikeLikeRecordBizIdObjIdUserIdKey := fmt.Sprintf("%s%v:%v:%v", cacheBeyondLikeLikeBizIdTargetIdUserIdPrefix, data.BizId, data.TargetId, data.UserId)
	beyondLikeLikeRecordIdKey := fmt.Sprintf("%s%v", cacheBeyondLikeLikeIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, beyondLikeLikeRecordBizIdObjIdUserIdKey, beyondLikeLikeRecordIdKey)
	return err
}

func (m *defaultLikeModel) FindOne(ctx context.Context, id int64) (*Like, error) {
	beyondLikeLikeIdKey := fmt.Sprintf("%s%v", cacheBeyondLikeLikeIdPrefix, id)
	var resp Like
	err := m.QueryRowCtx(ctx, &resp, beyondLikeLikeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeModel) FindOneByBizIdObjIdUserId(ctx context.Context, bizId string, objId int64, userId int64) (*Like, error) {
	beyondLikeLikeRecordBizIdObjIdUserIdKey := fmt.Sprintf("%s%v:%v:%v", cacheBeyondLikeLikeBizIdTargetIdUserIdPrefix, bizId, objId, userId)
	var resp Like
	err := m.QueryRowIndexCtx(ctx, &resp, beyondLikeLikeRecordBizIdObjIdUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `biz_id` = ? and `obj_id` = ? and `user_id` = ? limit 1", likeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, bizId, objId, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeModel) Insert(ctx context.Context, data *Like) (sql.Result, error) {
	beyondLikeLikeRecordBizIdObjIdUserIdKey := fmt.Sprintf("%s%v:%v:%v", cacheBeyondLikeLikeBizIdTargetIdUserIdPrefix, data.BizId, data.TargetId, data.UserId)
	beyondLikeLikeRecordIdKey := fmt.Sprintf("%s%v", cacheBeyondLikeLikeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, likeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.BizId, data.TargetId, data.UserId, data.Type)
	}, beyondLikeLikeRecordBizIdObjIdUserIdKey, beyondLikeLikeRecordIdKey)
	return ret, err
}

func (m *defaultLikeModel) Update(ctx context.Context, newData *Like) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	beyondLikeLikeRecordBizIdObjIdUserIdKey := fmt.Sprintf("%s%v:%v:%v", cacheBeyondLikeLikeBizIdTargetIdUserIdPrefix, data.BizId, data.TargetId, data.UserId)
	beyondLikeLikeRecordIdKey := fmt.Sprintf("%s%v", cacheBeyondLikeLikeIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, likeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.BizId, newData.TargetId, newData.UserId, newData.Type, newData.Id)
	}, beyondLikeLikeRecordBizIdObjIdUserIdKey, beyondLikeLikeRecordIdKey)
	return err
}

func (m *defaultLikeModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBeyondLikeLikeIdPrefix, primary)
}

func (m *defaultLikeModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLikeModel) tableName() string {
	return m.table
}
