// Copyright © 2017 Bartosz Wieczorek
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"strconv"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "resize an image",
	Long:  `as arguments use wanted 1:path, 2:width 3:quality`,
	Run: func(cmd *cobra.Command, args []string) {
		// open "test.jpg"
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}

		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		widthU64, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		widthUint := uint(widthU64)

		quality, err := strconv.Atoi(args[2])

		resizeImg := resize.Resize(widthUint, 0, img, resize.Lanczos3)

		newImg := args[1] + "q" + args[2] + args[0]

		newFile, err := os.Create(newImg)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		// write new image into file
		jpeg.Encode(newFile, resizeImg, &jpeg.Options{quality})

		fmt.Println("resize " + args[0] + " called")
		if _, err := os.Stat(newImg); os.IsNotExist(err) {
			log.Fatal("failure in creating " + newImg)
		}
		if _, err := os.Stat(newImg); err == nil {
			fmt.Println(newImg + " created")
		}
	},
}

func init() {
	RootCmd.AddCommand(resizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
