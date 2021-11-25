package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/manhhnv/sk-gcs/pkg/setting"
	"google.golang.org/api/iterator"
	"log"
	"strings"
	"sync"
)

var client *storage.Client

/*Public functions*/

func Setup() {
	ctx := context.Background()
	var err error
	client, err = storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func ListFoldersInRoot(ctx *context.Context) (error, []string) {
	it := client.Bucket(setting.StorageSetting.BucketName).Objects(*ctx, &storage.Query{
		Delimiter: "/",
		Prefix:    "",
	})

	var folders []string

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("ListFoldersInRoot error %v", err)
			return err, []string{}
		}
		folders = append(folders, attrs.Prefix)
	}
	return nil, folders
}

func ListFiles(name, tag string) interface{} {
	var wg sync.WaitGroup
	tag = strings.ToLower(tag)

	ctx := context.Background()

	err, folders := ListFoldersInRoot(&ctx)
	if err != nil {
		panic(err)
	}
	wg.Add(len(folders))

	resultsChannel := make(chan []interface{})

	filesInFolder := func(folder string) {
		defer wg.Done()
		var result []interface{}
		it := client.Bucket(setting.StorageSetting.BucketName).Objects(ctx, &storage.Query{
			Delimiter: "",
			Prefix:    folder,
		})
		for {
			attrs, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				panic(err)
			}
			metadata := attrs.Metadata
			if len(metadata) != 0 {
				tags := metadata["type"]
				if tags != "all" && tags != "" {
					if strings.Contains(tags, tag) || strings.Contains(strings.ToLower(attrs.Prefix), strings.ToLower(name)) {
						file := map[string]interface{}{}
						file["publicUrl"] = getPublicUrl(attrs.Name)
						file["name"] = attrs.Name
						file["type"] = tags
						result = append(result, file)
					}
				} else {

				}
			}
		}
		resultsChannel <- result
	}
	for _, folder := range folders {
		go filesInFolder(folder)
	}
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	var results []interface{}
	for r := range resultsChannel {
		results = append(results, r...)
	}
	return results
}

/** Private functions*/
func getPublicUrl(name string) string {
	publicUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", setting.StorageSetting.BucketName, name)
	return publicUrl
}
