package 策略模式

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func ExamplePayByBank() {
	payment := NewPayment("Ada", "123456", 888, &Bank{})
	payment.Pay()
}
