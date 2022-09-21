package cmd

import (
	"archiver/cmd/lib/vlc"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-lenght code",
	Run:   pack,
}

/*
	Проверяем на ошибку: Если путь до файла равен нулю, либо пуст -> error
	Открываем документ, котоырй отправил пользователь
	Читаем зашифрованные данные и приводим их к строке
	Пререзаписываем файл.
			1) Меняем формат на .vlc
			2) Удаляем фрагмент суффикс из конца строки, а после возвращаем копию данных (string.TrimSuffix)
			3) Меняем права на чтение + запись
*/

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not corrected. Path is empty")

func pack(_ *cobra.Command, args []string) {
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
	packed := vlc.Encode(string(data))

	fmt.Println(string(data)) // TODO: remove

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644) //0644 - пользователь может записывать и читать, остальные - читать
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
