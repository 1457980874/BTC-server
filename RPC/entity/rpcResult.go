package entity

type Data struct {
	id int
	err error
	result string
}

type RPCresult struct {
	code int
	msg string
	data Data
}
