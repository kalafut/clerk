class Account:
    def __init__(self, name):
        self.name = name
        self.children = None
        self.parent = None

root = Account("__root")

def find_or_create(name):
    """
	var child *Account
	var name string
	var ok bool

	idx := strings.Index(acctName, ":")
	if idx == -1 {
		name = acctName
	} else {
		name = acctName[0:idx]
	}

	if child, ok = acct.children[name]; !ok {
		child = &Account{
			Name:     name,
			children: make(map[string]*Account),
			parent:   acct,
		}
		acct.children[name] = child
	}

	if idx != -1 {
		child = child.FindOrAddAccount(acctName[idx+1:])
	}

	return child


    """

