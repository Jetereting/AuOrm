# AuOrm
基于gorm 实现对原生查询的[]map[string]string{}返回

支持增删改查以及事物操作

示例:
```golang
    datas, e := Raw(DB, "select * from users where user_id=?", 123)
	if e != nil {
		fmt.Println("err:", e)
	}
	fmt.Println("datas:", datas)
```

