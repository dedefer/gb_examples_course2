package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"
)

var (
	dataDirPath  = flag.String("data_dir", "lesson1/panic_real_example/data", "data_dir")
	flockDirPath = flag.String("flock_dir", "lesson1/panic_real_example/flock", "flock_dir")

	errFileLocked = errors.New("file locked")
)

func main() {
	files, err := ioutil.ReadDir(*dataDirPath)
	if err != nil {
		fmt.Printf("can't read dir %s: %s\n", *dataDirPath, err)
		return
	}

	for _, file := range files {
		if file.Mode().IsRegular() && path.Ext(file.Name()) == ".db" {
			fmt.Printf("processing file %s\n", file.Name())
			if err := process(file); err != nil {
				if errors.Is(err, errFileLocked) {
					fmt.Printf("file (%s) locked\n", file.Name())
				} else {
					fmt.Printf("error while file (%s) processing: %s\n", file.Name(), err)
				}
			}
		}
	}
}

func process(file os.FileInfo) (err error) {
	locked, err := lockFile(file.Name())
	if err != nil { // ошибка при взятии лока
		return err
	}
	if !locked { // кто-то уже обрабатывает файл
		return errFileLocked
	}

	// разблокируем файл в конце
	defer func() {
		// ловим панику и присваиваем ошибку если была паника
		if v := recover(); v != nil {
			buff := make([]byte, 1024)
			runtime.Stack(buff, false)
			err = fmt.Errorf("panic: %w\n%s", v, buff)
		}

		if err != nil { // разблокируем файл при ошибке
			if unlockErr := unlockFile(file.Name()); unlockErr != nil {
				fmt.Printf("can't unlock file (%s): %v\n", file.Name(), unlockErr)
			}
			return
		}

		if doneErr := doneFile(file.Name()); doneErr != nil { // помечаем файл обработанным если ошибок не было
			fmt.Printf("can't mark file (%s) as done: %v\n", file.Name(), doneErr)
		}
		if unlockErr := unlockFile(file.Name()); unlockErr != nil {
			fmt.Printf("can't unlock file (%s): %v\n", file.Name(), unlockErr)
		}
	}()

	return internalProcessing(file)
}

func internalProcessing(file os.FileInfo) error {
	// Внутреннюю логику пишет большая команда
	// люди могут и будут ошибаться поэтому
	// этот код может вызывать панику

	if rand.Intn(100) < 30 {
		panic(errors.New("don't hug me i'm scared"))
	}
	return nil
}

func lockfileName(name string) string {
	basename := path.Base(name)
	return path.Join(*flockDirPath, basename+".lock")
}
func donefileName(name string) string {
	basename := path.Base(name)
	return path.Join(*flockDirPath, basename+".done")
}

func lockFile(name string) (ok bool, err error) {
	_, err = os.Stat(donefileName(name))
	if err == nil { // файл уже обработан
		return false, nil
	}

	if err != nil && !os.IsNotExist(err) { // какая-то ошибка
		return false, err
	}

	// проверяем лок файл
	_, err = os.Stat(lockfileName(name))
	if err == nil { // лок файл уже существует
		return false, nil
	}

	if err != nil && !os.IsNotExist(err) { // какая-то ошибка
		return false, err
	}

	if _, err = os.Create(lockfileName(name)); err != nil { // какая-то ошибка при создании лок файла
		return false, err
	}

	return true, nil
}

func unlockFile(name string) error { // Удаляем лок файл
	return os.Remove(lockfileName(name))
}

func doneFile(name string) error { // Создаем файл-маркер, что мы закончили
	_, err := os.Create(donefileName(name))
	return err
}
