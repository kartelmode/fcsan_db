package console

var UnregisteredCommands []string = []string{
	"1. Register",
	"2. Log in",
	"3. Show products",
}

var UserCommands []string = []string{
	"1. Log out",
	"2. Show products",
	"3. Show orders",
	"4. Take product to basket",
	"5. Make an order",
	"6. Create organization",
	"7. Create product",
	"8. Show wallet",
	"9. Top up wallet balance",
	"10. Show my transactions",
	"11. Show products in the basket",
}

var AdminCommands []string = []string{
	"1. Log out",
	"2. Show products",
	"3. Show organization requests",
	"4. Approve organization request",
	"5. Reject organization request",
}

var Cmds [][]string = [][]string{
	UnregisteredCommands,
	UserCommands,
	AdminCommands,
}
