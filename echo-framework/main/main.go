/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"github.com/ragilmaulana/bootcamp/tugas-golang/echo-framework/config"
	"github.com/ragilmaulana/bootcamp/tugas-golang/echo-framework/router"
	"github.com/ragilmaulana/bootcamp/tugas-golang/echo-framework/service"
)

func main() {
	// Koneksi Database
	db, err := config.KoneksiDB()
	if err != nil {
		panic(err)
		// return &bank.User{}, err
	}

	// Membuat struct koneksi
	con := service.UserService{
		db,
	}



	router.GetRouter()

}
