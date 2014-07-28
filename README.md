go-enom-client
==============

A [Enom API Catalog](http://enom.com/APICommandCatalog/) client written in Go.

Library under heavy development, so please be patient if some things break.


# Usage

	# setup custom configs
	var config &enom.Config{}
	config.Username = "your username"
	config.Password = "your password"
	config.Url = "Enom API URL"
	config.ResponseType = "XML"

	# create client
	var client &enom.Client
	client = &enom.EnomClient{ Config: config }

	# perform domain check
	factory := client.CommandFactory()
	command := factory.DomainCheck("a-simple-domain.com")
	result, err := client.Execute(command)


# Available Commands

*	Domain Commands

	- Check
	- Purchase
	- Extend
	- DeleteRegistration
	- GetDomainInfo
	- GetPasswordBit
	- SetRegLock

*	Transfer-in Commands

	- TP_CreateOrder
	- TP_GetOrderStatuses
	- TP_GetOrderReview
	- TP_GetOrderDetail

*	Contact Commands

	- GetContacts
	- GetWhoisContact
	- Contacts

*	Nameserver Commands

	- RegisterNameserver
	- UpdateNameserver
	- CheckNSStatus

*	Report Commands

	- GetReport


# TODO

* Tests
* Map RRPCode / RRPText to some commands (CheckNSStatus, Extend, RegLock etc)
* Isolate common fields returned by Enom for each command (ExecTime, Done, Server etc)
* Return specific result fields to a identified section for each command (Enom response data is kinda entropia, no structured information)
* Parse Err\* fields (when command returns Err1, Err2 and so on)
* Parse TEXT response format
* More commands if needed


# Marcelo Pinheiro 2014

