package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vls",
	Short: "Pack file using variable-lenght code",
	Run:   pack,
}

const packedExtension = "vlc"

/*	Открываем документ, котоырй отправил пользователь
	Читаем зашифрованные данные и приводим их к строке
	Пререзаписываем файл.
			1) Меняем формат на .vlc
			2) Удаляем фрагмент суффикс из конца строки, а после возвращаем копию данных (string.TrimSuffix)
			3) Меняем права на чтение + запись

*/
func pack(_ *cobra.Command, args []string) {
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	//data -> Encode(data)
	packed := "" + string(data)

	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644) //0644 - пользователь может записывать и читать, остальные - читать
	if err != nil {
		handleErr(err)
	}
}

// Превращаем из myFile.txt -> myFile.vlc
func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
