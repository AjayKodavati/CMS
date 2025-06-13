package repository

import (
	"context"

	"github.com/AjayKodavati/CMS/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepository interface {
	CreateCategory(ctx context.Context, categoryName db.Category) (error)
	GetCategoryIDByName(ctx context.Context, categoryName string) (int, error)
}

type CategoriesRepositoryService struct {
	pool *pgxpool.Pool
}

func NewCategoriesRepository(pool *pgxpool.Pool) CategoriesRepository {
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

func (c *CategoriesRepositoryService) CreateCategory(ctx context.Context, category db.Category) (error) {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING category_id`
	_, err := c.pool.Exec(ctx, query, category.CategoryName)
	if err != nil {
		return err
	}
	return nil
}