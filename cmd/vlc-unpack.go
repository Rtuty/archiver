package cmd

import (
	"archiver/cmd/lib/vlc"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	//data -> Encode(data)
	packed := vlc.Decode(data)
	//packed := vlc.Encode(string(data))

	fmt.Println(string(data)) // TODO: remove

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644) //0644 - пользователь может записывать и читать, остальные - читать
	if err != nil {
		handleErr(err)
	}
}

// Превращаем из myFile.txt -> myFile.vlc
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcPackCmd)
}
