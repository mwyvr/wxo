package internal

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

func init() {
	// ensure the cache path exists
	path := GetCachePath()
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// GetCachePath returns the full path to the cache directory for the curent
// user. On *ux systems, this will likely be in in ~/.cache/wxo.
func GetCachePath() string {
	cp, err := os.UserCacheDir()
	if err != nil {
		// we can't go on
		panic(err)
	}
	return path.Join(cp, "wxo")
}

func makeCacheFilePath(uri string) string {
	return path.Join(GetCachePath(), hash(uri))
}

func GetData(uri string) ([]byte, error) {

	if uri == "" {
		panic(fmt.Errorf("uri is nil string, can't proceed"))
	}

	cachedPath := makeCacheFilePath(uri)

	if isCacheFileExpired(cachedPath) {
		// grab and write a new one
		var hc = &http.Client{
			Timeout: time.Second * 10,
		}
		resp, err := hc.Get(uri)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// We've got the body but it may be an error return
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("request error: %s", string(data))
		}
		// OK, we're good to go; write the new data to the cache
		fout, err := os.Create(cachedPath)
		if err != nil {
			return nil, err
		}
		fout.Write(data)
		err = fout.Close()
		if err != nil { // no point in going forward
			return nil, err
		}
	}
	// get the existing or newly created cache file
	fin, err := os.Open(cachedPath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(fin)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func isCacheFileExpired(filePath string) bool {
	fi, err := os.Stat(filePath)
	if errors.Is(err, fs.ErrPermission) {
		// we can't go on
		panic(err)
	}
	if errors.Is(err, fs.ErrNotExist) {
		return true
	}
	diff := time.Since(fi.ModTime())
	// TODO make cache expiry configurable.
	if diff.Seconds() >= (5 * 60) {
		// delete the existing file
		os.Remove(filePath)
		return true
	}
	return false
}

// Return a sha1 hash in hex; used to encode URIs for file caching
func hash(s string) string {

	h := sha1.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		// really no point in living
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
