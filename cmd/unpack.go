package cmd

import (
	"AlgoCompress/lib/compression"
	"AlgoCompress/lib/compression/vlc"
	"AlgoCompress/lib/compression/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file using",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrorEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New(shannon_fano.NewGenerator())
	default:
		cmd.PrintErr("unknown method")
	}

	filePath := args[0] //path to file
	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()
	data, err := io.ReadAll(r)

	if err != nil {
		handleErr(err)
	}
	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
