.PHONY: go/singly_linked_list
go/singly_linked_list:
	@go run go/singly_linked_list.go

.PHONY: c/singly_linked_list
c/singly_linked_list:
	@gcc -o c/singly_linked_list c/singly_linked_list.c && ./c/singly_linked_list
