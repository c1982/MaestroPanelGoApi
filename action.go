package maestropanel

type action struct{
	URL string
	Method string
}

var getListAction = action{"/domain/getlist", "GET"}
var createDomainAction = action{"/domain/create", "POST"}
var deleteDomainAction = action{"/domain/delete", "DELETE"}
var stopDomainAction = action{"/domain/stop", "POST"}
var passwordChangeDomainAction = action{"/domain/password", "POST"}
var addDominAliasAction = action{"/domain/adddomainalias", "POST"}
var deleteDominAliasAction = action{"/domain/deletedomainalias", "DELETE"}
var getDomainAliasesAction = action{"/domain/getdomainaliases", "GET"}
var addSubdomainAction = action{"/domain/addsubdomain", "POST"}
var deleteSubdomainAction = action{"/domain/deletesubdomain", "DELETE"}
var subDomainsAction = action{"/domain/getsubdomains","GET"}
var setSubDomainFTPAccountAction = action{"/domain/setsubdomainftpaccount","POST"}