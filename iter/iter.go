/*
 *          Copyright 2022, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package iter runs over files to process them.
package iter

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// FileProcessor is called when iterating over files to process a file.
type FileProcessor interface {
	ProcessFile(path string, info os.FileInfo, err error) error
}

type tFileProcessorDummy struct {
}

type tFileProcessorGo struct {
	wg        sync.WaitGroup
	proc      FileProcessor
	mutex     sync.Mutex
	errResult error
}

func (p *tFileProcessorDummy) ProcessFile(path string, info os.FileInfo, err error) error {
	if err == nil || os.IsNotExist(err) {
		fmt.Println(path)
		return nil
	}
	return err
}

func (p *tFileProcessorGo) Init(proc FileProcessor) {
	p.proc = proc
}

func (p *tFileProcessorGo) ProcessFile(path string, info os.FileInfo, err error) {
	err = p.proc.ProcessFile(path, info, err)
	if err != nil && p.errResult == nil {
		p.mutex.Lock()
		p.errResult = err
		p.mutex.Unlock()
	}
	p.wg.Done()
}

func (p *tFileProcessorGo) err() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.errResult
}

// IterateFilesFlat runs over files in directory dir calling proc.ProcessFile.
func IterateFilesFlat(dir string, proc FileProcessor) error {
	if proc == nil {
		proc = new(tFileProcessorDummy)
	}
	lengthDir := pathLengthAsPrefix(dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && lengthDir == len(path)-len(info.Name()) {
			return proc.ProcessFile(path, info, err)
		}
		return nil
	})
	return err
}

// IterateFilesRecr runs over files in directory dir calling proc.ProcessFile.
// Files in subdirectories are traversed, too.
func IterateFilesRecr(dir string, proc FileProcessor) error {
	if len(dir) > 0 {
		if proc == nil {
			proc = new(tFileProcessorDummy)
		}
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info != nil && !info.IsDir() {
				return proc.ProcessFile(path, info, err)
			}
			return nil
		})
		return err
	}
	return nil
}

// IterateFilesFlatGo runs over files in directory dir calling proc.ProcessFile.
// Call to proc.ProcessFile happens in a separate Goroutine.
func IterateFilesFlatGo(dir string, proc FileProcessor) error {
	var procGo tFileProcessorGo
	if proc == nil {
		proc = new(tFileProcessorDummy)
	}
	lengthDir := pathLengthAsPrefix(dir)
	procGo.Init(proc)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && lengthDir == len(path)-len(info.Name()) {
			procGo.wg.Add(1)
			go procGo.ProcessFile(path, info, err)
			return procGo.err()
		}
		return nil
	})
	procGo.wg.Wait()
	if err == nil {
		return procGo.err()
	}
	return err
}

// IterateFilesRecrGo runs over files in directory dir calling proc.ProcessFile.
// Files in subdirectories are traversed, too. Call to proc.ProcessFile happens
// in a separate Goroutine.
func IterateFilesRecrGo(dir string, proc FileProcessor) error {
	var procGo tFileProcessorGo
	if proc == nil {
		proc = new(tFileProcessorDummy)
	}
	procGo.Init(proc)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			procGo.wg.Add(1)
			go procGo.ProcessFile(path, info, err)
			return procGo.err()
		}
		return nil
	})
	procGo.wg.Wait()
	if err == nil {
		return procGo.err()
	}
	return err
}

func pathLengthAsPrefix(path string) int {
	pathTrimmed := filepath.Join(path, ".")
	if pathTrimmed != "." {
		return len(pathTrimmed) + 1
	}
	return 0
}
