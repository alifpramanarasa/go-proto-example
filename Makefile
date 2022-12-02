generate:
	protoc -I model/ --go_out=. model/*.proto