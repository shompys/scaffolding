package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	scaffoldpkg "github.com/shompys/scaffolding"
)

type Config struct {
	ModuleName string
	TargetPath string
}

func main() {
	name := flag.String("name", "", "nombre del módulo Go (requerido)")
	pathFlag := flag.String("path", ".", "directorio donde crear el proyecto (default: directorio actual)")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "error: -name es requerido")
		flag.Usage()
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error obteniendo directorio actual: %v\n", err)
		os.Exit(1)
	}

	targetPath := *pathFlag
	if targetPath == "." {
		nameBase := filepath.Base(*name)
		if filepath.Base(cwd) == nameBase {
			targetPath = cwd
		} else {
			targetPath = filepath.Join(cwd, nameBase)
		}
	} else if !filepath.IsAbs(targetPath) {
		targetPath = filepath.Join(cwd, targetPath)
	}

	config := Config{
		ModuleName: *name,
		TargetPath: targetPath,
	}

	if err := scaffold(config); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Proyecto \"%s\" creado en %s\n", config.ModuleName, config.TargetPath)

	if err := copyToClipboard(config.TargetPath); err != nil {
		fmt.Fprintf(os.Stderr, "aviso: no se pudo copiar al portapapeles: %v\n", err)
	} else {
		fmt.Println("(path copiado al portapapeles)")
	}
}

func copyToClipboard(path string) error {
	var cmd *exec.Cmd

	if _, err := exec.LookPath("xclip"); err == nil {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	} else if _, err := exec.LookPath("wl-copy"); err == nil {
		cmd = exec.Command("wl-copy")
	} else {
		return fmt.Errorf("ni xclip ni wl-copy encontrados")
	}

	cmd.Stdin = strings.NewReader(path)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func scaffold(cfg Config) error {
	binName := filepath.Base(cfg.ModuleName)

	if err := os.MkdirAll(cfg.TargetPath, 0755); err != nil {
		return fmt.Errorf("creando directorio destino: %w", err)
	}

	err := fs.WalkDir(scaffoldpkg.Templates, "go", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "go" {
			return nil
		}

		relPath := strings.TrimPrefix(path, "go/")

		dest := filepath.Join(cfg.TargetPath, relPath)

		if d.IsDir() {
			return os.MkdirAll(dest, 0755)
		}

		content, err := scaffoldpkg.Templates.ReadFile(path)
		if err != nil {
			return fmt.Errorf("leyendo template %s: %w", path, err)
		}

		replaced := strings.ReplaceAll(string(content), "{{BINARY_NAME}}", binName)

		if err := os.WriteFile(dest, []byte(replaced), 0644); err != nil {
			return fmt.Errorf("escribiendo %s: %w", dest, err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", cfg.ModuleName)
	cmd.Dir = cfg.TargetPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ejecutando go mod init: %w", err)
	}

	return nil
}
