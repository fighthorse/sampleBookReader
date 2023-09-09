// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fighthorse/sampleBookReader/domain/models/model"
)

func newGoadminUserPermission(db *gorm.DB, opts ...gen.DOOption) goadminUserPermission {
	_goadminUserPermission := goadminUserPermission{}

	_goadminUserPermission.goadminUserPermissionDo.UseDB(db, opts...)
	_goadminUserPermission.goadminUserPermissionDo.UseModel(&model.GoadminUserPermission{})

	tableName := _goadminUserPermission.goadminUserPermissionDo.TableName()
	_goadminUserPermission.ALL = field.NewAsterisk(tableName)
	_goadminUserPermission.UserID = field.NewInt32(tableName, "user_id")
	_goadminUserPermission.PermissionID = field.NewInt32(tableName, "permission_id")
	_goadminUserPermission.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminUserPermission.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminUserPermission.fillFieldMap()

	return _goadminUserPermission
}

type goadminUserPermission struct {
	goadminUserPermissionDo goadminUserPermissionDo

	ALL          field.Asterisk
	UserID       field.Int32
	PermissionID field.Int32
	CreatedAt    field.Time
	UpdatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (g goadminUserPermission) Table(newTableName string) *goadminUserPermission {
	g.goadminUserPermissionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goadminUserPermission) As(alias string) *goadminUserPermission {
	g.goadminUserPermissionDo.DO = *(g.goadminUserPermissionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goadminUserPermission) updateTableName(table string) *goadminUserPermission {
	g.ALL = field.NewAsterisk(table)
	g.UserID = field.NewInt32(table, "user_id")
	g.PermissionID = field.NewInt32(table, "permission_id")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *goadminUserPermission) WithContext(ctx context.Context) *goadminUserPermissionDo {
	return g.goadminUserPermissionDo.WithContext(ctx)
}

func (g goadminUserPermission) TableName() string { return g.goadminUserPermissionDo.TableName() }

func (g goadminUserPermission) Alias() string { return g.goadminUserPermissionDo.Alias() }

func (g goadminUserPermission) Columns(cols ...field.Expr) gen.Columns {
	return g.goadminUserPermissionDo.Columns(cols...)
}

func (g *goadminUserPermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goadminUserPermission) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["permission_id"] = g.PermissionID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminUserPermission) clone(db *gorm.DB) goadminUserPermission {
	g.goadminUserPermissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g goadminUserPermission) replaceDB(db *gorm.DB) goadminUserPermission {
	g.goadminUserPermissionDo.ReplaceDB(db)
	return g
}

type goadminUserPermissionDo struct{ gen.DO }

func (g goadminUserPermissionDo) Debug() *goadminUserPermissionDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminUserPermissionDo) WithContext(ctx context.Context) *goadminUserPermissionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminUserPermissionDo) ReadDB() *goadminUserPermissionDo {
	return g.Clauses(dbresolver.Read)
}

func (g goadminUserPermissionDo) WriteDB() *goadminUserPermissionDo {
	return g.Clauses(dbresolver.Write)
}

func (g goadminUserPermissionDo) Session(config *gorm.Session) *goadminUserPermissionDo {
	return g.withDO(g.DO.Session(config))
}

func (g goadminUserPermissionDo) Clauses(conds ...clause.Expression) *goadminUserPermissionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminUserPermissionDo) Returning(value interface{}, columns ...string) *goadminUserPermissionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goadminUserPermissionDo) Not(conds ...gen.Condition) *goadminUserPermissionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminUserPermissionDo) Or(conds ...gen.Condition) *goadminUserPermissionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminUserPermissionDo) Select(conds ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminUserPermissionDo) Where(conds ...gen.Condition) *goadminUserPermissionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminUserPermissionDo) Order(conds ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminUserPermissionDo) Distinct(cols ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminUserPermissionDo) Omit(cols ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminUserPermissionDo) Join(table schema.Tabler, on ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminUserPermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminUserPermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminUserPermissionDo) Group(cols ...field.Expr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminUserPermissionDo) Having(conds ...gen.Condition) *goadminUserPermissionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminUserPermissionDo) Limit(limit int) *goadminUserPermissionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminUserPermissionDo) Offset(offset int) *goadminUserPermissionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminUserPermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminUserPermissionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminUserPermissionDo) Unscoped() *goadminUserPermissionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminUserPermissionDo) Create(values ...*model.GoadminUserPermission) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminUserPermissionDo) CreateInBatches(values []*model.GoadminUserPermission, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminUserPermissionDo) Save(values ...*model.GoadminUserPermission) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminUserPermissionDo) First() (*model.GoadminUserPermission, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminUserPermission), nil
	}
}

func (g goadminUserPermissionDo) Take() (*model.GoadminUserPermission, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminUserPermission), nil
	}
}

func (g goadminUserPermissionDo) Last() (*model.GoadminUserPermission, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminUserPermission), nil
	}
}

func (g goadminUserPermissionDo) Find() ([]*model.GoadminUserPermission, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoadminUserPermission), err
}

func (g goadminUserPermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminUserPermission, err error) {
	buf := make([]*model.GoadminUserPermission, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminUserPermissionDo) FindInBatches(result *[]*model.GoadminUserPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminUserPermissionDo) Attrs(attrs ...field.AssignExpr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminUserPermissionDo) Assign(attrs ...field.AssignExpr) *goadminUserPermissionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminUserPermissionDo) Joins(fields ...field.RelationField) *goadminUserPermissionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goadminUserPermissionDo) Preload(fields ...field.RelationField) *goadminUserPermissionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goadminUserPermissionDo) FirstOrInit() (*model.GoadminUserPermission, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminUserPermission), nil
	}
}

func (g goadminUserPermissionDo) FirstOrCreate() (*model.GoadminUserPermission, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminUserPermission), nil
	}
}

func (g goadminUserPermissionDo) FindByPage(offset int, limit int) (result []*model.GoadminUserPermission, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goadminUserPermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goadminUserPermissionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goadminUserPermissionDo) Delete(models ...*model.GoadminUserPermission) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goadminUserPermissionDo) withDO(do gen.Dao) *goadminUserPermissionDo {
	g.DO = *do.(*gen.DO)
	return g
}