package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Ntrakiyski/screenshot-cli/internal/screenshot"
	"github.com/spf13/cobra"
)

var (
	urlFlag    string
	widthFlag  int
	heightFlag int
	keyFlag    string
	uploadFlag bool
	verboseFlag bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "screenshot",
		Short: "Screenshot URL to ImgBB URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			if verboseFlag {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
			}
			if urlFlag == "" {
				return fmt.Errorf("--url required")
			}
			if widthFlag < 100 || widthFlag > 3840 || heightFlag < 100 || heightFlag > 3840 {
				return fmt.Errorf("width/height 100-3840")
			}
			buf, err := screenshot.Capture(urlFlag, widthFlag, heightFlag)
			if err != nil {
				return err
			}
			fmt.Printf("Captured %d bytes (%dx%d)\\n", len(buf), widthFlag, heightFlag)

			if uploadFlag {
				if keyFlag == "" {
					keyFlag = os.Getenv("IMGBB_API_KEY")
					if keyFlag == "" {
						return fmt.Errorf("IMGBB_API_KEY or --key required")
					}
				}
				b64 := base64.StdEncoding.EncodeToString(buf)
				data := url.Values{}
				data.Set("image", b64)
				data.Set("expiration", "600")
				resp, err := http.PostForm("https://api.imgbb.com/1/upload?key="+keyFlag, data)
				if err != nil {
					return fmt.Errorf("upload: %w", err)
				}
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				var result struct {
					Success bool `json:"success"`
					Data    struct {
						UrlViewer string `json:"url_viewer"`
					} `json:"data"`
				}
				if err := json.Unmarshal(body, &result); err != nil {
					return fmt.Errorf("json: %w", err)
				}
				if !result.Success {
					return fmt.Errorf("upload failed: %s", string(body))
				}
				fmt.Printf("✅ ImgBB URL: %s\\n", result.Data.UrlViewer)
				return nil
			}
			return fmt.Errorf("use --upload to get URL")
		},
	}

	rootCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "URL (required)")
	rootCmd.Flags().IntVarP(&widthFlag, "width", "w", 1920, "Width")
	rootCmd.Flags().IntVar(&heightFlag, "height", 1080, "Height")
	rootCmd.Flags().StringVarP(&keyFlag, "key", "k", "", "ImgBB key")
	rootCmd.Flags().BoolVarP(&uploadFlag, "upload", "U", true, "Upload to ImgBB")
	rootCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Verbose")

	rootCmd.Execute()
}
