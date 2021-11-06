package dao

import (
	"fmt"

	"github.com/Spikii/go-web-microserve/models"
	"github.com/Spikii/go-web-microserve/utils"
)

//分页
func ArticlePage(pi int, ps int) ([]*models.Article, error) {
	articles := make([]*models.Article, 0, ps)
	sql := "select * from article limit ?,?"
	err := utils.Db.Select(&articles, sql, (pi-1)*ps,ps)
	if err != nil{
		return nil, err
	}
	return  articles, nil
}