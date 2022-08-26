package 抽象工厂模式

import "fmt"

//OrderMainDAO 为订单主记录
//OrderDetailDAO 为订单详情纪录
//DAOFactory DAO 抽象模式工厂接口
//RDBMainDAP 为关系型数据库的OrderMainDAO实现
//SaveOrderMain ...
//RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
// SaveOrderDetail ...
//RDBDAOFactory 是RDB 抽象工厂实现
//XMLMainDAO XML存储
//SaveOrderMain ...
//XMLDetailDAO XML存储
// SaveOrderDetail ...
//XMLDAOFactory 是RDB 抽象工厂实现

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderMainDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderMainDetail() OrderMainDetailDAO
}

type RDBMainDAP struct{}

func (RDBMainDAP) SaveOrderMain() {
	fmt.Print("RDBMainDAP SaveOrderMain")
}

type RDBDetailDAP struct{}

func (RDBDetailDAP) SaveOrderDetail() {
	fmt.Print("RDBDetailDAP SaveOrderDetail")
}

type XMLMainDAP struct{}

func (XMLMainDAP) SaveOrderMain() {
	fmt.Print("xml SaveOrderMain")
}

type XMLDetailDAP struct{}

func (XMLDetailDAP) SaveOrderDetail() {
	fmt.Print("xml SaveOrderDetail")
}

type RDBDAOFactory struct{}

func (RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAP{}
}
func (RDBDAOFactory) CreateOrderMainDetail() OrderMainDetailDAO {
	return &RDBDetailDAP{}
}

type XMLDAOFactory struct {
}

func (XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAP{}
}
func (XMLDAOFactory) CreateOrderMainDetail() OrderMainDetailDAO {
	return &XMLDetailDAP{}
}
