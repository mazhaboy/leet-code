package main

import (
	"fmt"
	"github.com/leet-code/tree"
)

func main() {

	t1 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	t2 := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &tree.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(tree.IsSameTree(t1, t2))
}

//const (
//	workers = 200
//	job     = 1000
//)
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//
//func main() {
//	requests := 100
//	ctx := context.Background()
//	ch := make(chan resp, requests)
//	wg := sync.WaitGroup{}
//	for i := 1; i <= requests; i++ {
//		wg.Add(1)
//		go func(iter int) {
//			defer wg.Done()
//			fmt.Println(iter)
//			DoReq(ctx, ch)
//		}(i)
//	}
//
//	wg.Wait()
//	for i := 1; i <= requests; i++ {
//		fmt.Println(<-ch)
//	}
//
//	log.Printf("%d Requests Completed", workers)
//}
//
//type resp struct {
//	Status int   `json:"status"`
//	Err    error `json:"err"`
//}
//
//func DoReq(ctx context.Context, ch chan<- resp) {
//	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
//	defer cancel()
//	time.Sleep(3 * time.Second)
//
//	posturl := "http://localhost:8115/api/cpms/v1/products/sku"
//
//	body := []byte(`{
//	  "sku": "41344",
//	  "price": 27900,
//	  "amount": 27900,
//	  "quantity": 1
//	}`)
//
//	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	r.Header.Add("Content-Type", "application/json")
//
//	client := &http.Client{}
//	res, err := client.Do(r)
//	if err != nil {
//		ch <- resp{
//			Status: res.StatusCode,
//			Err:    errors.New("cpms error"),
//		}
//	}
//
//	if res.StatusCode == http.StatusOK {
//		ch <- resp{
//			Status: 200,
//			Err:    nil,
//		}
//	}
//}
