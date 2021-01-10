package domain

var SourceDir string
var TargetDir string
const VersionStr = "0.1.1"
const AppName = "PDFnmbrr - pre-alpha"


// SetupInitialState initializes directories and other relevant settings
// todo: read user-specific configuration from persistent storage
func SetupInitialState() {
	SourceDir = GetUserHomeDirectory()
	TargetDir = GetUserHomeDirectory() + "_target"


}
