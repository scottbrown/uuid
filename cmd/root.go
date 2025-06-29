package cmd

import (
	"fmt"
	"os"

	"github.com/scottbrown/uuid/internal/generator"
	"github.com/spf13/cobra"
)

var (
	version = "dev"     // Default version, can be overridden by build flags
	build   = "unknown" // Default build, can be overridden by build flags
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUIDs from the command line",
	Long: `A simple CLI tool for generating UUIDs.
	
By default, generates UUIDv4. Use version flags to generate other UUID versions.
Use the timestamp flag (-t) to generate UUIDv7 from a specific timestamp.

SECURITY NOTE: UUIDv7 contains embedded timestamps that reveal timing information.
Use UUIDv4 when privacy is important.

Examples:
  uuid                        # Generate UUIDv4 (default)
  uuid -4                     # Generate UUIDv4 (explicit)
  uuid -6                     # Generate UUIDv6
  uuid -7                     # Generate UUIDv7 (contains timestamp)
  uuid -t 1234567890          # Generate UUIDv7 from Unix timestamp
  uuid -t 2023-06-14          # Generate UUIDv7 from date
  uuid -t "2023-06-14 10:30"  # Generate UUIDv7 from date-time`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check which version flag was used
		v4, _ := cmd.Flags().GetBool("4")
		v6, _ := cmd.Flags().GetBool("6")
		v7, _ := cmd.Flags().GetBool("7")
		timestamp, _ := cmd.Flags().GetString("timestamp")

		// Handle timestamp flag
		if timestamp != "" {
			// Validate that timestamp is only used with UUIDv7 (or no version specified)
			if v4 || v6 {
				fmt.Fprintf(os.Stderr, "Error: Timestamp flag (-t) is only supported with UUIDv7. Use 'uuid -t %s' or 'uuid -7 -t %s'.\n", timestamp, timestamp)
				os.Exit(1)
			}

			// Parse the timestamp
			parsedTime, err := generator.ParseTimestamp(timestamp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			// Generate UUIDv7 with the specified timestamp
			fmt.Println(generator.GenerateUUIDv7WithTimestamp(parsedTime))
			return
		}

		// Default to UUIDv4 if no version flag is specified
		if !v4 && !v6 && !v7 {
			v4 = true
		}

		// Generate and output the appropriate UUID
		if v7 {
			fmt.Println(generator.GenerateUUIDv7())
		} else if v6 {
			fmt.Println(generator.GenerateUUIDv6())
		} else if v4 {
			fmt.Println(generator.GenerateUUIDv4())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Version-specific flags
	rootCmd.Flags().BoolP("4", "4", false, "Generate UUIDv4 (default)")
	rootCmd.Flags().BoolP("6", "6", false, "Generate UUIDv6")
	rootCmd.Flags().BoolP("7", "7", false, "Generate UUIDv7 (contains timestamp)")

	// Timestamp flag for UUIDv7
	rootCmd.Flags().StringP("timestamp", "t", "", "Generate UUIDv7 from timestamp (Unix seconds/milliseconds, RFC3339, or ISO date)")

	// Make version flags mutually exclusive
	rootCmd.MarkFlagsMutuallyExclusive("4", "6", "7")

	// Set version for --version flag (combine version and build)
	if build != "unknown" && build != "" {
		rootCmd.Version = fmt.Sprintf("%s+%s", version, build)
	} else {
		rootCmd.Version = version
	}
}
