package impl

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
)

// 完成对象和数据库之间的转换

// 把Host对象保存到数据库里
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 把数据入库到 resource表和host表
	// 一次需要往2个表录入数据, 我们需要2个操作 要么都成功，要么都失败, 事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error %s", err)
	}

	// 通过Defer处理事务提交方式
	// 1. 无错误，则提交 commit事务
	// 2. 有报错，则rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Errorf("rollback error %s", err)
			}

		} else {
			if err := tx.Commit(); err != nil {
				i.l.Errorf("commit error %s", err)
			}
		}
	}()

	// 插入resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return fmt.Errorf("InsertResourceSQL: %s", err)
	}

	// 插入describe数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return fmt.Errorf("InsertDescribeSQL: %s", err)
	}

	return nil
}

func (i *HostServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 开启一个实物 tx
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// 通过Defer处理事务提交方式
	// 1. 无错误，则提交 commit事务
	// 2. 有报错，则rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Errorf("rollback error %s", err)
			}

		} else {
			if err := tx.Commit(); err != nil {
				i.l.Errorf("commit error %s", err)
			}
		}
	}()

	var (
		resStmt, hostStmt *sql.Stmt
	)

	// 更新resource表
	resStmt, err = tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return err
	}
	// 更新host表
	hostStmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return err
	}

	_, err = hostStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}
	return nil
}

func (i *HostServiceImpl) delete(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	if ins.Id == "" {
		return fmt.Errorf("Input params not has instance id")
	}

	// 开启一个实物 tx
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// 通过Defer处理事务提交方式
	// 1. 无错误，则提交 commit事务
	// 2. 有报错，则rollback事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Errorf("rollback error %s", err)
			}

		} else {
			if err := tx.Commit(); err != nil {
				i.l.Errorf("commit error %s", err)
			}
		}
	}()

	var (
		resStmt, hostStmt *sql.Stmt
	)

	// 删除resource表数据
	resStmt, err = tx.PrepareContext(ctx, deleteResourceSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.ExecContext(ctx, ins.Id)
	if err != nil {
		return err
	}

	// 删除host表数据
	hostStmt, err = tx.PrepareContext(ctx, deleteHostSQL)
	if err != nil {
		return err
	}
	_, err = hostStmt.ExecContext(ctx, ins.Id)
	if err != nil {
		return err
	}
	return nil
}
