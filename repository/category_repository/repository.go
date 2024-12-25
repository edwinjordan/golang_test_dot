package category_repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/edwinjordan/golang_test_dot.git/app/repository"
	"github.com/edwinjordan/golang_test_dot.git/entity"
	"github.com/edwinjordan/golang_test_dot.git/pkg/helpers"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB          *gorm.DB
	redisClient *redis.Client
}

func New(db *gorm.DB, redisClient *redis.Client) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		DB:          db,
		redisClient: redisClient,
	}
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (entity.CategoryResponse, error) {
	categoryData := &Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("category_id = ?", categoryId).
		First(&categoryData).Error
	//panic(err)
	if err != nil {
		return *categoryData.ToEntity(), errors.New("data kategori tidak ditemukan")
	}
	return *categoryData.ToEntity(), nil
}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context) []entity.CategoryResponse {
	var tempData []entity.CategoryResponse

	cacheKey := "categories"
	cachedData, err := repo.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &tempData)
		if err == nil {
			return tempData
		}
	}

	category := []Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err = tx.WithContext(ctx).Find(&category).Error
	helpers.PanicIfError(err)

	for _, v := range category {
		tempData = append(tempData, *v.ToEntity())
	}

	cachedDataBytes, err := json.Marshal(tempData)
	if err == nil {
		repo.redisClient.Set(ctx, cacheKey, cachedDataBytes, 0)
	}

	return tempData
}

func (repo *CategoryRepositoryImpl) Create(ctx context.Context, category entity.Category) entity.CategoryResponse {
	categoryData := &Category{}
	categoryData = categoryData.FromEntity(&category)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&categoryData).Error
	helpers.PanicIfError(err)

	// Hapus cache jika ada
	cacheKey := "categories"
	repo.redisClient.Del(ctx, cacheKey)

	return *categoryData.ToEntity()
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse {
	categoryData := &Category{}
	categoryData = categoryData.FromEntity(&category)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_id = ?", categoryId).Select(selectField).Updates(&categoryData).Error
	helpers.PanicIfError(err)

	// Hapus cache jika ada
	cacheKey := "categories"
	repo.redisClient.Del(ctx, cacheKey)

	return *categoryData.ToEntity()
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId string) {
	category := &Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_id = ?", categoryId).Delete(&category).Error
	helpers.PanicIfError(err)

	// Hapus cache jika ada
	cacheKey := "categories"
	repo.redisClient.Del(ctx, cacheKey)
}

func (repo *CategoryRepositoryImpl) Patch(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse {
	categoryData := &Category{}
	categoryData = categoryData.FromEntity(&category)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_id = ?", categoryId).Select(selectField).Updates(&categoryData).Error

	helpers.PanicIfError(err)

	// Hapus cache jika ada
	cacheKey := "categories"
	repo.redisClient.Del(ctx, cacheKey)
	return *categoryData.ToEntity()
}
