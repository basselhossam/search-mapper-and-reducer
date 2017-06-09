// Works on part of dataset and saves result in file as a reducer.

package main

// Import necessary package to accomplish your work.
import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"sort"
	"io/ioutil"
	"strings"
	"log"
)

type SortWikiResult struct {
	searchWiki map[string]int
	Keys       []string
}

func (sw *SortWikiResult) Len() int {
	// TODO: Implement Len function.
	return len(sw.searchWiki)
}

func (sw *SortWikiResult) Less(i, j int) bool {
	// TODO: Implement Less function.
	return sw.searchWiki[sw.Keys[i]] > sw.searchWiki[sw.Keys[j]]
}

func (sw *SortWikiResult) Swap(i, j int) {
	// TODO: Implement Swap function.
	sw.Keys[i],sw.Keys[j] = sw.Keys[j],sw.Keys[i]
}

func sortArticles(searchWiki map[string]int) []string {
	// TODO: Implement sortKeys function.
	var d SortWikiResult
	d.searchWiki = searchWiki
	for k := range searchWiki {
	    d.Keys = append(d.Keys, k)
	}
	sort.Sort(&d)
	return d.Keys
}

func getMappersFiles(jobID int) []string {
	// TODO: get a list of path of files were produced
	// by mappers in path /mnt/datanode/tmp/jobID.
	paths := []string{}
	for i := 1; i < 4; i++ {
		paths = append(paths,"/mnt/datanode/tmp/" + strconv.Itoa(jobID) + "/" + strconv.Itoa(i))
	}
	return paths
}

func readFile(path string, searchWiki map[string]int) {
	// TODO: read file which was written by the mapper and
	// add data to searchWiki.
	file, err := os.Open(path)
	if err != nil {
	    log.Fatal(err)
	}

	defer file.Close() // What does this line do?
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        s := strings.Fields(line)
        count,err := strconv.Atoi(s[1])
        if(err != nil){
        	log.Fatal(err)
        }
        searchWiki[s[0]] = count
	}
}

func main() {
	// TODO: get one arguments in format jobID.
	argsWithoutProg := os.Args[1:]
	jobID_str := argsWithoutProg[0]
	jobID, err := strconv.Atoi(jobID_str)
	if err != nil {
		fmt.Println("Main error")
	    log.Fatal(err)
	}
	// TODO: call getMappersFiles and save result in array.
	mappersFiles := getMappersFiles(jobID)
	// TODO: use readFile for files of mappers.
	var m = make(map[string]int)
	for i := 0; i < len(mappersFiles); i++ {
		readFile(mappersFiles[i],m)
	}
	// TODO: perform sort on results.
	sortedKeys := sortArticles(m)
	// TODO: save only results of top 100 in output file
	// in this path pattern /mnt/datanode/tmp/jobID/output.
	my_output := ""
	for i := 0; i < 100; i++ {
		if(m[sortedKeys[i]] > 0){
			my_output = my_output + sortedKeys[i]+"\n"
		}
	}
	out := []byte(my_output)
	err2 := ioutil.WriteFile("/mnt/datanode/tmp/"+jobID_str+"/output", out,0644)
	if err2 != nil {
    	panic(err)
	}
}

