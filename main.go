package main

import (
	. "go-design-pattern/design_pattern"
)

func main() {
	// 单例模式
	//single := new(Singleton)
	//instance := single.GetInstance()
	//fmt.Println("instance:", instance)

	// 工厂模式
	//factory := new(ParseFactory)
	//parser := factory.NewFactory("json")
	//parser.Parse("hello world")

	// 建造者模式
	//director := new(Director)
	//// concreteBuilder := new(BuilderA)
	//concreteBuilder := new(BuilderB)
	//director.NewBuilder(concreteBuilder)
	//director.Build()
	//director.Show()

	// 原型模式
	//protoA := new(ProtoA)
	//protoA.AppendNodes([]string{"hello", "world", "david"})
	//protoA.UpdateKey(18, "eighteen")
	//
	//protoB := protoA.Clone()
	//protoB.Print()

	// 对象池
	//pool := InitPool(8)
	//for i := 0; i < 7; i++ {
	//	go func(i int) {
	//		activeObject := pool.Loan()
	//		id := activeObject.GetId()
	//		fmt.Printf("i:%d, id:%s\n", i, id)
	//	}(i)
	//}
	//time.Sleep(2 * 1e9)
	//activeObject := pool.Loan()
	//id := activeObject.GetId()
	//fmt.Printf("last one, id:%s\n", id)
	//
	//// destroy
	//pool.Destroy(activeObject)
	//recoverObject := pool.Loan()
	//rid := recoverObject.GetId()
	//fmt.Printf("recover one, id:%s\n", rid)

	// 责任链
	//reception := new(Reception)
	//reception.SetNext(new(Doctor)).SetNext(new(MedicineRoom))
	//patient := Patient{Name: "张三", Charge: 0,
	//	HasPaid: true, HasChecked: false, HasMedicine: false}
	//reception.Execute(&patient)

	// 命令
	//tv := new(TV)
	//onCommand := OnCommand{Device: tv}
	//buttonOn := Button{Command: &onCommand}
	//buttonOn.Press()
	//
	//lights := new(Lights)
	//offCommand := OffCommand{Device: lights}
	//buttonOff := Button{Command: &offCommand}
	//buttonOff.Press()

	// 迭代器
	//users := make([]*User, 0)
	//users = append(users, &User{Name: "david", Age: 22})
	//users = append(users, &User{Name: "john", Age: 20})
	//users = append(users, &User{Name: "leon", Age: 18})
	//
	//userCollection := UserCollection{Users: users}
	//iter := userCollection.CreateIterator()
	//for iter.HasNext() {
	//	curUser := iter.GetNext()
	//	fmt.Printf("name: %s, age: %d\n", curUser.Name, curUser.Age)
	//}

	// 中介者
	//platformMediator := NewPlatformMediator()
	//passengerTrain := PassengerTrain{Mediator: platformMediator}
	//goodsTrain := GoodsTrain{Mediator: platformMediator}
	//oilTrain := OilTrain{Mediator: platformMediator}
	//
	//go func() {
	//	duration := 2 * time.Second
	//	platformMediator.ArrangeDepart(duration)
	//}()
	//
	//passengerTrain.RequestArrival()
	//goodsTrain.RequestArrival()
	//oilTrain.RequestArrival()
	//
	//time.Sleep(10 * time.Second)

	// 备忘录
	//originator := new(Originator)
	//originator.SetState("INIT")
	//fmt.Println("state:", originator.GetState())
	//
	//memoRecords := new(MemoRecords)
	//memoRecords.AddMemo(originator.CreateMemo())
	//originator.SetState("SUCCESS")
	//memoRecords.AddMemo(originator.CreateMemo())
	//originator.RestoreMemo(memoRecords.FindMemo(0))
	//
	//fmt.Println("state:", originator.GetState())

	// 空对象
	//college := CreateCollege()
	//departmentList := []string{"math", "computer science", "physic"}
	//totalNum := 0
	//
	//for _, requireDepart := range departmentList {
	//	depart := college.GetDepartment(requireDepart)
	//	totalNum += depart.GetStudentNum()
	//}
	//
	//fmt.Println("totalNum:", totalNum)

	// 观察者 -- Update方法 修改*Item为Good报错
	//iPhone := &Item{Price: 5999}
	//richMan := &Buyer{Id: "01", ExpectPrice: 5000}
	//poorMan := &Buyer{Id: "02", ExpectPrice: 2500}
	//
	//iPhone.RegisterObserver(richMan)
	//iPhone.RegisterObserver(poorMan)
	//
	//iPhone.UpdatePrice(4999)
	//iPhone.UpdatePrice(2499)

	// 状态
	//vendingMachine := NewVendingMachine(0)
	//
	//vendingMachine.AddItem(1)
	//vendingMachine.SelectItem()
	//vendingMachine.InsertMoney(120)
	//vendingMachine.DispenseItem()

	// 策略
	//cache := InitCache(3, "LRU")
	//cache.Put(1, "one")
	//cache.Put(2, "two")
	//cache.Put(3, "three")
	//
	//cache.Get(1)
	//cache.Put(4, "four")
	//cache.Print()

	// 模板
	//humanBuild := new(Human)
	//humanSkeleton := Skeleton{
	//	T: humanBuild,
	//}
	//humanSkeleton.Merge()
	//
	//rocketBuild := new(Rocket)
	//rocketSkeleton := Skeleton{
	//	T: rocketBuild,
	//}
	//rocketSkeleton.Merge()

	// 访问者
	//cube := NewCube(3, 4, 5)
	//areaCalc := &AreaVisitor{}
	//volumeCalc := &VolumeVisitor{}
	//
	//cube.Accept(areaCalc)
	//cube.Accept(volumeCalc)

	// 适配器
	//connect := new(ConnectClient)
	//mac := new(Mac)
	//windowsAdaptor := &WindowsAdaptor{Win: new(Windows)}
	//
	//connect.InsertUsb(mac)
	//connect.InsertUsb(windowsAdaptor)

	// 桥接
	//laserPrinter := new(LaserPrinter)
	//laserPrinter.SetBrand(new(Epson))
	//laserPrinter.PrintFile()
	//
	//inkPrinter := new(InkPrinter)
	//inkPrinter.SetBrand(new(Hp))
	//inkPrinter.PrintFile()

	// 组合
	//root := &Folder{Name: "root"}
	//tempFile := &File{Name: "temp"}
	//xmlPath := &Folder{Name: "xmlFolder"}
	//xmlFile := &File{Name: "xml"}
	//root.AddComponent(tempFile)
	//xmlPath.AddComponent(xmlFile)
	//root.AddComponent(xmlPath)
	//
	//root.Search("temp")

	// 装饰器
	//pizza := &PizzaA{}
	//fullPizza := &ChessTopping{
	//	Pizza: pizza,
	//}
	//fmt.Printf("pizza price: %d\n", fullPizza.GetPrice())
	//
	//pizzaB := &PizzaB{}
	//fullPizzaB := &TomatoTopping{
	//	Pizza: pizzaB,
	//}
	//fmt.Printf("pizza price: %d\n", fullPizzaB.GetPrice())

	// 外观
	//buyFacade := new(BuyFacade)
	//buyFacade.BuySomething()

	// 享元
	playerFactory := &PlayerFactory{}
	var playerType string
	for i := 0; i < 10; i++ {
		playerType = func() string {
			if i%2 == 0 {
				return "terrorist"
			} else {
				return "counterTerrorist"
			}
		}()
		newPlayer := playerFactory.CreatePlayer(playerType)
		newPlayer.Info()
	}

	// 代理
	//maxAllowed := 3
	//nginxServer := NewNginx(2)
	//for i := 0; i < maxAllowed+1; i++ {
	//	httpCode, body := nginxServer.HandleRequest("/create/user", "POST")
	//	fmt.Printf("code: %d, body: %s \n", httpCode, body)
	//}
}
