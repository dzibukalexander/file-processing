package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/core"
	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	"github.com/dzibukalexander/file-processing/internal/encryption/rsa"
	"github.com/dzibukalexander/file-processing/internal/logger"
)

func main() {
	if err := config.LoadConfig("config.json"); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
	logger.SetupLogger()
	log := logger.GetInstance()

	appCore := core.NewCore()
	log.Info("Application started")
	fmt.Println("File Processing CLI. Type 'exit' to quit.")
	fmt.Println("Commands: load, apply, process, save-pipeline, load-pipeline, gen-key, exit")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.ToLower(line) == "exit" {
			break
		}

		log.Infof("Executing command: %s", line)
		if err := handleCommand(appCore, line); err != nil {
			log.Errorf("Command failed: %v", err)
			fmt.Printf("Error: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Errorf("Error reading input: %v", err)
		fmt.Printf("Error reading input: %v\n", err)
	}

	log.Info("Application shutting down")
	fmt.Println("Exiting.")
}

func handleCommand(appCore *core.Core, line string) error {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}

	command := strings.ToLower(parts[0])
	args := parts[1:]

	switch command {
	case "help":
		printHelp()
		return nil
	case "load":
		if len(args) != 1 {
			return fmt.Errorf("load command requires a file path")
		}
		return appCore.Load(args[0])
	case "process":
		if len(args) != 1 {
			return fmt.Errorf("process command requires an output file path")
		}
		return appCore.ProcessFile(args[0])
	case "save-pipeline":
		if len(args) != 1 {
			return fmt.Errorf("save-pipeline command requires a file path")
		}
		return appCore.SavePipeline(args[0])
	case "load-pipeline":
		if len(args) != 1 {
			return fmt.Errorf("load-pipeline command requires a file path")
		}
		return appCore.LoadPipeline(args[0])
	case "apply":
		if len(args) < 1 {
			return fmt.Errorf("apply command requires an operation type")
		}
		op := strings.ToLower(args[0])
		params := make(map[string]string)
		for _, arg := range args[1:] {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid parameter format: %s", arg)
			}
			params[parts[0]] = parts[1]
		}
		return appCore.Apply(op, params)
	case "gen-key":
		if len(args) != 2 {
			return fmt.Errorf("gen-key command requires algorithm and path")
		}
		alg := strings.ToLower(args[0])
		path := args[1]

		switch alg {
		case "aes":
			generator := aes.AESEncryptor{}
			return generator.GenerateKey(path)
		case "rsa":
			generator := rsa.RSAEncryptor{}
			return generator.GenerateKey(path)
		default:
			return fmt.Errorf("unsupported algorithm for key generation: %s", alg)
		}
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  load <file_path>              - Load a file to process.")
	fmt.Println("  apply <operation> [params...] - Add a processing step to the pipeline.")
	fmt.Println("    compress type=<zip|gzip>")
	fmt.Println("    decompress type=<zip|gzip>")
	fmt.Println("    encrypt type=<aes|rsa> key_file=<path>")
	fmt.Println("    decrypt type=<aes|rsa> key_file=<path>")
	fmt.Println("    calculate type=<library|parser|regex>")
	fmt.Println("  process <output_path>           - Run the pipeline and save the result.")
	fmt.Println("  save-pipeline <file_path>     - Save the current pipeline to a file.")
	fmt.Println("  load-pipeline <file_path>     - Load a pipeline from a file.")
	fmt.Println("  gen-key <aes|rsa> <path>      - Generate a new encryption key.")
	fmt.Println("  help                            - Show this help message.")
	fmt.Println("  exit                            - Exit the application.")
}
