package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/cheggaaa/pb/v3"
	yaml "gopkg.in/yaml.v2"
)

type Data struct {
	Members []string `yaml:"members"`
}

func main() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	buf, err := ioutil.ReadFile("members.yml")
	if err != nil {
		panic(err)
	}

	var data Data
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("選ばれし者を探しています...")

	count := 2000
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

	fmt.Printf("選ばれしものは%vでした！！！", data.Members[r.Intn(len(data.Members))])
}
