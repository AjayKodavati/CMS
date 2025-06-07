package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepository interface {
	GetCategoryIDByName(ctx context.Context, categoryName string) (int, error)
}

type CategoriesRepositoryService struct {
	pool *pgxpool.Pool
}

func NewCategoriesRepository(pool *pgxpool.Pool) *CategoriesRepositoryService {
	return &CategoriesRepositoryService{
		pool: pool,
	}
}

func (c *CategoriesRepositoryService) GetCategoryIDByName(ctx context.Context, categoryName string) (int, error) {
	var categoryID int
	query := `SELECT category_id FROM categories WHERE category_name = $1`
	err := c.pool.QueryRow(ctx, query, categoryName).Scan(&categoryID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, nil // No category found with this name
		}
		return 0, err // Return the error if something else went wrong
	}
	return categoryID, nil
}