package product

import (
	"context"

	"github.com/minio/minio-go"
	"github.com/zelvann/minio-ds/internal/utils"
	"gorm.io/gorm"
)

type (
	ProductUsecase interface {
		CreateProduct(context.Context, CreateProductDTO) error
		GetProductByID(context.Context, string) (GetProductDTO, error)
		UpdateProduct(context.Context, UpdateProductDTO, string) error
		DeleteProduct(context.Context, string) error
	}

	productUsecase struct {
		productRepo ProductRepository
		mc          *minio.Client
	}
)

func NewProductUsecase(productRepo ProductRepository, mc *minio.Client) ProductUsecase {
	return &productUsecase{productRepo: productRepo, mc: mc}
}

func (u *productUsecase) CreateProduct(ctx context.Context, req CreateProductDTO) error {
	ext := utils.GetExtensions(req.ProductPicture.Filename)
	if ext != "jpg" && ext != "jpeg" && ext != "png" {
		return ErrFormatFileNotValid
	}

	exist, err := u.mc.BucketExists(PRODUCT_BUCKET)
	if err != nil {
		return ErrMinioSomethingWentWrong
	}

	if !exist {
		if err = u.mc.MakeBucket(PRODUCT_BUCKET, "ap-southeast-3"); err != nil {
			return ErrMinioSomethingWentWrong
		}
	}

	file, err := req.ProductPicture.Open()
	if err != nil {
		return ErrSomethingWentWrong
	}
	defer file.Close()

	product := ProductSchema{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		ExpireDate:  req.ExpireDate,
		ImageName:   req.ProductPicture.Filename,
	}

	_, err = u.mc.PutObject(PRODUCT_BUCKET, req.ProductPicture.Filename, file, req.ProductPicture.Size, minio.PutObjectOptions{
		ContentType: req.ProductPicture.Header.Get("Content-Type"),
	})

	if err != nil {
		return ErrMinioSomethingWentWrong
	}

	err = u.productRepo.Create(product)
	return err
}

func (u *productUsecase) GetProductByID(ctx context.Context, id string) (GetProductDTO, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return GetProductDTO{}, ErrProductNotFound
		} else {
			return GetProductDTO{}, ErrSomethingWentWrong
		}
	}

	return GetProductDTO{
		Name:        product.Name,
		Description: product.Description,
		ExpireDate:  product.ExpireDate.String(),
	}, nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, req UpdateProductDTO, id string) error {
	product := ProductSchema{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := u.productRepo.Update(id, product); err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrProductNotFound
		} else {
			return ErrSomethingWentWrong
		}
	}

	return nil

}

func (u *productUsecase) DeleteProduct(ctx context.Context, id string) error {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrProductNotFound
		} else {
			return ErrSomethingWentWrong
		}
	}

	if err = u.mc.RemoveObject(PRODUCT_BUCKET, product.ImageName); err != nil {
		return ErrMinioSomethingWentWrong
	}

	err = u.productRepo.Delete(id)
	if err != nil {
		return ErrSomethingWentWrong
	}

	return nil
}
