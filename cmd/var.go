package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// varCmd represents the var command
var varCmd = &cobra.Command{
	Use:   "var",
	Short: "Информация о варианте",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Усольцев Константин. Группа 4228. Вариант 9. Реализовать алгоритм шифрования TripleDES(3DES). ")
	},
}

func init() {
	rootCmd.AddCommand(varCmd)
}
