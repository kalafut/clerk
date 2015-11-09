// Common csv importer support
package main

type columnConfig struct {
	date        string
	posted      string
	code        string
	description string
	amount      string
	cost        string
	total       string
	note        string
}

type csvConfig struct {
	multiaccount bool
	invertAmount bool
	columns      columnConfig
}

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"
//Type,Trans Date,Post Date,Description,Amount
