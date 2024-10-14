package main

import "fmt"

type Banker struct {
}

func (p *Banker) Save() {
	fmt.Println("存款")
}

func (p *Banker) Transfer() {
	fmt.Println("转账")
}

func (p *Banker) Pay() {
	fmt.Println("支付")
}

// 扩充业务修改代码:
// 取款()
// 购买保险()
// ......

type AbstarctBanker interface {
	DoBusi()
}

type SaveBanker struct{}

func (p *SaveBanker) DoBusi() {
	fmt.Println("存款")
}

type TransferBanker struct{}

func (p *TransferBanker) DoBusi() {
	fmt.Println("转账")
}

type PayBanker struct{}

func (p *PayBanker) DoBusi() {
	fmt.Println("支付")
}

// 增加业务不用修改代码，增加新类型即可

func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()

	DoBusi(&SaveBanker{})
	DoBusi(&TransferBanker{})
	DoBusi(&PayBanker{})
}

func DoBusi(banker AbstarctBanker) {
	banker.DoBusi()
}
