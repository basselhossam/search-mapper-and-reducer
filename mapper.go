// Works on part of dataset and saves result in file as mapper.
 
package main

// Import necessary package to accomplish your work.
import (
	"os"
 	"bufio"
 	"strings"
	"strconv"
	"io/ioutil"
	"fmt"
	"log"
)

type Directories struct {
	Dirs      []string
}


func getFilesList(searchDir string) []string {
	// TODO: Reads the file contains a list of paths in path
	// /mnt/datanode/dataset/fileList.
	file, err := os.Open(searchDir)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data = []string{} 
	//i :=0
	for scanner.Scan() {
        line := scanner.Text()
        data = append(data,line)
        //i= i+1
    }
    return data
}

func searchInFile(path string, keyword string, searchMap map[string]int) {
	// TODO: open File and read it line by line and count
	// number of keyword and save result in map.
	out,err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("searchInFile error")
	    log.Fatal(err)

	}
	s := string(out)
	word := strings.ToLower(keyword)
	count := strings.Count(strings.ToLower(s),word)
	searchMap[path] = searchMap[path]+count

}

func main() {
	// TODO: get three arguments in format mapperID jobID keyword.
	argsWithoutProg := os.Args[1:]
	mapperID_str := argsWithoutProg[0]
	mapperID, err := strconv.Atoi(mapperID_str)
	if err != nil {
		fmt.Println("Main error")
	    log.Fatal(err)
	}
	jobID_str := argsWithoutProg[1]
	keyword := strings.ToLower(argsWithoutProg[2])
	
	// TODO: call getFilesList and save result in array
	// path for Wikipedia is at path /mnt/datanode/dataset.
	fileList := getFilesList("/mnt/datanode/dataset/fileList")
	
	// TODO: determine the start and the end of working data.
	end :=0
	size := int(len(fileList)/3)
	start := int((mapperID-1)*size)
	if(mapperID==3) {
		end=len(fileList)
	} else {
		end = start+size
	}
	
	// TODO: create a map with string key for path and int for
	// values.
	data := make(map[string]int)
	
	// TODO: loop on files and use call searchInFile.
	my_output := ""
	for i := start; i < end; i++ {
		searchInFile(fileList[i],keyword,data)
		my_output = my_output + fileList[i]+" "+strconv.Itoa(data[fileList[i]])+"\n"
	}
	
	// TODO: save results in any format in output file with name of
	// mapperID in this path pattern /mnt/datanode/tmp/jobID/mapperID.
	out := []byte(my_output)
	err2 := ioutil.WriteFile("/mnt/datanode/tmp/"+jobID_str+"/"+mapperID_str, out,0644)
	if err2 != nil {
    	panic(err)
	}

}

