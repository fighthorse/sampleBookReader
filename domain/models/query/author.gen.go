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

func newAuthor(db *gorm.DB, opts ...gen.DOOption) author {
	_author := author{}

	_author.authorDo.UseDB(db, opts...)
	_author.authorDo.UseModel(&model.Author{})

	tableName := _author.authorDo.TableName()
	_author.ALL = field.NewAsterisk(tableName)
	_author.ID = field.NewInt32(tableName, "id")
	_author.AuthorName = field.NewString(tableName, "author_name")
	_author.AuthorDesc = field.NewString(tableName, "author_desc")
	_author.AuthorDay = field.NewString(tableName, "author_day")
	_author.AuthorSex = field.NewString(tableName, "author_sex")

	_author.fillFieldMap()

	return _author
}

type author struct {
	authorDo authorDo

	ALL        field.Asterisk
	ID         field.Int32
	AuthorName field.String // 作者姓名
	AuthorDesc field.String // 作者简介
	AuthorDay  field.String // 作者入住时间
	AuthorSex  field.String // 作者性别

	fieldMap map[string]field.Expr
}

func (a author) Table(newTableName string) *author {
	a.authorDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a author) As(alias string) *author {
	a.authorDo.DO = *(a.authorDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *author) updateTableName(table string) *author {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.AuthorName = field.NewString(table, "author_name")
	a.AuthorDesc = field.NewString(table, "author_desc")
	a.AuthorDay = field.NewString(table, "author_day")
	a.AuthorSex = field.NewString(table, "author_sex")

	a.fillFieldMap()

	return a
}

func (a *author) WithContext(ctx context.Context) *authorDo { return a.authorDo.WithContext(ctx) }

func (a author) TableName() string { return a.authorDo.TableName() }

func (a author) Alias() string { return a.authorDo.Alias() }

func (a author) Columns(cols ...field.Expr) gen.Columns { return a.authorDo.Columns(cols...) }

func (a *author) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *author) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["author_name"] = a.AuthorName
	a.fieldMap["author_desc"] = a.AuthorDesc
	a.fieldMap["author_day"] = a.AuthorDay
	a.fieldMap["author_sex"] = a.AuthorSex
}

func (a author) clone(db *gorm.DB) author {
	a.authorDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a author) replaceDB(db *gorm.DB) author {
	a.authorDo.ReplaceDB(db)
	return a
}

type authorDo struct{ gen.DO }

func (a authorDo) Debug() *authorDo {
	return a.withDO(a.DO.Debug())
}

func (a authorDo) WithContext(ctx context.Context) *authorDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authorDo) ReadDB() *authorDo {
	return a.Clauses(dbresolver.Read)
}

func (a authorDo) WriteDB() *authorDo {
	return a.Clauses(dbresolver.Write)
}

func (a authorDo) Session(config *gorm.Session) *authorDo {
	return a.withDO(a.DO.Session(config))
}

func (a authorDo) Clauses(conds ...clause.Expression) *authorDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authorDo) Returning(value interface{}, columns ...string) *authorDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authorDo) Not(conds ...gen.Condition) *authorDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authorDo) Or(conds ...gen.Condition) *authorDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authorDo) Select(conds ...field.Expr) *authorDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authorDo) Where(conds ...gen.Condition) *authorDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authorDo) Order(conds ...field.Expr) *authorDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authorDo) Distinct(cols ...field.Expr) *authorDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authorDo) Omit(cols ...field.Expr) *authorDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authorDo) Join(table schema.Tabler, on ...field.Expr) *authorDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authorDo) LeftJoin(table schema.Tabler, on ...field.Expr) *authorDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authorDo) RightJoin(table schema.Tabler, on ...field.Expr) *authorDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authorDo) Group(cols ...field.Expr) *authorDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authorDo) Having(conds ...gen.Condition) *authorDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authorDo) Limit(limit int) *authorDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authorDo) Offset(offset int) *authorDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *authorDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authorDo) Unscoped() *authorDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authorDo) Create(values ...*model.Author) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authorDo) CreateInBatches(values []*model.Author, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authorDo) Save(values ...*model.Author) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authorDo) First() (*model.Author, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Take() (*model.Author, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Last() (*model.Author, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) Find() ([]*model.Author, error) {
	result, err := a.DO.Find()
	return result.([]*model.Author), err
}

func (a authorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Author, err error) {
	buf := make([]*model.Author, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authorDo) FindInBatches(result *[]*model.Author, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authorDo) Attrs(attrs ...field.AssignExpr) *authorDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authorDo) Assign(attrs ...field.AssignExpr) *authorDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authorDo) Joins(fields ...field.RelationField) *authorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authorDo) Preload(fields ...field.RelationField) *authorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authorDo) FirstOrInit() (*model.Author, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) FirstOrCreate() (*model.Author, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Author), nil
	}
}

func (a authorDo) FindByPage(offset int, limit int) (result []*model.Author, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authorDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authorDo) Delete(models ...*model.Author) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authorDo) withDO(do gen.Dao) *authorDo {
	a.DO = *do.(*gen.DO)
	return a
}
