package cmd

import (
	"AlgoCompress/lib/compression"
	"AlgoCompress/lib/compression/vlc"
	"AlgoCompress/lib/compression/vlc/table/shannon_fano"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file using",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrorEmptyPath = errors.New("path to new file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrorEmptyPath)
	}
	filePath := args[0] //path to file

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New(shannon_fano.NewGenerator())
	default:
		cmd.PrintErr("unknown method")
	}

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()
	data, err := io.ReadAll(r)

	if err != nil {
		handleErr(err)
	}
	packed := encoder.Encode(string(data))
	// data -> Encode(data)

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}

}
