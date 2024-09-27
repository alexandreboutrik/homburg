package initializers

// ---------------
// Configuration
// ---------------

var TESTING_EP bool = true // enable the testing endpoints, for dev purposes only
var LOGGING_EP bool = true // enable the logging endpoints, required for legal purposes

var REQUIRE_PGA bool = false            // require PGA beforing showing anything
var FILTER_KEYWORDS bool = true         // enable the 'filter keywords' system, see keywords.go
var ALLOW_ANONYMOUS_REPORT bool = false // allow anonymous reporting (without PGA)

// ---------------
// Email
// ---------------

var SenderEmail string = ""
var SenderPassword string = ""

// ---------------
// Database & Authentication
// ---------------

// WARNING: DO NOT EDIT
var HAT_DB_PWD string = "3e23d61dde"
