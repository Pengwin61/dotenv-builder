package main

// func printLatestVersion(kv2 *api.KVv2, ctx context.Context, path string) {
// 	key, err := kv2.Get(ctx, path)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("---")
// 	fmt.Println("CURRENT DATA")
// 	for k, v := range key.Data {
// 		fmt.Println(k, v)
// 	}
// }

// func printTagretVersion(kv2 *api.KVv2, ctx context.Context, path string, version int) {
// 	old, err := kv2.GetVersion(ctx, path, version)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("---")
// 	fmt.Println("VERSION:", version)
// 	for k, v := range old.Data {
// 		fmt.Println(k, v)
// 	}
// 	fmt.Println("---")
// }
