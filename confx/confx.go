package confx

var DefaultConfig = &TestConfig{
	A: 0.5,
	B: "prod",
	C: true,
}

type TestConfig struct {
	A float64
	B string
	C bool
}

func getTestConfig() *TestConfig {
	// read from config center
	// if err != nil {
	//		logs.Error()
	//		return DefaultConfig
	// }
	// res, err := parse()
	// if err != nil {
	// 		logs.Error()
	//		return DefaultConfig
	// }
	// return res
	return nil
}

func GetA() float64 {
	return getTestConfig().A
}

func GetB() string {
	return getTestConfig().B
}

func GetC() bool {
	return getTestConfig().C
}
