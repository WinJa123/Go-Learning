package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите путь к папке, которую нужно скопировать: ")
	srcInput, _ := reader.ReadString('\n')
	src := strings.TrimSpace(srcInput)

	// Проверяем, существует ли введённая папка
	info, err := os.Stat(src)
	if os.IsNotExist(err) {
		fmt.Printf("Ошибка: папка '%s' не найдена.\n", src)
		return
	}
	if !info.IsDir() {
		fmt.Printf("Ошибка: '%s' не является папкой.\n", src)
		return
	}

	// Путь к .exe на флешке
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка: не удалось определить путь к программе.")
		return
	}
	exeDir := filepath.Dir(exePath)
	dst := filepath.Join(exeDir, "InDataLoad")

	// Копирование
	err = copyDir(src, dst)
	if err != nil {
		fmt.Printf("Ошибка при копировании: %v\n", err)
		return
	}
	fmt.Printf("Папка '%s' скопирована в '%s'\n", src, dst)
}

// функции copyDir и copyFile остаются без изменений (как в предыдущем ответе)

// copyDir и copyFile такие же, как в предыдущем ответе
func copyDir(src, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return fmt.Errorf("не удалось создать папку назначения: %w", err)
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("не удалось прочитать исходную папку: %w", err)
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return fmt.Errorf("ошибка копирования %s: %w", srcPath, err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	info, err := os.Stat(src)
	if err == nil {
		os.Chmod(dst, info.Mode())
	}
	return nil
}
