package database

var (
	indexDir        = "index"
	idDir           = "ID"
	dbEndpoint      = "db.json"
	modulesEndpoint = "modules.json"
	vulnsEndpoint   = "vulns.json"
)

func IsIndexEndpoint(filename string) bool {
	return filename == dbEndpoint ||
		filename == modulesEndpoint ||
		filename == vulnsEndpoint
}
