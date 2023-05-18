package my_function

import (
	"io"
	"log"
	"os"
)

// リソースのクリーンアップ?
func Defer() {
	if len(os.Args) < 2 {
		log.Fatal("ファイルが指定されていない")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data := make([]byte, 2048)
	for {
		// 読み込んだバイト数
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
