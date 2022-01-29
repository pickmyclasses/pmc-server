package es

//func Init() {
//	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
//	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
//	var err error
//	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
//		elastic.SetTraceLog(logger))
//	if err != nil {
//		panic(err)
//	}
//
//	exists, err := global.EsClient.IndexExists(model.Course{}.GetIndexName()).Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	if !exists {
//		_, err = global.EsClient.CreateIndex(model.Class{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
//		if err != nil {
//			panic(err)
//		}
//	}
//}
