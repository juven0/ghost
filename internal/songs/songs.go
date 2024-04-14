package songs

import (
	"io/fs"
	"path/filepath"
)

func GetSongList(path string, ch chan []Song) {
	var allSongs []Song
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// if !info.IsDir() || filepath.Ext(path) == ".mp4" {
		// 	allSongs = append(allSongs, Song{info, path})
		// }
		if !info.IsDir() {
			allSongs = append(allSongs, Song{info, path})
		}
		return nil
	})
	ch <- allSongs
	return
}
