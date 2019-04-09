package common

import "os"

const EnvDev = "dev"
const EnvTest = "test"
const EnvProduct = "product"

const LocationSaas = "saas"
const LocationLocal = "local"

func GetRunEnv() (env string) {
	env = os.Getenv("GIN_RUN_ENV")
	if env != "" {
		return env
	} else {
		return EnvDev
	}
}

func GetRunLocation() (location string) {
	location = os.Getenv("GIN_RUN_LOCATION")
	if location != "" {
		return location
	} else {
		return LocationLocal
	}
}

func IsCronENV() bool {
	if GetRunEnv() == EnvProduct {
		cron := os.Getenv("GIN_RUN_CRON")
		if cron == "true" {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func IsPprofEnv() bool {
	if GetRunEnv() == EnvDev {
		return true
	}
	if os.Getenv("GIN_RUN_PPROF") == "true" {
		return true
	}
	return false
}
