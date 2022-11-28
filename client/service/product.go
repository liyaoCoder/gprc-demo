package service

import "context"

/**
* @Author:liyao
* @Description:
* @Version:
* @File:product
* @Date: 2022/11/28-17:37
 */
var Product productService

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, req *ProductReq) (*ProductResp, error) {
	prodoctStock := &ProductResp{
		ProductStock: req.ProductId,
	}
	return prodoctStock, nil
}
func (p *productService) mustEmbedUnimplementedProductServiceServer() {

}
