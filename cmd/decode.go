package cmd

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use: "decode",
	Run: func(cmd *cobra.Command, args []string) {
		key1, key2, key3 := getKeys(cmd)
		showKeys, _ := cmd.Flags().GetBool("sk")
		if showKeys {
			fmt.Printf("Сгенерированные ключи:\nkey1: %d\nkey2: %d\nkey3: %d\n", key1, key2, key3)
		}
		fmt.Printf("encode: %s\n", des3Decode(args[0], key1, key2, key3))
	},
}

func des3Decode(text string, key1, key2, key3 uint64) string {
	data, _ := base64.StdEncoding.DecodeString(text)

	data = desDecode(data, key3)
	data = desEncode(data, key2)
	data = desDecode(data, key1)

	return string(unpad(data))
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().Uint64("key1", rand.Uint64(), "Ключ шифрования 1")
	decodeCmd.Flags().Uint64("key2", rand.Uint64(), "Ключ шифрования 2")
	decodeCmd.Flags().Uint64("key3", rand.Uint64(), "Ключ шифрования 3")
	decodeCmd.Flags().Bool("sk", false, "Вывести сгенерированные ключи в консоль")
}
