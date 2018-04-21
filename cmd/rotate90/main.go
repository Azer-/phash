// transformation invariant image phash
// based on https://github.com/pippy360/transformationInvariantImageSearch
package main

import (
	"fmt"
	"os"

	"github.com/azr/phash/geometry"

	_ "net/http/pprof"

	"github.com/azr/phash/cmd"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println("Usage: dtc path/to/image.jpg")
		os.Exit(1)
	}

	img, _ := cmd.OpenImageFromPath(os.Args[1])
	img90 := geometry.InPlaceRotation90(img)
	cmd.WriteImageToPath(img90, os.Args[1]+".90")
}
