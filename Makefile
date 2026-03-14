proto-gen:
	protoc -I=./proto --go_out=./backend/pb --go_opt=paths=source_relative ./proto/addressbook.proto
