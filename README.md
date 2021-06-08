## fineract-client
simple fineract client for easy dev & test.

### install binaries
```
# creates the binaries in this project under your $GOPATH
./install.sh
```

### loan product commands
@TODO

### loan account commands
```
# creates a new loan account
loanaccount create

# creates, approves and disburses a new loan account
loanaccount open

# makes a loan repayment
loanaccount repay -id=<loanId> -p=<principal> -i=<interest> -f=<fee> -a=<amount> -d=<date>
# eg.
loanaccount repay -id=64 -i=75000 -d="31 March 2021"
```

### project commands
```
# creates a new project
project create -n="New Project Name"

# updates the project with id
project update -id=123
```