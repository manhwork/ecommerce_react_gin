package repo

type ProductRepo struct {
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (pr *ProductRepo) FindAllProduct() string {
	return "get all products"
}
