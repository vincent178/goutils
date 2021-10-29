package cursor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type FileCursor struct {
	Name string
	f    *os.File

	mu   sync.Mutex
	data map[string]int
}

func NewFileCursor(name string) (*FileCursor, error) {
	data := make(map[string]int)

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		if len(split) != 2 {
			log.Fatal("invalid data format")
		}
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])
		val, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("invalid data format", err)
		}
		data[key] = val
	}

	return &FileCursor{
		f:    f,
		data: data,
	}, nil
}

func (c *FileCursor) Get(cursorName string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[cursorName]
	if !ok {
		return -1
	}
	return val
}

func (c *FileCursor) Save() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.data {
		_, err := c.f.WriteString(fmt.Sprintf("%s,%d", k, v))
		if err != nil {
			return err
		}
	}

	err := c.f.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (c *FileCursor) Update(cursorName string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[cursorName] = val
}

func (c *FileCursor) Close()  error {
	return c.f.Close()
}