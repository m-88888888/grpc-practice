

import (
	"context"
	"fmt"
	"net"

	pb "github.com/m-88888888/grpc-practice/pb/product/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"

type ServerUnary struct {
	pb.UnimplementedProductServer
}

func server() error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "missing port.")
	}
	s := grpc.NewServer()
	var server ServerUnary
	pb.RegisterProductServer(s, &server)
	if err := s.Serve(listen); err != nil {
		return errors.Wrap(err, "Failed to start server.")
	}
	return nil
}

func (s *ServerUnary) DeleteProduct(ctx context.Context, in *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	id := in.GetId()
	fmt.Println(id)
	db := initDB()
	product := deleteProduct(db)
	// 削除したProductデータを返す
	reply := fmt.Sprintf(product)
	return &pb.DeleteProductReply{
		Message: reply,
	}, nil
}

// Product is struct
type Product struct {
	ID              int `gorm:"primary_key"`
	Code            string
	Price           uint
	ProductServices []ProductService `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ProductService is struct
type ProductService struct {
	ID        int `gorm:"primary_key"`
	ProductID int
	Name      string
}

func initDB() (db *gorm.DB) {
	DNS := "root:@tcp(localhost:3306)/product?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(DNS)

	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &ProductService{})

	fmt.Println("db connection is successful")
	return
}

func deleteProduct(db *gorm.DB) (product Product) {
	var product Product
	db.Debug().Find(&product)
	var productService ProductService
	db.Debug().Find(&productService)
	fmt.Println(product)
	fmt.Println(productService)
	db.Debug().Select("ProductServices").Delete(&product)
	return product
}
