package esLib

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"minijwc-kefu/lib"
	"os"
)

type QA struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

const indexName = "jwc_qa"
const csvPath = "./questions.csv"
const mapping = `
{
    "mappings": {
        "properties": {
            "question": {
                "type": "text",
                "analyzer": "ik_smart"
            },
            "answer": {
                "type": "text",
                "analyzer": "ik_smart"
            }
        }
    }
}
`

//NewElastic 新建一个es客户端，并自动添加索引
func NewElastic() *elastic.Client {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	client, err := elastic.NewSimpleClient(
		elastic.SetURL(lib.Config.ElasticAddr),
		elastic.SetBasicAuth(lib.Config.ElasticUser, lib.Config.ElasticPass))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(lib.Config.ElasticAddr).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	return client
}

func InsertCsv(c *elastic.Client) {
	csvData := ReadCsv()
	for i, strings := range csvData {
		if i != 0 {
			qa := QA{
				Question: strings[1],
				Answer:   strings[2],
			}
			put1, err := c.Index().
				Index(indexName).
				Id(strings[0]).
				BodyJson(qa).
				Do(context.Background())
			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("Indexed %s %s to index %s, type %s\n", indexName, put1.Id, put1.Index, put1.Type)
		}

	}
}

func GetEsAns(c *elastic.Client, question string) *[]QA {
	multiQuery := elastic.NewMultiMatchQuery(question, "question^5", "answer")
	searchResult, err := c.Search().
		Index(indexName).  // search in index "twitter"
		Query(multiQuery). // specify the query
		//Sort("question", true). // sort by "user" field, ascending
		Size(5).                 // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d %s \n", searchResult.TotalHits(), indexName)

	// Here's how you iterate through results with full control over each step.
	var res []QA

	// Iterate through results
	for _, hit := range searchResult.Hits.Hits {
		// hit.Index contains the name of the index

		// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
		var qa QA
		err := json.Unmarshal(hit.Source, &qa)
		if err != nil {
			// Deserialization failed
		}
		res = append(res, qa)

		// Work with tweet
		//fmt.Printf("Tweet by %s: %s\n", qa.Question, qa.Answer)
	}
	return &res
}

//ReadCsv csv文件读取
func ReadCsv() [][]string {
	//打开文件(只读模式)，创建io.read接口实例
	var opencast *os.File
	var err error
	opencast, err = os.Open(csvPath)
	if err != nil {
		opencast, err = os.Open("." + csvPath)
	}
	if err != nil {
		log.Panicln("打开csv失败", err)
	}
	defer opencast.Close()

	//创建csv读取接口实例
	csvFile := csv.NewReader(opencast)

	//获取一行内容，一般为第一行内容
	//read, _ := csvFile.Read() //返回切片类型：[chen  hai wei]
	//log.Println(read)

	//读取所有内容
	ReadAll, err := csvFile.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	log.Println(ReadAll)

	return ReadAll
	/*
	  说明：
	   1、读取csv文件返回的内容为切片类型，可以通过遍历的方式使用或Slicer[0]方式获取具体的值。
	   2、同一个函数或线程内，两次调用Read()方法时，第二次调用时得到的值为每二行数据，依此类推。
	   3、大文件时使用逐行读取，小文件直接读取所有然后遍历，两者应用场景不一样，需要注意。
	*/

}
