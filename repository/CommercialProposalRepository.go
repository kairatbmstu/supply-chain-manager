package repository

import (
	"database/sql"
	"log"

	"example.com/m/v2/domain"
	"github.com/huandu/go-sqlbuilder"
)

type CommercialProposalRepository struct {
	db *sql.DB
}

func NewCommercialProposalRepository(db *sql.DB) *CommercialProposalRepository {
	return &CommercialProposalRepository{db: db}
}

func (r *CommercialProposalRepository) Create(cp *domain.CommercialProposal) error {
	// Generate SQL query for inserting a new commercial proposal
	query := sqlbuilder.NewInsertBuilder()
	query.InsertInto("CommercialProposal").
		Cols("Id", "ProcessDate", "CommercialProposalStatus").
		Values(cp.Id, cp.ProcessDate, cp.CommercialProposalStatus)

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to create commercial proposal:", err)
		return err
	}

	return nil
}

func (r *CommercialProposalRepository) GetById(id int64) (*domain.CommercialProposal, error) {
	// Generate SQL query for selecting a commercial proposal by id
	query := sqlbuilder.NewSelectBuilder()
	query.Select("*").
		From("CommercialProposal").
		Where(query.Equal("Id", id))

	// Execute the SQL query
	row := r.db.QueryRow(query.Build())
	cp, err := r.scanCommercialProposalRow(row)
	if err != nil {
		log.Println("Failed to get commercial proposal by id:", err)
		return nil, err
	}

	return cp, nil
}

func (r *CommercialProposalRepository) Update(cp *domain.CommercialProposal) error {
	// Generate SQL query for updating an existing commercial proposal
	query := sqlbuilder.NewUpdateBuilder()
	query.Update("CommercialProposal").
		Set(query.Assign("ProcessDate", cp.ProcessDate)).
		Set(query.Assign("CommercialProposalStatus", cp.CommercialProposalStatus)).
		Where(query.Equal("Id", cp.Id))

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to update commercial proposal:", err)
		return err
	}

	return nil
}

func (r *CommercialProposalRepository) Delete(cp *domain.CommercialProposal) error {
	// Generate SQL query for deleting a commercial proposal
	query := sqlbuilder.NewDeleteBuilder()
	query.DeleteFrom("CommercialProposal").
		Where(query.Equal("Id", cp.Id))

	// Execute the SQL query
	_, err := r.db.Exec(query.Build())
	if err != nil {
		log.Println("Failed to delete commercial proposal:", err)
		return err
	}

	return nil
}

func (r *CommercialProposalRepository) GetAll() ([]*domain.CommercialProposal, error) {
	// Generate SQL query for selecting all commercial proposals
	query := sqlbuilder.NewSelectBuilder()
	query.Select("*").
		From("CommercialProposal")

	// Execute the SQL query
	rows, err := r.db.Query(query.Build())
	if err != nil {
		log.Println("Failed to get all commercial proposals:", err)
		return nil, err
	}
	defer rows.Close()

	var commercialProposals []*domain.CommercialProposal

	// Iterate over the rows and scan the commercial proposals
	for rows.Next() {
		cp, err := r.scanCommercialProposalRows(rows)
		if err != nil {
			log.Println("Failed to scan commercial proposal:", err)
			return nil, err
		}
		commercialProposals = append(commercialProposals, cp)
	}

	return commercialProposals, nil
}

// Helper

func (r *CommercialProposalRepository) scanCommercialProposalRow(row *sql.Row) (*domain.CommercialProposal, error) {
	cp := &domain.CommercialProposal{}

	err := row.Scan(
		&cp.Id,
		&cp.ProcessDate,
		&cp.CommercialProposalStatus,
	)
	if err != nil {
		return nil, err
	}

	// Additional logic to populate the remaining fields of CommercialProposal

	return cp, nil
}

func (r *CommercialProposalRepository) scanCommercialProposalRows(rows *sql.Rows) (*domain.CommercialProposal, error) {
	cp := &domain.CommercialProposal{}

	err := rows.Scan(
		&cp.Id,
		&cp.ProcessDate,
		&cp.CommercialProposalStatus,
	)
	if err != nil {
		return nil, err
	}

	// Additional logic to populate the remaining fields of CommercialProposal

	return cp, nil
}
