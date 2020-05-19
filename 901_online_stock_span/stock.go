package stock

type StockSpanner struct {
}

func Construct() *StockSpanner {
    return &StockSpanner{}
}

func (ss *StockSpanner) Next(price int) int {
    return -1
}
