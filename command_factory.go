package enom

import (
	"strings"

	"enom-client/response"
)

type CommandFactory struct {
	Config *Config;
}

func (f CommandFactory) DomainCheck(domain string) *Command {
	parser := new(response.DomainCheck)
	command := f.CreateCommandWithDomain(domain, "Check", parser)

	return command
}

func (f CommandFactory) Purchase(domain string, nameservers []string) *Command {
	parser := new(response.Purchase)
	command := f.CreateCommandWithDomain(domain, "Purchase", parser)

	return command
}

func (f CommandFactory) CreateTransferOrder(domain string, authInfo string) *Command {
	parser := new(response.CreateTransferOrder)
	command := f.CreateCommand("TP_CreateOrder", parser)

	sld, tld := f.ParseDomain(domain)

	command.AddParam("SLD1", sld)
	command.AddParam("TLD1", tld)
	command.AddParam("AuthInfo1", authInfo)
	command.AddParam("DomainCount", "1")
	command.AddParam("OrderType", "Autoverification")

	return command
}

func (f CommandFactory) WhoisContact(domain string) *Command {
	parser := new(response.WhoisContact)
	command := f.CreateCommandWithDomain(domain, "GetWhoisContact", parser)

	return command
}

func (f CommandFactory) Renew(domain string, period string) *Command {
	parser := new(response.Extend)
	command := f.CreateCommandWithDomain(domain, "Extend", parser)

	command.AddParam("NumYears", period)

	return command
}

func (f CommandFactory) Delete(domain string) *Command {
	parser := new(response.Delete)
	command := f.CreateCommandWithDomain(domain, "DeleteRegistration", parser)

	return command
}

func (f CommandFactory) DomainInfo(domain string) *Command {
	parser := new(response.DomainInfo)
	command := f.CreateCommandWithDomain(domain, "GetDomainInfo", parser)

	return command
}

func (f CommandFactory) PasswordBit(domain string) *Command {
	parser := new(response.PasswordBit)
	command := f.CreateCommandWithDomain(domain, "GetPasswordBit", parser)

	return command
}

func (f CommandFactory) SetRegLock(domain string, enable bool) *Command {
	parser := new(response.RegLock)
	command := f.CreateCommandWithDomain(domain, "SetRegLock", parser)

	var value string
	if enable {
		value = "1"
	} else {
		value = "0"
	}

	command.AddParam("UnlockRegistrar", value)

	return command
}

func (f CommandFactory) Contacts(domain string) *Command {
	parser := new(response.Contacts)
	command := f.CreateCommandWithDomain(domain, "GetContacts", parser)

	return command
}

func (f CommandFactory) SetContact(domain string, contactData map[string]string) *Command {
	parser := new(response.SetContact)
	command := f.CreateCommandWithDomain(domain, "Contacts", parser)

	for key, val := range contactData {
		command.AddParam(key, val)
	}

	return command
}

func (f CommandFactory) RegisterNameserver(domain string, name string, ip string) *Command {
	parser := new(response.RegisterNameserver)
	command := f.CreateCommandWithDomain(domain, "RegisterNameserver", parser)

	command.AddParam("Add", "true")
	command.AddParam("NsName", name)
	command.AddParam("Ip", ip)

	return command
}

func (f CommandFactory) UpdateNameserver(domain string, name string, oldIp string, newIp string) *Command {
	parser := new(response.UpdateNameserver)
	command := f.CreateCommandWithDomain(domain, "UpdateNameserver", parser)

	command.AddParam("NS", name)
	command.AddParam("OldIp", oldIp)
	command.AddParam("NewIp", newIp)

	return command
}

func (f CommandFactory) CheckNSStatus(domain string, host string) *Command {
	parser := new(response.CheckNSStatus)
	command := f.CreateCommandWithDomain(domain, "CheckNSStatus", parser)
	command.AddParam("CheckNSName", host)

	return command
}

func (f CommandFactory) Report() *Command {
	parser := new(response.Report)
	command := f.CreateCommand("GetReport", parser)

	return command
}

func (f CommandFactory) TransferOrderStatuses() *Command {
	parser := new(response.TransferOrderStatuses)
	command := f.CreateCommand("TP_GetOrderStatuses", parser)

	return command
}

func (f CommandFactory) TransferOrderReview(id string) *Command {
	parser := new(response.TransferOrderReview)
	command := f.CreateCommand("TP_GetOrderReview", parser)

	command.AddParam("TransferOrderID", id)

	return command
}

func (f CommandFactory) TransferOrderDetail(id string) *Command {
	parser := new(response.TransferOrderDetail)
	command := f.CreateCommand("TP_GetOrderDetail", parser)

	command.AddParam("TransferOrderDetailID", id)

	return command
}

func (f CommandFactory) CreateCommandWithDomain(domain string, commandName string, parser response.ResponseParser) *Command {
	sld, tld := f.ParseDomain(domain)

	command := f.CreateCommand(commandName, parser)
	command.AddParam("SLD", sld)
	command.AddParam("TLD", tld)

	return command
}

func (f CommandFactory) CreateCommand(commandName string, parser response.ResponseParser) *Command {
	command := &Command{ Name: commandName, Parser: parser }
	command.SetDefaultArgs(f.Config)

	return command
}

func (f CommandFactory) ParseDomain(domain string) (string, string) {
	fragments := strings.Split(domain, ".")
	slice := fragments[0 : 2]

	return slice[0], slice[1]
}

