## Migrar tabla de productos

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
	log.Fatalf("product.Migrate: %v", err)
}
```

## Migrar tabla de invoiceheader

```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

## Migrar tabla de invoiceitem

```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

# Crear un producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m := &product.Model{
Name:         "Door",
Observations: "Madera",
Price:        40,
}
if err := serviceProduct.Create(m); err != nil {
log.Fatalf("product.Create: %v", err)
}

fmt.Printf("%+v\n", m)
```

# Consultar Productos

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

ms, err := serviceProduct.GetAll()
if err != nil {
log.Fatalf("product.GetAll: %v", err)
}
for _, m := range ms {
fmt.Printf("ID: %d, Name: %s, Observations: %s, Price: %d, CreateAt: %s, UpdatedAt: %s\n",
m.ID, m.Name, m.Observations, m.Price, m.CreateAt.Format(time.RFC822), m.UpdateAt.Format(time.RFC822))
}
```

# Consultar un solo producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m, err := serviceProduct.GetByID(2)
switch {
case errors.Is(err, sql.ErrNoRows):
fmt.Println("No product found with this id")
case err != nil:
log.Fatalf("product.GetById: %v", err)
default:
fmt.Printf("ID: %d, Name: %s, Observations: %s, Price: %d, CreateAt: %s, UpdatedAt: %s\n",
m.ID, m.Name, m.Observations, m.Price, m.CreateAt.Format(time.RFC822), m.UpdateAt.Format(time.RFC822))
}
```

# Actualizar un producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

m := &product.Model{
ID:           8,
Name:         "Mouse",
Observations: "XD",
Price:        10,
}
err := serviceProduct.Update(m)
if err != nil {
log.Fatalf("product.Update: %v", err)
}
```

# Eliminar un producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

err := serviceProduct.Delete(2)
if err != nil {
	log.Fatalf("product.Delete: %v", err)
}
```

# Crear Factura (tx)

```go
storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
storageInvoice := storage.NewPsqlInvoice(
	storage.Pool(),
	storageHeader,
	storageItems,
)

m := &invoice.Model{
	Header: &invoiceheader.Model{
		Client: "Luis",
	},
	Items: invoiceitem.Models{
        &invoiceitem.Model{ProductID: 3},
        &invoiceitem.Model{ProductID: 4},
	},
}

serviceInvoice := invoice.NewService(storageInvoice)
if err := serviceInvoice.Create(m); err != nil {
	log.Fatalf("invoice.Create: %v", err)
}
```
