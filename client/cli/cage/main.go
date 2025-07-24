package main

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

//||------------------------------------------------------------------------------------------------||
//|| Main
//||------------------------------------------------------------------------------------------------||

func main() {

	//||------------------------------------------------------------------------------------------------||
	//|| Main / Help
	//||------------------------------------------------------------------------------------------------||

	root := &cobra.Command{
		Use:   "cage",
		Short: "cage: centralized CLI for deployment builds",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build the Loader Script
	//||------------------------------------------------------------------------------------------------||

	root.AddCommand(&cobra.Command{
		Use:   "build-ts",
		Short: "Compile TypeScript to browser JS via esbuild",
		Run: func(cmd *cobra.Command, args []string) {
			runExec("go", "run", "./cli/cachegate/generate.go")
		},
	})

	//||------------------------------------------------------------------------------------------------||
	//|| Translate I18n Files
	//||------------------------------------------------------------------------------------------------||

	root.AddCommand(&cobra.Command{
		Use:   "translate",
		Short: "Generate translated JSON files from i18n/en.json",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot get working dir:", err)
				os.Exit(1)
			}
			translatePath := filepath.Join(wd, "cli", "translate", "translate.go")
			runExec("go", "run", translatePath)
		},
	})

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Age-Gate Cache
	//||------------------------------------------------------------------------------------------------||

	root.AddCommand(&cobra.Command{
		Use:   "build-gate",
		Short: "Generate age-gate HTML files for zones/lang",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot get working dir:", err)
				os.Exit(1)
			}
			genPath := filepath.Join(wd, "cli", "cachegate", "generate.go")
			runExec("go", "run", genPath)
		},
	})

	//||------------------------------------------------------------------------------------------------||
	//|| Loader
	//||------------------------------------------------------------------------------------------------||

	root.AddCommand(&cobra.Command{
		Use:   "build-loader",
		Short: "Bundle & minify client/loader TS into cache/loader.js",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot get working dir:", err)
				os.Exit(1)
			}

			loaderEntry := filepath.Join(wd, "loader", "src", "index.ts")
			outFile := filepath.Join(wd, "cache", "loader.js")

			if err := os.MkdirAll(filepath.Dir(outFile), 0755); err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot create cache dir:", err)
				os.Exit(1)
			}

			// invoke esbuild via npx with full minification
			runExec("npx", "esbuild",
				loaderEntry,
				"--bundle",
				"--outfile="+outFile,
				"--platform=browser",
				"--format=iife",
				"--minify", // minify syntax, whitespace & identifiers
				"--define:process.env.NODE_ENV=\"production\"", // optional: drop dev-only code
				//"--sourcemap=external", // if you want an external map
			)

			fmt.Println("✅ Created minified", outFile)
		},
	})

	//||------------------------------------------------------------------------------------------------||
	//|| Loader
	//||------------------------------------------------------------------------------------------------||

	root.AddCommand(&cobra.Command{
		Use:   "dev-loader",
		Short: "Bundle client/loader TS into cache/loader.js (no minify, dev mode)",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot get working dir:", err)
				os.Exit(1)
			}

			loaderEntry := filepath.Join(wd, "loader", "src", "index.ts")
			outFile := filepath.Join(wd, "cache", "loader.dev.js")

			if err := os.MkdirAll(filepath.Dir(outFile), 0755); err != nil {
				fmt.Fprintln(os.Stderr, "❌ cannot create cache dir:", err)
				os.Exit(1)
			}

			// invoke esbuild via npx without minification
			runExec("npx", "esbuild",
				loaderEntry,
				"--bundle",
				"--outfile="+outFile,
				"--platform=browser",
				"--format=iife",
				"--sourcemap=inline", // inline map for dev
			)

			fmt.Println("✅ Created dev bundle (un‑minified) at", outFile)
		},
	})

	//||------------------------------------------------------------------------------------------------||
	//|| 404
	//||------------------------------------------------------------------------------------------------||

	if err := root.Execute(); err != nil {
		fmt.Println("[x] Command not recognized.", err)
		os.Exit(1)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Run Helper
//||------------------------------------------------------------------------------------------------||

func runExec(name string, args ...string) {
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running %s %v: %v\n", name, args, err)
		os.Exit(1)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| runExecDir runs a command in the specified working directory.
//||------------------------------------------------------------------------------------------------||

func runExecDir(dir, binary string, args ...string) {
	cmd := exec.Command(binary, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running %s %v in %s: %v\n", binary, args, dir, err)
		os.Exit(1)
	}
}
