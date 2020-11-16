package main

import (
	w "golang.org/x/sys/windows"
)

var dll_dir = `dll/`

var dll_list = []string{
	`gdbus.exe`,
	`libatk-1.0-0.dll`,
	`libbrotlicommon.dll`,
	`libbrotlidec.dll`,
	`libbz2-1.dll`,
	`libcairo-2.dll`,
	`libcairo-gobject-2.dll`,
	`libdatrie-1.dll`,
	`libepoxy-0.dll`,
	`libexpat-1.dll`,
	`libffi-7.dll`,
	`libfontconfig-1.dll`,
	`libfreetype-6.dll`,
	`libfribidi-0.dll`,
	`libgcc_s_seh-1.dll`,
	`libgdk-3-0.dll`,
	`libgdk_pixbuf-2.0-0.dll`,
	`libgio-2.0-0.dll`,
	`lib-2.0-0.dll`,
	`libgmodule-2.0-0.dll`,
	`libgobject-2.0-0.dll`,
	`libgraphite2.dll`,
	`libgtk-3-0.dll`,
	`libharfbuzz-0.dll`,
	`libiconv-2.dll`,
	`libintl-8.dll`,
	`libpango-1.0-0.dll`,
	`libpangocairo-1.0-0.dll`,
	`libpangoft2-1.0-0.dll`,
	`libpangowin32-1.0-0.dll`,
	`libpcre-1.dll`,
	`libpixman-1-0.dll`,
	`libpng16-16.dll`,
	`libstdc++-6.dll`,
	`libthai-0.dll`,
	`libwinpthread-1.dll`,
	`zlib1.dll`,
}

func init(){
	for _,v := range dll_list {
		w.NewLazyDLL(dll_dir + v).Load()
	}
}