package repository

import (
	"database/sql"
	"log"

	"example.com/m/v2/domain"
	"github.com/huandu/go-sqlbuilder"
)

type ProductTemplateRepository struct {
	db *sql.DB
}

func NewProductTemplateRepository(db *sql.DB) *ProductTemplateRepository {
	return &ProductTemplateRepository{db: db}
}

func (r *ProductTemplateRepository) Create(pt *domain.ProductTemplate) error {
	// Generate SQL query for inserting a new product template
	query := sqlbuilder.NewInsertBuilder()
	query.InsertInto("ProductTemplate").
		Cols(
			"ItemNameKz",
			"ItemNameRu",
			"ItemNameEn",
			"DescriptionKz",
			"DescriptionRu",
			"DescriptionEn",
			"ExternalId",
			"Barcode",
			// Add the remaining columns here
		).
		Values(
			pt.ItemNameKz,
			pt.ItemNameRu,
			pt.ItemNameEn,
			pt.DescriptionKz,
			pt.DescriptionRu,
			pt.DescriptionEn,
			pt.ExternalId,
			pt.Barcode,
			// Add the remaining values here
		)

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to create product template:", err)
		return err
	}

	return nil
}

func (r *ProductTemplateRepository) GetById(id int64) (*domain.ProductTemplate, error) {
	// Generate SQL query for selecting a product template by id
	query := sqlbuilder.NewSelectBuilder()
	query.Select("*").
		From("ProductTemplate").
		Where(query.Equal("ExternalId", id))

	// Execute the SQL query
	row := r.db.QueryRow(query.Build())
	pt, err := r.scanProductTemplate(row)
	if err != nil {
		log.Println("Failed to get product template by id:", err)
		return nil, err
	}

	return pt, nil
}

func (r *ProductTemplateRepository) Update(pt *domain.ProductTemplate) error {
	// Generate SQL query for updating an existing product template
	query := sqlbuilder.NewUpdateBuilder()
	query.Update("ProductTemplate").
		Set(query.Assign("ItemNameKz", pt.ItemNameKz)).
		Set(query.Assign("ItemNameRu", pt.ItemNameRu)).
		Set(query.Assign("ItemNameEn", pt.ItemNameEn)).
		Set(query.Assign("DescriptionKz", pt.DescriptionKz)).
		Set(query.Assign("DescriptionRu", pt.DescriptionRu)).
		Set(query.Assign("DescriptionEn", pt.DescriptionEn)).
		Set(query.Assign("Barcode", pt.Barcode)).
		// Add the remaining columns to be updated here
		Where(query.Equal("ExternalId", pt.ExternalId))

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to update product template:", err)
		return err
	}

	return nil
}

func (r *ProductTemplateRepository) Delete(pt *domain.ProductTemplate) error {
	// Generate SQL query for deleting a product template
	query := sqlbuilder.NewDeleteBuilder()
	query.DeleteFrom("ProductTemplate").
		Where(query.Equal("ExternalId", pt.ExternalId))

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to delete product template:", err)
		return err
	}

	return nil
}

func (r *ProductTemplateRepository) GetAll() ([]*domain.ProductTemplate, error) {
	// Generate SQL query for selecting all product templates
	query := sqlbuilder.NewSelectBuilder()
	query.Select("*").
		From("ProductTemplate")

	// Execute the SQL query
	rows, err := r.db.Query(query.Build())
	if err != nil {
		log.Println("Failed to get all product templates:", err)
		return nil, err
	}
	defer rows.Close()

	var productTemplates []*domain.ProductTemplate

	// Iterate over the rows and scan the product templates
	for rows.Next() {
		pt, err := r.scanProductTemplateRows(rows)
		if err != nil {
			log.Println("Failed to scan product template:", err)
			return nil, err
		}
		productTemplates = append(productTemplates, pt)
	}

	return productTemplates, nil
}

// Helper function to scan a single product template from a row
func (r *ProductTemplateRepository) scanProductTemplate(row *sql.Row) (*domain.ProductTemplate, error) {
	pt := &domain.ProductTemplate{}

	err := row.Scan(
		&pt.ItemNameKz,
		&pt.ItemNameRu,
		&pt.ItemNameEn,
		&pt.DescriptionKz,
		&pt.DescriptionRu,
		&pt.DescriptionEn,
		&pt.ExternalId,
		&pt.Barcode,
		// Add the remaining fields to be scanned here
	)
	if err != nil {
		return nil, err
	}

	return pt, nil
}

func (r *ProductTemplateRepository) scanProductTemplateRows(rows *sql.Rows) (*domain.ProductTemplate, error) {
	pt := &domain.ProductTemplate{}

	err := rows.Scan(
		&pt.ItemNameKz,
		&pt.ItemNameRu,
		&pt.ItemNameEn,
		&pt.DescriptionKz,
		&pt.DescriptionRu,
		&pt.DescriptionEn,
		&pt.ExternalId,
		&pt.Barcode,
		// Add the remaining fields to be scanned here
	)
	if err != nil {
		return nil, err
	}

	// Additional logic to populate the remaining fields of ProductTemplate

	return pt, nil
}
