package cmd

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Шифрование строки",
	Long:  "Шифрование строки своим или автоматически сгенерированным ключом шифрования.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key1, key2, key3 := getKeys(cmd)
		showKeys, _ := cmd.Flags().GetBool("sk")
		info, _ := cmd.Flags().GetBool("info")
		if showKeys {
			fmt.Printf("Сгенерированные ключи:\nkey1: %d\nkey2: %d\nkey3: %d\n", key1, key2, key3)
		}
		fmt.Printf("encode: %s\n", des3Encode(args[0], key1, key2, key3, info))
	},
}

func des3Encode(text string, key1, key2, key3 uint64, info bool) string {
	data := pad([]byte(text))

	data = desEncode(data, key1)
	data = desDecode(data, key2)
	data = desEncode(data, key3)

	if info {
		correlationInfo(pad([]byte(text)), data)
		distributionInfo(data)
	}

	return base64.StdEncoding.EncodeToString(data)
}

func correlationInfo(input, output []byte) {
	// количество битов
	n := len(input) * 8

	var sumX, sumY float64
	var numerator, denominator float64

	for i := 0; i < len(input); i++ {
		for b := 7; b >= 0; b-- {
			x := float64((input[i] >> b) & 1)
			y := float64((output[i] >> b) & 1)

			sumX += x
			sumY += y
		}
	}

	meanX := sumX / float64(n)
	meanY := sumY / float64(n)

	var x2, y2 float64
	for i := 0; i < len(input); i++ {
		for b := 7; b >= 0; b-- {
			x := float64((input[i] >> b) & 1)
			y := float64((output[i] >> b) & 1)

			numerator += (x - meanX) * (y - meanY)
			x2 += math.Pow(x-meanX, 2)
			y2 += math.Pow(y-meanY, 2)
		}
	}
	denominator = math.Sqrt(x2 * y2)

	if denominator == 0 {
		fmt.Println("Correlation undefined (division by zero)")
		return
	}

	r := numerator / denominator
	fmt.Printf("Коэффициент корреляции: %.6f\n", r)
}

func distributionInfo(output []byte) {
	zero := 0
	one := 0
	for _, b := range output {
		for i := 0; i < 8; i++ {
			if (b & (1 << uint(7-i))) != 0 {
				one++
			} else {
				zero++
			}
		}
	}
	fmt.Printf(
		"Вероятности распределения битов в выходных данных:\n0: %f\n1: %f\n",
		float64(zero)/float64(len(output)*8),
		float64(one)/float64(len(output)*8),
	)
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().Uint64("key1", rand.Uint64(), "Ключ шифрования 1")
	encodeCmd.Flags().Uint64("key2", rand.Uint64(), "Ключ шифрования 2")
	encodeCmd.Flags().Uint64("key3", rand.Uint64(), "Ключ шифрования 3")
	encodeCmd.Flags().Bool("sk", false, "Вывести сгенерированные ключи в консоль")
	encodeCmd.Flags().Bool("info", false, "Вывести информацию о корреляции между входными и выходными данными")
}
