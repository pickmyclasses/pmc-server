package main

import migrate "pmc_server/migrate_scripts"

func main() {
	err := migrate.Emphasis()
	if err != nil {
		panic(err)
	}
}
