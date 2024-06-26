package variantsrvc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"persia_atlas/server/db/rawsql"
	"persia_atlas/server/models"
	"persia_atlas/server/models/network"
	"persia_atlas/server/pagination"
)

type VariantServiceImpl struct {
	db *gorm.DB
}

func NewVariantService(db *gorm.DB) VariantService {
	return &VariantServiceImpl{
		db: db,
	}
}

func (vs VariantServiceImpl) CreateVariant(p *models.Variant) error {
	res := vs.db.Create(p)
	return res.Error
}

func (vs VariantServiceImpl) GetVariantsPaginated(c *gin.Context) *pagination.PaginatedResponse {
	var scopes []func(db *gorm.DB) *gorm.DB

	if searchQuery := c.Query("search"); searchQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.
				Where("title ILIKE ?", "%"+searchQuery+"%").
				Or("dkp = ?", searchQuery)
		})
	}
	if priceStepQuery := c.Query("price_step"); priceStepQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("price_step = ?", priceStepQuery)
		})
	}
	if isActiveQuery := c.Query("is_active"); isActiveQuery != "" {
		var isActive bool
		if isActiveQuery == "true" {
			isActive = true
		} else {
			isActive = false
		}
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", isActive)
		})
	}
	if hasCompetitionQuery := c.Query("has_competition"); hasCompetitionQuery != "" {
		var hasCompetition bool
		if hasCompetitionQuery == "true" {
			hasCompetition = true
		} else {
			hasCompetition = false
		}
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("has_competition = ?", hasCompetition)
		})
	}
	if selectorIdQuery := c.Query("selector_id"); selectorIdQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("selector_id = ?", selectorIdQuery)
		})
	}
	if actualProductIdQuery := c.Query("actual_product_id"); actualProductIdQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("actual_product_id = ?", actualProductIdQuery)
		})
	}
	if orderByQuery := c.Query("actual_product_id"); orderByQuery != "" {
		orders := []string{"price_min", "is_active", "has_competition"}
		if slices.Contains(orders, orderByQuery) {
			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
				return db.Order(orderByQuery)
			})
		}
	}

	chain := vs.db.
		Model(models.Variant{}).
		Preload("Product.Type.SelectorType").
		Preload("Selector.SelectorType").
		Preload("ActualProduct.Brand").
		Preload(clause.Associations)
	var variants []network.VariantSerializer
	paginator := pagination.NewPageNumberPaginator(c, chain, scopes, &variants)
	response := paginator.GetPaginatedResponse()

	return response
}

func (vs VariantServiceImpl) GetVariantById(id uint) *network.VariantSerializer {
	var variantSerializer network.VariantSerializer
	vs.db.
		Model(models.Variant{}).
		Preload("Product.Type.SelectorType").
		Preload("Selector.SelectorType").
		Preload("ActualProduct.Brand").
		Preload(clause.Associations).
		First(&variantSerializer, id) // Preload runs separate query for each table!
	//if variantSerializer.ID == 0 {
	//	return nil
	//}

	var vTest network.VariantSerializer
	//vs.db.
	//	Table("variants").
	//	Select("p.id").
	//	Joins("INNER JOIN products as p on variants.product_id = p.id").
	//	First(&vTest, id) // you may as well write raw sql!
	vs.db.Raw(rawsql.SqlVariant, id).Scan(&vTest) // scan doesn't work with nested structs
	pretty.Println("variant to VariantSerializer:", vTest)

	var nested network.VariantSerializerNested
	vs.db.Raw(rawsql.SqlVariant, id).Scan(&nested)
	pretty.Println("variant to serializer nested:", nested)

	var result map[string]any
	vs.db.Raw(rawsql.SqlVariantNoName, id).Scan(&result)
	pretty.Println("variant to map:", result)

	var result2 network.VariantScanner
	vs.db.Raw(rawsql.SqlVariant, id).Scan(&result2)
	pretty.Println("variant to VariantScanner:", result2)
	if result2.Id == 0 {
		return nil
	}
	var res2 = result2.Serialize()
	pretty.Println("variant to Serialize:", res2)

	var result3 network.VariantScannerGoConvention
	vs.db.Raw(rawsql.SqlVariant2, id).Scan(&result3)
	pretty.Println("variant to VariantScannerGoConvention:", result3)
	var res3 = result3.Serialize()
	pretty.Println("VariantScannerGoConvention to Serialize:", res3) // <-- BEST SOLUTION SO FAR

	tt := vs.db.Raw(rawsql.SqlVariantNoName, id)
	rows, _ := tt.Rows()
	cols, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	fmt.Println("cols:", cols)
	fmt.Println("column types:", colTypes[0].Name())
	fmt.Println("column types:", colTypes[0].DatabaseTypeName())
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for _, colName := range cols {
			// value is in columns[i] of interface type.
			// How to extract it from here?
			// ....
			fmt.Println("col name:", colName)
		}
	}

	//return &variantSerializer
	return &res3
}
