package main

import "grabber/handler"

const URLSITE = "http://vpustotu.ru/moderation/"

func main() {
	handler.ParseSiteHandler(URLSITE)
}
